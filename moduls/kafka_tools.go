package moduls
/*

 */

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "os/exec"
    "strings"
)


/*
1. 查询kafka topic信息（输入kafka url,topic名字）返回topic信息。
2. 添加topic副本功能（）
 */

//读取配置文件,转成map数据类姓：
func InitConfig() map[string]string {
    config := make(map[string]string)

    //默认配置文件放在:程序目录下的/config/kafka.conf
    Pwd,err1 := os.Getwd()
    if err1 != nil{fmt.Println("读取配置文件时获取当前路径失败：")}
    Pwd = Pwd+"/config/kafka.cfg"

    f, err := os.Open(Pwd)
    defer f.Close()
    if err != nil {
        panic(err)
    }

    r := bufio.NewReader(f)
    for {
        b, _, err := r.ReadLine()
        if err != nil {
            if err == io.EOF {
                break
                             }
            panic(err)
                      }

        //去掉字符串前后空白字符：
        s := strings.TrimSpace(string(b))

        //返回=再字符串中的位置，int类型
        index := strings.Index(s, "=")

        //如果‘=’的位置<0 ，说明位置异常，跳过这行字符串。
        if index < 0 {
            continue
        }

        //切割 =左边的字符作为key,右边的字符作为value,并设置map变量的kv数据。
        key := strings.TrimSpace(s[:index])
        if len(key) == 0 {
            continue
        }
        value := strings.TrimSpace(s[index+1:])
        if len(value) == 0 {
            continue
        }
        config[key] = value
    }

    //for循环结束，config文件的dv读取完毕，返回map数据
    return config
}

//查询topic的方法：
func Select_topic_info() string{
    //读取配置文件：
    config := InitConfig()
    cmd := exec.Command("/bin/bash",config["kafka_home"]+"/bin/kafka-topics.sh","--describe","--bootstrap-server","10.0.0.13:9092")
    fmt.Println(config["kafka_config"]+"/bin/kafka-topics.sh")
    /*
    var   Stdin,Stdout bytes.Buffer
    cmd.Stdout = & Stdout
    cmd.Stdin = & Stdin
    _ = cmd.Run()
    */
    out, _ := cmd.CombinedOutput()
    _ = cmd.Wait()

    //fmt.Println("topic 查询：",out,"错误信息：",err)
    return string(out)
}

//创建topic的方法：
func Create_topic(topicname string,server_add string,parttion_number string,replication_number string) string{
    config := InitConfig()
    cmd := exec.Command("/bin/bash",config["kafka_home"]+"/bin/kafka-topics.sh","--create","--bootstrap-server",server_add,parttion_number,replication_number,"--topic",topicname)
    out, _ := cmd.CombinedOutput()
    _ = cmd.Wait()
   return  string(out)
}
