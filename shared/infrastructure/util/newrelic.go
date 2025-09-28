package util

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"team-maker-api/shared/infrastructure/config"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type HTTPHandler struct {
	Config   config.Config
	Method   string
	URL      string
	ClientID string
	PassKey  string
	Payload  interface{}
}

type NewRelicLogs struct {
	RequestUrl     string `json:"request_url"`
	RequestBody    string `json:"request_body"`
	ResponseStatus string `json:"response_status"`
	ResponseBody   string `json:"response_body"`
}

type BodyDump struct {
	RequestBody  interface{} `json:"request_body"`
	ResponseBody RespBody    `json:"response_body"`
	Url          string      `json:"url"`
}

type RespBody struct {
	Code        string      `json:"errorCode"`
	CodeMessage string      `json:"traceId"`
	Data        interface{} `json:"data"`
	Message     string      `json:"errorMessage"`
}

type DispatchRespBody struct {
	Code        int         `json:"code"`
	CodeMessage string      `json:"code_message"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
}

var CODE_SUCCESS_RESPONSE = []string{"200", "201"}

// BACKUP REAL USING GINCONTEXT
// func SendOutgoingLogToNewRelic(c *gin.Context, config config.Config, url string, vendorName string, taskCode string, reqBody interface{}, resBody *resty.Response, severity string, app *newrelic.Application) {
func SendOutgoingLogToNewRelic(c *context.Context, config config.Config, url string, vendorName string, taskCode string, reqBody interface{}, resBody DispatchRespBody, severity string, app *newrelic.Application) {
	type RequestBodyLogs struct {
		VendorName     string `json:"vendor_name"`
		TaskCode       string `json:"task_code"`
		PayloadRequest string `json:"payload_request"`
		Purpose        string `json:"purpose"`
	}

	payloadRequest, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	if config.NewRelic.Enable && config.NewRelic.EnableOutgoingLogging {
		payloadRequest := RequestBodyLogs{
			VendorName:     vendorName,
			TaskCode:       taskCode,
			PayloadRequest: "[" + string(payloadRequest) + "]",
			Purpose:        "dispatch-outgoing-logging",
		}

		requestBody, err := json.Marshal(payloadRequest)
		if err != nil {
			fmt.Println("unable to marshal newrelic outgoing request body", err)
			return
		}

		// ignore send logs to new relic if response success and enable_logging_on_success_outgoing is false
		if !config.NewRelic.EnableLoggingOnSuccessOutgoing && StringContains(CODE_SUCCESS_RESPONSE, strconv.Itoa(resBody.Code)) {
			return
		}

		respBody, err := json.Marshal(resBody)
		if err != nil {
			fmt.Println(err)
			return
		}

		newRelicLogs := NewRelicLogs{
			RequestUrl:     url,
			RequestBody:    string(requestBody),
			ResponseStatus: strconv.Itoa(resBody.Code),
			ResponseBody:   string(respBody),
		}

		newRelicLogsJson, err := json.Marshal(newRelicLogs)
		if err != nil {
			fmt.Println("marshal error newrelic outgoing logs >>> " + err.Error())
			return
		}

		loc, _ := time.LoadLocation("Asia/Jakarta")
		now := time.Now().In(loc)
		app.RecordLog(newrelic.LogData{
			Timestamp: now.UnixNano() / int64(time.Millisecond),
			Message:   string(newRelicLogsJson),
			Severity:  severity,
		})
	}
}

func SendIncomingLogToNewRelic(config config.Config, url string, reqBody interface{}, resBody RespBody, responseCodeStr string, severity string, app *newrelic.Application) {
	type RequestBodyLogs struct {
		PayloadRequest string `json:"payload_request"`
		Purpose        string `json:"purpose"`
	}

	payloadRequest, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	respBody, err := json.Marshal(resBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	if config.NewRelic.Enable && config.NewRelic.EnableIncomingLogging {
		payloadRequest := RequestBodyLogs{
			PayloadRequest: "[" + string(payloadRequest) + "]",
			Purpose:        "incoming-logging",
		}

		requestBody, err := json.Marshal(payloadRequest)
		if err != nil {
			fmt.Println("unable to marshal newrelic incoming request body", err)
			return
		}

		// ADD COMMENTS
		// ignore send logs to new relic if response success and enable_logging_on_success_incoming is false
		if !config.NewRelic.EnableLoggingOnSuccessIncoming && responseCodeStr == "" {
			return
		}

		newRelicLogs := NewRelicLogs{
			RequestUrl:     url,
			RequestBody:    string(requestBody),
			ResponseStatus: resBody.Code,
			ResponseBody:   string(respBody),
		}

		newRelicLogsJson, err := json.Marshal(newRelicLogs)
		if err != nil {
			fmt.Println("marshal error newrelic incoming logs >>> " + err.Error())
			return
		}

		fmt.Println("send logs to newrelic incoming logs")
		loc, _ := time.LoadLocation("Asia/Jakarta")
		now := time.Now().In(loc)
		app.RecordLog(newrelic.LogData{
			Timestamp: now.UnixNano() / int64(time.Millisecond),
			Message:   string(newRelicLogsJson),
			Severity:  severity,
		})
	}
}
