package main

import (
	"os"
	"net/http"
	"io/ioutil"
	"zl/recruit_fetcher/fetcher"
	"log"
)

func main() {
	cookie := getCookie()
	fhttp := fetcher.FHttp{ C: &http.Client{} }
	result := fhttp.Get("https://github.com/", fetcher.Option{Cookie: cookie})
	doc := &fetcher.Doc{Body: result}
	eles := doc.GetByClassName(".header-nav-current-user")
	log.Print(eles.Text())
}

func getCookie() string{
	filePath := "./cookie.txt"
	file, err := os.Open(filePath)
	fetcher.HandleErr(err)
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	fetcher.HandleErr(err)

	return string(data)
}