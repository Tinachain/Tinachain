package web

import (
	"fmt"
	"net/http"
	"time"

	log4plus "github.com/Tinachain/Tina/chain/boker/common/log4go"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type ResponseCommon struct {
	Result int    `json:"result"`
	Msg    string `json:"msg"`
}

type transactionResponse struct {
	hash string `json:"hash"`
}

func HttpError(w http.ResponseWriter, result int, msg string) {
	w.Write([]byte(fmt.Sprintf("{\"result\":%d,\"msg\":\"%s\"}", result, msg)))
}

type Web struct {
	LastCheckManagerAddress time.Time
	ManagerAddress          string
}

func (b *Web) handlerWrap(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		}
		h.ServeHTTP(w, r)
	})
}

func (b *Web) StartWeb(listenString string) {
	go b.startWeb(listenString)
}

func (b *Web) startWeb(listenString string) {

	//用户路由;
	m := mux.NewRouter()

	//****************Node Interface
	// View
	m.Handle("/node/candidates", b.handlerWrap(http.HandlerFunc(UsergetCandidates)))
	m.Handle("/node/candidate", b.handlerWrap(http.HandlerFunc(UsergetCandidate)))

	//****************Block Interface
	m.Handle("/block/blockNumber", b.handlerWrap(http.HandlerFunc(BlockgetBlockNumber)))
	m.Handle("/block/txFromHash", b.handlerWrap(http.HandlerFunc(BlockgetTx)))

	//****************Stock Interface
	m.Handle("/stock/stocks", b.handlerWrap(http.HandlerFunc(StockgetStocks)))
	m.Handle("/stock/stock", b.handlerWrap(http.HandlerFunc(StockgetStock)))
	m.Handle("/stock/stockManager", b.handlerWrap(http.HandlerFunc(StockgetOwner)))

	n := negroni.Classic()
	n.UseHandler(m)
	log4plus.Debug("(b *Web) startWeb listen=%s", listenString)
	n.Run(listenString)
}

func (b *Web) Run(listenString string) {

	go b.StartWeb(listenString)
}

func New() *Web {
	b := &Web{}
	return b
}
