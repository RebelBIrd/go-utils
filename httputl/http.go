package httputl

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

func DoHttp(methodType MethodType, url string, header map[string]string, body map[string]interface{}, response *HttpResponse) {
	response.IsSuccessChan = make(chan bool)
	doNetWork(HttpParam{
		Method: methodType,
		Url:    url,
		Header: header,
		Body:   body,
		Result: response,
	})
}

func doNetWork(param HttpParam) {
	var client = &http.Client{}
	var bodyStr = ""
	if len(param.Body) > 0 {
		for key, value := range param.Body {
			var val string
			switch value.(type) {
			case string:
				val = value.(string)
			case int:
				val = strconv.Itoa(value.(int))
			default:
				v, _ := json.Marshal(value)
				val = string(v)
			}
			bodyStr += key + "=" + val + "&"
		}
		bodyStr = bodyStr[0 : len(bodyStr)-1]
	}
	var method = ""
	switch param.Method {
	case GET:
		method = "GET"
		break
	case POST:
		method = "POST"
		break
	case PUT:
		method = "PUT"
		break
	case DELETE:
		method = "DELETE"
		break
	}
	req, err := http.NewRequest(method, param.Url, strings.NewReader(bodyStr))
	if err != nil {
		param.Result.Err = err
		param.Result.IsSuccessChan <- false
	} else {
		if len(param.Header) > 0 {
			for key, value := range param.Header {
				req.Header.Add(key, value)
			}
		}
		resp, err := client.Do(req)
		if err != nil {
			param.Result.Err = err
			param.Result.IsSuccessChan <- false
		} else {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				param.Result.Err = err
				param.Result.IsSuccessChan <- false
			} else {
				err = json.Unmarshal(body, param.Result.Result)
				if err != nil {
					param.Result.Err = err
					param.Result.IsSuccessChan <- false
				} else {
					param.Result.IsSuccessChan <- true
				}
			}
		}
	}
}

func DownloadFile(url string, savePath *string, channel chan<- error, processChan chan<- int)  {
	var (
		fSize int64
		buf = make([]byte, 32 * 1024)
		written int64
	)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		channel <- err
	} else {
		if !strings.HasSuffix(*savePath, "/") {
			*savePath += "/"
		}
		*savePath += path.Base(url)
		fSize, err = strconv.ParseInt(res.Header.Get("Content-Length"), 10, 32)
		if err != nil {
			fmt.Println(err)
		}
		f, err := os.Create(*savePath)
		if err != nil {
			fmt.Println(err.Error())
			channel <- err
		}
		defer res.Body.Close()
		_, err = io.Copy(f, res.Body)
		if err != nil {
			fmt.Println(err.Error())
			channel <- err
		} else {
			channel <- nil
		}
		for {
			nr, er := res.Body.Read(buf)
			if nr > 0 {
				nw, ew := f.Write(buf[0:nr])
				if nw > 0 {
					written += int64(nw)
				}
				if ew != nil {
					channel <- ew
					break
				}
				if nr != nw {
					channel <- io.ErrShortWrite
					break
				}
			}
			if er != nil {
				if er != io.EOF {
					channel <- er
					break
				} else {
					close(processChan)
					channel <- nil
					break
				}
			}
			if processChan != nil {
				processChan <- int(written/fSize * 100)
			}
		}
	}
}
