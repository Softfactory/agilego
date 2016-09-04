package dl_test

import (
	"bitbucket.org/binet/go-gnuplot/pkg/gnuplot"
	"fmt"
	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
	"github.com/gonum/stat/distuv"
	"image/color"
	"math"
	"testing"
)

func TestLinear(t *testing.T) {
	// linear()
	// step()
	// sigmoid()
	//	tanh()
	// relu1()
	var linear = func(x float64) float64 { return x }
	gonumPlot("linear", linear, -4, 4, -4, 4)

	scale := []float64{-5, 5, 0, 1}

	var step = func(x float64) float64 {
		if x >= 0.5 {
			return 1
		} else {
			return 0
		}
	}
	gonumPlot("step", step, scale...)

	var sigmoid = func(x float64) float64 { return 1.0 / (1.0 + math.Exp(-x)) }
	gonumPlot("Sigmoid", sigmoid, scale...)

	gonumPlot("Tanh", math.Tanh, -5, 5, -1, 1)

	gonumPlot("Normal", distuv.UnitNormal.Prob)

	//
}

func linear() {
	gnuplot("x")
}

func step() {
	gnuplot("x>0.5?1:0")
}

func sigmoid() {
	gnuplot("1/(1+exp(-x))")
}

func tanh() {
	gnuplot("tanh(x)")
}

func relu1() {
	gnuplot("log(1+exp(x))")
}

func relu2() {
	gnuplot("x<0?0:x")
}

func gnuplot(strFunction string) {
	fname := ""
	persist := false
	debug := true

	p, err := gnuplot.NewPlotter(fname, persist, debug)
	if err != nil {
		errStr := fmt.Sprintf("** err: %v\n", err)
		panic(errStr)
	}
	defer p.Close()

	var cmd = "plot " + strFunction

	//Linear
	p.CheckedCmd(cmd)
	p.CheckedCmd("q")
	return
}

func gonumPlot(title string, f func(float64) float64, scale ...float64) {
	// plot 구조체에 제목을 설정합니다.
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = title
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	//확률분포함수를 지정하고 plot을 정의합니다.
	pdf := plotter.NewFunction(f)
	pdf.Color = color.RGBA{R: 255, A: 255}
	pdf.Dashes = []vg.Length{vg.Points(2), vg.Points(2)}
	pdf.Width = vg.Points(2)

	p.Add(pdf)
	if scale != nil && len(scale) == 4 {
		p.X.Min = scale[0]
		p.X.Max = scale[1]
		p.Y.Min = scale[2]
		p.Y.Max = scale[3]
	} else {
		p.X.Min = -1
		p.X.Max = 1
		p.Y.Min = -1
		p.Y.Max = 1
	}

	var fileName = title + ".png"
	//PNG 파일로 저장합니다.
	if err := p.Save(10*vg.Inch, 10*vg.Inch, fileName); err != nil {
		panic(err)
	}
}
