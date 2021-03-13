package addresses

import (
	"github.com/urfave/cli"
	
	"github.com/p9c/pod/app/save"
	"github.com/p9c/pod/cmd/node/state"
	"github.com/p9c/pod/pkg/pod"
	"github.com/p9c/pod/pkg/wallet"
	wm "github.com/p9c/pod/pkg/wallet/waddrmgr"
)

// RefillMiningAddresses adds new addresses to the mining address pool for the miner
// todo: make this remove ones that have been used or received a payment or mined
func RefillMiningAddresses(w *wallet.Wallet, cfg *pod.Config, stateCfg *state.Config) {
	if w == nil {
		dbg.Ln("trying to refill without a wallet")
		return
	}
	if cfg == nil {
		dbg.Ln("config is empty")
		return
	}
	var miningAddressLen int
	if cfg.MiningAddrs != nil {
		dbg.Ln("miningAddressLen", len(*cfg.MiningAddrs))
		miningAddressLen = len(*cfg.MiningAddrs)
	} else {
		dbg.Ln("miningaddrs slice is missing")
		cfg.MiningAddrs = new(cli.StringSlice)
	}
	toMake := 99 - miningAddressLen
	if miningAddressLen >= 99 {
		toMake = 0
	}
	if toMake < 1 {
		dbg.Ln("not making any new addresses")
		return
	}
	dbg.Ln("refilling mining addresses")
	account, e := w.AccountNumber(
		wm.KeyScopeBIP0044,
		"default",
	)
	if e != nil  {
		Error("error getting account number ", err)
	}
	for i := 0; i < toMake; i++ {
		addr, e := w.NewAddress(
			account, wm.KeyScopeBIP0044,
			true,
		)
		if e ==  nil {
			// add them to the configuration to be saved
			*cfg.MiningAddrs = append(*cfg.MiningAddrs, addr.EncodeAddress())
			// add them to the active mining address list so they
			// are ready to use
			stateCfg.ActiveMiningAddrs = append(stateCfg.ActiveMiningAddrs, addr)
		} else {
			Error("error adding new address ", err)
		}
	}
	if save.Pod(cfg) {
		dbg.Ln("saved config with new addresses")
	} else {
		Error("error adding new addresses", err)
	}
}