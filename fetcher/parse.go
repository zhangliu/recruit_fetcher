package fetcher

import (
	 "github.com/PuerkitoBio/goquery"
	"strings"
)

type Doc struct {
	Body string
}

func (d *Doc) GetByID(id string) *goquery.Selection {
	doc := getDoc(d.Body)
	ele := doc.Find(id).First()
	return ele
}

func (d *Doc) GetByClassName(className string) *goquery.Selection {
	doc := getDoc(d.Body)
	eles := doc.Find(className)
	return eles
}

func getDoc(text string) *goquery.Document {
	reader := strings.NewReader(text)
	doc, err := goquery.NewDocumentFromReader(reader)
	HandleErr(err)
	return doc
}

