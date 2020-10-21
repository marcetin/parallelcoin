package dap

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/p9c/pod/pkg/gui/wallet/dap/mod"
	"github.com/p9c/pod/pkg/gui/wallet/dap/res"
	"github.com/p9c/pod/pkg/gui/wallet/dap/win"
	"github.com/p9c/pod/pkg/gui/wallet/nav"
	"github.com/p9c/pod/pkg/gui/wallet/theme"
)

var (
	noReturn = func(gtx C) D { return D{} }
)

type (
	D = layout.Dimensions
	C = layout.Context
	W = layout.Widget
)
type dap struct {
	boot mod.Dap
}

func NewDap(title string, rc interface{}) dap {
	//if cfg.Initial {
	//	fmt.Println("running initial setup")
	//}
	d := mod.Dap{
		Rc:   RcInit(rc),
		Apps: make(map[string]mod.Sap),
	}

	d.UI = &mod.UserInterface{
		Theme: theme.NewTheme(),
		//mob:   make(chan bool),
	}
	w := map[string]*win.Window{
		"main": &win.Window{
			W: app.NewWindow(
				app.Size(unit.Dp(1024), unit.Dp(800)),
				app.Title(title),
			)},
	}
	d.UI.W = &win.Windows{
		W: w,
	}
	n := &nav.Navigation{
		Name:         "Navigacion",
		Bg:           d.UI.Theme.Colors["NavBg"],
		ItemIconSize: unit.Px(24),
	}
	d.UI.N = n

	s := &mod.Settings{
		//Dir: appdata.Dir("dap", false),
	}
	d.S = s

	d.UI.R = res.Resposnsivity(0, 0)
	Debug("New DAP", d)

	return dap{boot: d}
}

func (d *dap) NewSap(s mod.Sap) {
	d.boot.Apps[s.Title] = s
	return
}
func (d *dap) BOOT() *mod.Dap {
	return &d.boot
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
}

func RcInit(w interface{}) (r *mod.RcVar) {
	b := mod.Boot{
		IsBoot:     true,
		IsFirstRun: false,
		IsBootMenu: false,
		IsBootLogo: false,
		IsLoading:  false,
	}
	// d := models.DuoUIdialog{
	//	Show:   true,
	//	Ok:     func() { r.Dialog.Show = false },
	//	Cancel: func() { r.Dialog.Show = false },
	//	Title:  "Dialog!",
	//	Text:   "Dialog text",
	// }
	//l := new(model.DuoUIlog)

	r = &mod.RcVar{
		Worker: w,
		//db:          new(DuoUIdb),
		Boot: &b,
		//AddressBook: new(model.DuoUIaddressBook),
		//Status: &model.DuoUIstatus{
		//	Node: &model.NodeStatus{},
		//	Wallet: &model.WalletStatus{
		//		WalletVersion: make(map[string]btcjson.VersionResult),
		//		LastTxs:       &model.DuoUItransactionsExcerpts{},
		//	},
		//	Kopach: &model.KopachStatus{},
		//},
		//Dialog:   &model.DuoUIdialog{},
		//Settings: settings(cx),
		//Log:      l,
		//Quit:  make(chan struct{}),
		//Ready: make(chan struct{}),
	}
	//r.db.DuoUIdbInit(r.cx.DataDir)
	return
}