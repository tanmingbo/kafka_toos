package main

import (
    "fmt"
    "kafka_tools/moduls"
    "kafka_tools/view"
)

func main(){

moduls.NetworkInfo()
config := moduls.InitConfig()
fmt.Println("配置文件正常",config["arm"])
fmt.Println("配置文件正常",config["kafka_home"])


// 运行http
view.Start()

}

