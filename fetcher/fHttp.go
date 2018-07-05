package fetcher

import (
	"net/http"
	"io/ioutil"
)

type FHttp struct {
	C *http.Client
}

type Option struct {
	Cookie string
}

func (fh *FHttp) Get(url string, opt Option) string{
	req, err := http.NewRequest("GET", url, nil)
	HandleErr(err)

	req.Header.Set("cookie", opt.Cookie)

	res, err := fh.C.Do(req)
	HandleErr(err)
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	HandleErr(err)

	return string(data)
}