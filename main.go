// act
package main

import (
	"mtcomm/db/mysql"
	logger "mtcomm/log"
	"fmt"
)

const (
	token        = "eyJ0eXAiOiJKV1QiLCJhbGciOiJTSEEyNTYifQ__.eyJpYXQiOjE1NDE0MDg4NjYsImV4cCI6MTU0NDAwMDg2NiwidWlkIjoyMDEwMjd9.c6390a41d86f211758da4941468c0222d8343a66221c7f2cb22a9a1a5c50c22a"
	bottlesUrl      = "http://api.syplp.com/bottles/201027"
	userBaseUrl            = "http://api.syplp.com/users/"


	XmThrowBotterUrl       = "http://api.xmplp.com/bottles/293567"
	XmToken="eyJ0eXAiOiJKV1QiLCJhbGciOiJTSEEyNTYifQ__.eyJpYXQiOjE1NDE2MDI4ODAsImV4cCI6MTU0NDE5NDg4MCwidWlkIjoyOTM1Njd9.3212729ea5cf4cd375f129d78aa07f6cb6963702e1f8ac4611924f6511e622cb"
	bucket_name_mtalk = "qx-mtalk-test"//type=1&content=cccccc
)

var(
	ip="127.0.0.1:3306"
	user="root"
	pwd="root"
	dbName="botter_user"

	mysqlClient mysql.MysqlClient
)
func init(){
	logger.SetDefaultLogLevel(1)
	mysqlClient = mysql.NewMysqlClient(&mysql.MysqlInfo{
		UserName:     "root",               //"mtalk",
		Password:     "shQX,34537916Mm123", //"shQX34537916Mm123",
		IP:           "192.168.10.163",     //"rm-uf68s9b57zdc4dhx9o.mysql.rds.aliyuncs.com",
		Port:         "3306",
		DatabaseName: "print",
		Logger:       logger.GetDefaultLogger(),
	})
}

func main() {
	req,_:=mysqlClient.SearchMutiRows(&mysql.Stmt{Sql: "SELECT * FROM USER", Args: []interface{}{}})
	fmt.Println(req)
	//bottlesInfo(token, bottlesUrl, userBaseUrl)
	//userInfo(token, userBaseUrl)

}

type Request struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    RequestMap `json:"data"`
}

type RequestMap struct {
	Content    string `json:"content"`
	Id         int    `json:"id"`
	City       string `json:"city"`
	Avatar     string `json:"avatar"`
	Province   string `json:"province"`
	Area       string `json:"area"`
	Type       int    `json:"type"`
	Difference int    `json:"difference"`
	User_id    int    `json:"user_id"`
	Nickname   string `json:"nickname"`
	Phone      string `json:"phone"`
	Sex        int    `json:"sex"`
}

