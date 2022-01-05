package main

import (
	routes "api/router"
	"github.com/appleboy/gofight/v2"
	"github.com/buger/jsonparser"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)
var IniToken string
func init(){
	r := gofight.New()
	r.POST("/api/login").
		SetDebug(true).
		SetJSON(gofight.D{
			"UID": "test123",
		}).
		Run(routes.NewRouter(), func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
			data := []byte(response.Body.String())
			ResultMessage, _ := jsonparser.GetString(data, "ResultMessage")
			IniToken=ResultMessage
		})
}
func TestSignupStatusCreated(t *testing.T) {
	r := gofight.New()
	r.POST("/api/signup").
		SetDebug(true).
		SetJSON(gofight.D{
			"UID": time.Now(),
			"name": "0000",
			"birthday": "2021-11-11",
			"country": "USA",
  			"state": "州名",
   			"gender": "male",
		}).
		Run(routes.NewRouter(), func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
			data := []byte(response.Body.String())
			ResultCode, _ := jsonparser.GetInt(data, "ResultCode")
			ResultMessage, _ := jsonparser.GetString(data, "ResultMessage")

			assert.Equal(t, http.StatusCreated, response.Code)
			assert.Equal(t, int64(0) , ResultCode)
			assert.Equal(t,"success", ResultMessage)
		})
}
func TestSignupBadRequest(t *testing.T) {
	r := gofight.New()
	r.POST("/api/signup").
		SetDebug(true).
		SetJSON(gofight.D{
			"UID": time.Now(),
			"birthday": "2021-11-11",
			"country": "USA",
  			"state": "州名",
   			"gender": "male",
		}).
		Run(routes.NewRouter(), func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
			data := []byte(response.Body.String())
			ResultCode, _ := jsonparser.GetInt(data, "ResultCode")

			assert.Equal(t, http.StatusBadRequest, response.Code)
			assert.Equal(t, int64(2) , ResultCode)
		})
}
func TestLoginStatusCreated(t *testing.T) {
	r := gofight.New()
	r.POST("/api/login").
		SetDebug(true).
		SetJSON(gofight.D{
			"UID": "test123",
		}).
		Run(routes.NewRouter(), func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
			data := []byte(response.Body.String())
			ResultCode, _ := jsonparser.GetInt(data, "ResultCode")
			ResultMessage, _ := jsonparser.GetString(data, "ResultMessage")
		
			assert.Equal(t, http.StatusCreated, response.Code)
			assert.Equal(t, int64(0) , ResultCode)
			assert.NotEmpty(t, ResultMessage)
		})
}
func TestLoginBadRequest(t *testing.T) {
	r := gofight.New()
	r.POST("/api/login").
		SetDebug(true).
		SetJSON(gofight.D{
			"UID": "none",
		}).
		Run(routes.NewRouter(), func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
			data := []byte(response.Body.String())
			ResultCode, _ := jsonparser.GetInt(data, "ResultCode")
			ResultMessage, _ := jsonparser.GetString(data, "ResultMessage")

			assert.Equal(t, http.StatusBadRequest, response.Code)
			assert.Equal(t, int64(5) , ResultCode)
			assert.NotEmpty(t, ResultMessage)
		})
}
func TestFeedbackStatusCreated(t *testing.T){
	r := gofight.New()
	r.POST("/api/feedBack").
		SetDebug(true).
		SetHeader(gofight.H{
			"Authorization": IniToken,
		}).
		SetJSON(gofight.D{
    		"score": 1,
    		"feedback": "好玩喔",
    		"type":"feel",
		}).
		Run(routes.NewRouter(), func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
			data := []byte(response.Body.String())
			ResultCode, _ := jsonparser.GetInt(data, "ResultCode")
			ResultMessage, _ := jsonparser.GetString(data, "ResultMessage")

			assert.Equal(t, http.StatusCreated, response.Code)
			assert.Equal(t, int64(0) , ResultCode)
			assert.Equal(t,"success", ResultMessage)
		})
}
func TestFeedbackBadRequest(t *testing.T){
	r := gofight.New()
	r.POST("/api/feedBack").
		SetDebug(true).
		SetHeader(gofight.H{
			"Authorization": IniToken,
		}).
		SetJSON(gofight.D{
    		"feedback": "好玩喔",
		}).
		Run(routes.NewRouter(), func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
			data := []byte(response.Body.String())
			ResultCode, _ := jsonparser.GetInt(data, "ResultCode")
			ResultMessage, _ := jsonparser.GetString(data, "ResultMessage")

			assert.Equal(t, http.StatusBadRequest, response.Code)
			assert.Equal(t, int64(2) , ResultCode)
			assert.NotEmpty(t, ResultMessage)
		})
}
func TestRecordStatusCreated(t *testing.T){
	r := gofight.New()
	r.POST("/api/record").
		SetDebug(true).
		SetHeader(gofight.H{
			"Authorization": IniToken,
		}).
		SetJSON(gofight.D{
    		"type":"feel",
			"playsec":123,
		}).
		Run(routes.NewRouter(), func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
			data := []byte(response.Body.String())
			ResultCode, _ := jsonparser.GetInt(data, "ResultCode")
			ResultMessage, _ := jsonparser.GetString(data, "ResultMessage")

			assert.Equal(t, http.StatusCreated, response.Code)
			assert.Equal(t, int64(0) , ResultCode)
			assert.Equal(t,"success", ResultMessage)
		})
}
func TestRecordBadRequest(t *testing.T){
	r := gofight.New()
	r.POST("/api/record").
		SetDebug(true).
		SetHeader(gofight.H{
			"Authorization": IniToken,
		}).
		SetJSON(gofight.D{
			"playsec":123,
		}).
		Run(routes.NewRouter(), func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
			data := []byte(response.Body.String())
			ResultCode, _ := jsonparser.GetInt(data, "ResultCode")
			ResultMessage, _ := jsonparser.GetString(data, "ResultMessage")

			assert.Equal(t, http.StatusBadRequest, response.Code)
			assert.Equal(t, int64(2) , ResultCode)
			assert.Equal(t,"Invalid parameters.", ResultMessage)
		})
}
func TestMemberInfoStatusOK(t *testing.T){
	r := gofight.New()
	r.GET("/api/memberInfo").
		SetDebug(true).
		SetHeader(gofight.H{
			"Authorization": IniToken,
		}).
		Run(routes.NewRouter(), func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
			data := []byte(response.Body.String())
			ResultCode, _ := jsonparser.GetInt(data, "ResultCode")
			ResultMessage, _,_,_ := jsonparser.Get(data, "ResultMessage")

			assert.Equal(t, http.StatusOK, response.Code)
			assert.Equal(t, int64(0) , ResultCode)
			assert.NotEmpty(t, string(ResultMessage))
		})
}
func TestMemberInfoBadRequest(t *testing.T){
	r := gofight.New()
	r.GET("/api/memberInfo").
		SetDebug(true).
		SetHeader(gofight.H{
			"Authorization": "none",
		}).
		Run(routes.NewRouter(), func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
			data := []byte(response.Body.String())
			ResultCode, _ := jsonparser.GetInt(data, "ResultCode")

			assert.Equal(t, http.StatusBadRequest, response.Code)
			assert.Equal(t, int64(3) , ResultCode)
		})
}
func TestPlayAnalysisStatusOK(t *testing.T){
	r := gofight.New()
	r.GET("/api/playAnalysis").
		SetDebug(true).
		SetQuery(gofight.H{
			"startdate":"2021-01-01",
			"enddate":"2023-01-01",

		}).
		SetHeader(gofight.H{
			"Authorization": IniToken,
		}).
		Run(routes.NewRouter(), func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
			data := []byte(response.Body.String())
			ResultCode, _ := jsonparser.GetInt(data, "ResultCode")
			ResultMessage, _,_,_ := jsonparser.Get(data, "ResultMessage")

			assert.Equal(t, http.StatusOK, response.Code)
			assert.Equal(t, int64(0) , ResultCode)
			assert.NotEmpty(t, string(ResultMessage))
		})
}
func TestPlayAnalysisBadRequest(t *testing.T){
	r := gofight.New()
	r.GET("/api/playAnalysis").
		SetDebug(true).
		SetQuery(gofight.H{
			"startdate":"2021-01-01",
		}).
		SetHeader(gofight.H{
			"Authorization": IniToken,
		}).
		Run(routes.NewRouter(), func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
			data := []byte(response.Body.String())
			ResultCode, _ := jsonparser.GetInt(data, "ResultCode")

			assert.Equal(t, http.StatusBadRequest, response.Code)
			assert.Equal(t, int64(2) , ResultCode)
		})
}