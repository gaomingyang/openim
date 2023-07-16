package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"openim/dao"
	"strings"
)

type GroupCreateRequest struct {
	GroupName string `form:"group_name" json:"group_name" binding:"required"` // 长度最少1位
}

func GroupCreate(c *gin.Context) {
	var request GroupCreateRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "parameters error"})
		return
	}
	groupName := request.GroupName
	groupName = strings.TrimSpace(groupName)
	if groupName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "group_name can't be empty"})
		return
	}

	groupId, err := dao.GroupCreate(groupName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "create group error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "group_id": groupId})
}
