package goini

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

/**
 读取ini文件配置
 @param filename 文件路径
 @param keyname 读取键值 格式 : section->key
 @return value,error
 eg .
   if val,err:= GetValue("/usr/local/etc/php/7.2/php.ini","Session->session.name") ; err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("val is ",val)
	}
 */

func GetValue(filename,keyname string) (string, error) {
	if filename == "" || keyname == "" {
		return "", errors.New("filename or keyname is empty")
	}
	dest := strings.Split(keyname,"->")
	if len(dest) != 2 {
		return "", errors.New("wrong format keyname")
	}
	destSection,destKey := dest[0],dest[1]
	file,err := os.Open(filename)
	if err != nil {
		return  "",err
	}
	reader := bufio.NewReader(file)
	var sectionName string
	for {
		//读取一行
		lineText ,err := reader.ReadString('\n')
		if err != nil {
			return "" ,err
		}
		lineText = strings.TrimSpace(lineText)
		if lineText == "" {
			continue
		}
		if lineText[0] == ';' {
			continue
		}
		if lineText[0] == '[' && lineText[len(lineText) - 1] == ']' {
			sectionName = lineText[1:len(lineText) - 1]
		} else if sectionName == destSection {
			orgString := strings.Split(lineText,"=")
			if len(orgString) != 2 {
				continue
			}
			key, val := strings.TrimSpace(orgString[0]),strings.TrimSpace(orgString[1])
			if key == destKey {
				return val,nil
			}
		}
	}
}


