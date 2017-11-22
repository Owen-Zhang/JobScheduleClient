package api

import "github.com/gin-gonic/gin"
import "model"
import "net/http"
import "fmt"

//新增任务
func (this *ApiServer) newtask(c *gin.Context) {
	fmt.Println("check before")
	var newRequest model.Task_New
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
		fmt.Println("check after")
		flag := this.controller.NewTask(&newRequest)
		response.Success = flag
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
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

//停止任务
func (this *ApiServer) stoptask(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

//删除任务
func (this *ApiServer) deletetask(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}
