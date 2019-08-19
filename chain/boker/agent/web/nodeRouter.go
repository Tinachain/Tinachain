package web

import (
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"

	"github.com/Tinachain/Tina/chain/boker/agent/business"
	log4plus "github.com/Tinachain/Tina/chain/boker/common/log4go"
	ethcommon "github.com/Tinachain/Tina/chain/common"
)

//****************Chainware Node Interface

type registerCandidateRequest struct {
	Address     string `json:"address"`
	Description string `json:"description"`
	Team        string `json:"team"`
	Name        string `json:"name"`
	Tickets     int64  `json:"tickets"`
}

func UserregisterCandidate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	log4plus.Info("cmdRouter.go UserregisterCandidate")

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	req := &registerCandidateRequest{}
	err := json.Unmarshal(bodyBytes, req)
	if err != nil {

		log4plus.Error("cmdRouter.go UserregisterCandidate err=%s", err)
		HttpError(w, -1, err.Error())
		return
	}
	log4plus.Debug("cmdRouter.go UserregisterCandidate address=%s Description=%s Team=%s Name=%s Tickets=%d",
		req.Address,
		req.Description,
		req.Team,
		req.Name,
		req.Tickets)

	hash, err := business.ChainClient.RegisterCandidate(req.Description, req.Team, req.Name, new(big.Int).SetInt64(req.Tickets))
	if err != nil {

		bytes, _ := json.Marshal(&ResponseCommon{0, ""})
		w.Write(bytes)
	}
	log4plus.Info("cmdRouter.go UserregisterCandidate RegisterCandidate hash=%s", hash.String())

	resp := &transactionResponse{}
	resp.hash = hash.String()

	bytes, _ := json.Marshal(resp)
	w.Write(bytes)
}

type voteRequest struct {
	AddrVoter     string `json:"addrvoter"`
	AddrCandidate string `json:"addrcandidate"`
	Tokens        int64  `json:"tokens"`
}

/*func Uservote(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	log4plus.Info("cmdRouter.go Uservote")

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	req := &voteRequest{}
	err := json.Unmarshal(bodyBytes, req)
	if err != nil {

		log4plus.Error("cmdRouter.go Uservote err=%s", err)
		HttpError(w, -1, err.Error())
		return
	}
	log4plus.Debug("cmdRouter.go Uservote addrvoter=%s addrcandidate=%s tokens=%d",
		req.AddrVoter,
		req.AddrCandidate,
		req.Tokens)

	hash, err := business.ChainClient.Vote(ethcommon.StringToAddress(req.AddrVoter), ethcommon.StringToAddress(req.AddrCandidate), req.Tokens)
	if err != nil {

		log4plus.Error("cmdRouter.go Uservote Vote err=%s", err)
		bytes, _ := json.Marshal(&ResponseCommon{0, ""})
		w.Write(bytes)
	}
	log4plus.Info("cmdRouter.go Uservote Vote hash=%s", hash.String())

	resp := &transactionResponse{}
	resp.hash = hash.String()

	bytes, _ := json.Marshal(resp)
	w.Write(bytes)
}*/

type cancelAllVotesRequest struct {
	Address string `json:"address"`
}

func UsercancelAllVotes(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	log4plus.Info("cmdRouter.go UserCancelAllVotes")

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	req := &cancelAllVotesRequest{}
	err := json.Unmarshal(bodyBytes, req)
	if err != nil {

		log4plus.Error("cmdRouter.go UserCancelAllVotes err=%s", err)
		HttpError(w, -1, err.Error())
		return
	}
	log4plus.Debug("cmdRouter.go UserCancelAllVotes address=%s", req.Address)

	hash, err := business.ChainClient.CancelAllVotes(ethcommon.StringToAddress(req.Address))
	if err != nil {

		bytes, _ := json.Marshal(&ResponseCommon{0, ""})
		w.Write(bytes)
	}
	resp := &transactionResponse{}
	resp.hash = hash.Str()

	bytes, _ := json.Marshal(resp)
	w.Write(bytes)
}

type candidatesResponse struct {
	Addresses []string `json:"addresses"`
	Tickets   []int64  `json:"tickets"`
}

func UsergetCandidates(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	log4plus.Info("cmdRouter.go UsergetCandidates")

	if business.ChainClient == nil {
		log4plus.Error("cmdRouter.go UsergetCandidates Failed chainclient is nil")
	}

	log4plus.Info("cmdRouter.go UsergetCandidates chainclient GetCandidates")
	candidates, err := business.ChainClient.GetCandidates()
	if err != nil {

		log4plus.Error("cmdRouter.go UsergetCandidates chainclient GetCandidates is Failed")
		bytes, _ := json.Marshal(&ResponseCommon{0, ""})
		w.Write(bytes)
	}

	log4plus.Info("cmdRouter.go UsergetCandidates len(address)=%d len(tickets)=%d", len(candidates.Addresses), len(candidates.Tickets))
	if len(candidates.Addresses) != len(candidates.Tickets) {

		bytes, _ := json.Marshal(&ResponseCommon{0, ""})
		w.Write(bytes)
	}

	log4plus.Info("cmdRouter.go UsergetCandidates response candidates")
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

	log4plus.Info("cmdRouter.go UsergetCandidate")

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	req := &candidateRequest{}
	err := json.Unmarshal(bodyBytes, req)

	if err != nil {

		log4plus.Error("cmdRouter.go UsergetCandidate", "err", err)
		HttpError(w, -1, err.Error())
		return
	}
	log4plus.Debug("cmdRouter.go UsergetCandidate", "address", req.Address)

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
