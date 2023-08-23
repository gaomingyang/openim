package dao

type GroupJoinInfo struct {
	Id      int64 `db:"id" json:"id"`
	GroupId int64 `db:"group_id" json:"group_id"`
	UserId  int64 `db:"user_id" json:"user_id"`
}

// 创建申请记录
func ApplyJoinGroup(groupId int64, userId int64) (err error) {
	row := GroupJoinInfo{
		GroupId: groupId,
		UserId:  userId,
	}
	err = DBInstance.Table("group_join_request").Create(&row).Error
	return
}
