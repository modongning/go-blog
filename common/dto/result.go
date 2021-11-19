package dto

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ResultData struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Timestamp int64       `json:"timestamp"`
	TraceId   string      `json:"traceId"`
	Data      interface{} `json:"data"`
}

type Result struct {
	Ctx *gin.Context
}

func NewResult(ctx *gin.Context) *Result {
	return &Result{Ctx: ctx}
}

func (r *Result) Success(data interface{}) {
	if nil == data {
		data = gin.H{}
	}
	res := ResultData{}
	res.Msg = "success"
	res.Code = 200
	res.Timestamp = time.Now().UnixNano()
	res.TraceId = ""
	res.Data = data

	r.Ctx.JSON(http.StatusOK, res)
}

func (r *Result) Fail(msg string) {
	r.FailByCode(40000, msg)
}

func (r *Result) FailByCode(code int, msg string) {
	res := ResultData{}
	res.Msg = msg
	res.Code = code
	res.Timestamp = time.Now().UnixNano()
	res.TraceId = ""
	res.Data = nil

	r.Ctx.JSON(http.StatusOK, res)
}
