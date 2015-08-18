/*
 * http.go
 *
 *  Created on: 13/08/2015
 *      Author: miaohong(miaohong01@baidu.com)
 */

package proxy

import (
    "net/http"
    "io/ioutil"
    "fmt"
    "time"
    "bytes"
    "../glog"
)

type ReqHttp struct {
	httpClient *http.Client
	method     string  
	url        string
	header     http.Header   
}

func NewReqHttp(url string, method string, timeout int) *ReqHttp {
	client := http.Client {
	    Timeout: time.Duration(5 * time.Second),
	}

	return &ReqHttp {
		method     : method,
		url        : url,
		httpClient : &client,
	}
}

func (r *ReqHttp) AddHeader(key, val string) {
	r.header.Add(key, val)
}

func (r *ReqHttp) DoGetData() error {
	var err error
	request, err := http.NewRequest(r.method, r.url, nil)
	if err != nil {
		glog.Error(err.Error())
		return err
	}

	//add header 
	request.Header = r.header

	response, err := r.httpClient.Do(request)
	if err != nil {
		glog.Error(err.Error())
		return err
	}

    if response.StatusCode == 200 {
        body, err := ioutil.ReadAll(response.Body)
        if err != nil {
 			glog.Error(err.Error())
			return err       	
        }
        bodystr := string(body);
        fmt.Println(bodystr)
    }

    return err
}

func (r *ReqHttp) DoPostData(body []byte) error {
	var err error

	request, err := http.NewRequest(r.method, r.url, bytes.NewReader(body))
	if err != nil {
		glog.Error(err.Error())
		return err
	}

	//add header 
	request.Header = r.header

	response, err := r.httpClient.Do(request)

	if err != nil {
		glog.Error(err.Error())
		return err
	}
	
	if response.StatusCode == 200 {
        body, err := ioutil.ReadAll(response.Body)
        if err != nil {
 			glog.Error(err.Error())
			return err       	
        }
        bodystr := string(body);
        fmt.Println(bodystr)
    }

	return err
}
 