package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("test start")
	file := "/home/dev/data/mrepo/cocos-go/tests/files/test.lua"
	bytes, err := ioutil.ReadFile(file)
    if err != nil {
        fmt.Print(err)
    }

	str := string(bytes)
	fmt.Println(str)

	fmt.Print("test end")
}