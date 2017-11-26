package api

import (
	"net/http"
	"jobworker/ctrl"
	"model"
	"github.com/gin-gonic/gin"
)

//新增任务
func (this *ApiServer) newtask(c *gin.Context) {
	var newRequest model.TaskNew
	if err := c.ShouldBindJSON(&newRequest); err == nil {
		response := &model.WorkerResponse{
			Success: false,
			Message: "",
		}
		if newRequest.Id == "" {
			response.Message = "task id is empty"
			c.JSON(http.StatusBadRequest, response)
			return
		} else if newRequest.ZipFileUrl == "" {
			response.Message = "task id is empty"
			c.JSON(http.StatusBadRequest, response)
			return
		}

		this.controller.Actionlist <- ctrl.Action{
			ActionType: 1,
			Id:         newRequest.Id,
			ZipFileUrl: newRequest.ZipFileUrl,
		}

		response.Success = true
		c.JSON(http.StatusOK, response)

	} else {
		c.JSON(http.StatusBadRequest, &model.WorkerResponse{
			Success: false,
			Message: err.Error(),
		})
	}
}

//运行任务
func (this *ApiServer) starttask(c *gin.Context) {
	var newRequest model.TaskNew
	if err := c.ShouldBindJSON(&newRequest); err == nil {
		response := &model.WorkerResponse{
			Success: false,
			Message: "",
		}
		if newRequest.Id == "" {
			response.Message = "task id is empty"
			c.JSON(http.StatusBadRequest, response)
			return
		}

		this.controller.Actionlist <- ctrl.Action{
			ActionType: 2,
			Id:         newRequest.Id,
			ZipFileUrl: newRequest.ZipFileUrl,
		}

		response.Success = true
		c.JSON(http.StatusOK, response)

	} else {
		c.JSON(http.StatusBadRequest, &model.WorkerResponse{
			Success: false,
			Message: err.Error(),
		})
	}
}

//停止任务
func (this *ApiServer) stoptask(c *gin.Context) {
	id := c.Param("id")

	this.controller.Actionlist <- ctrl.Action{
		ActionType: 3,
		Id:         id,
	}

	c.JSON(http.StatusOK, &model.WorkerResponse{
		Success: true,
		Message: "",
	})
}

//删除任务
func (this *ApiServer) deletetask(c *gin.Context) {
	id := c.Param("id")

	this.controller.Actionlist <- ctrl.Action{
		ActionType: 4,
		Id:         id,
	}

	c.JSON(http.StatusOK, &model.WorkerResponse{
		Success: true,
		Message: "",
	})
}
