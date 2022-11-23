package function

import (
	"faya/cronlist/jobstatus"
	"faya/list"
	"fmt"
	"time"
)

func Prefill() {
	jobname := "prefill"
	l := list.Get()
	ll := len(l)
	i := 0
	js := jobstatus.GetJobStatus(jobname)
	js.Status = "prefilling"
	js.ProgressDown = len(l)
	

	for _,o := range l {
		fmt.Println("prefilling ", i, ll)
		list.RiK(o)
		list.Min(o)
		//list.GetBk(o)

		i = i + 1
		js.ProgressUp += 1
		js.LastTime = time.Now()
		jobstatus.SetJobStatus(jobname, js)
	}
}
