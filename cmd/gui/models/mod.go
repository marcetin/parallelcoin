package models

import (
	"github.com/p9c/pod/cmd/gui/theme"
	"github.com/p9c/pod/cmd/gui/widget"
	"github.com/p9c/pod/pkg/gio/app"
	"github.com/p9c/pod/pkg/gio/layout"
	"github.com/p9c/pod/pkg/pod"
	"image/color"
	"time"
)

type DuoUI struct {
	DuoUIboot          *Boot
	DuoUIwindow        *app.Window
	DuoUIcontext       *layout.Context
	DuoUItheme         *theme.DuoUItheme
	DuoUIconstraints   *layout.Constraints
	DuoUIico           map[string]*theme.DuoUIicon
	DuoUIcomponents    *DuoUIcomponents
	DuoUIconfiguration *DuoUIconfiguration
	Quit               chan struct{}
	Ready              chan struct{}
	IsReady            bool
	DuoUIready         chan struct{}
	DuoUIisReady       bool
	CurrentPage        string
}

type DuoUIcomponents struct {
	View   DuoUIcomponent
	Header DuoUIcomponent
	Footer DuoUIcomponent
	// Intro              DuoUIcomponent
	Logo DuoUIcomponent
	// Log                DuoUIcomponent
	Body           DuoUIcomponent
	Sidebar        DuoUIcomponent
	Menu           DuoUIcomponent
	Content        DuoUIcomponent
	Overview       DuoUIcomponent
	Send           DuoUIcomponent
	SendButtons    DuoUIcomponent
	Receive        DuoUIcomponent
	ReceiveButtons DuoUIcomponent
	Status         DuoUIcomponent
	StatusItem     DuoUIcomponent
	History        DuoUIcomponent
	AddressBook    DuoUIcomponent
	Explorer       DuoUIcomponent
	Network        DuoUIcomponent
	Console        DuoUIcomponent
	// ConsoleOutput      DuoUIcomponent
	// ConsoleInput       DuoUIcomponent
	Settings DuoUIcomponent
}

type DuoUIcomponent struct {
	Layout layout.Flex
	Inset  layout.Inset
}

type Boot struct {
	IsBoot     bool   `json:"boot"`
	IsFirstRun bool   `json:"firstrun"`
	IsBootMenu bool   `json:"menu"`
	IsBootLogo bool   `json:"logo"`
	IsLoading  bool   `json:"loading"`
	IsScreen   string `json:"screen"`
}

type DuoUIconfiguration struct {
	Abbrevation        string
	PrimaryTextColor   color.RGBA
	SecondaryTextColor color.RGBA
	PrimaryBgColor     color.RGBA
	SecondaryBgColor   color.RGBA
	Navigations        map[string]*theme.DuoUIthemeNav
	Tabs               DuoUIconfTabs
	Settings           DuoUIsettings
}

type DuoUIconfTabs struct {
	Current  string
	TabsList map[string]*widget.Button
}

type DuoUIalert struct {
	Time      time.Time   `json:"time"`
	Title     string      `json:"title"`
	Message   interface{} `json:"message"`
	AlertType string      `json:"type"`
}

type DuoUIsettings struct {
	//db DuoUIdb
	//Display mod.DisplayConfig `json:"display"`
	Daemon DaemonConfig `json:"daemon"`
}

type DaemonConfig struct {
	Config  *pod.Config `json:"config"`
	Schema  pod.Schema  `json:"schema"`
	Widgets map[string]interface{}
}
