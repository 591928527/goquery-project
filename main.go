// 爬取豆瓣电影 TOP250
package main

import (
	"log"

	"qoquery/model"
	"qoquery/parse"
	"qoquery/pkg/request"
)

var (
	BaseUrl = "https://movie.douban.com/top250"
)

// 新增数据
func Add(movies []parse.DoubanMovie) {
	for index, movie := range movies {
		if err := model.DB.Create(&movie).Error; err != nil {
			log.Printf("db.Create index: %d, err : %v", index, err)
		} else {
			log.Printf("add data success!")
		}
	}

}

// 开始爬取
func Start() {
	var movies []parse.DoubanMovie
	pages := parse.GetPages(BaseUrl)
	for _, page := range pages {
		doc, err := request.Get(BaseUrl + page.Url)
		if err != nil {
			log.Println(err)
		}
		movies = append(movies, parse.ParseMovies(doc)...)
	}
	Add(movies)
}

func main() {
	Start()
	defer model.DB.Close()
}
