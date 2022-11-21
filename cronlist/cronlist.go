package cronlist

import (
	"encoding/json"
	"faya/db"
	"faya/function"
	"faya/serve"
	"faya/strategy"
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

type JobStatus struct {
	Status string `json:"status"`
	LastTime time.Time `json:"tm"`
	ProgressUp int `json:"pgu"`
	ProgressDown int `json:"pgd"`
}


type JobListUnit struct {
	Name string
	Func func()
}
var JobList []JobListUnit = []JobListUnit{
JobListUnit{Name: "ss1", Func: Ss1Job},
JobListUnit{Name: "prefill", Func: PrefillJob},
}

func GetJobCacheKey(jn string) string{
	return "job_status_" + jn
}
func GetJobStatus(jobname string) *JobStatus {
	ck := GetJobCacheKey(jobname)
	ct, succ := db.SimpleGet(ck)
	if succ {
		var js JobStatus
		ct_byte := []byte(ct)
		err := json.Unmarshal(ct_byte, &js)
		if err != nil {
			fmt.Println("dbdata unmarshal error", err.Error)
			fmt.Println(ct)
			return nil
		}else {
			return &js
		}
	} else {
		return nil
	}
}
func SetJobStatus(jobname string, js *JobStatus) {
	ck := GetJobCacheKey(jobname)

	ct, err := json.Marshal(js)
	if err != nil {
		fmt.Println("set job status err ", jobname, err)
		panic(err)
	}

	db.SimpleInsert(ck, string(ct))

}


func Ss1Job(){
	jobname := "ss1"
	fmt.Println("ss1 job begin")
	js := GetJobStatus(jobname)

	timeTenMinutesBefore := time.Now().Add(time.Minute * -10)
	if (js != nil && js.Status == "complete" && js.LastTime.After(timeTenMinutesBefore)){
		fmt.Println("ss1 is calced in 10 minutes, won't calc again")
	}else {
		jb := &JobStatus{
			Status : "started",
			LastTime : time.Now(),
			ProgressUp : 0,
			ProgressDown : 1,
		}
		SetJobStatus(jobname, jb)

		ss := &serve.ServeStrategy1{}
		ss.Run()

		jb.Status = "complete"
		jb.LastTime = time.Now()
		jb.ProgressUp += 1
		SetJobStatus(jobname, jb)
	}
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
