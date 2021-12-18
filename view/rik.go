package view

import (
	"faya/list"
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/pplcc/plotext/custplotter"
	"gonum.org/v1/plot"
)
func test(x interface{}) interface{} {
	return x
}

func PlotRik(code string){
	fmt.Println("plot:", code)
	rikData := list.RiKCode(code)
	Plot(list.GetObj(code), rikData)
}
func Plot(o list.TimeObject, data []*list.RiKUnit) {
	//test(data)
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
	p :=  plot.New()
	p.Title.Text = o.Code 
// 	p.X.Label.Text = "X [mm]"
// 	p.Y.Label.Text = "Y [A.U.]"
// 	p.X.Label.Position = draw.PosRight
// 	p.Y.Label.Position = draw.PosTop
// 	p.X.Min = -10
// 	p.X.Max = +10
// 	p.Y.Min = -10
// 	p.Y.Max = +10

	p.X.Tick.Marker = plot.TimeTicks{Format: "2006-01-02\n15:04:05"}

	bars, err := custplotter.NewCandlesticks(arr)
	if err != nil {
		log.Panic(err)
	}

	p.Add(bars)

	tmp_file_name := "test.png"
	err = p.Save(800, 600, tmp_file_name)
	if err != nil {
		log.Fatalf("could not save plot: %+v", err)
	}
	/*
	infile, err := os.Open(tmp_file_name)
    if err != nil {
        // replace this with real error handling
        panic(err)
    }
    defer infile.Close()
	*/
	fmt.Println("will open ", tmp_file_name)
	cmd := exec.Command("open", tmp_file_name)
	if err := cmd.Run(); err != nil {
        fmt.Println("Error: ", err)
    }
}
