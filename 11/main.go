package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
)

func main() {
	url := "https://qiita.com/"
	// Getリクエスト
	res, _ := http.Get(url)
	defer res.Body.Close()

	// 読み取り
	buf, _ := ioutil.ReadAll(res.Body)

	// 文字コード判定
	det := chardet.NewTextDetector()
	detRslt, _ := det.DetectBest(buf)
	fmt.Println(detRslt.Charset)

	bReader := bytes.NewReader(buf)
	reader, _ := charset.NewReaderLabel(detRslt.Charset, bReader)

	// HTMLパース
	doc, _ := goquery.NewDocumentFromReader(reader)

	rslt := doc.Find("title").Text()
	fmt.Println(rslt)

	res2 := doc.Find("div.css-16lfj6j")
	res2.Each(func(i int, s *goquery.Selection) {
		fmt.Printf("%s\t%s\n", s.Find("div.css-1038633").Text(), s.Find("p.css-1xfnzth").Text())
	})
}
