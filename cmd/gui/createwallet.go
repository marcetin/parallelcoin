package gui

import (
	"encoding/hex"
	"fmt"
	"os"
	"time"
	
	l "gioui.org/layout"
	"github.com/urfave/cli"
	
	"github.com/p9c/pod/app/save"
	"github.com/p9c/pod/pkg/blockchain/chaincfg/netparams"
	"github.com/p9c/pod/pkg/blockchain/fork"
	"github.com/p9c/pod/pkg/gui"
	"github.com/p9c/pod/pkg/util/hdkeychain"
	"github.com/p9c/pod/pkg/wallet"
)

const slash = string(os.PathSeparator)

func (wg *WalletGUI) CreateWalletPage(gtx l.Context) l.Dimensions {
	return wg.Fill(
		"PanelBg", l.Center, 0, 0, wg.Inset(
			0.5,
			wg.Flex().
				SpaceAround().
				Flexed(0.5, gui.EmptyMaxHeight()).
				Rigid(
					func(gtx l.Context) l.Dimensions {
						return wg.VFlex().
							AlignMiddle().
							SpaceSides().
							Rigid(
								wg.H4("create new wallet").
									Color("PanelText").
									Fn,
							).
							Rigid(
								wg.Inset(
									0.25,
									wg.passwords["passEditor"].Fn,
								).
									Fn,
							).
							Rigid(
								wg.Inset(
									0.25,
									wg.passwords["confirmPassEditor"].Fn,
								).
									Fn,
							).
							Rigid(
								wg.Inset(
									0.25,
									wg.inputs["walletSeed"].Fn,
								).
									Fn,
							).
							Rigid(
								wg.Inset(
									0.25,
									func(gtx l.Context) l.Dimensions {
										// gtx.Constraints.Min.X = int(wg.TextSize.Scale(16).V)
										return wg.CheckBox(
											wg.bools["testnet"].SetOnChange(
												wg.createWalletTestnetToggle,
											),
										).
											IconColor("Primary").
											TextColor("DocText").
											Text("Use testnet?").
											Fn(gtx)
									},
								).Fn,
							).
							Rigid(
								wg.Body1("your seed").
									Color("PanelText").
									Fn,
							).
							Rigid(
								func(gtx l.Context) l.Dimensions {
									gtx.Constraints.Max.X = int(wg.TextSize.Scale(22).V)
									return wg.Caption(wg.inputs["walletSeed"].GetText()).
										Font("go regular").
										TextScale(0.66).
										Fn(gtx)
								},
							).
							Rigid(
								wg.Inset(
									0.5,
									func(gtx l.Context) l.Dimensions {
										gtx.Constraints.Max.X = int(wg.TextSize.Scale(32).V)
										gtx.Constraints.Min.X = int(wg.TextSize.Scale(16).V)
										return wg.CheckBox(
											wg.bools["ihaveread"].SetOnChange(
												func(b bool) {
													dbg.Ln("confirmed read", b)
													// if the password has been entered, we need to copy it to the variable
													if wg.createWalletPasswordsMatch() {
														wg.cx.Config.Lock()
														*wg.cx.Config.WalletPass = wg.passwords["confirmPassEditor"].GetPassword()
														wg.cx.Config.Unlock()
													}
												},
											),
										).
											IconColor("Primary").
											TextColor("DocText").
											Text(
												"I have stored the seed and password safely " +
													"and understand it cannot be recovered",
											).
											Fn(gtx)
									},
								).Fn,
							).
							Rigid(
								func(gtx l.Context) l.Dimensions {
									if !wg.createWalletInputsAreValid() {
										gtx = gtx.Disabled()
									}
									return wg.Flex().
										Rigid(
											wg.Button(wg.clickables["createWallet"]).
												Background("Primary").
												Color("Light").
												SetClick(
													func() {
														go wg.createWalletAction()
													},
												).
												CornerRadius(0).
												Inset(0.5).
												Text("create wallet").
												Fn,
										).
										Fn(gtx)
								},
							).
							
							Fn(gtx)
					},
				).
				Flexed(0.5, gui.EmptyMaxWidth()).Fn,
		).Fn,
	).Fn(gtx)
}

func (wg *WalletGUI) createWalletPasswordsMatch() bool {
	return wg.passwords["passEditor"].GetPassword() != "" ||
		wg.passwords["confirmPassEditor"].GetPassword() != "" ||
		len(wg.passwords["passEditor"].GetPassword()) >= 8 ||
		wg.passwords["passEditor"].GetPassword() ==
			wg.passwords["confirmPassEditor"].GetPassword()
}

func (wg *WalletGUI) createWalletInputsAreValid() bool {
	var b []byte
	var e error
	seedValid := true
	if b, e = hex.DecodeString(wg.inputs["walletSeed"].GetText()); err.Chk(e) {
		seedValid = false
	} else if len(b) != 0 && len(b) < hdkeychain.MinSeedBytes ||
		len(b) > hdkeychain.MaxSeedBytes {
		seedValid = false
	}
	return wg.createWalletPasswordsMatch() && seedValid && wg.bools["ihaveread"].GetValue()
}

func (wg *WalletGUI) createWalletAction() {
	// wg.NodeRunCommandChan <- "stop"
	dbg.Ln("clicked submit wallet")
	*wg.cx.Config.WalletFile = *wg.cx.Config.DataDir +
		string(os.PathSeparator) + wg.cx.ActiveNet.Name +
		string(os.PathSeparator) + wallet.DbName
	dbDir := *wg.cx.Config.WalletFile
	loader := wallet.NewLoader(wg.cx.ActiveNet, dbDir, 250)
	seed, _ := hex.DecodeString(wg.inputs["walletSeed"].GetText())
	pass := []byte(wg.passwords["passEditor"].GetPassword())
	*wg.cx.Config.WalletPass = string(pass)
	dbg.Ln("password", string(pass))
	save.Pod(wg.cx.Config)
	w, e := loader.CreateNewWallet(
		pass,
		pass,
		seed,
		time.Now(),
		false,
		wg.cx.Config,
		nil,
	)
	dbg.Ln("*** created wallet")
	if err.Chk(e) {
		// return
	}
	// dbg.Ln("refilling mining addresses")
	// addresses.RefillMiningAddresses(
	// 	w,
	// 	wg.cx.Config,
	// 	wg.cx.StateCfg,
	// )
	// wrn.Ln("done refilling mining addresses")
	w.Stop()
	dbg.Ln("shutting down wallet", w.ShuttingDown())
	w.WaitForShutdown()
	dbg.Ln("starting main app")
	*wg.cx.Config.Generate = true
	*wg.cx.Config.GenThreads = 1
	*wg.cx.Config.NodeOff = false
	*wg.cx.Config.WalletOff = false
	save.Pod(wg.cx.Config)
	// if *wg.cx.Config.Generate {
	// 	wg.miner.Start()
	// }
	*wg.noWallet = false
	// wg.node.Start()
	// if e = wg.writeWalletCookie(); err.Chk(e) {
	// }
	// wg.wallet.Start()
}

func (wg *WalletGUI) createWalletTestnetToggle(b bool) {
	dbg.Ln("testnet on?", b)
	// if the password has been entered, we need to copy it to the variable
	if wg.passwords["passEditor"].GetPassword() != "" ||
		wg.passwords["confirmPassEditor"].GetPassword() != "" ||
		len(wg.passwords["passEditor"].GetPassword()) >= 8 ||
		wg.passwords["passEditor"].GetPassword() ==
			wg.passwords["confirmPassEditor"].GetPassword() {
		*wg.cx.Config.WalletPass = wg.passwords["confirmPassEditor"].GetPassword()
		dbg.Ln("wallet pass", *wg.cx.Config.WalletPass)
	}
	if b {
		wg.cx.ActiveNet = &netparams.TestNet3Params
		fork.IsTestnet = true
	} else {
		wg.cx.ActiveNet = &netparams.MainNetParams
		fork.IsTestnet = false
	}
	inf.Ln("activenet:", wg.cx.ActiveNet.Name)
	dbg.Ln("setting ports to match network")
	*wg.cx.Config.Network = wg.cx.ActiveNet.Name
	// _, routeableAddress, _ := routeable.GetInterface()
	*wg.cx.Config.P2PListeners = cli.StringSlice{
		fmt.Sprintf(
			"0.0.0.0:" + wg.cx.ActiveNet.DefaultPort,
		),
	}
	address := fmt.Sprintf(
		"127.0.0.1:%s",
		wg.cx.ActiveNet.RPCClientPort,
	)
	*wg.cx.Config.RPCListeners = cli.StringSlice{address}
	*wg.cx.Config.RPCConnect = address
	address = fmt.Sprintf("127.0.0.1:" + wg.cx.ActiveNet.WalletRPCServerPort)
	*wg.cx.Config.WalletRPCListeners = cli.StringSlice{address}
	*wg.cx.Config.WalletServer = address
	*wg.cx.Config.NodeOff = false
	save.Pod(wg.cx.Config)
}
