/*
 * file.go
 *
 *  Created on: 13/08/2015
 *      Author: miaohong(miaohong01@baidu.com)
 */

package parser

import (
 	"os"
)

type DataFile struct {
	filePath    string
}


func (f *DataFile)isFileExist() bool {
	_, err := os.Stat(f.filePath)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}