package v1

import "time"

type DataNode struct {
	Id        int `json:"id"`
	DeviceId  string `json:"deviceId"`
	SensorId  string `json:"sensorId"`
	DataType  string `json:"dataType"`
	Timestamp string `json:"timestamp"`
	Metadata  string `json:"metadata"`
}

type NodeSlice [] DataNode

func (a NodeSlice) Len() int {
	return len(a)
}
func (a NodeSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a NodeSlice) Less(i, j int) bool {
	time1 := a[i].Timestamp
	time2 := a[j].Timestamp
	t1, err := time.Parse("2006-01-02 15:04:05", time1)
	if err != nil {
		return true
	}
	t2, err := time.Parse("2006-01-02 15:04:05", time2)
	if err != nil {
		return false
	}
	return t1.Unix() > t2.Unix()
	//return !t1.Before(t2)
}
