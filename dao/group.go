package dao

type Group struct {
	Id        int64  `db:"id" json:"id"`
	GroupName string `db:"group_name" json:"group_name"`
}

func GroupCreate(groupName string) (groupId int64, err error) {
	group := Group{GroupName: groupName}
	result := DBInstance.Create(&group)
	groupId = group.Id
	err = result.Error
	return
}
