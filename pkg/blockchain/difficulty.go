package blockchain

import (
	"encoding/hex"
	"math/big"
	"strings"
	"time"
	
	"github.com/p9c/pod/pkg/blockchain/fork"
	chainhash "github.com/p9c/pod/pkg/blockchain/chainhash"
)

var (
	// ScryptPowLimit is
	ScryptPowLimit = scryptPowLimit
	// ScryptPowLimitBits is
	ScryptPowLimitBits = BigToCompact(&scryptPowLimit)
	// bigOne is 1 represented as a big.Int. It is defined here to avoid the overhead of creating it multiple times.
	bigOne = big.NewInt(1)
	// oneLsh256 is 1 shifted left 256 bits. It is defined here to avoid the overhead of creating it multiple times.
	oneLsh256      = new(big.Int).Lsh(bigOne, 256)
	scryptPowLimit = func() big.Int {
		mplb, _ := hex.DecodeString(
			"000000039fcaa04ac30b6384471f337748ef5c87c7aeffce5e51770ce6283137,")
		return *big.NewInt(0).SetBytes(mplb) // AllOnes.Rsh(&AllOnes, 0)
	}()
)

// CalcNextRequiredDifficulty calculates the required difficulty for the block after the end of the current best chain
// based on the difficulty retarget rules. This function is safe for concurrent access.
func (b *BlockChain) CalcNextRequiredDifficulty(algo string) (difficulty uint32, e error) {
	b.ChainLock.Lock()
	difficulty, e = b.CalcNextRequiredDifficultyFromNode(
		b.BestChain.
			Tip(), algo, false,
	)
	// F.Ln("CalcNextRequiredDifficulty", difficulty)
	b.ChainLock.Unlock()
	return
}

// calcEasiestDifficulty calculates the easiest possible difficulty that a block can have given starting difficulty bits
// and a duration.
//
// It is mainly used to verify that claimed proof of work by a block is sane as compared to a known good checkpoint.
func (b *BlockChain) calcEasiestDifficulty(bits uint32, duration time.Duration) uint32 {
	// Convert types used in the calculations below.
	durationVal := int64(duration / time.Second)
	adjustmentFactor := big.NewInt(b.params.RetargetAdjustmentFactor)
	// Since easier difficulty equates to higher numbers, the easiest difficulty for a given duration is the largest
	// value possible given the number of retargets for the duration and starting difficulty multiplied by the max
	// adjustment factor.
	newTarget := fork.CompactToBig(bits)
	for durationVal > 0 && newTarget.Cmp(b.params.PowLimit) < 0 {
		newTarget.Mul(newTarget, adjustmentFactor)
		durationVal -= b.maxRetargetTimespan
	}
	// Limit new value to the proof of work limit.
	if newTarget.Cmp(b.params.PowLimit) > 0 {
		newTarget.Set(b.params.PowLimit)
	}
	return BigToCompact(newTarget)
}

// CalcNextRequiredDifficultyFromNode calculates the required difficulty for the block after the passed previous block node
// based on the difficulty retarget rules.
//
// This function differs from the exported CalcNextRequiredDifficulty in that the exported version uses the current best
// chain as the previous block node while this function accepts any block node.
func (b *BlockChain) CalcNextRequiredDifficultyFromNode(lastNode *BlockNode, algoname string, l bool,) (
	newTargetBits uint32,
	e error,
) {
	nH := lastNode.height + 1
	cF := fork.GetCurrent(nH)
	newTargetBits = fork.GetMinBits(algoname, nH)
	// Tracef("CalcNextRequiredDifficultyFromNode %08x", newTargetBits)
	switch cF {
	// Legacy difficulty adjustment
	case 0:
		// F.Ln("before hardfork")
		return b.CalcNextRequiredDifficultyHalcyon(lastNode, algoname, l)
	// Plan 9 from Crypto Space
	case 1:
		bits, ok := lastNode.Diffs.Load().(Diffs)
		if bits == nil || !ok {
			lastNode.Diffs.Store(make(Diffs))
		}
		version := fork.GetAlgoVer(algoname, lastNode.height+1)
		if bits[version] == 0 {
			bits, e = b.CalcNextRequiredDifficultyPlan9Controller(lastNode)
			if e != nil  {
				E.Ln(e)
				return
			}
			// D.Ln(bits, reflect.TypeOf(bits))
			b.DifficultyBits.Store(bits)
			// D.F("got difficulty %d %08x %+v", version, (*b.DifficultyBits)[version], *bits)
		}
		newTargetBits = bits[version]
		return
	}
	return
}

// RightJustify takes a string and right justifies it by a width or crops it
func RightJustify(s string, w int) string {
	sw := len(s)
	diff := w - sw
	if diff > 0 {
		s = strings.Repeat(" ", diff) + s
	} else if diff < 0 {
		s = s[:w]
	}
	return s
}

// BigToCompact converts a whole number N to a compact representation using an unsigned 32-bit number.
//
// The compact representation only provides 23 bits of precision, so values larger than (2^23 - 1) only encode the most
// significant digits of the number. See CompactToBig for details.
func BigToCompact(n *big.Int) uint32 {
	// No need to do any work if it's zero.
	if n.Sign() == 0 {
		return 0
	}
	// Since the base for the exponent is 256, the exponent can be treated as the number of bytes. So, shift the number
	// right or left accordingly. This is equivalent to: mantissa = mantissa / 256^(exponent-3)
	var mantissa uint32
	exponent := uint(len(n.Bytes()))
	if exponent <= 3 {
		mantissa = uint32(n.Bits()[0])
		mantissa <<= 8 * (3 - exponent)
	} else {
		// Use a copy to avoid modifying the caller's original number.
		tn := new(big.Int).Set(n)
		mantissa = uint32(tn.Rsh(tn, 8*(exponent-3)).Bits()[0])
	}
	// When the mantissa already has the sign bit set, the number is too large to fit into the available 23-bits, so
	// divide the number by 256 and increment the exponent accordingly.
	if mantissa&0x00800000 != 0 {
		mantissa >>= 8
		exponent++
	}
	// Pack the exponent, sign bit, and mantissa into an unsigned 32-bit int and return it.
	compact := uint32(exponent<<24) | mantissa
	if n.Sign() < 0 {
		compact |= 0x00800000
	}
	return compact
}

// CalcWork calculates a work value from difficulty bits.
// Bitcoin increases the difficulty for generating a block by decreasing the
// value which the generated hash must be less than.
// This difficulty target is stored in each block header using a compact
// representation as described in the documentation for CompactToBig.
// The main chain is selected by choosing the chain that has the most proof
// of work (highest difficulty).
// Since a lower target difficulty value equates to higher actual difficulty,
// the work value which will be accumulated must be the inverse of the
// difficulty.  Also,
// in order to avoid potential division by zero and really small floating
// point numbers, the result adds 1 to the denominator and multiplies the
// numerator by 2^256.
func CalcWork(bits uint32, height int32, algover int32) *big.Int {
	// Return a work value of zero if the passed difficulty bits represent a negative number. Note this should not
	// happen in practice with valid blocks, but an invalid block could trigger it.
	difficultyNum := CompactToBig(bits)
	// To make the difficulty values correlate to number of hash operations, multiply this difficulty base by the
	// nanoseconds/hash figures in the fork algorithms list
	if difficultyNum.Sign() <= 0 {
		return big.NewInt(0)
	}
	denominator := new(big.Int).Add(difficultyNum, bigOne)
	r := new(big.Int).Div(oneLsh256, denominator)
	return r
}

// CompactToBig converts a compact representation of a whole number N to an unsigned 32-bit number. The representation
// is similar to IEEE754 floating point numbers.
/*
   Like IEEE754 floating point, there are three basic components: the sign, the exponent, and the mantissa. They are
   broken out as follows:

   	* the most significant 8 bits represent the unsigned base 256 exponent

   	* bit 23 (the 24th bit) represents the sign bit

   	* the least significant 23 bits represent the mantissa

   	-------------------------------------------------
   	|   Exponent     |    Sign    |    Mantissa     |
   	-------------------------------------------------
   	| 8 bits [31-24] | 1 bit [23] | 23 bits [22-00] |
   	-------------------------------------------------

   The formula to calculate N is:

   	N = (-1^sign) * mantissa * 256^(exponent-3)

   This compact form is only used in bitcoin to encode unsigned 256-bit numbers which represent difficulty targets, thus
   there really is not a need for a sign bit, but it is implemented here to stay consistent with bitcoind.
*/
func CompactToBig(compact uint32) *big.Int {
	// Extract the mantissa, sign bit, and exponent.
	mantissa := compact & 0x007fffff
	isNegative := compact&0x00800000 != 0
	exponent := uint(compact >> 24)
	// Since the base for the exponent is 256, the exponent can be treated as the number of bytes to represent the full
	// 256-bit number. So, treat the exponent as the number of bytes and shift the mantissa right or left accordingly.
	// This is equivalent to: N = mantissa * 256^( exponent-3)
	var bn *big.Int
	if exponent <= 3 {
		mantissa >>= 8 * (3 - exponent)
		bn = big.NewInt(int64(mantissa))
	} else {
		bn = big.NewInt(int64(mantissa))
		bn.Lsh(bn, 8*(exponent-3))
	}
	// Make it negative if the sign bit is set.
	if isNegative {
		bn = bn.Neg(bn)
	}
	return bn
}

// HashToBig converts a chainhash.Hash into a big. Int that can be used to perform math comparisons.
func HashToBig(hash *chainhash.Hash) *big.Int {
	// A Hash is in little-endian, but the big package wants the bytes in big-endian, so reverse them.
	buf := *hash
	blen := len(buf)
	for i := 0; i < blen/2; i++ {
		buf[i], buf[blen-1-i] = buf[blen-1-i], buf[i]
	}
	// buf := hash.CloneBytes()
	return new(big.Int).SetBytes(buf[:])
}
