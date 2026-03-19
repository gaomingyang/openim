package dao

type GroupMember struct {
	Id      int64  `db:"id" json:"id"`
	GroupId int64  `db:"group_id" json:"group_id"`
	UserId  int64  `db:"user_id" json:"user_id"`
	Role    string `db:"role" json:"role"`
}

func GetGroupMembers(groupId int) (groupMembers []GroupMember, err error) {
	err = DBInstance.Table("group_members").Where("group_id=?", groupId).Find(&groupMembers).Error
	return
}

// 查询用户是否在某个组里
func CheckExistGroupMember(groupID, userID int64) (exist bool, err error) {
	var groupMembers []GroupMember
	err = DBInstance.Table("group_members").Where("group_id=? and user_id = ?", groupID, userID).Find(&groupMembers).Error
	if len(groupMembers) > 0 {
		exist = true
	}
	return
}
