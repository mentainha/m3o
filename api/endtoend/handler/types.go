package handler

import (
	alertpb "m3o.dev/api/alerts/proto"
	balancepb "m3o.dev/api/balance/proto"
	custpb "m3o.dev/api/customers/proto"
	pubpb "m3o.dev/api/publicapi/proto"
)

type Endtoend struct {
	custSvc  custpb.CustomersService
	alertSvc alertpb.AlertsService
	balSvc   balancepb.BalanceService
	pubSvc   pubpb.PublicapiService
	email    string
}

type mailinMessage struct {
	Headers  map[string]interface{} `json:"headers"`
	Envelope map[string]interface{} `json:"envelope"`
	Plain    string                 `json:"plain"`
	Html     string                 `json:"html"`
}

type otp struct {
	Token string `json:"token"`
	Time  int64  `json:"time"`
}

type checkResult struct {
	Time   int64  `json:"time"`
	Passed bool   `json:"passed"`
	Error  string `json:"error"`
}

type MailinResponse struct{}

type apiExamples map[string][]apiExample // map of endpoint name to list of examples

type apiExample struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Request     interface{} `json:"request"`
	Response    interface{} `json:"response"`
	RunCheck    bool        `json:"run_check"`
}
