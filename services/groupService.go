package services

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"openim/dao"
	"strings"
)

type GroupInfoRequest struct {
	Id string `form:"id" json:"id" binding:"required"` // 长度最少1位
}
type CreateGroupRequest struct {
	GroupName string `form:"group_name" json:"group_name" binding:"required"` // 长度最少1位
}

func GroupInfo(c *gin.Context) {
	var request GroupInfoRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "parameters error"})
		return
	}
	group, err := dao.GroupInfo(request.Id)
	if err != nil {
		log.Println("get group info err:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "get group info error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "group": group})
}

func GroupMembers(c *gin.Context) {

}

func CreateGroup(c *gin.Context) {
	var request CreateGroupRequest
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

func JoinGroup(c *gin.Context) {

}

func QuitGroup(c *gin.Context) {

}
