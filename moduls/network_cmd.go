package  moduls
/*
测试exec包的使用方式，与项目关系不大。
 */

import (
    "bytes"
    "fmt"
    "os/exec"
    "regexp"
)

func  NetworkInfo(){
    Stdout := bytes.Buffer{}
    CmdProject := exec.Command("ipconfig")
    CmdProject.Stdout  = &Stdout
    CmdProject.Run()
    //fmt.Printf("命令缓冲区的 数据是：",Stdout.String())

    Info_str := Stdout.String()
    compile,_ := regexp.Compile("IPv4.*")
    IPv4 := compile.FindAllString(Info_str,-1)
    fmt.Println(IPv4)
}


