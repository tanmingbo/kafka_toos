package  moduls

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
    "os/exec"
    "strings"
)


//修改分区副本的方法：
func Alter_replication(server_add string,topicname string,parrtion int,replication []int) string{

    P := []Partition{
        {topicname,parrtion,replication},
    }


    r := Replication{
        1,
        P,
    }
    //获取到修改replication的json数据了
    j,_ := json.Marshal(r)
    JSON := string(j)
    //struct是大写，使用的格式必须小写：
    JSON  = strings.ToLower(JSON)
    fmt.Println("得到的json数据",JSON)

    //json写入到文件
    dir,_ := os.Getwd()
    fmt.Println(dir)
    Json_name := "alter_replication.json"
    //file_name := dir+Json_name

    f,err  := os.Create(Json_name)
    if err != nil{
        log.Fatal("文件创建失败！")
    }else {
       _,err = f.WriteString(JSON)
       if err != nil {log.Fatal("文件写入失败",err)}
    }

    //json已准备好，可以执行命令：
    config := InitConfig()
    cmd := exec.Command("/bin/bash",config["kafka_home"]+"/bin/kafka-reassign-partitions.sh","--zookeeper","10.0.0.1:2181","--reassignment-json-file",Json_name,"--execute")
    out,err := cmd.CombinedOutput()
    if err  != nil {fmt.Println("修改副本命令执行失败：",err)}
    _ = cmd.Wait()

    //fmt.Println("topic 查询：",out,"错误信息：",err)
    return string(out)


}



