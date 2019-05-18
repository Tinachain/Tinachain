// support project web.go
package web

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	log4plus "github.com/alecthomas/log4go"
)

type Web struct {
	Ip   string
	Port int
}

func New() *Web {

	downloadPtr := new(Download)
	downloadPtr.loadConfig()
	return downloadPtr
}

func (w *Web) httpDomain() {

	if origin := req.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
}

func (w *Web) getBlock(w http.ResponseWriter, req *http.Request) error {

	w.httpDomain()

	req.ParseForm()
	_, cpid_found := req.Form["cpid"]
	if !(cpid_found) {

		urlerror := getUrlError(PARAM_LACK, "参数不完整，请重新输入")
		fmt.Fprint(w, urlerror)
		log4plus.Info("参数不完整，请重新输入")
		return
	}

	return nil
}

func (w *Web) getTransition(w http.ResponseWriter, req *http.Request) error {

	w.httpDomain()

	req.ParseForm()
	_, cpid_found := req.Form["cpid"]
	if !(cpid_found) {

		urlerror := getUrlError(PARAM_LACK, "参数不完整，请重新输入")
		fmt.Fprint(w, urlerror)
		log4plus.Info("参数不完整，请重新输入")
		return
	}

	return nil
}

func (w *Web) getNumber(w http.ResponseWriter, req *http.Request) error {

	w.httpDomain()

	req.ParseForm()
	_, cpid_found := req.Form["cpid"]
	if !(cpid_found) {

		urlerror := getUrlError(PARAM_LACK, "参数不完整，请重新输入")
		fmt.Fprint(w, urlerror)
		log4plus.Info("参数不完整，请重新输入")
		return
	}

	return nil
}

func (w *Web) getTransitionCount(w http.ResponseWriter, req *http.Request) error {

	w.httpDomain()

	req.ParseForm()
	_, cpid_found := req.Form["cpid"]
	if !(cpid_found) {

		urlerror := getUrlError(PARAM_LACK, "参数不完整，请重新输入")
		fmt.Fprint(w, urlerror)
		log4plus.Info("参数不完整，请重新输入")
		return
	}

	return nil
}

func (w *Web) getWord(w http.ResponseWriter, req *http.Request) error {

	w.httpDomain()

	req.ParseForm()
	_, cpid_found := req.Form["cpid"]
	if !(cpid_found) {

		urlerror := getUrlError(PARAM_LACK, "参数不完整，请重新输入")
		fmt.Fprint(w, urlerror)
		log4plus.Info("参数不完整，请重新输入")
		return
	}

	return nil
}

func (w *Web) getPicture(w http.ResponseWriter, req *http.Request) error {

	w.httpDomain()

	req.ParseForm()
	_, cpid_found := req.Form["cpid"]
	if !(cpid_found) {

		urlerror := getUrlError(PARAM_LACK, "参数不完整，请重新输入")
		fmt.Fprint(w, urlerror)
		log4plus.Info("参数不完整，请重新输入")
		return
	}

	return nil
}

func (w *Web) goHttp() error {

	//得到区块信息
	http.HandleFunc("/http/getBlock", w.getBlock)
	//得到交易信息
	http.HandleFunc("/http/getTransition", w.getTransition)
	//得到区块数量
	http.HandleFunc("/http/getNumber", w.getNumber)
	//得到交易数量
	http.HandleFunc("/http/getTransitionCount", w.getTransitionCount)
	//得到文字内容
	http.HandleFunc("/http/getWord", w.getWord)
	//得到图片内容
	http.HandleFunc("/http/getPicture", w.getPicture)

	//服务器要监听的主机地址和端口号
	listenString := w.Ip + ":" + w.Port
	err := http.ListenAndServe(listenString, nil)
	if err != nil {
		log4plus.Info("ListenAndServe error: %s", err.Error())
		return errors.New("Listen And Start Server Failed")
	}
	return nil
}
