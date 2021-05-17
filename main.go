package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
	"time"
)

func main() {
	startUrl := "http://www.xbiquge.la/"
	resp, err := http.Get(startUrl)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var bookUrlArr []string
	bookUrlList := document.Find("#newscontent > div.l > ul > li")
	bookUrlList.Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Find("a").Attr("href")
		bookUrlArr = append(bookUrlArr, url)
	})
	for _, v := range bookUrlArr {
		go GetEveryBookInfo(v)
	}
	time.Sleep(time.Minute)
}

type BookInfo struct {
	BookName       string
	BookIndexUrl   string
	BookNewChapter string
	BookAuthor     string
}

func GetEveryBookInfo(bookUrl string) {
	bookInfo := BookInfo{}
	resp, err := http.Get(bookUrl)
	if err != nil {
		fmt.Println(err)
	}
	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	bookInfo.BookName = document.Find("#info > h1").Text()
	bookInfo.BookIndexUrl = bookUrl
	bookInfo.BookAuthor = strings.Replace(document.Find("#info > p:nth-child(2)").Text(), "作    者：", "", 1)
	bookInfo.BookNewChapter = document.Find("#info > p:nth-child(5) > a").Text()
	fmt.Println(bookInfo)
}
