package main

import (
	"fmt"
	"mmrights/global"
	_ "mmrights/initialize"
)

func main(){
	fmt.Println("start server")
	fmt.Println(fmt.Sprintf("appname %v", global.AppName))
}