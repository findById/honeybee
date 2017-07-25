package v1

import (
	"honeybee/repository"
	"log"
	"errors"
)

func init() {
	sql := "create table if not exists sensor (" +
		"id integer not null primary key auto_increment" +
		", deviceId varchar(50)" +
		", sensorId varchar(50)" +
		", name varchar(50)" +
		", type varchar(50)" +
		");"
	_, err := repository.GetDBM().Exec(sql)
	if err != nil {
		log.Println(err)
	}
}

func SaveSensor(item *Sensor) {
	sql := "INSERT INTO sensor (deviceId,sensorId,name,type) VALUES (?,?,?,?)"
	_, err := repository.GetDBM().Exec(sql, item.DeviceId, item.SensorId, item.SensorName, item.SensorType)
	if err != nil {
		log.Println(err)
	}
}

func FindSensorById(deviceId, sensorId string) ([]Sensor, error) {
	sql := "SELECT * FROM `sensor` WHERE deviceId=? AND sensorId=?"
	rows, err := repository.GetDBM().Query(sql, deviceId, sensorId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []Sensor
	count := -1
	for rows.Next() {
		count++

		var item Sensor
		rows.Scan(&item.Id, &item.DeviceId, &item.SensorId, &item.SensorName, &item.SensorType)
		list = append(list, item)
	}
	if count == -1 {
		return nil, errors.New("no data")
	}
	return list, nil
}

func FindSensorByDeviceId(deviceId string) ([]Sensor, error) {
	sql := "SELECT * FROM `sensor` WHERE deviceId=?"
	rows, err := repository.GetDBM().Query(sql, deviceId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []Sensor
	count := -1
	for rows.Next() {
		count++

		var item Sensor
		rows.Scan(&item.Id, &item.DeviceId, &item.SensorId, &item.SensorName, &item.SensorType)
		list = append(list, item)
	}
	if count == -1 {
		return nil, errors.New("no data")
	}
	return list, nil
}
