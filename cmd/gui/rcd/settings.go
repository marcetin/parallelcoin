package rcd

import (
	"fmt"
	"github.com/p9c/pod/cmd/gui/mvc/controller"
	"github.com/p9c/pod/cmd/gui/mvc/model"
	"github.com/p9c/pod/pkg/conte"
	"github.com/p9c/pod/pkg/pod"
)

func (r *RcVar) SaveDaemonCfg() {
	//save.Pod(r.Settings.Daemon.Config)
}

func settings(cx *conte.Xt) *model.DuoUIsettings {

	settings := &model.DuoUIsettings{
		Abbrevation: "DUO",
		Tabs: &model.DuoUIconfTabs{
			Current:  "wallet",
			TabsList: make(map[string]*controller.Button),
		},
		Daemon: &model.DaemonConfig{
			Config: cx.Config,
			Schema: pod.GetConfigSchema(cx.Config, cx.ConfigMap),
		},
	}
	// Settings tabs

	settingsFields := make(map[string]interface{})
	for _, group := range settings.Daemon.Schema.Groups {
		settings.Tabs.TabsList[group.Legend] = new(controller.Button)
		for _, field := range group.Fields {
			switch field.Type {
			case "array":
				settingsFields[field.Label] = new(controller.Button)
			case "input":
				settingsFields[field.Label] = &controller.Editor{
					SingleLine: true,
				}
				//if field.Value != nil {
				switch field.InputType {
				case "text":
					(settingsFields[field.Label]).(*controller.Editor).SetText(fmt.Sprint(*cx.ConfigMap[field.Model].(*string)))
					//(settingsFields[field.Label]).(*controller.Editor).SetText(fmt.Sprint(*field.Value.(*string)))
				case "number":
					//(settingsFields[field.Label]).(*controller.Editor).SetText(fmt.Sprint(*field.Value.(*int)))
				case "decimal":
					//(settingsFields[field.Label]).(*controller.Editor).SetText(fmt.Sprint(*field.Value.(*float64)))
				case "time":
					//(settingsFields[field.Label]).(*controller.Editor).SetText(fmt.Sprint(*field.Value.(*time.Duration)))
				}
				//}
			case "switch":
				settingsFields[field.Label] = new(controller.CheckBox)
				(settingsFields[field.Label]).(*controller.CheckBox).SetChecked(*cx.ConfigMap[field.Model].(*bool))
			case "radio":
				settingsFields[field.Label] = new(controller.Enum)
			default:
				settingsFields[field.Label] = new(controller.Button)
			}
		}
	}
	settings.Daemon.Widgets = settingsFields
	return settings
}
