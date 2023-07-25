package dao

type Group struct {
	Id        int64  `db:"id" json:"id"`
	GroupName string `db:"group_name" json:"group_name"`
	GroupInfo string `db:"group_info" json:"group_info"`
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
