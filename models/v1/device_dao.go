package v1

import (
	"honeybee/repository"
	"log"
	"errors"
)

func init() {
	sql := "create table if not exists device (" +
		"id integer not null primary key auto_increment" +
		", deviceId varchar(50)" +
		", type varchar(50)" +
		", locate varchar(50)" +
		");"
	_, err := repository.GetDBM().Exec(sql)
	if err != nil {
		log.Println(err)
	}
}

func SaveDevice(device *Device) {
	sql := "INSERT INTO device (deviceId,type,locate) VALUES (?,?,?)"
	_, err := repository.GetDBM().Exec(sql, device.DeviceId, device.DeviceType, device.Locate)
	if err != nil {
		log.Println(err)
	}
}

func FindDeviceById(deviceId string) ([]Device, error) {
	sql := "SELECT * FROM device WHERE deviceId=?"
	rows, err := repository.GetDBM().Query(sql, deviceId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var list []Device
	count := -1
	for rows.Next() {
		count++

		var item Device
		rows.Scan(&item.Id, &item.DeviceId, &item.DeviceType, &item.Locate)
		list = append(list, item)
	}
	if count == -1 {
		return nil, errors.New("no data")
	}
	return list, nil
}
