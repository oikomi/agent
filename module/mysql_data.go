/*
 * mysql_data.go
 *
 *  Created on: 31/08/2015
 *      Author: miaohong(miaohong01@baidu.com)
 */


package module

import (
	// "github.com/bitly/go-simplejson"

	// "../glog"
)


type MysqlData struct {
	TotalReqCount       int64
	AverageRespTime     float64
	Top5Slow            []*MysqlSlowData
}

func NewMysqlData() *MysqlData {
	return &MysqlData {
		Top5Slow  : make([]*MysqlSlowData, 0),
	}
}

type MysqlSlowData struct {
	Time                int64
	Script              string
	SqlDuration         map[string]float64
}

func NewMysqlSlowData() *MysqlSlowData {
	return &MysqlSlowData {
		SqlDuration : make(map[string]float64),
	}
}

