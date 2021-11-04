package main

import (
	"fmt"
	"mmfile/global"
	_ "mmfile/initialize"
)

func main(){
	fmt.Println("start server")
	fmt.Println(fmt.Sprintf("appname %v", global.AppName))
}