package model

type User struct {
	Id string `gorm:"primary_key;auto_increment:false;not null;type:varchar(30)"`
	Name string `gorm:"not null;type:varchar(10)"`
	Gender int `gorm:"default:0"`
	Avatar string `gorm:"type:varchar(100)"`
	SelfGroup Group
	SelfGroupId int32  `gorm:"ForeignKey:id"`// foreignkey? id or Id
	Groups []Group `gorm:"many2many:user_group;ForeignKey:user_id"`
	Tasks []Task `gorm:"many2many:user_task;ForeignKey:user_id"`
}

type Group struct {
	Id int32 `gorm:"primary_key;auto_increment;not null"`
	Name string `gorm:""`
	CreatorId string `gorm:""`
	Type int `gorm:""`
	Tasks []Task `gorm:"ForeignKey:group_id"`
}

type Task struct {
	Id int32 `gorm:"primary_key;not null"`
	GroupId int32 `gorm:""`
	Name string `gorm:""`
	PublisherId string `gorm:""`
	LeaderId string `gorm:""`
	StartDate string `gorm:""`
	EndDate string `gorm:""`
	Readonly bool `gorm:""`
	Type int `gorm:""`
	Desciption string `gorm:""`
	Done bool `gorm:""`
}

type UserGroup struct {
	UserId string `gorm:"primary_key;not null;type:varchar(30)"`
	GroupId int32 `gorm:"primary_key;not null"`
}

type UserTask struct {
	UserId string `gorm:"primary_key;not null;type:varchar(30)"`
	TaskId int32 `gorm:"primary_key;not null"`
}
