package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jessevdk/go-flags"

	"github.com/p9c/pod/app/appdata"
	"github.com/p9c/pod/pkg/blockchain/chaincfg/netparams"
	"github.com/p9c/pod/pkg/blockchain/wire"
	database "github.com/p9c/pod/pkg/database"
	_ "github.com/p9c/pod/pkg/database/ffldb"
)

const (
	minCandidates        = 1
	maxCandidates        = 20
	defaultNumCandidates = 5
	defaultDbType        = "ffldb"
)

var (
	podHomeDir      = appdata.Dir("pod", false)
	defaultDataDir  = filepath.Join(podHomeDir, "data")
	knownDbTypes    = database.SupportedDrivers()
	activeNetParams = &netparams.MainNetParams
)

// config defines the configuration options for findcheckpoint. See loadConfig for details on the configuration load
// process.
type config struct {
	DataDir        string `short:"b" long:"datadir" description:"Location of the pod data directory"`
	DbType         string `long:"dbtype" description:"Database backend to use for the Block Chain"`
	TestNet3       bool   `long:"testnet" description:"Use the test network"`
	RegressionTest bool   `long:"regtest" description:"Use the regression test network"`
	SimNet         bool   `long:"simnet" description:"Use the simulation test network"`
	NumCandidates  int    `short:"n" long:"numcandidates" description:"Max num of checkpoint candidates to show {1-20}"`
	UseGoOutput    bool   `short:"g" long:"gooutput" description:"Display the candidates using Go syntax that is ready to insert into the btcchain checkpoint list"`
}

// validDbType returns whether or not dbType is a supported database type.
func validDbType(
	dbType string) bool {
	for _, knownType := range knownDbTypes {
		if dbType == knownType {
			return true
		}
	}
	return false
}

// netName returns the name used when referring to a bitcoin network. At the time of writing, pod currently places
// blocks for testnet version 3 in the data and log directory "testnet", which does not match the Name field of the
// chaincfg parameters. This function can be used to override this directory name as "testnet" when the passed active
// network matches wire.TestNet3. A proper upgrade to move the data and log directories for this network to "testnet3"
// is planned for the future, at which point this function can be removed and the network parameter's name used instead.
func netName(
	chainParams *netparams.Params) string {
	switch chainParams.Net {
	case wire.TestNet3:
		return "testnet"
	default:
		return chainParams.Name
	}
}

// loadConfig initializes and parses the config using command line options.
func loadConfig() (*config, []string, error) {
	// Default config.
	cfg := config{
		DataDir:       defaultDataDir,
		DbType:        defaultDbType,
		NumCandidates: defaultNumCandidates,
	}
	// Parse command line options.
	parser := flags.NewParser(&cfg, flags.Default)
	remainingArgs, e := parser.Parse()
	if e != nil  {
				if e, ok := err.(*flags.Error); !ok || e.Type != flags.ErrHelp {
			parser.WriteHelp(os.Stderr)
		}
		return nil, nil, e
	}
	// Multiple networks can't be selected simultaneously.
	funcName := "loadConfig"
	numNets := 0
	// Count number of network flags passed; assign active network netparams while we're at it
	if cfg.TestNet3 {
		numNets++
		activeNetParams = &netparams.TestNet3Params
	}
	if cfg.RegressionTest {
		numNets++
		activeNetParams = &netparams.RegressionTestParams
	}
	if cfg.SimNet {
		numNets++
		activeNetParams = &netparams.SimNetParams
	}
	if numNets > 1 {
		str := "%s: The testnet, regtest, and simnet netparams can't be " +
			"used together -- choose one of the three"
		e := fmt.Errorf(str, funcName)
		_, _ = fmt.Fprintln(os.Stderr, err)
		parser.WriteHelp(os.Stderr)
		return nil, nil, e
	}
	// Validate database type.
	if !validDbType(cfg.DbType) {
		str := "%s: The specified database type [%v] is invalid -- " +
			"supported types %v"
		e := fmt.Errorf(str, "loadConfig", cfg.DbType, knownDbTypes)
		_, _ = fmt.Fprintln(os.Stderr, err)
		parser.WriteHelp(os.Stderr)
		return nil, nil, e
	}
	// Append the network type to the data directory so it is "namespaced" per network. In addition to the block
	// database, there are other pieces of data that are saved to disk such as address manager state. All data is
	// specific to a network, so namespacing the data directory means each individual piece of serialized data does not
	// have to worry about changing names per network and such.
	cfg.DataDir = filepath.Join(cfg.DataDir, netName(activeNetParams))
	// Validate the number of candidates.
	if cfg.NumCandidates < minCandidates || cfg.NumCandidates > maxCandidates {
		str := "%s: The specified number of candidates is out of " +
			"range -- parsed [%v]"
		e = fmt.Errorf(str, "loadConfig", cfg.NumCandidates)
		_, _ = fmt.Fprintln(os.Stderr, err)
		parser.WriteHelp(os.Stderr)
		return nil, nil, e
	}
	return &cfg, remainingArgs, nil
}
