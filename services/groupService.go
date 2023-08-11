package services

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"openim/dao"
	"strings"
)

// get user group List
type GroupListRequest struct {
	UserId string `form:"user_id" json:"user_id" binding:"required"` // 长度最少1位
}

// get group infomation
type GroupInfoRequest struct {
	Id string `form:"id" json:"id" binding:"required"` // 长度最少1位
}

// create a new group
type CreateGroupRequest struct {
	GroupName string `form:"group_name" json:"group_name" binding:"required"` // 长度最少1位
}

// get group members infomation
type GroupMembersRequest struct {
	GroupId int `form:"group_id" json:"group_id" binding:"required"` // 长度最少1位
}

func GroupList(c *gin.Context) {
	var request GroupListRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "parameters error"})
		return
	}
	groups, err := dao.UserGroupList(request.UserId)
	if err != nil {
		log.Println("get group list err:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "get group list error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "groups": groups})
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
	var request GroupMembersRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "parameters error"})
		return
	}
	groupMembers, err := dao.GetGroupMembers(request.GroupId)
	if err != nil {
		log.Println("get group members err:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "get group members error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "group_members": groupMembers})
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
