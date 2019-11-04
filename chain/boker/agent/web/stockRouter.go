package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Tinachain/Tina/chain/boker/agent/business"
	log4plus "github.com/Tinachain/Tina/chain/boker/common/log4go"
)

//****************Stock Interface
type stocksResponse struct {
	Account []string `json:"accounts"`
	Number  []uint64 `json:"numbers"`
	State   []uint64 `json:"states"`
}

func StockgetStocks(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	log4plus.Info("blockRouter.go StockgetStocks")

	if business.ChainClient == nil {
		log4plus.Error("stockRouter.go StockgetStocks Failed chainclient is nil")
	}

	log4plus.Info("stockRouter.go StockgetStocks chainclient GetStocks")
	stocks, err := business.ChainClient.GetStocks()
	if err != nil {

		log4plus.Error("stockRouter.go StockgetStocks chainclient GetStocks is Failed")
		bytes, _ := json.Marshal(&ResponseCommon{0, ""})
		w.Write(bytes)
	}

	log4plus.Info("stockRouter.go StockgetStocks response GetStocks")
	resp := &stocksResponse{}

	for _, v := range stocks.Stock {

		resp.Account = append(resp.Account, v.Account.String())
		resp.Number = append(resp.Number, v.Number)
		resp.State = append(resp.State, uint64(v.State))
	}

	bytes, _ := json.Marshal(resp)
	w.Write(bytes)
}

type stockRequest struct {
	Address string `json:"address"`
}
type stockResponse struct {
	Account string `json:"account"`
	Number  uint64 `json:"number"`
	State   uint64 `json:"state"`
}

func StockgetStock(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	log4plus.Info("stockRouter.go StockgetStock")

	if business.ChainClient == nil {
		log4plus.Error("stockRouter.go StockgetStock Failed chainclient is nil")
	}

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	req := &stockRequest{}
	err := json.Unmarshal(bodyBytes, req)

	if err != nil {

		log4plus.Error("stockRouter.go StockgetStock", "err", err)
		HttpError(w, -1, err.Error())
		return
	}

	log4plus.Info("stockRouter.go StockgetStock chainclient GetStock hash=%s", req.Address)
	stock, err := business.ChainClient.GetStock(req.Address)
	if err != nil {

		log4plus.Error("stockRouter.go StockgetStock chainclient GetStock is Failed")
		bytes, _ := json.Marshal(&ResponseCommon{0, ""})
		w.Write(bytes)
	}

	resp := &stockResponse{}
	resp.Account = stock.Account.String()
	resp.Number = stock.Number
	resp.State = uint64(stock.State)

	bytes, _ := json.Marshal(resp)
	w.Write(bytes)
}

type ownerResponse struct {
	Owner string `json:"owner"`
}

func StockgetOwner(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	log4plus.Info("stockRouter.go StockgetOwner")

	if business.ChainClient == nil {
		log4plus.Error("stockRouter.go StockgetOwner Failed chainclient is nil")
	}

	log4plus.Info("stockRouter.go StockgetOwner chainclient GetOwner")
	address, err := business.ChainClient.GetOwner()
	if err != nil {

		log4plus.Error("stockRouter.go StockgetOwner chainclient GetOwner is Failed")
		bytes, _ := json.Marshal(&ResponseCommon{0, ""})
		w.Write(bytes)
	}

	log4plus.Info("stockRouter.go StockgetOwner response GetOwner")
	resp := &ownerResponse{}
	resp.Owner = address.String()

	bytes, _ := json.Marshal(resp)
	w.Write(bytes)
}

type gasResponse struct {
	Gas uint64 `json:"gas"`
}

func StockgetGas(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	log4plus.Info("stockRouter.go StockgetGas")

	if business.ChainClient == nil {
		log4plus.Error("stockRouter.go StockgetGas Failed chainclient is nil")
	}

	log4plus.Info("stockRouter.go StockgetGas chainclient GetGas")
	gas, err := business.ChainClient.GetGas()
	if err != nil {

		log4plus.Error("stockRouter.go StockgetGas chainclient GetGas is Failed")
		bytes, _ := json.Marshal(&ResponseCommon{0, ""})
		w.Write(bytes)
	}

	log4plus.Info("stockRouter.go StockgetGas response GetGas")
	resp := &gasResponse{}
	resp.Gas = gas

	bytes, _ := json.Marshal(resp)
	w.Write(bytes)
}
