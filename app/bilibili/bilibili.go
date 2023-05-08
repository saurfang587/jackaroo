package Bilibili

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// 调用AllIdInformation 和 GetIndex 两个函数实现
// -----------------------------------------
// 用于存储所有的id信息
var AllId []string
var receive1 Context1
var receive Context

//用来收集所有的页面信息

// 抓取所有id信息
func AllIdInformation() {
	url := "https://jobs.bilibili.com/api/campus/position/positionList"
	//构造请求体
	fmt.Println("爬虫启动，正在爬取所有职位id：")
	for i := 1; i <= 19; i++ {
		fmt.Printf("正在爬取第%d页\n", i)
		str := strconv.Itoa(i)
		requestBody := map[string]interface{}{
			"pageSize":         10,
			"pageNum":          str,
			"positionName":     "",
			"postCode":         []interface{}{},
			"postCodeList":     []interface{}{},
			"workLocationList": []interface{}{},
			"workTypeList":     []string{"0"},
			"positionTypeList": []string{"0"},
			"deptCodeList":     []interface{}{},
			"recruitType":      nil,
		}
		//转化为byte[]类型
		resq, err := json.Marshal(requestBody)
		if err != nil {
			fmt.Println("转化失败，读取所有id信息失败，请重新检验")
			return
		}
		Fetch(url, "POST", resq)
	}
	//fmt.Println(AllId)
	fmt.Println("提取完毕")
}

// 抓取职位信息，主要作用：抓取到页面的id，为之后的每一页信息爬取做铺垫
func Fetch(url string, Method string, requestBody []byte) {
	req, err := http.NewRequest(Method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println(err)
		return
	}
	// 添加请求头部信息
	req.Header.Add("authority", "jobs.bilibili.com")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Add("content-type", "application/json;charset=UTF-8")
	req.Header.Add("origin", "https://jobs.bilibili.com")
	req.Header.Add("referer", "https://jobs.bilibili.com/campus/positions?type=0&page=2")
	req.Header.Add("sec-ch-ua", `"Chromium";v="112", "Google Chrome";v="112", "Not:A-Brand";v="99"`)
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", `"Windows"`)
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
	req.Header.Add("x-appkey", "ops.ehr-api.auth")
	req.Header.Add("x-channel", "campus")
	req.Header.Add("x-csrf", "728e3c6c-e981-412e-be94-ab1f76557972")
	req.Header.Add("x-usertype", "2")

	// 发送请求并获取响应
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// 处理响应数据
	err = json.NewDecoder(resp.Body).Decode(&receive)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	//将所有的id信息添加到Allid数据中
	for i := 0; i < 10; i++ {
		//AllId = append(AllId, receive.Data.List[i].Id)
		if len(receive.Data.List) == i {
			return
		}
		AllId = append(AllId, receive.Data.List[i].Id)
	}
}

// 用来爬取每页的信息
func Fetch1(url string, Method string, requestBody []byte) {
	req, err := http.NewRequest(Method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println(err)
		return
	}
	// 添加请求头部信息
	req.Header.Add("authority", "jobs.bilibili.com")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Add("content-type", "application/json;charset=UTF-8")
	req.Header.Add("origin", "https://jobs.bilibili.com")
	req.Header.Add("referer", "https://jobs.bilibili.com/campus/positions?type=0&page=2")
	req.Header.Add("sec-ch-ua", `"Chromium";v="112", "Google Chrome";v="112", "Not:A-Brand";v="99"`)
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", `"Windows"`)
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
	req.Header.Add("x-appkey", "ops.ehr-api.auth")
	req.Header.Add("x-channel", "campus")
	req.Header.Add("x-csrf", "728e3c6c-e981-412e-be94-ab1f76557972")
	req.Header.Add("x-usertype", "2")

	// 发送请求并获取响应
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// 处理响应数据
	err = json.NewDecoder(resp.Body).Decode(&receive1)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	//写进数据库
	//err1 := DB.AutoMigrate(&receive1.Data)
	//if err1 != nil {
	//	return
	//}
	//err2 := DB.Create(&receive1.Data).Error
	//if err2 != nil {
	//	fmt.Println("插入数据失败")
	//	return
	//}
	fmt.Println(receive1)
}

func GetIndex() {
	//err := DB.AutoMigrate(&AllInformation)
	//if err != nil {
	//	return
	//}
	for i := 0; i < len(AllId); i++ {
		url := "https://jobs.bilibili.com/api/campus/position/detail/" + AllId[i]
		Fetch1(url, "GET", nil)
	}
}
