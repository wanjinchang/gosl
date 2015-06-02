// Copyright 2015 Dorival Pedroso and Raul Durand. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fun

import (
	"math"
	"testing"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
	"github.com/cpmech/gosl/plt"
)

func Test_fun01(tst *testing.T) {

	//verbose()
	chk.PrintTitle("fun01. Decreasing Reference Model")

	ya := 1.0
	yb := -0.5
	λ1 := 1.0

	o, err := New("ref-dec-gen", []*Prm{
		&Prm{N: "bet", V: 5.0},
		&Prm{N: "a", V: -λ1},
		&Prm{N: "b", V: -1.0},
		&Prm{N: "c", V: ya},
		&Prm{N: "A", V: 0.0},
		&Prm{N: "B", V: λ1},
		&Prm{N: "xini", V: 0.0},
		&Prm{N: "yini", V: yb},
	})
	if err != nil {
		tst.Errorf("test failed: %v\n")
		return
	}

	tmax := 3.0
	xcte := []float64{0, 0, 0}
	if false {
		plt.Reset()
		withG, withH, save, show := true, true, false, true
		PlotT(o, "/tmp/gosl", "ref-dec-gen-01.png", 0.0, tmax, xcte, 201, "", withG, withH, save, show, func() {
			plt.Plot([]float64{0, tmax}, []float64{ya, ya - λ1*tmax}, "'k-'")
			plt.Equal()
		})
	}
	//
	sktol := 1e-10
	dtol := 1e-10
	dtol2 := 1e-10
	ver := chk.Verbose
	CheckT(tst, o, 0.0, tmax, xcte, 11, nil, sktol, dtol, dtol2, ver)
}

func Test_fun02(tst *testing.T) {

	//verbose()
	chk.PrintTitle("fun02. Dec Ref Model (specialised)")

	ya := 1.0
	yb := -50.0
	λ1 := 1.0

	o, err := New("ref-dec-sp1", []*Prm{
		&Prm{N: "bet", V: 5.0},
		&Prm{N: "lam1", V: λ1},
		&Prm{N: "ya", V: ya},
		&Prm{N: "yb", V: yb},
	})
	if err != nil {
		tst.Errorf("test failed: %v\n")
		return
	}

	tmin := 0.0
	tmax := 300.0
	//tmax := 140.0
	xcte := []float64{0, 0, 0}
	if false {
		plt.Reset()
		withG, withH, save, show := true, true, false, true
		PlotT(o, "/tmp/gosl", "ref-dec-sp1-01.png", tmin, tmax, xcte, 201, "lw=2,color='orange'", withG, withH, save, show, func() {
			plt.Plot([]float64{0, tmax}, []float64{ya, ya - λ1*tmax}, "'k--'")
			plt.Equal()
		})
	}

	if true {
		//if false {
		sktol := 1e-10
		dtol := 1e-10
		dtol2 := 1e-10
		ver := chk.Verbose
		CheckT(tst, o, tmin, tmax, xcte, 11, nil, sktol, dtol, dtol2, ver)
	}
}

func Test_fun03(tst *testing.T) {

	//verbose()
	chk.PrintTitle("fun03. add, cte, srmps")

	cte, err := New("cte", []*Prm{&Prm{N: "C", V: 30}})
	if err != nil {
		tst.Errorf("test failed: %v\n")
		return
	}

	srmps, err := New("srmps", []*Prm{
		&Prm{N: "ca", V: 0},
		&Prm{N: "cb", V: 1},
		&Prm{N: "ta", V: 0},
		&Prm{N: "tb", V: 1},
	})
	if err != nil {
		tst.Errorf("test failed: %v\n")
		return
	}

	add, err := New("add", []*Prm{
		&Prm{N: "a", V: 1},
		&Prm{N: "b", V: 1},
		&Prm{N: "fa", Fcn: cte},
		&Prm{N: "fb", Fcn: srmps},
	})
	if err != nil {
		tst.Errorf("test failed: %v\n")
		return
	}

	tmin := 0.0
	tmax := 1.0
	xcte := []float64{0, 0, 0}
	if false {
		withG, withH, save, show := true, true, false, true
		plt.Reset()
		PlotT(cte, "/tmp/gosl", "fun-cte-01.png", tmin, tmax, xcte, 41, "", withG, withH, save, show, nil)
		plt.Reset()
		PlotT(srmps, "/tmp/gosl", "fun-srmps-01.png", tmin, tmax, xcte, 41, "", withG, withH, save, show, nil)
		plt.Reset()
		PlotT(add, "/tmp/gosl", "fun-add-01.png", tmin, tmax, xcte, 41, "", withG, withH, save, show, nil)
	}

	if true {
		//if false {
		sktol := 1e-10
		dtol := 1e-10
		dtol2 := 1e-9
		ver := chk.Verbose
		tskip := []float64{tmin, tmax}
		CheckT(tst, cte, tmin, tmax, xcte, 11, nil, sktol, dtol, dtol2, ver)
		io.Pf("\n")
		CheckT(tst, srmps, tmin, tmax, xcte, 11, tskip, sktol, dtol, dtol2, ver)
		io.Pf("\n")
		CheckT(tst, add, tmin, tmax, xcte, 11, tskip, sktol, dtol, dtol2, ver)
	}
}

func Test_fun04(tst *testing.T) {

	//verbose()
	chk.PrintTitle("fun04. lin")

	lin, err := New("lin", []*Prm{
		&Prm{N: "m", V: 0.5},
		&Prm{N: "ts", V: 0},
	})
	if err != nil {
		tst.Errorf("test failed: %v\n")
		return
	}

	tmin := 0.0
	tmax := 1.0
	xcte := []float64{0, 0, 0}
	if false {
		plt.Reset()
		withG, withH, save, show := true, true, false, true
		PlotT(lin, "/tmp/gosl", "fun-lin-01.png", tmin, tmax, xcte, 11, "", withG, withH, save, show, nil)
	}

	if true {
		//if false {
		sktol := 1e-10
		dtol := 1e-10
		dtol2 := 1e-10
		ver := chk.Verbose
		CheckT(tst, lin, tmin, tmax, xcte, 11, nil, sktol, dtol, dtol2, ver)
	}
}

func Test_fun05(tst *testing.T) {

	//verbose()
	chk.PrintTitle("fun05. zero and one")

	io.Pforan("Zero(666,nil) = %v\n", Zero.F(666, nil))
	io.Pforan("One(666,nil)  = %v\n", One.F(666, nil))
	chk.Scalar(tst, "zero", 1e-17, Zero.F(666, nil), 0)
	chk.Scalar(tst, "one ", 1e-17, One.F(666, nil), 1)
}

func Test_fun06(tst *testing.T) {

	//verbose()
	chk.PrintTitle("fun06. pts")

	fun, err := New("pts", []*Prm{
		&Prm{N: "t0", V: 0.00}, {N: "y0", V: 0.50},
		&Prm{N: "t1", V: 1.00}, {N: "y1", V: 0.20},
		&Prm{N: "t2", V: 2.00}, {N: "y2", V: 0.20},
		&Prm{N: "t3", V: 3.00}, {N: "y3", V: 0.05},
		&Prm{N: "t4", V: 4.00}, {N: "y4", V: 0.01},
		&Prm{N: "t5", V: 5.00}, {N: "y5", V: 0.00},
	})
	if err != nil {
		tst.Errorf("test failed: %v\n")
		return
	}

	tmin := -1.0
	tmax := 6.0
	xcte := []float64{0, 0, 0}
	if false {
		plt.Reset()
		withG, withH, save, show := true, true, false, true
		PlotT(fun, "/tmp/gosl", "fun-pts-01.png", tmin, tmax, xcte, 8, "'o-', clip_on=0", withG, withH, save, show, nil)
	}

	if true {
		tmin = 0.01
		tmax = 4.99
		//if false {
		sktol := 1e-10
		dtol := 1e-10
		dtol2 := 1e-10
		ver := chk.Verbose
		CheckT(tst, fun, tmin, tmax, xcte, 11, nil, sktol, dtol, dtol2, ver)
	}
}

func Test_fun07(tst *testing.T) {

	//verbose()
	chk.PrintTitle("fun07. exc1")

	fun, err := New("exc1", []*Prm{
		&Prm{N: "A", V: 200},
		&Prm{N: "b", V: 2},
	})
	if err != nil {
		tst.Errorf("test failed: %v\n")
		return
	}

	tmin := 0.0
	tmax := 1.0
	xcte := []float64{0, 0, 0}
	if false {
		plt.Reset()
		withG, withH, save, show := true, true, false, true
		PlotT(fun, "/tmp/gosl", "fun-exc1-01.png", tmin, tmax, xcte, 41, "'o-'", withG, withH, save, show, nil)
	}

	if true {
		sktol := 1e-10
		dtol := 1e-7
		dtol2 := 1e-6
		ver := chk.Verbose
		CheckT(tst, fun, tmin, tmax, xcte, 11, nil, sktol, dtol, dtol2, ver)
	}
}

func Test_fun08(tst *testing.T) {

	//verbose()
	chk.PrintTitle("fun08. exc2")

	fun, err := New("exc2", []*Prm{
		&Prm{N: "ta", V: 5},
		&Prm{N: "A", V: 3},
		&Prm{N: "b", V: 0.2},
	})
	if err != nil {
		tst.Errorf("test failed: %v\n")
		return
	}

	tmin := 0.0
	tmax := 7.0
	xcte := []float64{0, 0, 0}
	if false {
		plt.Reset()
		withG, withH, save, show := true, true, false, true
		PlotT(fun, "/tmp/gosl", "fun-exc2-01.png", tmin, tmax, xcte, 41, "'o-'", withG, withH, save, show, nil)
	}

	if true {
		sktol := 1e-10
		dtol := 1e-10
		dtol2 := 1e-10
		ver := chk.Verbose
		CheckT(tst, fun, tmin, tmax, xcte, 11, nil, sktol, dtol, dtol2, ver)
	}
}

func Test_fun09(tst *testing.T) {

	//verbose()
	chk.PrintTitle("fun09. cos")

	fun, err := New("cos", []*Prm{
		&Prm{N: "a", V: 10},
		&Prm{N: "b", V: math.Pi},
		&Prm{N: "c", V: 1.0},
	})
	if err != nil {
		tst.Errorf("test failed: %v\n")
		return
	}

	tmin := 0.0
	tmax := 2.0
	xcte := []float64{0, 0, 0}
	if false {
		plt.Reset()
		withG, withH, save, show := true, true, false, true
		PlotT(fun, "/tmp/gosl", "fun-cos-01.png", tmin, tmax, xcte, 41, "'.-'", withG, withH, save, show, nil)
	}

	if true {
		sktol := 1e-10
		dtol := 1e-8
		dtol2 := 1e-7
		ver := chk.Verbose
		CheckT(tst, fun, tmin, tmax, xcte, 11, nil, sktol, dtol, dtol2, ver)
	}
}

func Test_fun10(tst *testing.T) {

	//verbose()
	chk.PrintTitle("fun10. rmp")

	fun, err := New("rmp", []*Prm{
		&Prm{N: "ta", V: 1},
		&Prm{N: "tb", V: 2},
		&Prm{N: "ca", V: 0.5},
		&Prm{N: "cb", V: -1.5},
	})
	if err != nil {
		tst.Errorf("test failed: %v\n")
		return
	}

	tmin := 0.0
	tmax := 3.0
	xcte := []float64{0, 0, 0}
	if false {
		plt.Reset()
		withG, withH, save, show := true, true, false, true
		PlotT(fun, "/tmp/gosl", "fun-rmp-01.png", tmin, tmax, xcte, 4, "'.-'", withG, withH, save, show, nil)
	}

	if true {
		sktol := 1e-10
		dtol := 1e-12
		dtol2 := 1e-17
		ver := chk.Verbose
		CheckT(tst, fun, tmin, tmax, xcte, 11, nil, sktol, dtol, dtol2, ver)
	}
}

func Test_fun11(tst *testing.T) {

	//verbose()
	chk.PrintTitle("fun11. ref-inc-rl1")

	fun, err := New("ref-inc-rl1", []*Prm{
		&Prm{N: "lam0", V: 0.001},
		&Prm{N: "lam1", V: 1.2},
		&Prm{N: "alp", V: 0.01},
		&Prm{N: "bet", V: 10},
	})
	if err != nil {
		tst.Errorf("test failed: %v\n")
		return
	}

	tmin := 0.0
	tmax := 1.0
	xcte := []float64{0, 0, 0}
	if false {
		plt.Reset()
		withG, withH, save, show := true, true, false, true
		PlotT(fun, "/tmp/gosl", "fun-ref-inc-rl1-01.png", tmin, tmax, xcte, 41, "'.-'", withG, withH, save, show, nil)
	}

	if true {
		sktol := 1e-10
		dtol := 1e-10
		dtol2 := 1e-10
		ver := chk.Verbose
		CheckT(tst, fun, tmin, tmax, xcte, 11, nil, sktol, dtol, dtol2, ver)
	}
}

func Test_fun12(tst *testing.T) {

	//verbose()
	chk.PrintTitle("fun12. mul")

	cos, err := New("cos", []*Prm{
		&Prm{N: "a", V: 1},
		&Prm{N: "b/pi", V: 2},
		&Prm{N: "c", V: 1},
	})
	if err != nil {
		tst.Errorf("test failed: %v\n")
		return
	}

	lin, err := New("lin", []*Prm{
		&Prm{N: "m", V: 0.5},
		&Prm{N: "ts", V: 0},
	})
	if err != nil {
		tst.Errorf("test failed: %v\n")
		return
	}

	mul, err := New("mul", []*Prm{
		&Prm{N: "fa", Fcn: cos},
		&Prm{N: "fb", Fcn: lin},
	})
	if err != nil {
		tst.Errorf("test failed: %v\n")
		return
	}

	tmin := 0.0
	tmax := 1.0
	xcte := []float64{0, 0, 0}
	//if true {
	if false {
		withG, withH, save, show := true, true, false, true
		plt.Reset()
		PlotT(cos, "/tmp/gosl", "fun-cos-12.png", tmin, tmax, xcte, 41, "", withG, withH, save, show, nil)
		plt.Reset()
		PlotT(lin, "/tmp/gosl", "fun-lin-12.png", tmin, tmax, xcte, 41, "", withG, withH, save, show, nil)
		plt.Reset()
		PlotT(mul, "/tmp/gosl", "fun-mul-12.png", tmin, tmax, xcte, 41, "", withG, withH, save, show, nil)
	}

	if true {
		//if false {
		sktol := 1e-10
		dtol := 1e-9
		dtol2 := 1e-8
		ver := chk.Verbose
		tskip := []float64{tmin, tmax}
		CheckT(tst, cos, tmin, tmax, xcte, 11, nil, sktol, dtol, dtol2, ver)
		io.Pf("\n")
		CheckT(tst, lin, tmin, tmax, xcte, 11, tskip, sktol, dtol, dtol2, ver)
		io.Pf("\n")
		CheckT(tst, mul, tmin, tmax, xcte, 11, tskip, sktol, dtol, dtol2, ver)
	}
}

func Test_fun13(tst *testing.T) {

	//verbose()
	chk.PrintTitle("fun13. mul")

	pulse, err := New("pulse", []*Prm{
		&Prm{N: "ca", V: 0.2},
		&Prm{N: "cb", V: 2.0},
		&Prm{N: "ta", V: 1.0},
		&Prm{N: "tb", V: 2.5},
	})
	if err != nil {
		tst.Errorf("test failed: %v\n")
		return
	}

	tmin := 0.0
	tmax := 5.0
	xcte := []float64{0, 0, 0}
	//if true {
	if false {
		withG, withH, save, show := true, true, false, true
		plt.Reset()
		PlotT(pulse, "/tmp/gosl", "fun-pulse-13.png", tmin, tmax, xcte, 61, "", withG, withH, save, show, nil)
	}

	if true {
		//if false {
		sktol := 1e-17
		dtol := 1e-10
		dtol2 := 1e-10
		ver := chk.Verbose
		tskip := []float64{1, 4}
		CheckT(tst, pulse, tmin, tmax, xcte, 11, tskip, sktol, dtol, dtol2, ver)
	}
}
