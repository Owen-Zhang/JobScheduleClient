package api

import "github.com/gin-gonic/gin"

//新增任务
func (this *ApiServer) newtask(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

//运行任务
func (this *ApiServer) starttask(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

//停止任务
func (this *ApiServer) stoptask(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

//删除任务
func (this *ApiServer) deletetask(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

//D:\Program\MongoDB\bin\mongod.exe --dbpath D:\Program\MongoDB\data\db --logpath D:\Program\MongoDB\data\log\mongodb.log --service