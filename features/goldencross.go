package features

import (
	"faya/list"
	"strconv"
)

//last day is [0]
func GetDayAvg(rklist []*list.RiKUnit, length int) {
	if length <= 0 {
		return
	}
	featureName := "Day" + strconv.Itoa(length)

	ll := len(rklist)
	sum := 0.0
	head := 0
	tail := 0
	//queue is [head, tail)

	for i := 0; i < ll; i = i + 1{
		rk := rklist[i]
		//add up to length
		for ; tail - head < length && tail < ll;{
			sum += rklist[tail].Close
			tail += 1
		}

		if tail - head  == length {
			//here need ensure sumy == length
			rk.Features[featureName] = sum / float64(tail - head)
		}else {
			rk.Features[featureName] = 0
		}

		//pop it
		sum -= rklist[head].Close
		head += 1
	}
}


