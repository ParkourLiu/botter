// act
package main

import (
	"mtcomm/db/mysql"
	logger "mtcomm/log"
	"time"
	"github.com/robfig/cron"
)

const (
	//随缘平台
	token       = "eyJ0eXAiOiJKV1QiLCJhbGciOiJTSEEyNTYifQ__.eyJpYXQiOjE1NDE0MDg4NjYsImV4cCI6MTU0NDAwMDg2NiwidWlkIjoyMDEwMjd9.c6390a41d86f211758da4941468c0222d8343a66221c7f2cb22a9a1a5c50c22a"
	bottlesUrl  = "http://api.syplp.com/bottles/201027"
	userBaseUrl = "http://api.syplp.com/users/"

	//星梦平台
	XmThrowBottleUrl = "http://api.xmplp.com/bottles/293567" //扔捞接口（捞Get,扔POST）
	XmToken          = "eyJ0eXAiOiJKV1QiLCJhbGciOiJTSEEyNTYifQ__.eyJpYXQiOjE1NDE2MDI4ODAsImV4cCI6MTU0NDE5NDg4MCwidWlkIjoyOTM1Njd9.3212729ea5cf4cd375f129d78aa07f6cb6963702e1f8ac4611924f6511e622cb"
	BottleContent    = "?type=1&content="                                                             //扔参数，post请求时跟在接口后，参数需要url.QueryEscape(BottleContent)处理
	UpdateUserUrl    = "http://api.xmplp.com/users/293567?nickname=#&sign=#&province=#&city=#&area=#" //#为需要替换的参数
)

var (
	mysqlClient mysql.MysqlClient
	timeFlag    int
)

func init() {
	timeFlag = time.Now().Hour()
	logger.SetDefaultLogLevel(1)
	//mysqlClient = mysql.NewMysqlClient(&mysql.MysqlInfo{
	//	UserName:     "root",
	//	Password:     "root",
	//	IP:           "127.0.0.1",
	//	Port:         "3306",
	//	DatabaseName: "botter_user",
	//	Logger:       logger.GetDefaultLogger(),
	//})
}

func main() {
	c := cron.New()

	//AddFunc

	//拉取远端数据到学校端来(每个半点拉取一次，对应云端搜集数据，所以不可更改)
	spec := "* 0/5 * * * ?"
	c.AddFunc(spec, cronFunc)

	//启动计划任务
	c.Start()

	//关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer c.Stop()
	select {}
}

func cronFunc() {
	//XmThrowBottle("午睡。。刚睡醒。。。")
	//RequestMap := &RequestMap{
	//	Nickname: "吹过",
	//	Sign:     "青春，像一把风吹过",
	//	Province: "湖南",
	//	City:     "长沙",
	//	Area:     "",
	//}
	//UpdateUser(RequestMap)
	//req, _ := mysqlClient.SearchMutiRows(&mysql.Stmt{Sql: "SELECT * FROM USER", Args: []interface{}{}})
	//fmt.Println(req)
	bottlesInfo(token, bottlesUrl, userBaseUrl)
	//userInfo(token, userBaseUrl)
}
