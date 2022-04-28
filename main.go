package main

import (
	// 	"faya/strategy"

	"faya/filter"
	"faya/function"
	"faya/list"
	"faya/strategy"
	"faya/view"
	"fmt"
	"os"
	"strconv"
	"time"
	//sadas
)

//Faya

func main() {
	fmt.Println("faya")
	fmt.Println(os.Args)
	if len(os.Args) < 2{
		fmt.Println("help")
		return
	}
		
	switch os.Args[1] {
	case "prefill":
		fmt.Println("prefill")
		function.Prefill()
	case "review":
		strategy.ZtReview()
		strategy.Day5DowngradeViewer()
	case "chi":
		l := list.Get()
		l = filter.HoldFilter(l)
		function.Chi(l)
	case "view":
		code := os.Args[2]
		//view.PlotRik(code)
		view.PlotRikMin(code)
	case "viewrik":
		code := os.Args[2]
		view.PlotRik(code)
	case "viewmin":
		code := os.Args[2]
		now := time.Now()
		targetDate := now
		if len(os.Args) > 3 {
			det, err := strconv.Atoi(os.Args[3])
			if err != nil {
				panic(err)
			}
			targetDate = now.AddDate(0, 0, det)
		}
		y, m, d := targetDate.Date()
		tm := fmt.Sprintf("%d-%02d-%02d", y, m, d)

		fmt.Println("hisotry zt review", now, "targetTm:", tm)
		view.PlotMin(code, tm)

	case "bigchi":
		fmt.Println("bigchi")
		fmt.Println("ensure time now is in opening time")
		now := time.Now()
		targetDate := now
		if len(os.Args) > 2 {
			det, err := strconv.Atoi(os.Args[2])
			if err != nil {
				panic(err)
			}
			targetDate = now.AddDate(0, 0, det)
		}
		y, m, d := targetDate.Date()
		tm := fmt.Sprintf("%d-%02d-%02d", y, m, d)
		fmt.Println("hisotry zt review", now, "targetTm:", tm)
		ar := strategy.HistoryZtOnly(tm)
		function.ShowChi(ar)
	case "dingban":
		fmt.Println("dingban")
		//function.Dingban()
	case "dingpan":
		fmt.Println("dingpan")
	case "fengdan":
		fmt.Println("fengdan")
		list.FengdanCode("600180")


	case "history":
		fmt.Println("hisotry zt review")
		// TODO remember to correct the timezone and when you exec it before opening time
		now := time.Now()
		targetDate := now
		if len(os.Args) > 2 {
			det, err := strconv.Atoi(os.Args[2])
			if err != nil {
				panic(err)
			}
			targetDate = now.AddDate(0, 0, det)
		}
		y, m, d := targetDate.Date()
		tm := fmt.Sprintf("%d-%02d-%02d", y, m, d)
		fmt.Println("hisotry zt review", now, "targetTm:", tm)

		strategy.HistoryZtReview(tm)

	default:
		fmt.Println("default", os.Args)
	}


	//function.TuishiProducer()
// 
// 	view.PlotMin("000665", "2022-01-12")

// 	l = filter.Filter300(l)

// 	l = filter.RecentZtFilter(l)
// 	//view.Plot(l[0])


/*
	for _, o := range l{
		//list.GetBkCode(o.Code)
		p := list.RiKCodeReverse(o.Code)
		if len(p) <= 0 || p[0].Date != "2022-01-10" {
			fmt.Println(o)
			if len(p) > 1 {
				fmt.Println(p[0])
			}
		}
	}
 	*/

// 	list.GetBkCode("301111")
 	//list.MinCode("301111")
// 	strategy.Day5Viewer()
// 	strategy.LianXuXiaoYangXian()

	/*
	a := list.RiKCode("300364")
	features.GetDay5(a)
	features.GetDay5Det(a)
// 	fmt.Println(t)
	for i:= 0; i < 10 && i < len(a); i = i + 1{
		fmt.Println(a[i])
	}
	*/

}
