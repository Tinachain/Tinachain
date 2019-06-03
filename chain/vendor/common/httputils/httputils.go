package httputils

import (
	"bytes"
	"common/utils"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"
)

func HttpRequestForm(r *http.Request, param string) (value string, err error) {
	if len(r.Form[param]) <= 0 {
		return "", fmt.Errorf("param %s not found!", param)
	}
	return strings.TrimSpace(r.Form[param][0]), nil
}

func PathOfUrl(remoteUrl string) string {
	urlInfo, err := url.Parse(remoteUrl)
	if err != nil {
		return ""
	}
	return urlInfo.Path
}

func ClientAddress(req *http.Request) string {
	ip := req.Header.Get("X-Real-IP")
	if ip == "" {
		ip = req.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip = req.Header.Get("Remote_addr")
			if ip == "" {
				ip = strings.Split(req.RemoteAddr, ":")[0]
			}
		}
	}
	return ip
}

func httpClient(timeout int) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Proxy:           utils.GetHttpProxy(),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Dial: func(netw, addr string) (net.Conn, error) {
				to := time.Duration(timeout) * time.Second
				conn, err := net.DialTimeout(netw, addr, to)
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(to))
				return conn, nil
			},
		},
	}
}

func httpGet(remote string, timeout int, headers http.Header, resBody io.Writer) (size int64, err error) {
	request, err := http.NewRequest("GET", remote, nil)
	if err != nil {
		return 0, err
	}
	request.Close = true
	request.Header.Set("Connection", "close") // 完成后断开连接
	if headers != nil {
		for key, header := range headers {
			request.Header.Set(key, header[0])
		}
	}

	response, err := httpClient(timeout).Do(request)
	if err != nil {
		return 0, err
	}
	// 保证I/O正常关闭
	defer response.Body.Close()
	size, err = io.Copy(resBody, response.Body)
	if err != nil {
		return 0, err
	}

	if http.StatusOK != response.StatusCode {
		err = fmt.Errorf("%s", response.Status)
	}

	return size, err
}

func HttpGet(remote string, timeout int, headers http.Header) (string, error) {
	buf := new(bytes.Buffer)
	_, err := httpGet(remote, timeout, headers, buf)
	if err != nil {
		return buf.String(), err
	}
	return buf.String(), nil
}

// download to buffer
// remote 远端文件路径
// timeout 下载超时
// buf
func DownloadBuffer(remote string, timeout int, buf *bytes.Buffer) (int64, error) {
	written, err := httpGet(remote, timeout, nil, buf)
	if err != nil {
		return 0, fmt.Errorf("DownloadBuffer : %s", err.Error())
	}

	return written, nil
}

// download to file
// localFile 本地保存路径
// remote 远端文件路径
// timeout 下载超时
func DownloadFile(remote, localFile string, timeout int) (int64, error) {
	localPath := path.Dir(localFile)
	err := os.MkdirAll(localPath, os.ModePerm)
	if err != nil {
		return 0, err
	}

	file, err := os.Create(localFile)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	written, err := httpGet(remote, timeout, nil, file)
	if err != nil {
		return 0, fmt.Errorf("DownloadFile : %s", err.Error())
	}

	return written, nil
}
