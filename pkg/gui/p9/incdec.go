package p9

import (
	"fmt"

	l "gioui.org/layout"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

type IncDec struct {
	th                *Theme
	nDigits           int
	min, max          int
	amount            int
	current           int
	changeHook        func(n int)
	inc, dec          *Clickable
	color, background string
	inactive          string
}

// IncDec is a simple increment/decrement for a number setting
func (th *Theme) IncDec() (out *IncDec) {
	out = &IncDec{
		th: th,
		// nDigits:    nDigits,
		// min:        min,
		// max:        max,
		// current:    current,
		// changeHook: changeHook,
		inc: th.Clickable(),
		dec: th.Clickable(),
		// color:      color,
		// background: background,
		// inactive:   inactive,
		amount: 1,
	}
	return
}

func (in *IncDec) Amount(n int) *IncDec {
	in.amount = n
	return in
}

func (in *IncDec) ChangeHook(fn func(n int)) *IncDec {
	in.changeHook = fn
	return in
}

func (in *IncDec) SetCurrent(current int) *IncDec {
	in.current = current
	return in
}

func (in *IncDec) Max(max int) *IncDec {
	in.max = max
	return in
}

func (in *IncDec) Min(min int) *IncDec {
	in.min = min
	return in
}

func (in *IncDec) NDigits(nDigits int) *IncDec {
	in.nDigits = nDigits
	return in
}

func (in *IncDec) Color(color string) *IncDec {
	in.color = color
	return in
}

func (in *IncDec) Background(color string) *IncDec {
	in.background = color
	return in
}
func (in *IncDec) Inactive(color string) *IncDec {
	in.inactive = color
	return in
}

func (in *IncDec) Fn(gtx l.Context) l.Dimensions {
	out := in.th.Flex().AlignMiddle()
	incColor, decColor := in.color, in.color
	if in.current == in.min {
		decColor = in.inactive
	}
	if in.current == in.max {
		incColor = in.inactive
	}
	if in.current == in.min {
		out.Rigid(
			in.th.Inset(0.25,
				in.th.Icon().Color(decColor).Scale(Scales["H5"]).Src(&icons.ContentRemove).Fn,
			).Fn,
		)
	} else {
		out.Rigid(in.th.Inset(0.25,
			in.th.ButtonLayout(in.inc.SetClick(func() {
				in.current -= in.amount
				if in.current < in.min {
					in.current = in.min
				} else {
					in.changeHook(in.current)
				}
			})).Background("Transparent").Embed(
				in.th.Icon().Color(decColor).Scale(Scales["H5"]).Src(&icons.ContentRemove).Fn,
			).Fn,
		).Fn,
		)
	}
	cur := fmt.Sprintf("%"+fmt.Sprint(in.nDigits)+"d", in.current)
	out.Rigid(in.th.Body1(cur).Color(in.color).Font("go regular").Fn)
	if in.current == in.max {
		out.Rigid(
			in.th.Inset(0.25,
				in.th.Icon().Color(incColor).Scale(Scales["H5"]).Src(&icons.ContentAdd).Fn,
			).Fn,
		)
	} else {
		out.Rigid(
			in.th.Inset(0.25,
				in.th.ButtonLayout(in.dec.SetClick(func() {
					in.current += in.amount
					if in.current > in.max {
						in.current = in.max
					} else {
						in.changeHook(in.current)
					}
				})).Background("Transparent").Embed(
					in.th.Icon().Color(incColor).Scale(Scales["H5"]).Src(&icons.ContentAdd).Fn,
				).Fn,
			).Fn,
		)
	}
	return out.Fn(gtx)
}
