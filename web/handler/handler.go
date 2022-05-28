package handler

import (
	"demo/dal"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ApiResult struct{
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const(
	CodeNormalSuccess = iota + 100
	CodeNormalFail
	CodeInternalError
)

var(
	EmptyPage dal.Page
)

func InitMiddleWare(router *gin.Engine){
	router.Use(optionsResponse, setCORSHeader)
}

func setCORSHeader(c *gin.Context){
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
}

func optionsResponse(c *gin.Context){
	if c.Request.Method == http.MethodOptions {
		c.Status(http.StatusOK)
	}else{
		c.Next()
	}
}

func GetPage(c *gin.Context) (dal.Page, error) {
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		return EmptyPage, err
	}
	pageNum, err := strconv.Atoi(c.Query("pageNum"))
	if err != nil {
		return EmptyPage, err
	}
	p := dal.Page{
		PageSize: pageSize,
		PageNum:  pageNum,
	}
	return p, nil
}

func Success(data interface{}) ApiResult{
	return ApiResult{
		Success: true,
		Status:  CodeNormalSuccess,
		Data:    data,
	}
}

func Fail() ApiResult{
	return ApiResult{
		Success: false,
		Status:  CodeNormalFail,
	}
}

func InternalError(err error) ApiResult{
	return ApiResult{
		Success: false,
		Status:  CodeInternalError,
		Message: err.Error(),
	}
}

func CommonResult(data interface{}, ok bool, err error) ApiResult {
	if err != nil{
		return InternalError(err)
	}
	if !ok {
		return Fail()
	}
	return Success(data)
}