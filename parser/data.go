/*
 * data.go
 *
 *  Created on: 13/08/2015
 *      Author: miaohong(miaohong01@baidu.com)
 */


package parser

import (
 	"sort"
 	"strings"
 	"strconv"
 	"../glog"
)

type StatsSummary struct {
	AverageRespTime     float64
	TotalReqCount       int
	Top5Slow            []*Stats
}

func NewStatsSummary() *StatsSummary {
	return &StatsSummary {
		Top5Slow : make([]*Stats, 0),
	}
}


func StatsParse(path string) (*StatsSummary, error) {
	s := NewStats()
	statsSummary, err := s.parse(path)
	if err != nil {
		glog.Error(err.Error())
		return nil, err
	}

	return statsSummary, nil
}


type Stats struct {
	Time      string
	Script    string
	Host      string
	Client    string
	Method    string
	Status    string
	Duration  string
}

func NewStats() *Stats{
	return &Stats {

	}
}

func (s *Stats) buildStats(data []string) *Stats {
	if len(data) == 7 {
		return &Stats {
			Time     : strings.TrimSpace(data[0]),
			Script   : strings.TrimSpace(data[1]),
			Host     : strings.TrimSpace(data[2]),
			Client   : strings.TrimSpace(data[3]),
			Method   : strings.TrimSpace(data[4]),
			Status   : strings.TrimSpace(data[5]),
			Duration : strings.TrimSpace(data[6]),
		}
	}

	return nil
}


func (s *Stats) parse(path string) (*StatsSummary, error) {
	var err error
	var statsSummary *StatsSummary
	var averageRespTime float64
	statsList := make([]*Stats, 0)

	var totalRespTime float64
	df := NewDataFile(path)

	df.GetAllContent()
	dataList := strings.Split(string(df.content), "\n")

	totalReqCount := len(dataList)

	for _, data := range dataList {
		lineList := strings.Split(string(data), ",")
		stats := s.buildStats(lineList)
		if stats != nil {
			statsList = append(statsList, stats)
		}
		
		if len(lineList) == 7 {
			tmpRespTime , err := strconv.ParseFloat(strings.TrimSpace(lineList[6]), 64)
			if err != nil {
				glog.Error(err.Error())
				return nil, err
			}
			totalRespTime += tmpRespTime

		}

	}

	sort.Sort(StatsWrapper{statsList, func (p, q *Stats) bool {
        tmp1 , err := strconv.ParseFloat(q.Duration, 64) 
		if err != nil {
			glog.Error(err.Error())
			
		}
        tmp2 , err := strconv.ParseFloat(p.Duration, 64)
		if err != nil {
			glog.Error(err.Error())
			
		}

        return tmp1 < tmp2
    }})

	averageRespTime = (totalRespTime / (float64)(totalReqCount))

	if len(statsList) <= 5 {
		statsSummary = NewStatsSummary()
		statsSummary.AverageRespTime = averageRespTime
		statsSummary.TotalReqCount = totalReqCount
		statsSummary.Top5Slow = statsList

		//&StatsSummary {totalRespTime, totalReqCount, statsList}
	} else {
		//statsSummary = &StatsSummary {totalRespTime, totalReqCount, statsList[0:5]}
		statsSummary = NewStatsSummary()
		statsSummary.AverageRespTime = averageRespTime
		statsSummary.TotalReqCount = totalReqCount
		for i:=0; i<5; i++ {
			statsSummary.Top5Slow = append(statsSummary.Top5Slow, statsList[i])
		}
		
	}


	return statsSummary, err

}

func EventsParse(path string) ([]*Events, error) {
	e := NewEvents()
	eventsList, err := e.parse(path)
	if err != nil {
		glog.Error(err.Error())
		return nil, err
	}

	return eventsList, nil
}

type Events struct {
	Time     string
	Script   string
	Msg      string
	Trace    string
}

func NewEvents() *Events {
	return &Events {

	}
}

func (e *Events) buildEvents(data []string) *Events{
	if len(data) == 4 {
		return &Events {
			Time : strings.TrimSpace(data[0]),
			Script : strings.TrimSpace(data[1]),
			Msg : strings.TrimSpace(data[2]),
			Trace : strings.TrimSpace(data[3]),
		}
	}

	return nil
}


func (e *Events) parse(path string) ([]*Events, error){
	eventsList := make([]*Events, 0)
	df := NewDataFile(path)
	df.GetAllContent()
	dataList := strings.Split(string(df.content), "\n")
	for _, data := range dataList {
		//glog.Info(data)
		lineList := strings.Split(string(data), "+")
		events := e.buildEvents(lineList)
		if events != nil {
			eventsList = append(eventsList, events)
		}

	}

	return eventsList ,nil
}