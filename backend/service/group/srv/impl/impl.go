package impl

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/micro/go-micro/util/log"
	"github.com/sjtu-miniapp/dolphin/service/group/pb"
)

type Group struct {
	SqlDb *sql.DB
}

func (g Group) UserInGroup(ctx context.Context, request *pb.UserInGroupRequest, response *pb.UserInGroupResponse) error {
	db := g.SqlDb
	rows, err := db.QueryContext(ctx, "SELECT * FROM `user_group` WHERE `user_id` = ? and `group_id` = ?", request.UserId, request.GroupId)
	if err != nil {
		return err
	}
	response.Ok = rows.Next()
	return nil
}

func (g Group) GetGroupByUserId(ctx context.Context, request *pb.GetGroupByUserIdRequest, response *pb.GetGroupByUserIdResponse) error {
	db := g.SqlDb
	rows, err := db.QueryContext(ctx, "SELECT `id`, `creator_id`, `name` FROM `user_group` JOIN `group` ON `group_id`=`id` WHERE `user_id`= ?", request.UserId)
	if err != nil {
		return err
	}

	for rows.Next() {
		r := new(pb.GetGroupByUserIdResponse_Group)
		err = rows.Scan(&r.Id, &r.CreatorId, &r.Name)
		if err != nil {
			return err
		}
		response.Groups = append(response.Groups, r)
	}
	return nil
}

func (g Group) CreateGroup(ctx context.Context, request *pb.CreateGroupRequest, response *pb.CreateGroupResponse) error {
	db := g.SqlDb
	result, err := db.ExecContext(ctx, "INSERT INTO `group`(`name`,`creator_id`,`type`) VALUES(?, ?, ?)", request.Name, request.CreatorId, request.Type)
	if err != nil {
		return err
	}
	if id, err := result.LastInsertId(); err != nil {
		return err
	} else {
		response.Id = uint32(id)
	}
	err = g.AddUser(ctx, &pb.AddUserRequest{
		GroupId:              response.Id,
		UserIds:              []string{request.CreatorId},
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

func (g Group) AddUser(ctx context.Context, request *pb.AddUserRequest, response *pb.AddUserResponse) error {
	db := g.SqlDb

	if len(request.UserIds) == 0 {
		return nil
	}
	pg := new(pb.GetGroupResponse)
	err := g.GetGroup(ctx, &pb.GetGroupRequest{
		Id:                   request.GroupId,
	}, pg)
	if err != nil {
		return err
	}
	if pg.Type == 1 {
		return fmt.Errorf("can't add users to an individual group")
	}
	sql1 := fmt.Sprintf("INSERT INTO `user_group`(`user_id`, `group_id`) VALUES('%s', %d)", request.UserIds[0], request.GroupId)
	log.Info(sql1)
	for _, v := range request.UserIds[1:] {
		sql1 = sql1 + fmt.Sprintf(", ('%s', %d)", v, request.GroupId)
	}
	_, err = db.ExecContext(ctx, sql1)
	log.Info(sql1)
	return err
}

// all values are prepared in api
func (g Group) UpdateGroup(ctx context.Context, request *pb.UpdateGroupRequest, response *pb.UpdateGroupResponse) error {
	db := g.SqlDb
	_, err := db.ExecContext(ctx, "UPDATE `group` SET `name`=? WHERE `id`=?", request.Name, request.Id)
	return err
}

func (g Group) DeleteGroup(ctx context.Context, request *pb.DeleteGroupRequest, response *pb.DeleteGroupResponse) error {
	db := g.SqlDb
	pg := new(pb.GetGroupResponse)
	err := g.GetGroup(ctx, &pb.GetGroupRequest{
		Id:                   request.Id,
	}, pg)
	if err != nil {
		return err
	}
	if pg.Type == 1 {
		return fmt.Errorf("can't delete an individual group")
	}
	_, err = db.ExecContext(ctx, "DELETE FROM `group` WHERE `id`=?",
		request.Id)
	return err
}

func (g Group) GetGroup(ctx context.Context, request *pb.GetGroupRequest, response *pb.GetGroupResponse) error {
	db := g.SqlDb
	errChan := make(chan error)
	go func() {
		rows, err := db.QueryContext(ctx, "SELECT `name`, `type`, `creator_id` FROM `group` WHERE `id`=?", request.Id)
		if err != nil {
			errChan <- err
			return
		}
		if rows.Next() {
			err = rows.Scan(&response.Name, &response.Type, &response.CreatorId)
			if err != nil {
				errChan <- err
				return
			}
		}
		errChan <- err
	}()
	go func() {
		rows, err := db.QueryContext(ctx, "SELECT `user`.`id`, `user`.`name` FROM `user_group` JOIN `group` JOIN `user` ON `group_id` = `group`.`id` AND `user_id` = `user`.id WHERE `group_id` = ?", request.Id)
		if err != nil {
			errChan <- err
			return
		}
		user := new(pb.User)
		for rows.Next() {
			err = rows.Scan(&user.Id, &user.Name)
			if err != nil {
				errChan <- err
				return
			}
			response.Users = append(response.Users, user)
		}
		errChan <- err
	}()
	if <-errChan != nil {
		return <-errChan
	}
	if <-errChan != nil {
		return <-errChan
	}
	return nil
}
