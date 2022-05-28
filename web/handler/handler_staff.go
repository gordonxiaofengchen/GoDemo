package handler

import (
	"demo/biz"
	"demo/dal"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type staffHandler struct{

}

var StaffHandler staffHandler

func InitStaffHandler(router *gin.Engine){
	staffRouter := router.Group("/staff")
	{
		staffRouter.GET("", StaffHandler.QueryAll)
		staffRouter.GET("/:id", StaffHandler.GetById)
		/*staffRouter.POST("/")
		staffRouter.PUT("/")
		staffRouter.PATCH("/")
		staffRouter.DELETE("/")*/
	}
}

func (staffHandler) QueryAll(c *gin.Context){
	p, err := GetPage(c)
	if err != nil{
		c.JSON(http.StatusOK, InternalError(err))
	}
	ok, err := biz.StaffBiz.QueryAll(&p)
	r := CommonResult(p, ok, err)
	c.JSON(http.StatusOK, r)
}

func (staffHandler) GetById(c *gin.Context){
	id := c.Param("id")
	staffs, ok, err := biz.StaffBiz.QueryById(id)
	r := CommonResult(staffs, ok, err)
	c.JSON(http.StatusOK, r)
}