package view

import (
	"faya/list"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/pplcc/plotext"
	"github.com/pplcc/plotext/custplotter"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

func PlotMin(code string, date string){
	fmt.Println("plot:", code)
	rikData := list.RiKCode(code)
	Plot(list.GetObj(code), rikData)
}
func Plot(o list.TimeObject, data []*list.RiKUnit) {
	//test(data)
	//prepare data
	arr := make(custplotter.TOHLCVs, len(data))
	for i := range data {
		tm, err := time.Parse("2006-01-02", data[i].Date)
		if err != nil {
			panic(err)
		}
		arr[i].T = float64(tm.Unix())
		arr[i].O = data[i].Open
		arr[i].H = data[i].High
		arr[i].L = data[i].Low
		arr[i].C = data[i].Close
		arr[i].V = float64(data[i].Amount)
	}

	fmt.Println(o)
	// prepare for p1 rik
	p1 :=  plot.New()
	p1.Title.Text = o.Code 

	p1.X.Tick.Marker = plot.TimeTicks{Format: "2006-01-02\n15:04:05"}

	bars, err := custplotter.NewCandlesticks(arr)
	if err != nil {
		log.Panic(err)
	}

	p1.Add(bars)

	//for p2
	p2 := plot.New()
	p2.X.Label.Text = "Time"
	p2.Y.Label.Text = "Volume"
	p2.X.Tick.Marker = plot.TimeTicks{Format: "2006-01-02\n15:04:05"}

	vBars, err := custplotter.NewVBars(arr)
	if err != nil {
		log.Panic(err)
	}
	// p2.Y.Padding += (candlesticks.CandleWidth - vBars.LineStyle.Width) / 2
	p2.Add(vBars)


	// combine p1 and p2
	plotext.UniteAxisRanges([]*plot.Axis{&p1.X, &p2.X})

	// create a table with one column and two rows
	table := plotext.Table{
		RowHeights: []float64{2, 1}, // 2/3 for candlesticks and 1/3 for volume bars
		ColWidths:  []float64{1},
	}

	// see align_test.go for another example on how to construct this structure using loops
	plots := [][]*plot.Plot{[]*plot.Plot{p1}, []*plot.Plot{p2}}

	img := vgimg.New(800, 600)
	dc := draw.New(img)

	canvases := table.Align(plots, dc)
	plots[0][0].Draw(canvases[0][0])
	plots[1][0].Draw(canvases[1][0])




	tmp_file_name := "test.png"
	w, err := os.Create(tmp_file_name)
	if err != nil {
		panic(err)
	}

	png := vgimg.PngCanvas{Canvas: img}
	if _, err := png.WriteTo(w); err != nil {
		panic(err)
	}


	//Open file
	fmt.Println("will open ", tmp_file_name)
	cmd := exec.Command("open", tmp_file_name)
	if err := cmd.Run(); err != nil {
        fmt.Println("Error: ", err)
    }
}
