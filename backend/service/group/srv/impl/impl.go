package impl

import (
	"context"
	"database/sql"
	"github.com/sjtu-miniapp/dolphin/service/group/pb"
)

type Group struct {
	SqlDb *sql.DB
}

func (g Group) UserInGroup(ctx context.Context, request *pb.UserInGroupRequest, response *pb.UserInGroupResponse) error {
	db := g.SqlDb
	rows, err := db.QueryContext(ctx, "SELECT COUNT(*) FROM `user_group` WHERE `user_id` = ? and `group_id` = ?", request.UserId, request.GroupId)
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
	panic("implement me")
}

func (g Group) AddUser(ctx context.Context, request *pb.AddUserRequest, response *pb.AddUserResponse) error {
	//db := g.SqlDb
	panic("implement me")
}

func (g Group) UpdateGroup(ctx context.Context, request *pb.UpdateGroupRequest, response *pb.UpdateGroupResponse) error {
	//db := g.SqlDb
	panic("implement me")
}

func (g Group) DeleteGroup(ctx context.Context, request *pb.DeleteGroupRequest, response *pb.DeleteGroupResponse) error {
	//db := g.SqlDb
	panic("implement me")
}

func (g Group) GetGroup(ctx context.Context, request *pb.GetGroupRequest, response *pb.GetGroupResponse) error {
	//db := g.SqlDb
	//rows, err := db.QueryContext(ctx, sql1)
	//if err != nil {
	//	return err
	//}
	//if rows.Next() {
	//	err = rows.Scan(&response.Name)
	//	if err != nil {
	//		return err
	//	}
	//}
	return nil
}
