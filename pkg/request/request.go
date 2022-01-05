package request

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"qoquery/pkg/upload"
	"qoquery/pkg/util"

	"github.com/PuerkitoBio/goquery"
)

func Get(url string) (*goquery.Document, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.25 Safari/537.36 Core/1.70.3756.400 QQBrowser/10.5.4039.400")

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Fatal("query topic failed", err.Error())
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	return doc, err
}

func GetImg(url string, filename string) (string, bool) {

	r, _ := http.NewRequest("GET", url, nil)
	r.Header.Set("Referer", "https://movie.douban.com/top250?start=225&filter=")

	resp, err := (&http.Client{}).Do(r)

	// resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	path, _ := os.Getwd()
	filename = upload.GetImageFullPath() + util.EncodeMD5(filename) + ".jpg"
	pathfile := path + filename
	err = ioutil.WriteFile(pathfile, bytes, 0666)
	if err != nil {
		return "", false
	} else {
		return filename, true
	}

}
