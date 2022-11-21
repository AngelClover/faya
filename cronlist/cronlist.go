package cronlist

import (
	"faya/function"
	"faya/serve"
	"faya/strategy"
	"fmt"

	"github.com/robfig/cron/v3"
)

func Ss1Job(){
	fmt.Println("ss1 job begin")
	ss := &serve.ServeStrategy1{}
	ss.Run()
	fmt.Println("ss1 job end")
}

func PrefillJob(){
	fmt.Println("prefill job begin")
	function.Prefill()
	strategy.ZtReview()
	strategy.Day5DowngradeViewer()
	fmt.Println("prefill job end")
}

func Init() {
	fmt.Println("cron job list initializing start")
	c := cron.New()
	//c.AddFunc("TZ=Asia/Tokyo 30 04 * * * *", func() { fmt.Println("Runs at 04:30 Tokyo time every day") })

	c.AddFunc("TZ=Asia/Shanghai 15,20,25,30,40,50 9 * * 1-5", Ss1Job)
	c.AddFunc("TZ=Asia/Shanghai /10 10 * * 1-5", Ss1Job)
	c.AddFunc("TZ=Asia/Shanghai 0-40/10 11 * * 1-5", Ss1Job)
	//11:40 for morning closing
	c.AddFunc("TZ=Asia/Shanghai /10 13-14 * * 1-5", Ss1Job)
	c.AddFunc("TZ=Asia/Shanghai 0,10 15 * * 1-5", Ss1Job)
	//15:10 for afternoon closing

	c.AddFunc("TZ=Asia/Shanghai 0 16-23 * * 1-5", Ss1Job)
	//each hour run for test

	c.AddFunc("TZ=Asia/Shanghai 0 16-23 * * 1-5", PrefillJob)
	c.AddFunc("TZ=Asia/Shanghai 37 0-8 * * 1-5", PrefillJob)

	c.Start()

	fmt.Println("cron job list initialized end")
}
