package main

import (
	"mtcomm/db/mysql"
	"bytes"
	"github.com/golang/go/src/pkg/fmt"
)

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

func s_user(r *RequestMap) ([]map[string]string, error) {
	sqlBuffer := bytes.Buffer{}
	sqlBuffer.WriteString("select * from `user` where 1=1")
	if r.Id != 0 {
		sqlBuffer.WriteString(" and id='")
		sqlBuffer.WriteString(fmt.Sprint(r.Id))
		sqlBuffer.WriteString("'")
	}
	sql := sqlBuffer.String()
	result, err := mysqlClient.SearchMutiRows(&mysql.Stmt{Sql: sql, Args: []interface{}{}})
	if err != nil {
		return result, err
	}
	return result, nil
}
