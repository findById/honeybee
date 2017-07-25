package v1

import (
	"honeybee/repository"
	"time"
	"log"
	"errors"
	"sort"
)

func init() {
	sql := "create table if not exists data_node (" +
		"id integer not null primary key auto_increment" +
		", deviceId varchar(50)" +
		", sensorId varchar(50)" +
		", dataType varchar(10)" +
		", timestamp varchar(50)" +
		", metadata varchar(200)" +
		");"
	_, err := repository.GetDBM().Exec(sql)
	if err != nil {
		log.Println(err)
	}

	ticker := time.NewTicker(time.Hour * 24 * 15)
	go func() {
		for range ticker.C {
			// CleanData(20000)
		}
	}()
}

func Save(node *DataNode) {
	sql := "INSERT INTO data_node (deviceId,sensorId,dataType,timestamp,metadata) VALUES (?,?,?,?,?)"
	_, err := repository.GetDBM().Exec(sql, node.DeviceId, node.SensorId, node.DataType, node.Timestamp, node.Metadata)
	if err != nil {
		log.Println(err)
	}
}

func Find(deviceId, sensorId string, size int) ([]DataNode, error) {
	sql := "SELECT * FROM data_node WHERE deviceId=? AND sensorId=? ORDER BY id DESC LIMIT 0, ?"
	rows, err := repository.GetDBM().Query(sql, deviceId, sensorId, size)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var list []DataNode
	count := -1
	for rows.Next() {
		count++

		var item DataNode
		rows.Scan(&item.Id, &item.DeviceId, &item.SensorId, &item.DataType, &item.Timestamp, &item.Metadata)
		//tmp, err := strconv.ParseInt(item.Timestamp, 10, 0)
		//if err != nil {
		//	tmp = time.Now().Unix()
		//}
		//// item.Timestamp = time.Unix(tmp, 0).Format("2006-01-02 15:04:05")
		//// t, err := time.Parse("2006-01-02 15:04:05", item.Timestamp)
		//item.Timestamp = time.Unix(tmp, 0).Format("2006-01-02 15:04:05")

		list = append(list, item)
	}
	if count == -1 {
		return nil, errors.New("no data")
	}
	sort.Sort(sort.Reverse(NodeSlice(list)))
	return list, nil
}

// 删除老数据
func CleanData(size int) {
	deviceList, err := findDeviceType()
	if err != nil {
		log.Println(err)
		return
	}
	if deviceList == nil || len(deviceList) <= 0 {
		return
	}

	for _, deviceId := range deviceList {
		sensorList, err := findSensorType(deviceId)
		if err != nil {
			log.Println(err)
			continue
		}
		if sensorList == nil || len(sensorList) <= 0 {
			continue
		}
		for _, sensorId := range sensorList {
			deleteData(deviceId, sensorId, size-1)
		}
	}
}

func findDeviceType() ([]string, error) {
	device := "SELECT deviceId FROM data_node GROUP BY deviceId"
	rows, err := repository.GetDBM().Query(device)
	if err != nil {
		return nil, err
	}
	var deviceList []string
	for rows.Next() {
		var deviceId string
		rows.Scan(&deviceId)
		deviceList = append(deviceList, deviceId)
	}
	return deviceList, nil
}

func findSensorType(deviceId string) ([]string, error) {
	sensor := "SELECT sensorId FROM data_node WHERE deviceId=? GROUP BY sensorId"
	rows, err := repository.GetDBM().Query(sensor, deviceId)
	if err != nil {
		return nil, err
	}
	var sensorList []string
	for rows.Next() {
		var deviceId string
		rows.Scan(&deviceId)
		sensorList = append(sensorList, deviceId)
	}
	return sensorList, nil
}

func deleteData(deviceId, sensorId string, size int) {
	del := "DELETE tb FROM data_node AS tb,(" +
		"SELECT id,deviceId,sensorId FROM data_node WHERE deviceId=? AND sensorId=? ORDER BY id DESC LIMIT ?,1" +
		") AS tmp " +
		"WHERE tb.id<tmp.id AND tb.deviceId=tmp.deviceId AND tb.sensorId=tmp.sensorId"
	res, _ := repository.GetDBM().Exec(del, deviceId, sensorId, size)
	affected, err := res.RowsAffected()
	if err != nil {
		return
	}
	log.Printf("%v item deleted", affected)
}
