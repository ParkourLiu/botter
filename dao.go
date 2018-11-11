package main

import "mtcomm/db/mysql"

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
	Nickname   string `json:"nickname"`
	Phone      string `json:"phone"`
	Sex        int    `json:"sex"`
	Sign       string `json:"sign"`
}

func i_user(r *RequestMap) error {
	sql := "INSERT INTO `user` SET `user_id`=?,`sex`=?,`nickname`=?,`phone`=?,`province`=?,`city`=?,`area`=?,`sign`=?,createTime=NOW(),`updateTime`=NOW() ON DUPLICATE KEY UPDATE `sex`=?,`nickname`=?,`phone`=?,`province`=?,`city`=?,`area`=?,`sign`=?,`updateTime`=NOW()"
	return mysqlClient.Execute(&mysql.Stmt{Sql: sql, Args: []interface{}{r.Id, r.Sex, r.Nickname, r.Phone, r.Province, r.City, r.Area, r.Sign, r.Sex, r.Nickname, r.Phone, r.Province, r.City, r.Area, r.Sign}})
}
