package impl

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sjtu-miniapp/dolphin/service/database/model"
	"github.com/sjtu-miniapp/dolphin/service/group/pb"
)

type Group struct {
	SqlDb *gorm.DB
}

func (g Group) UserInGroup(ctx context.Context, request *pb.UserInGroupRequest, response *pb.UserInGroupResponse) error {
	if request.GroupId == nil || request.UserId == nil {
		return fmt.Errorf("nil pointer")
	}
	db := g.SqlDb
	err := db.Find(&model.UserGroup{
		UserId:  *request.UserId,
		GroupId: *request.GroupId,
	}).Error
	response.Ok = new(bool)
	if err == nil {
		*response.Ok = true
		return nil
	} else if err.Error() == "record not found" {
		*response.Ok = false
		return nil
	} else {
		*response.Ok = false
		return err
	}
}

func (g Group) GetGroupByUserId(ctx context.Context, request *pb.GetGroupByUserIdRequest, response *pb.GetGroupByUserIdResponse) error {
	if request.UserId == nil {
		return fmt.Errorf("nil pointer")
	}
	db := g.SqlDb
	user := model.User{
		Id:          *request.UserId,
	}
	if err := db.Preload("Groups").Find(&user).Error; err != nil {
		if err.Error() != "record not found" {
			return err
		}
	}
	for _, v := range user.Groups {
		group := &pb.GetGroupByUserIdResponse_Group{
			Id:                   &v.Id,
			Name:                 v.Name,
			CreatorId:            &v.CreatorId,
		}
		response.Groups = append(response.Groups, group)
	}
	return nil
}

func (g Group) CreateGroup(ctx context.Context, request *pb.CreateGroupRequest, response *pb.CreateGroupResponse) error {
	if request.CreatorId == nil || request.Name == nil || request.Type == nil {
		return fmt.Errorf("nil pointer")
	}
	db := g.SqlDb
	if *request.Type == 1 {
		return fmt.Errorf("don't create inidvidual group")
	}
	group := &model.Group{
		Name:      request.Name,
		CreatorId: *request.CreatorId,
		Type:      *request.Type,
	}
	if err := db.Create(&group).Error; err != nil {
		return err
	}
	if group.Id == 0 {
		return fmt.Errorf("created group id == 0")
	}
	response.Id = &group.Id

	err := g.AddUser(ctx, &pb.AddUserRequest{
		GroupId:              response.Id,
		UserIds:              []string{*request.CreatorId},
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

func (g Group) AddUser(ctx context.Context, request *pb.AddUserRequest, response *pb.AddUserResponse) error {
	if request.GroupId == nil {
		return fmt.Errorf("nil pointer")
	}
	db := g.SqlDb

	if len(request.UserIds) == 0 {
		return nil
	}
	pg := new(pb.GetGroupResponse)
	err := g.GetGroup(ctx, &pb.GetGroupRequest{
		Id:                   request.GroupId,
	}, pg)
	if err != nil {
		// record not found
		return err
	}
	if *pg.Type == 1 {
		return fmt.Errorf("can't add users to an individual group")
	}
	tx := db.Begin()
	for _, v := range request.UserIds {
		err = tx.Create(&model.UserGroup{
			UserId:  v,
			GroupId: *request.GroupId,
		}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

// all values are prepared in api
func (g Group) UpdateGroup(ctx context.Context, request *pb.UpdateGroupRequest, response *pb.UpdateGroupResponse) error {
	if request.Id == nil || request.Name == nil{
		return fmt.Errorf("nil pointer")
	}
	db := g.SqlDb
	err := db.Model(&model.Group{
		Id:        *request.Id,
	}).Update("name", *request.Name).Error
	return err
}

func (g Group) DeleteGroup(ctx context.Context, request *pb.DeleteGroupRequest, response *pb.DeleteGroupResponse) error {
	if request.Id == nil {
		return fmt.Errorf("nil pointer")
	}
	db := g.SqlDb
	pg := new(pb.GetGroupResponse)
	err := g.GetGroup(ctx, &pb.GetGroupRequest{
		Id:                   request.Id,
	}, pg)
	if err != nil {
		return err
	}
	if *pg.Type == 1 {
		return fmt.Errorf("can't delete an individual group")
	}
	err = db.Delete(&model.Group{
		Id:        *request.Id,
	}).Error
	if err != nil {
		return err
	}
	return err
}

func (g Group) GetGroup(ctx context.Context, request *pb.GetGroupRequest, response *pb.GetGroupResponse) error {
	if request.Id == nil {
		return fmt.Errorf("nil pointer")
	}
	db := g.SqlDb
	group := model.Group{
		Id:        *request.Id,
	}
	if err := db.Preload("Users").Find(&group).Error; err != nil {
		return err
	}
	response.Name = group.Name
	response.Type = &group.Type
	response.CreatorId = &group.CreatorId

	for _, v := range group.Users {
		user := &pb.User{
			Id:                   &v.Id,
			Name:                 &v.Name,
			Avatar:               v.Avatar,
			Gender:               v.Gender,
			SelfGroupId:          v.SelfGroupId,
		}
		response.Users = append(response.Users, user)
	}
	return nil
}
