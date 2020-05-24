package model

import "time"

type User struct {
	Id string `gorm:"primary_key;auto_increment:false;not null;type:varchar(30)"`
	Name string `gorm:"not null;type:varchar(20)"`
	Gender *int32
	Avatar *string `gorm:"type:varchar(156)"`
	// associated
	SelfGroup *Group
	SelfGroupId *int32  `gorm:"foreign_key:id;default null"`
	// associated
	Groups []*Group `gorm:"many2many:user_group;foreign_key:user_id"`
	// associated
	Tasks []*Task `gorm:"many2many:user_task;foreign_key:user_id"`
}

type Group struct {
	Id int32 `gorm:"primary_key;auto_increment;not null"`
	Name *string `gorm:"default:'';type:varchar(20);unique_index:idx_name_creatorid"`
	// associated
	CreatorId string `gorm:"foreign_key:id;unique_index:idx_name_creatorid;not null;type:varchar(30)"`
	Type int32 `gorm:"default:0;not null"`
	Users []*User `gorm:"many2many:user_group;foreign_key:group_id"`
 	Tasks []*Task `gorm:"foreign_key:group_id"`
}

type Task struct {
	Id int32 `gorm:"primary_key;auto_increment;not null"`
	GroupId int32 `gorm:"foreign_key:id;not null;unique_index:idx_name_groupid"`
	Name *string `gorm:"type:varchar(20);default:'';unique_index:idx_name_groupid"`
	PublisherId string `gorm:"foreign_key:id;not null;type:varchar(30)"`
	LeaderId *string `gorm:"foreign_key:id;type:varchar(30)"`
	StartDate  *time.Time
	EndDate *time.Time
	Readonly bool `gorm:"default:false;not null"`
	Type int32 `gorm:"default:0;not null"`
	Desciption *string `gorm:"default:'';type:varchar(255)"`
	Done bool `gorm:"default:false;not null"`
}

type UserGroup struct {
	UserId string `gorm:"primary_key;not null;type:varchar(30)"`
	GroupId int32 `gorm:"primary_key;not null"`
}

type UserTask struct {
	UserId string `gorm:"primary_key;not null;type:varchar(30)"`
	TaskId int32 `gorm:"primary_key;not null"`
	Done bool `gorm:"default:false;not null"`
	DoneTime *time.Time
}
