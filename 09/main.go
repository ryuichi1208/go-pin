package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://ryuichi1208.hateblo.jp/archive/2022"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic("htmlの取得に失敗しました")

	}
	res := doc.Find("a.entry-title-link")
	res.Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())

	})
}
