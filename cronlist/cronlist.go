package cronlist

import (
	"faya/cronlist/jobstatus"
	"faya/function"
	"faya/serve/servestrategy"
	"faya/strategy"
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)


type JobListUnit struct {
	Name string
	Func func()
}
var JobList []JobListUnit = []JobListUnit{
JobListUnit{Name: "ss1", Func: Ss1Job},
JobListUnit{Name: "prefill", Func: PrefillJob},
}


func Ss1Job(){
	jobname := "ss1"
	fmt.Println("ss1 job begin")
	js := jobstatus.GetJobStatus(jobname)

	timeTenMinutesBefore := time.Now().Add(time.Minute * -4)
	if (js != nil && js.Status == "complete" && js.LastTime.After(timeTenMinutesBefore)){
		fmt.Println("ss1 is calced in 10 minutes, won't calc again")
	}else {
		jb := &jobstatus.JobStatus{
			Status : "started",
			LastTime : time.Now(),
			ProgressUp : 0,
			ProgressDown : 1,
		}
		jobstatus.SetJobStatus(jobname, jb)

		ss := &servestrategy.ServeStrategy1{}
		ss.Run()

		jb.Status = "complete"
		jb.LastTime = time.Now()
		jb.ProgressUp += 1
		jobstatus.SetJobStatus(jobname, jb)
	}
	fmt.Println("ss1 job end")
}

func PrefillJob(){
	jobname := "prefill"
	fmt.Println("prefill job begin")
	js := jobstatus.GetJobStatus(jobname)

	timeHourBefore := time.Now().Add(time.Hour * -1)
	if (js != nil && js.Status == "complete" && js.LastTime.After(timeHourBefore)){
		fmt.Println("prefill is calced in 10 minutes, won't calc again")
	}else {
		jb := &jobstatus.JobStatus{
			Status : "started",
			LastTime : time.Now(),
			ProgressUp : 0,
			ProgressDown : 3,
		}
		jobstatus.SetJobStatus(jobname, jb)

		function.Prefill()
		jb.Status = "prefilled"
		jb.LastTime = time.Now()
		jb.ProgressUp += 1
		jobstatus.SetJobStatus(jobname, jb)

		strategy.ZtReview()
		jb.Status = "ztreviewed"
		jb.LastTime = time.Now()
		jb.ProgressUp += 1
		jobstatus.SetJobStatus(jobname, jb)

		strategy.Day5DowngradeViewer()
		jb.Status = "complete"
		jb.LastTime = time.Now()
		jb.ProgressUp += 1
		jobstatus.SetJobStatus(jobname, jb)
	}
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
	c.AddFunc("TZ=Asia/Shanghai 0 0-8 * * 1-5", PrefillJob)

	c.Start()

	fmt.Println("cron job list initialized end")
}

func JobRun(jobname string){
	for _,j := range JobList {
		if j.Name == jobname {
			j.Func()
			return
		}
	}
	fmt.Println("jobname: ", jobname, " didn't found")
	return 
}
