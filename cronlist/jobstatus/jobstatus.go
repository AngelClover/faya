package jobstatus

import (
	"encoding/json"
	"faya/db"
	"fmt"
	"time"
)

type JobStatus struct {
	Status string `json:"status"`
	LastTime time.Time `json:"tm"`
	ProgressUp int `json:"pgu"`
	ProgressDown int `json:"pgd"`
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
