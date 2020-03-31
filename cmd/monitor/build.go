package monitor

import (
	"gioui.org/layout"
	"github.com/p9c/pod/pkg/gel"
)

func (s *State) BuildButtons() layout.FlexChild {
	return Rigid(func() {
		if s.WindowWidth >= 360 || !s.Config.FilterOpen {
			s.FlexH(Rigid(func() {
				bg, fg := "PanelBg", "PanelText"
				if s.Config.BuildOpen {
					bg, fg = "DocBg", "DocText"
				}
				//s.TextButton("Build", "Secondary", 23,
				//	fg, bg, s.BuildFoldButton)
				s.IconButton("Build", fg, bg, &s.BuildFoldButton)
				for s.BuildFoldButton.Clicked(s.Gtx) {
					Debug("run mode folder clicked")
					if !s.Config.BuildOpen {
						s.Config.FilterOpen = false
						s.Config.SettingsOpen = false
					}
					s.Config.BuildOpen = !s.Config.BuildOpen
					s.SaveConfig()
				}
			}),
			)
		}
	})
}

func (s *State) BuildPage() layout.FlexChild {
	if !s.Config.BuildOpen {
		return Flexed(0, func() {})
	}
	var weight float32 = 0.5
	switch {
	case s.Config.BuildZoomed:
		weight = 1
	case s.WindowHeight <= 800 && s.WindowWidth <= 800:
		weight = 1
	case s.WindowHeight <= 600 && s.WindowWidth > 800:
		weight = 1
	}
	return Flexed(weight, func() {
		cs := s.Gtx.Constraints
		s.Rectangle(cs.Width.Max, cs.Height.Max, "DocBg", "ff")
		s.FlexV(Rigid(func() {
			s.Rectangle(cs.Width.Max, cs.Height.Max, "DocBg", "ff")
			s.Inset(4, func() {})
		}), Rigid(func() {
			s.FlexH(Rigid(func() {
				s.Label("Build Configuration")
			}), Spacer(), Rigid(func() {
				if !(s.WindowHeight <= 800 && s.WindowWidth <= 800 ||
					s.WindowHeight <= 600 && s.WindowWidth > 800) {
					ic := "zoom"
					if s.Config.BuildZoomed {
						ic = "minimize"
					}
					s.IconButton(ic, "DocText", "DocBg",
						&s.BuildZoomButton)
					for s.BuildZoomButton.Clicked(s.Gtx) {
						Debug("settings panel fold button clicked")
						s.Config.BuildZoomed = !s.Config.BuildZoomed
						s.SaveConfig()
					}
				}
			}), Spacer(), Rigid(func() {
				s.IconButton("foldIn", "DocText", "DocBg",
					&s.BuildCloseButton)
				for s.BuildCloseButton.Clicked(s.Gtx) {
					Debug("settings panel close button clicked")
					s.Config.BuildOpen = false
					s.SaveConfig()
				}
			}),
			)
		}), Flexed(1, func() {
			cs := s.Gtx.Constraints
			s.Rectangle(cs.Width.Max, cs.Height.Max, "PanelBg", "ff")
			s.FlexV(Flexed(1, func() {
				s.Inset(8, func() {
					// cs := s.Gtx.Constraints
					// s.Rectangle(cs.Width.Max, cs.Height.Max, "DocBg")
					//if s.Config.BuildOpen {
					s.BuildConfigPage()
					//}
				})
			}))
		}), Rigid(func() {
			s.Rectangle(cs.Width.Max, cs.Height.Max, "DocBg", "ff")
			s.Inset(4, func() {})
		}),
		)
	})
}

func (s *State) BuildConfigPage() {
	s.FlexV(
		//s.FlexH(
		Rigid(func() {
			s.Inset(4, func() {
				s.FlexH(
					Rigid(func() {
						s.Inset(8,
							s.Text("Run in", "PanelText", "PanelBg", "Primary", "h6"),
						)
					}), Rigid(func() {
						if s.RunningInRepo {
							fg, bg := "DocText", "DocBg"
							if s.Config.RunInRepo {
								fg, bg = "ButtonText", "ButtonBg"
							}
							s.TextButton("repo", "Primary", 16,
								fg, bg, &s.RunningInRepoButton)
							for s.RunningInRepoButton.Clicked(s.Gtx) {
								if !s.Config.Running {
									s.Config.RunInRepo = true
									s.CannotRun = false
									s.SaveConfig()
								}
							}
						}
					}), Rigid(func() {
						fg, bg := "DocText", "DocBg"
						if !s.Config.RunInRepo {
							fg, bg = "ButtonText", "ButtonBg"
						}
						s.TextButton("profile", "Primary", 16,
							fg, bg, &s.RunFromProfileButton)
						for s.RunFromProfileButton.Clicked(s.Gtx) {
							if !s.Config.Running {
								s.Config.RunInRepo = false
								s.CannotRun = false
								s.SaveConfig()
							}
						}
					}), Rigid(func() {
						txt := "run pod in its repository"
						if !s.Config.RunInRepo {
							txt = "not implemented"
							s.CannotRun = true
						}
						s.Inset(8,
							s.Text(txt, "PanelText", "PanelBg", "Primary", "h6"),
						)
					}),
				)
			})
		}),
		Rigid(func() {
			s.Inset(4, func() {
				s.FlexH(Rigid(func() {
					s.Inset(8,
						s.Text("Use Go version", "PanelText", "PanelBg", "Primary", "h6"),
					)
				}), Rigid(func() {
					if s.HasGo {
						fg, bg := "DocText", "DocBg"
						if s.Config.UseBuiltinGo {
							fg, bg = "ButtonText", "ButtonBg"
						}
						s.TextButton("builtin", "Primary", 16,
							fg, bg, &s.UseBuiltinGoButton)
						for s.UseBuiltinGoButton.Clicked(s.Gtx) {
							if !s.Config.RunInRepo {
								s.Config.UseBuiltinGo = true
								s.CannotRun = false
								if !s.HasGo {
									s.CannotRun = true
								}
							}
						}
					}
				}), Rigid(func() {
					fg, bg := "DocText", "DocBg"
					if !s.Config.UseBuiltinGo {
						fg, bg = "ButtonText", "ButtonBg"
					}
					s.TextButton("install new", "Primary", 16,
						fg, bg, &s.InstallNewGoButton)
					for s.InstallNewGoButton.Clicked(s.Gtx) {
						if !s.Config.RunInRepo {
							s.Config.UseBuiltinGo = false
							s.CannotRun = false
							if !s.HasOtherGo {
								s.CannotRun = true
							}
						}
					}
				}), Rigid(func() {
					txt := "build using built in go"
					if !s.Config.UseBuiltinGo {
						txt = "not implemented"
						s.CannotRun = true
					}
					s.Inset(8,
						s.Text(txt, "PanelText", "PanelBg", "Primary", "h6"),
					)
				}),
				)
			})
		}), Rigid(func() {
			s.Inset(4, func() {
				s.FlexH(
					Rigid(func() {
						s.Inset(8,
							s.Text("Log entry click command", "PanelText",
								"PanelBg",
								"Primary", "h6"),
						)
					}), Rigid(func() {
						ww := len(s.Config.ClickCommand)
						//if ww < 12 {
						//	ww = 12
						//}
						s.Gtx.Constraints.Width.Max = ww*10 + 30
						s.Gtx.Constraints.Width.Min = ww*10 + 30
						s.Editor(&s.CommandEditor, ww, func(e gel.EditorEvent) {
							if e != nil {
								txt := s.CommandEditor.Text()
								if s.Config.ClickCommand == txt {
									return
								}
								s.Config.ClickCommand = txt
								Debug(s.Config.ClickCommand)
								s.SaveConfig()
							}
						})()
					}), Rigid(func() {
						s.Inset(8,
							s.Text("When a log entry is clicked run this"+
								" command with variables substituted for"+
								" values from the log entry:\n\n"+
								"$1 is the source code file location\n"+
								"$2 is the line number", "PanelText",
								"PanelBg",
								"Primary", "h6"),
						)
					}),
				)
			})
		}),
	)
}
