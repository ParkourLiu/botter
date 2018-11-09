package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"net/http"
	"encoding/json"
	"bytes"
	"net/url"
)

func userInfo(token string, userBaseUrl string) {
	for i := 122500; ; i++ {
		userUrl := userBaseUrl + fmt.Sprint(i)
		user := Get(userUrl, token)
		if user.Status == 200 {
			fmt.Println(user.Data.User_id, "	", user.Data.Sex, "	", user.Data.Nickname, "	", user.Data.Phone, "	", user.Data.Province, "	", user.Data.City, "	", user.Data.Area)
			fd, err := os.OpenFile("C:/Users/Parkour/go/src/test2/a_fffffff/user.txt", os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				fmt.Println(err)
			}
			fd_content := strings.Join([]string{fmt.Sprint(i), fmt.Sprint(user.Data.Sex), user.Data.Nickname, user.Data.Phone, user.Data.Province, user.Data.City, user.Data.Area, "\n"}, "	")
			buf := []byte(fd_content)
			fd.Write(buf)
			fd.Close()
		}
		if i%500 == 0 {
			fmt.Println("已到", i)
		}

	}
}
func bottlesInfo(token string, bottlesUrl string, userBaseUrl string) {
	for ; ; {
		bottles := Get(bottlesUrl, token)
		//fmt.Println(bottles)
		if bottles.Status != 200 {
			//fmt.Println("爬光啦，等待。。。。")
			time.Sleep(5 * time.Second)
			continue
		}
		userUrl := userBaseUrl + fmt.Sprint(bottles.Data.User_id)
		user := Get(userUrl, token)
		fmt.Println(user.Data.Sex, "	", user.Data.Phone, "	", bottles.Data.Nickname, "	", bottles.Data.Province, "	", bottles.Data.City, "	", bottles.Data.Content)
		fd, err := os.OpenFile("C:/Users/Parkour/go/src/test2/a_fffffff/data.txt", os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println(err)
		}
		fd_content := strings.Join([]string{fmt.Sprint(user.Data.Sex), user.Data.Phone, bottles.Data.Nickname, bottles.Data.Province, bottles.Data.City, bottles.Data.Area, bottles.Data.Content, "\n"}, "	")
		buf := []byte(fd_content)
		fd.Write(buf)
		fd.Close()
	}
}

func Get(url string, token string) *Request {
	req, _ := http.NewRequest("GET", url, nil)
	headMap := req.Header
	headMap["token"] = append(headMap["token"], token)
	client := &http.Client{}
	resp, _ := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}

	var request *Request
	if err := json.NewDecoder(resp.Body).Decode(&request); err != nil {
		fmt.Println("错误：", err)
	}
	return request
}

func Post(method string, url string, param string, token string) *Request {
	fmt.Println("url:", url)
	fmt.Println("param:", param)
	req, _ := http.NewRequest(method, url, bytes.NewReader([]byte(param)))
	headMap := req.Header
	headMap["token"] = append(headMap["token"], token)
	client := &http.Client{}
	resp, _ := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	var request *Request
	if err := json.NewDecoder(resp.Body).Decode(&request); err != nil {
		fmt.Println("错误：", err)
	}
	return request
}

//json转map
func json2Map(str string) (map[string]interface{}, error) {
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(str), &dat); err == nil {
		return dat, nil
	} else {
		return make(map[string]interface{}), err
	}
}

//修改用户信息
func UpdateUser(requestMap *RequestMap) bool {
	param := []string{}
	param = append(param, requestMap.Nickname) //名字
	param = append(param, requestMap.Sign)     //个性签名
	param = append(param, requestMap.Province) //省份
	param = append(param, requestMap.City)     //城市
	param = append(param, requestMap.Area)     //区

	upUserUrl := UpdateUserUrl
	for _, v := range param {
		upUserUrl = strings.Replace(upUserUrl, "#", url.QueryEscape(v), 1)
	}
	req := Post("PUT", upUserUrl, "{}", XmToken)
	fmt.Println(req)
	if req.Status == 200 {
		return true
	}
	return false
}

//扔瓶子
func XmThrowBottle(content string) {
	reqUrl := XmThrowBottleUrl + BottleContent + url.QueryEscape(content)
	req := Post("POST", reqUrl, "{}", XmToken)
	fmt.Println(req)
}
