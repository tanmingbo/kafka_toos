package view

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "kafka_tools/moduls"
    "net/http"
    "strconv"
)

func Start(){
    app := gin.Default()
    app.LoadHTMLGlob("templates/*")

    app.GET("/",index)
    app.POST("/create_topic",create_topic)
    app.POST("alter_topic",alter_topic)
    app.POST("topiclist",topiclist)

    app.Run(":8081")
}

//管理页面默认界面：
func index(c *gin.Context){
    c.HTML(http.StatusOK,"index.html",gin.H{
        "status":"yes",
    })
}

//查询topic的方法：
func topiclist(c *gin.Context){
    topic_list := moduls.Select_topic_info()

    fmt.Println("http处的输出：",topic_list)
    c.HTML(http.StatusOK,"topiclist.html",gin.H{
        "topiclist" : topic_list,
    })
}

//创建topic的方法：
func create_topic(c *gin.Context){
    topic_name,_ :=  c.GetPostForm("topicname")
    server_add,_ :=  c.GetPostForm("server_add")
    parttion_num,_ :=  c.GetPostForm("parttion_number")
    replication_num,_ :=  c.GetPostForm("replication_number")
    fmt.Println(topic_name,server_add,parttion_num,replication_num)

    _ = moduls.Create_topic(topic_name,server_add,parttion_num,replication_num)
    topic_list := moduls.Select_topic_info()
    c.HTML(http.StatusOK,"topiclist.html",gin.H{
        "topiclist" : topic_list,
    })
}

//修改副本的方法：
func alter_topic(c *gin.Context)  {
    server_add,_ :=  c.GetPostForm("server_add")
    topicname,_  :=  c.GetPostForm("topicname")
    parttion,_   :=  c.GetPostForm("parttion")
    replicas1,_   :=  c.GetPostForm("replication_number")

    replicas_int := moduls.Comma(replicas1)
    parttion_int,_ := strconv.Atoi(parttion)

    _ = moduls.Alter_replication(server_add,topicname,parttion_int,replicas_int)


}

