package initialize

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"mmfile/global"
)

func initConf(path ...string) *viper.Viper {
	fmt.Println("初始化配置文件")
	var config string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		config = global.ConfigFile
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	v.WatchConfig()

	if err := v.Unmarshal(&global.ServerConf); err != nil {
		fmt.Println(fmt.Sprintf("json unmarshal serverconf 配置文件 err %v", err))
		panic(err)
	}

	fmt.Println("配置文件加载成功")

	return v
}
