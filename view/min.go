package view

import (
	"faya/list"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/pplcc/plotext"
	"github.com/pplcc/plotext/custplotter"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

func PlotMin(code string, date string){
	fmt.Println("plot:", code, date)
	//rikData := list.RiKCode(code)
	minData := list.MinCodeDate(code, date)
	if len(minData) > 0{
		PlotMinDate(list.GetObj(code), minData)
	}else {
		fmt.Println("no min data for ", code, date)
	}
}
func PlotMinDate(o list.TimeObject, min []*list.MinUnit) {

	//for min
	arrmin := make(custplotter.TOHLCVs, len(min))
	fmt.Println(len(min))
	for i := range min {
// 		tm, err := time.Parse("2006-01-02 15:59:59", min[i].DateTime + ":00")
// 		if err != nil {
// 			panic(err)
// 		}
		//loc := time.Location("Asia/shanghai")
		loc := time.FixedZone("UTC+8", +8*60*60)
		l := strings.Split(min[i].DateTime, " ")
		dl := strings.Split(l[0], "-")
		tl := strings.Split(l[1], ":")
// 		fmt.Println(dl)
		dy, err := strconv.Atoi(dl[0])
		dm, err := strconv.Atoi(dl[1])
		dd, err := strconv.Atoi(dl[2])
// 		fmt.Println(tl)
		rh, err := strconv.Atoi(tl[0])
		rm, err := strconv.Atoi(tl[1])
		//rs, err := strconv.Atoi(tl[2])
		rs := 0
		if err != nil {
			fmt.Println(err)
		}
		tm := time.Date(dy, time.Month(dm), dd, rh, rm, rs, 0, loc)

		arrmin[i].T = float64(tm.Unix())
		arrmin[i].O = min[i].Open
		arrmin[i].H = min[i].High
		arrmin[i].L = min[i].Low
		arrmin[i].C = min[i].Close
		arrmin[i].V = float64(min[i].Amount)
	}

	fmt.Println(o)
	addFont()

	// prepare for p1 min
	p1 :=  plot.New()
	titleText := o.Code + " " + o.Name
	if len(min) > 1  {
// 		loc := time.FixedZone("UTC+8", +8*60*60)
		l := strings.Split(min[0].DateTime, " ")
		dl := strings.Split(l[0], "-")
		titleText += " " + dl[0] + "-" + dl[1] + "-" + dl[2] 
	}
	p1.Title.Text = titleText

	p1.X.Tick.Marker = plot.TimeTicks{Format: "2006-01-02\n15:04:05"}

	barsmin, err := custplotter.NewCandlesticks(arrmin)
	if err != nil {
		log.Panic(err)
	}

	p1.Add(barsmin)
	//for p2
	p2 := plot.New()
	p2.X.Label.Text = "Time"
	p2.Y.Label.Text = "Volume"
	p2.X.Tick.Marker = plot.TimeTicks{Format: "2006-01-02\n15:04:05"}

	vBarsmin, err := custplotter.NewVBars(arrmin)
	if err != nil {
		log.Panic(err)
	}
	// p2.Y.Padding += (candlesticks.CandleWidth - vBarsmin.LineStyle.Width) / 2
	p2.Add(vBarsmin)


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
