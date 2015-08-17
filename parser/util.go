/*
 * util.go
 *
 *  Created on: 13/08/2015
 *      Author: miaohong(miaohong01@baidu.com)
 */


package parser

import (
	"sort"
)

 type StatsWrapper struct {
    stats []*Stats
    by func(p, q * Stats) bool
}

func (sw StatsWrapper) Len() int {    // 重写 Len() 方法
    return len(sw.stats)
}
func (sw StatsWrapper) Swap(i, j int){     // 重写 Swap() 方法
    sw.stats[i], sw.stats[j] = sw.stats[j], sw.stats[i]
}
func (sw StatsWrapper) Less(i, j int) bool {    // 重写 Less() 方法
    return sw.by(sw.stats[i], sw.stats[j])
}
 
type SortBy func(p, q *Stats) bool

func SortStats(stats []*Stats, by SortBy){  
    sort.Sort(StatsWrapper{stats, by})
}

