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
	Number uint64 `json:"number"`
	Timer  uint64 `json:"timer"`
}

func StockgetStocks(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	log4plus.Info("blockRouter.go BlockgetBlockNumber")

	if business.ChainClient == nil {
		log4plus.Error("blockRouter.go BlockgetBlockNumber Failed chainclient is nil")
	}

	log4plus.Info("blockRouter.go BlockgetBlockNumber chainclient GetBlockNumber")
	block, err := business.ChainClient.GetBlockNumber()
	if err != nil {

		log4plus.Error("blockRouter.go BlockgetBlockNumber chainclient GetBlockNumber is Failed")
		bytes, _ := json.Marshal(&ResponseCommon{0, ""})
		w.Write(bytes)
	}

	log4plus.Info("blockRouter.go BlockgetBlockNumber response GetBlockNumber")
	resp := &blockNumberResponse{}
	resp.Number = block.Number
	resp.Timer = block.Timer

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
	block, err := business.ChainClient.GetOwner()
	if err != nil {

		log4plus.Error("stockRouter.go StockgetOwner chainclient GetOwner is Failed")
		bytes, _ := json.Marshal(&ResponseCommon{0, ""})
		w.Write(bytes)
	}

	log4plus.Info("stockRouter.go StockgetOwner response GetOwner")
	resp := &ownerResponse{}
	resp.Owner = block.Number
	resp.Timer = block.Timer

	bytes, _ := json.Marshal(resp)
	w.Write(bytes)
}
