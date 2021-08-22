package http

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"github.com/mymmsc/gox/api"
	"github.com/mymmsc/gox/logger"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	urlpkg "net/url"
	"strings"
)

func HttpGet0(url string) ([]byte, error) {
	logger.Debugf("url=[%s]\n", url)
	res, err := http.Get(url)
	if err != nil {
		logger.Errorf("url=[%s], err=[%+v]\n",  url, err)
		return nil, err
	}
	//defer res.Body.Close()
	defer api.CloseQuietly(res.Body)
	data, err := ioutil.ReadAll(transform.NewReader(res.Body, simplifiedchinese.GBK.NewDecoder()))
	if err != nil {
		logger.Errorf("url=[%s], err=[%+v]\n",  url, err)
		return nil, err
	}
	//sret, err := json.Marshal(res)
	// request={}, params=[{}], http-status=[{}], body=[{}], message=[{}], acrossTime={}
	status := res.StatusCode
	response := data
	logger.Infof("url=[%s], http-status=[%d], response=[%s], error=[%+v]\n", url, status, response, err)
	return data, nil
}

func HttpRequest(url string, method string) ([]byte, error) {
	u, err := urlpkg.Parse(url)
	if err != nil {
		return nil, err
	}
	header := make(map[string]string)
	header["Accept"] = "Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"
	header["Accept-Encoding"] = "gzip, deflate"
	header["Accept-Language"] = "zh-CN,zh;q=0.9,en;q=0.8"
	header["Cache-Control"] = "no-cache"
	header["Connection"] = "keep-alive"
	header["Host"] = u.Host
	header["Pragma"] = "no-cache"
	header["Upgrade-Insecure-Requests"] = "1"
	header["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36"

	client := &http.Client{}
	request, err := http.NewRequest(strings.ToUpper(method), url, nil)
	if err != nil {
		return nil, err
	}
	for key, v := range header {
		request.Header.Add(key, v)
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	contentEncoding := response.Header.Get("Content-Encoding")
	var reader io.ReadCloser = nil
	if len(contentEncoding) > 0 {
		contentEncoding = strings.ToLower(contentEncoding)

		switch contentEncoding {
		case "gzip":
			reader, err = gzip.NewReader(bytes.NewBuffer(body))
			if err != nil {
				logger.Error(err)
				reader = nil
			}
		case "deflate":
			reader = flate.NewReader(bytes.NewReader(body))
		}
	}
	if reader != nil {
		defer reader.Close()
		body, err = ioutil.ReadAll(reader)
		if err != nil {
			return nil, err
		}
	}
	return body, nil
}

func HttpGet(url string) ([]byte, error) {
	return HttpRequest(url, "get")
}