package  moduls

import (
    "strconv"
    "strings"
)

//json整体结构
type Replication struct {
    Version int
    Partitions []Partition
}

//json单个parttion：
type  Partition struct {
    Topic string
    Partition int
    Replicas []int
}


//一个[]str 转  []int的工具
func Conv_strToint(S []string)[]int{
    const  n   =  20
    var I  [n]int
    for i:=0 ;i< len(S) ;i++{
        I[i],_ = strconv.Atoi(S[i])
    }
    return I[0:len(S)]
}


// 分隔逗号的工具：例如  string"1,2,3" 转成  comma[1,3,4]
func Comma(c string)  []int {
    comma_str := strings.Split(c,",")
    comma_int := Conv_strToint(comma_str)
    return comma_int
}