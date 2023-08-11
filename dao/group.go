package dao

type GroupId struct {
	Id int64 `db:"id" json:"id"`
}
type Group struct {
	Id        int64  `db:"id" json:"id"`
	GroupName string `db:"group_name" json:"group_name"`
	GroupInfo string `db:"group_info" json:"group_info"`
}

// UserGroupList : one user's groups list
func UserGroupList(userId string) (groups []Group, err error) {
	var groupIds []int
	groupIdResult := DBInstance.Table("user_groups").Select("group_id").Where("user_id=?", userId).Find(&groupIds)
	err = groupIdResult.Error
	// log.Printf("groupIds:%+v\n", groupIds)
	if err != nil {
		return
	}
	groupsResult := DBInstance.Table("groups").Where("id in ?", groupIds).Find(&groups)
	err = groupsResult.Error
	return
}

func GroupInfo(id string) (group Group, err error) {
	err = DBInstance.Table("groups").First(&group, id).Error
	return
}

func GroupCreate(groupName string) (groupId int64, err error) {
	group := Group{GroupName: groupName}
	result := DBInstance.Create(&group)
	groupId = group.Id
	err = result.Error
	return
}
