package main

import (
	"fmt"
	"io/ioutil"
)

var jsonfiles = [6]string{
	"windrealpowerdata.json",
	"windtowerdata.json",
	"windtowerstatus.json",
	"lightrealpowerdata.json",
	"lightobservationdata.json",
	"inverterstatus.json"}

func GetJsonFromFile(inx uint) (string, error) {
	if inx > 5 {
		return "", StringErr("error inx")
	}
	dat, err := ioutil.ReadFile(jsonfiles[inx])
	if err != nil {
		fmt.Println("Read file ", jsonfiles[inx], "error")
		return "", err
	}
	return string(dat), nil
}
