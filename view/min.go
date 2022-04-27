package view

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"faya/list"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/pplcc/plotext"
	"github.com/pplcc/plotext/custplotter"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/font"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"

	"golang.org/x/image/font/opentype"
)

func PlotMin(code string, date string){
	fmt.Println("plot:", code, date)
	rikData := list.RiKCode(code)
	minData := list.MinCode(code)
	PlotMinDate(list.GetObj(code), rikData, minData)
}
func untargz(name string, r io.Reader) ([]byte, error) {
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, fmt.Errorf("could not create gzip reader: %v", err)
	}
	defer gr.Close()

	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				return nil, fmt.Errorf("could not find %q in tar archive", name)
			}
			return nil, fmt.Errorf("could not extract header from tar archive: %v", err)
		}

		if hdr == nil || hdr.Name != name {
			continue
		}

		buf := new(bytes.Buffer)
		_, err = io.Copy(buf, tr)
		if err != nil {
			return nil, fmt.Errorf("could not extract %q file from tar archive: %v", name, err)
		}
		return buf.Bytes(), nil
	}
}

func addFont(){
	/*
	fontName := "simhei"
	ttfBytes, err := ioutil.ReadFile("../font/SimHei.ttf")
	if err != nil {
		fmt.Println("read ttf file error:", err)
	}
    font, err := truetype.Parse(ttfBytes)
	if err != nil {
		fmt.Println("truetype parse ttf file error:", err)
	}
	simhei := font.Font{Typeface: "simhei"}
	font.DefaultCache.Add([]font.Face{
		{
			Font: simhei,
			Face: fontTTF,
		},
	})
	//vg.AddFont(fontName, font)
    plot.DefaultFont = simhei
    plotter.DefaultFont = simhei
	*/
		// download font from debian
	const url = "http://http.debian.net/debian/pool/main/f/fonts-ipafont/fonts-ipafont_00303.orig.tar.gz"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("could not download IPA font file: %+v", err)
	}
	defer resp.Body.Close()

	ttf, err := untargz("IPAfont00303/ipam.ttf", resp.Body)
	if err != nil {
		log.Fatalf("could not untar archive: %+v", err)
	}

	fontTTF, err := opentype.Parse(ttf)
	if err != nil {
		log.Fatal(err)
	}
	mincho := font.Font{Typeface: "Mincho"}
	font.DefaultCache.Add([]font.Face{
		{
			Font: mincho,
			Face: fontTTF,
		},
	})
	if !font.DefaultCache.Has(mincho) {
		log.Fatalf("no font %q!", mincho.Typeface)
	}
	plot.DefaultFont = mincho
	plotter.DefaultFont = mincho

}
func PlotMinDate(o list.TimeObject, data []*list.RiKUnit, min []*list.MinUnit) {
	//test(data)
	//prepare data
	//for rik
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

	//support chinese word in title
	addFont()

	// prepare for p1 rik
	p1 :=  plot.New()
	p1.Title.Text = o.Code  + " " + o.Name

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
	// prepare for p3 min
	p3 :=  plot.New()
	rightTitleText := o.Code
	if len(min) > 1  {
// 		loc := time.FixedZone("UTC+8", +8*60*60)
		l := strings.Split(min[0].DateTime, " ")
		dl := strings.Split(l[0], "-")
// 		tl := strings.Split(l[1], ":")
// 		fmt.Println(dl)
// 		dy, err := strconv.Atoi(dl[0])
// 		dm, err := strconv.Atoi(dl[1])
// 		dd, err := strconv.Atoi(dl[2])
// 		if err != nil {
// 			fmt.Println(err)
// 		}
		rightTitleText += " " + dl[0] + "-" + dl[1] + "-" + dl[2] 
	}
	p3.Title.Text = rightTitleText

	p3.X.Tick.Marker = plot.TimeTicks{Format: "2006-01-02\n15:04:05"}

	barsmin, err := custplotter.NewCandlesticks(arrmin)
	if err != nil {
		log.Panic(err)
	}

	p3.Add(barsmin)
	//for p4
	p4 := plot.New()
	p4.X.Label.Text = "Time"
	p4.Y.Label.Text = "Volume"
	p4.X.Tick.Marker = plot.TimeTicks{Format: "2006-01-02\n15:04:05"}

	vBarsmin, err := custplotter.NewVBars(arrmin)
	if err != nil {
		log.Panic(err)
	}
	// p2.Y.Padding += (candlesticks.CandleWidth - vBarsmin.LineStyle.Width) / 2
	p4.Add(vBarsmin)


	// combine p1 and p2
	plotext.UniteAxisRanges([]*plot.Axis{&p1.X, &p2.X})
// 	plotext.UniteAxisRanges([]*plot.Axis{&p3.X, &p4.X})

	// create a table with one column and two rows
	table := plotext.Table{
		RowHeights: []float64{2, 1}, // 2/3 for candlesticks and 1/3 for volume bars
		ColWidths:  []float64{1,1},
	}

	// see align_test.go for another example on how to construct this structure using loops
	plots := [][]*plot.Plot{[]*plot.Plot{p1, p3}, []*plot.Plot{p2, p4}}

	img := vgimg.New(800, 600)
	dc := draw.New(img)

	canvases := table.Align(plots, dc)
	plots[0][0].Draw(canvases[0][0])
	plots[1][0].Draw(canvases[1][0])
	plots[0][1].Draw(canvases[0][1])
	plots[1][1].Draw(canvases[1][1])




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
