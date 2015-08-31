/*
 * memcache.go
 *
 *  Created on: 31/08/2015
 *      Author: miaohong(miaohong01@baidu.com)
 */


package module

import (
	"fmt"
	"strconv"
	"strings"

	"../glog"
)


var memcacheMonitorFuncs []string

func init() {
	memcacheMonitorFuncs = []string{"Memcache->set(", "Memcache->get("}
}


type MemcacheMonitor struct {
	rawData          string
	funcDuration     []map[string]float64
}

func NewMemcacheMonitor(rawData string) *MemcacheMonitor {
	return &MemcacheMonitor {
		rawData      : rawData,
		funcDuration : make([]map[string]float64, 0),
	}
}

func (m *MemcacheMonitor) Parse() error {
	//fmt.Println(m.rawData)

	rawDataList := strings.Split(m.rawData, "+")

	for _, v := range rawDataList {
		
		for _, vv := range memcacheMonitorFuncs {
			if strings.Contains(v, vv) && strings.Contains(v, "wt") {
				fd := make(map[string]float64)

				indexWt := strings.Index(v, "wt")
				indexTotal := strings.Index(v, "total")
				
				tmpTotal, err := strconv.ParseFloat(strings.TrimSpace(strings.Split(v[indexTotal:], ":")[1]), 64)
				if err != nil {
					glog.Error(err.Error())
					return err
				}

				fd[v[:indexWt]] = tmpTotal
				//fmt.Println(m)

				m.funcDuration = append(m.funcDuration, fd)
				fmt.Println(m.funcDuration)
			}
		}
	}

	return nil
}
