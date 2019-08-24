package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Tinachain/Tina/chain/boker/agent/business"
	log4plus "github.com/Tinachain/Tina/chain/boker/common/log4go"
	ethcommon "github.com/Tinachain/Tina/chain/common"
)

//****************Node Interface
type candidatesResponse struct {
	Addresses []string `json:"addresses"`
	Tickets   []int64  `json:"tickets"`
}

func UsergetCandidates(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	log4plus.Info("nodeRouter.go UsergetCandidates")

	if business.ChainClient == nil {
		log4plus.Error("nodeRouter.go UsergetCandidates Failed chainclient is nil")
	}

	log4plus.Info("nodeRouter.go UsergetCandidates chainclient GetCandidates")
	candidates, err := business.ChainClient.GetCandidates()
	if err != nil {

		log4plus.Error("nodeRouter.go UsergetCandidates chainclient GetCandidates is Failed")
		bytes, _ := json.Marshal(&ResponseCommon{0, ""})
		w.Write(bytes)
	}

	log4plus.Info("nodeRouter.go UsergetCandidates len(address)=%d len(tickets)=%d", len(candidates.Addresses), len(candidates.Tickets))
	if len(candidates.Addresses) != len(candidates.Tickets) {

		bytes, _ := json.Marshal(&ResponseCommon{0, ""})
		w.Write(bytes)
	}

	log4plus.Info("nodeRouter.go UsergetCandidates response candidates")
	resp := &candidatesResponse{}
	for i, v := range candidates.Addresses {

		resp.Addresses = append(resp.Addresses, v.String())
		resp.Tickets = append(resp.Tickets, candidates.Tickets[i].Int64())
	}

	bytes, _ := json.Marshal(resp)
	w.Write(bytes)

}

type candidateRequest struct {
	Address string `json:"addrCandidate"`
}

type candidateResponse struct {
	Description string `json:"description"`
	Team        string `json:"team"`
	Name        string `json:"name"`
	Tickets     int64  `json:"tickets"`
}

func UsergetCandidate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	log4plus.Info("nodeRouter.go UsergetCandidate")

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	req := &candidateRequest{}
	err := json.Unmarshal(bodyBytes, req)

	if err != nil {

		log4plus.Error("nodeRouter.go UsergetCandidate", "err", err)
		HttpError(w, -1, err.Error())
		return
	}
	log4plus.Debug("nodeRouter.go UsergetCandidate", "address", req.Address)

	candidate, err := business.ChainClient.GetCandidate(ethcommon.StringToAddress(req.Address))
	if err != nil {

		bytes, _ := json.Marshal(&ResponseCommon{0, ""})
		w.Write(bytes)
	}

	resp := &candidateResponse{}

	resp.Description = candidate.Description
	resp.Team = candidate.Team
	resp.Name = candidate.Name
	resp.Tickets = candidate.Tickets.Int64()

	bytes, _ := json.Marshal(resp)
	w.Write(bytes)
}
