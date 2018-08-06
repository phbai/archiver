package main

import (
	// "fmt"
	// "time"

	"github.com/phbai/archiver/util"
)

type Data struct {
	Result Result `json:"result"`
}

type Result struct {
	Docs []Post `json:"docs"`
	Total int `json:"total"`
	Limit int `json:"limit"`
	Page int `json:"page"`
	Pages int `json:"pages"`
}

type Post struct {
	ID string `json:"_id"`
	Title string `json:"title"`
	Link string `json:"link"`
	Viewkey string `json:"viewkey"`
	Author string `json:"author"`
	AuthorURL string `json:"author_url"`
	Thumbnail string `json:"thumbnail"`
	Duration string `json:"duration"`
	Time string `json:"time"`
	Views string `json:"views"`
	Favorites string `json:"favorites"`
	Comments string `json:"comments"`
	Points string `json:"points"`
	Description string `json:"description"`
	HasRead bool `json:"hasRead"`
}

func main() {
	url := "https://api.prpr.io/search"
	Data := &Data{} // or &Foo{}
	util.GetJson(url, Data)
	posts := Data.Result.Docs

	tasks := []Task{}

	// go util.Spinner(100 * time.Millisecond, "正在下载所有文件")

	for _, v := range posts[0:4] {
		task := &Task{Name: v.Title, URL: v.Link}
		tasks = append(tasks, *task)
		go task.Start()
	}

	for _, t := range tasks {
		<- t.IsFinished
	}
	// fileUrl := "https://s3.didiyunapi.com/marisa/3-3.mp4"

	// go util.Spinner(100 * time.Millisecond, "正在下载" + fileUrl)
	// err := util.DownloadFile("backup/3-3.mp4", fileUrl)
	// if err != nil {
	// 		panic(err)
	// }
	// fmt.Println("  \033[32m下载成功\033[0m")
}