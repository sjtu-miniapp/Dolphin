package impl

import (
	"context"
	"database/sql"
	"github.com/sjtu-miniapp/dolphin/service/group/pb"
)


type Group struct {
	SqlDb *sql.DB
}

func (g Group) GetGroupByUserId(ctx context.Context, request *pb.GetGroupByUserIdRequest, response *pb.GetGroupByUserIdResponse) error {
	panic("implement me")
}

func (g Group) CreateGroup(ctx context.Context, request *pb.CreateGroupRequest, response *pb.CreateGroupResponse) error {
	panic("implement me")
}

func (g Group) AddUser(ctx context.Context, request *pb.AddUserRequest, response *pb.AddUserResponse) error {
	panic("implement me")
}

func (g Group) UpdateGroup(ctx context.Context, request *pb.UpdateGroupRequest, response *pb.UpdateGroupResponse) error {
	panic("implement me")
}

func (g Group) DeleteGroup(ctx context.Context, request *pb.DeleteGroupRequest, response *pb.DeleteGroupResponse) error {
	panic("implement me")
}

func (g Group) GetGroup(ctx context.Context, request *pb.GetGroupRequest, response *pb.GetGroupResponse) error {
	db := g.SqlDb
	rows, err := db.QueryContext(ctx, sql1)
	if err != nil {
		return err
	}
	if rows.Next() {
		err = rows.Scan(&response.Name)
		if err != nil {
			return err
		}
	}
	return nil
}
