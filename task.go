package main

import (
	"fmt"
	// "time"

	"./util"
)
/**
* 创建
* 判断是否处理过了 如果未处理 通过Parse解析播放url 并进行下载
* 下载成功后更新处理状态 如果下载失败则重新进行下载
* 初始化 →  解析      → 下载   → 完成
* 未开始 →  解析地址中 → 下载中 → 完成
*/
type Task struct {
	Name string
	URL string
	Status string
	VideoURL string
	IsFinished chan bool
}

type Response struct {
	Result string `json:"result"`
}

func (t *Task) Init() {
	t.Status = "初始化中"
}

func (t *Task) Parse() {
	url := "https://api.prpr.io/parse?url=" + t.URL
	response := &Response{} // or &Foo{}
	util.GetJson(url, response)

	t.VideoURL = response.Result
	t.Status = "解析地址中"
}

func (t *Task) Download() {
	filename := t.Name + ".mp4"
	// fmt.Println("正在下载" + filename)
	// go util.Spinner(1000 * time.Millisecond, "正在下载" + filename)
	err := util.DownloadFile("backup/" + filename, t.VideoURL)
	if err != nil {
		t.Status = "下载出错"
		t.IsFinished <- true
	}
	fmt.Println(filename + "  \033[32m下载成功\033[0m")
	t.IsFinished <- true
}

func (t *Task) Start() {
	t.Parse()
	t.Download()
}