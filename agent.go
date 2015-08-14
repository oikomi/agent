/*
 * agent.go
 *
 *  Created on: 13/08/2015
 *      Author: miaohong(miaohong01@baidu.com)
 */

 package main

import (
 	"fmt"
 	"flag"
	"./glog"
)


/*
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
const char* build_time(void) {
	static const char* psz_build_time = "["__DATE__ " " __TIME__ "]";
	return psz_build_time;
}
*/
import "C"

var (
	buildTime = C.GoString(C.build_time())
)

func BuildTime() string {
	return buildTime
}

const VERSION string = "0.0.1"

func version() {
	fmt.Printf("php agent version %s \n", VERSION)
}

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

var InputConfFile = flag.String("conf_file", "agent.json", "input conf file name") 


 func main() {
	version()
	fmt.Printf("built on %s\n", BuildTime())
	flag.Parse()
	cfg := NewAgentConfig(*InputConfFile)
	err := cfg.LoadConfig()
	if err != nil {
		glog.Error(err.Error())
		return
	}
	

}