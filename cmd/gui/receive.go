package gui

import (
	"fmt"
	
	l "gioui.org/layout"
	"gioui.org/text"
	"github.com/atotto/clipboard"
	
	"github.com/p9c/pod/pkg/gui"
)

const inputWidth float32 = 20
const Break1 = 48

func (wg *WalletGUI) ReceivePage() l.Widget {
	return func(gtx l.Context) l.Dimensions {
		if wg.State != nil {
			// Debug(wg.State.isAddress)
			// Debug(wg.State.isAddress.Load())
			if wg.State.isAddress.Load() {
				ad := wg.State.currentReceivingAddress.Load()
				wg.currentReceiveAddress = ad.EncodeAddress()
				// var err error
				// // Debug(ad.ScriptAddress())
				// var conv []byte
				// if conv, err = bech32.ConvertBits(ad.ScriptAddress(), 8, 5, true); Check(err) {
				// }
				// if bech, err = bech32.Encode("pc", conv); Check(err) {
				// }
			}
		}
		if wg.ReceiveAddressbook == nil {
			wg.ReceiveAddressbook = wg.Inset(0.25, wg.H1("addressbook").Alignment(text.End).Fn).Fn
		}
		
		var widgets []l.Widget
		header := wg.Flex().Flexed(1,
			wg.Inset(
				0.25,
				wg.H6("Receive Address History").Alignment(text.Middle).Fn,
			).Fn,
		).Fn
		widgets = append(widgets,
			wg.Flex().AlignMiddle().Flexed(1, wg.Inset(0.5, gui.EmptySpace(0, 0)).Fn).Fn,
			header)
		avail := len(wg.addressbookClickables)
		req := len(wg.State.receiveAddresses)
		if req > avail {
			for i := 0; i < req-avail; i++ {
				wg.addressbookClickables = append(wg.addressbookClickables, wg.WidgetPool.GetClickable())
			}
		}
		for x := range wg.State.receiveAddresses {
			j := x
			i := len(wg.State.receiveAddresses) - 1 - x
			widgets = append(widgets, func(gtx l.Context) l.Dimensions {
				return wg.Inset(0.25,
					wg.ButtonLayout(wg.addressbookClickables[i].SetClick(func() {
						qrText := fmt.Sprintf(
							"parallelcoin:%s?amount=%8.8f&message=%s",
							wg.State.receiveAddresses[i].Address,
							wg.State.receiveAddresses[i].Amount.ToDUO(),
							wg.State.receiveAddresses[i].Message,
						)
						Debug("clicked receive address list item", j)
						if err := clipboard.WriteAll(qrText); Check(err) {
						}
					})).
						Background("PanelBg").
						Embed(
							wg.Inset(0.25,
								wg.VFlex().
									Rigid(
										wg.Flex().AlignBaseline().
											Rigid(
												wg.Caption(wg.State.receiveAddresses[i].Address).
													Font("go regular").Fn,
											).
											Flexed(1,
												wg.Body1(wg.State.receiveAddresses[i].Amount.String()).
													Alignment(text.End).Fn,
											).
											Fn,
									).
									Rigid(
										wg.Body1(wg.State.receiveAddresses[i].Message).Fn,
									).
									Fn,
							).
								Fn,
						).
						Fn,
				).Fn(gtx)
			})
		}
		// assemble the list for the small, scrolling list view
		smallWidgets := append([]l.Widget{
			wg.Inset(
				0.25,
				wg.Body2("Scan to send or click to copy").Alignment(text.Middle).Fn,
			).Fn,
			wg.Flex().SpaceSides().
				Rigid(
					wg.ButtonLayout(
						wg.currentReceiveCopyClickable.SetClick(
							func() {
								qrText := fmt.Sprintf(
									"parallelcoin:%s?amount=%s&message=%s",
									wg.State.currentReceivingAddress.Load().EncodeAddress(),
									wg.inputs["receiveAmount"].GetText(),
									wg.inputs["receiveMessage"].GetText(),
								)
								Debug("clicked qr code copy clicker")
								if err := clipboard.WriteAll(qrText); Check(err) {
								}
							},
						),
					).
						// CornerRadius(0.5).
						// Corners(gui.NW | gui.SW | gui.NE).
						Background("white").
						Embed(
							wg.Inset(
								0.125,
								wg.Image().Src(*wg.currentReceiveQRCode).Scale(1).Fn,
							).Fn,
						).Fn,
				).
				Fn,
			
			// wg.Inset(
			// 	0.25,
			// 	wg.Caption(wg.currentReceiveAddress).Alignment(text.Middle).Font("go regular").Fn,
			// ).Fn,
			func(gtx l.Context) l.Dimensions {
				// gtx.Constraints.Max.X, gtx.Constraints.Min.X = int(wg.TextSize.V * 17),  int(wg.TextSize.V * 17)
				return wg.Inset(0.25, wg.Fill("DocBg", l.Center, 0, 0, wg.inputs["receiveAmount"].Fn).Fn).Fn(gtx)
			},
			func(gtx l.Context) l.Dimensions {
				// gtx.Constraints.Max.X, gtx.Constraints.Min.X = int(wg.TextSize.V * 17),  int(wg.TextSize.V * 17)
				return wg.Inset(0.25, wg.Fill("DocBg", l.Center, 0, 0, wg.inputs["receiveMessage"].Fn).Fn).Fn(gtx)
			},
			wg.Inset(0.25,
				wg.ButtonLayout(
					wg.currentReceiveRegenClickable.SetClick(
						func() {
							Debug("clicked regenerate button")
							wg.currentReceiveGetNew.Store(true)
						},
					),
				).
					Background("Primary").
					Embed(
						wg.Inset(
							0.25,
							wg.H6("regenerate").Color("Light").Fn,
						).Fn,
					).
					Fn,
			).Fn,
		}, widgets...)
		le := func(gtx l.Context, index int) l.Dimensions {
			return smallWidgets[index](gtx)
		}
		return wg.Responsive(
			*wg.Size, gui.Widgets{
				{Size: 0,
					Widget:
					wg.Fill(
						"PanelBg", l.W, 0, 0,
						wg.Inset(
							0.25,
							wg.lists["receive"].
								Vertical().
								Length(len(smallWidgets)).
								ListElement(le).Fn,
						).Fn,
					).
						Fn,
				},
				{
					Size: Break1,
					Widget:
					wg.Fill(
						"PanelBg", l.W, wg.TextSize.V, 0,
						wg.Flex().AlignMiddle().Rigid(
							wg.VFlex().AlignMiddle().
								Rigid(
									wg.VFlex().AlignMiddle().
										Rigid(
											wg.Inset(
												0.25,
												wg.Body2("Scan to send or click to copy").Alignment(text.Middle).Fn,
											).Fn,
										).
										Rigid(
											wg.currentReceiveQR,
										).
										// Rigid(
										// 	wg.Inset(
										// 		0.25,
										// 		wg.Caption(wg.currentReceiveAddress).Font("go regular").Fn,
										// 	).Fn,
										// ).
										Rigid(
											wg.Inset(
												0.25,
												func(gtx l.Context) l.
												Dimensions {
													gtx.Constraints.Max.X, gtx.Constraints.Min.X = int(wg.TextSize.V*inputWidth), int(wg.TextSize.V*inputWidth)
													return wg.inputs["receiveAmount"].Fn(gtx)
												},
											).Fn,
										).
										Rigid(
											wg.Inset(
												0.25,
												func(gtx l.Context) l.Dimensions {
													gtx.Constraints.Max.X, gtx.Constraints.Min.X = int(wg.TextSize.V*inputWidth), int(wg.TextSize.V*inputWidth)
													return wg.inputs["receiveMessage"].Fn(gtx)
												},
											).Fn,
										).
										Fn,
								).
								
								
								Rigid(
									wg.Inset(
										0.25,
										func(gtx l.Context) l.Dimensions {
											gtx.Constraints.Max.X, gtx.Constraints.Min.X = int(wg.TextSize.V*inputWidth), int(wg.TextSize.V*inputWidth)
											return wg.ButtonLayout(
												wg.currentReceiveRegenClickable.SetClick(
													func() {
														Debug("clicked regenerate button")
														wg.currentReceiveGetNew.Store(true)
													},
												),
											).
												// CornerRadius(0.5).Corners(gui.NW | gui.SW | gui.NE).
												Background("Primary").
												Embed(
													wg.Inset(
														0.5,
														wg.H6("regenerate").Color("Light").Fn,
													).Fn,
												).
												Fn(gtx)
										},
									).
										Fn,
								).
								Fn,
						).
							Rigid(
								wg.VFlex().
									Rigid(header).
									Flexed(1,
										wg.Fill(
											"DocBg", l.Center, wg.TextSize.V, 0,
											wg.Inset(
												0.25,
												wg.ReceiveAddressbook,
											).Fn,
										).Fn,
									).
									Fn,
							).
							Fn,
					).
						Fn,
				},
				{
					Size: 64,
					Widget:
					wg.Fill(
						"PanelBg", l.W, wg.TextSize.V, 0,
						wg.Flex().AlignMiddle().Rigid(
							wg.VFlex().AlignMiddle().
								Rigid(
									wg.VFlex().AlignMiddle().
										Rigid(
											wg.Inset(
												0.25,
												wg.Body2("Scan to send or click to copy").Alignment(text.Middle).Fn,
											).Fn,
										).
										Rigid(
											wg.currentReceiveQR,
										).
										Rigid(
											wg.Inset(
												0.25,
												func(gtx l.Context) l.
												Dimensions {
													gtx.Constraints.Max.X, gtx.Constraints.Min.X = int(wg.TextSize.V*inputWidth), int(wg.TextSize.V*inputWidth)
													return wg.inputs["receiveAmount"].Fn(gtx)
												},
											).Fn,
										).
										Rigid(
											wg.Inset(
												0.25,
												func(gtx l.Context) l.Dimensions {
													gtx.Constraints.Max.X, gtx.Constraints.Min.X = int(wg.TextSize.V*inputWidth), int(wg.TextSize.V*inputWidth)
													return wg.inputs["receiveMessage"].Fn(gtx)
												},
											).Fn,
										).
										Fn,
								).
								Rigid(
									wg.Inset(
										0.25,
										func(gtx l.Context) l.Dimensions {
											gtx.Constraints.Max.X, gtx.Constraints.Min.X = int(wg.TextSize.V*inputWidth), int(wg.TextSize.V*inputWidth)
											return wg.ButtonLayout(
												wg.currentReceiveRegenClickable.SetClick(
													func() {
														Debug("clicked regenerate button")
														wg.currentReceiveGetNew.Store(true)
													},
												),
											).
												// CornerRadius(0.5).Corners(gui.NW | gui.SW | gui.NE).
												Background("Primary").
												Embed(
													wg.Inset(
														0.5,
														wg.H6("regenerate").Color("Light").Fn,
													).Fn,
												).
												Fn(gtx)
										},
									).
										Fn,
								).
								Fn,
						).
							Rigid(
								wg.VFlex().
									Rigid(header).
									Flexed(1,
										wg.Fill(
											"DocBg", l.Center, wg.TextSize.V, 0,
											wg.Inset(
												0.25,
												wg.ReceiveAddressbook,
											).Fn,
										).Fn,
									).
									Fn,
							).
							Fn,
					).
						Fn,
				},
				{
					Size: 96,
					Widget:
					wg.Fill(
						"PanelBg", l.W, wg.TextSize.V, 0,
						wg.Flex().AlignMiddle().Rigid(
							wg.VFlex().AlignMiddle().
								Rigid(
									wg.Flex().AlignMiddle().
										Rigid(
											wg.VFlex().AlignMiddle().
												Rigid(
													wg.Inset(
														0.25,
														wg.Body2("Scan to send or click to copy").Alignment(text.Middle).Fn,
													).Fn,
												).
												Rigid(
													wg.currentReceiveQR,
												).
												// Rigid(
												// 	wg.Inset(
												// 		0.25,
												// 		wg.Caption(wg.currentReceiveAddress).Font("go regular").Fn,
												// 	).Fn,
												// ).
												Fn,
										).
										Rigid(
											wg.VFlex().AlignMiddle().
												Rigid(
													wg.Inset(
														0.25,
														func(gtx l.Context) l.
														Dimensions {
															gtx.Constraints.Max.X, gtx.Constraints.Min.X = int(wg.TextSize.V*inputWidth), int(wg.TextSize.V*inputWidth)
															return wg.inputs["receiveAmount"].Fn(gtx)
														},
													).Fn,
												).
												Rigid(
													wg.Inset(
														0.25,
														func(gtx l.Context) l.Dimensions {
															gtx.Constraints.Max.X, gtx.Constraints.Min.X = int(wg.TextSize.V*inputWidth), int(wg.TextSize.V*inputWidth)
															return wg.inputs["receiveMessage"].Fn(gtx)
														},
													).Fn,
												).
												Rigid(
													wg.Inset(
														0.25,
														func(gtx l.Context) l.Dimensions {
															gtx.Constraints.Max.X, gtx.Constraints.Min.X = int(wg.TextSize.V*inputWidth), int(wg.TextSize.V*inputWidth)
															return wg.ButtonLayout(
																wg.currentReceiveRegenClickable.SetClick(
																	func() {
																		Debug("clicked regenerate button")
																		wg.currentReceiveGetNew.Store(true)
																	},
																),
															).
																// CornerRadius(0.5).Corners(gui.NW | gui.SW | gui.NE).
																Background("Primary").
																Embed(
																	wg.Inset(
																		0.5,
																		wg.H6("regenerate").Color("Light").Fn,
																	).Fn,
																).
																Fn(gtx)
														},
													).
														Fn,
												).Fn,
										).
										Fn,
								).
								Fn,
						).
							Rigid(
								wg.VFlex().
									Rigid(header).
									Flexed(1,
										wg.Fill(
											"DocBg", l.Center, wg.TextSize.V, 0,
											wg.Inset(
												0.25,
												wg.ReceiveAddressbook,
											).Fn,
										).Fn,
									).
									Fn,
							).
							Fn,
					).
						Fn,
				},
			},
		).
			Fn(gtx)
	}
}
