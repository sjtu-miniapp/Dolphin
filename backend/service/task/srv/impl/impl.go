package impl

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sjtu-miniapp/dolphin/service/task/pb"
	"log"
)


type Task struct {
	SqlDb *sql.DB
}

func (g Task) GetTaskMeta(ctx context.Context, request *pb.GetTaskMetaRequest, response *pb.GetTaskMetaResponse) error {
	panic("implement me")
}

func (g Task) GetTaskPeolple(ctx context.Context, requset *pb.GetTaskPeopleRequset, response *pb.GetTaskPeopleResponse) error {
	panic("implement me")
}

func (g Task) GetTaskMetaByGroupId(ctx context.Context, request *pb.GetTaskMetaByGroupIdRequest, response *pb.GetTaskMetaByGroupIdResponse) error {
	panic("implement me")
}

func (g Task) CreateTask(ctx context.Context, request *pb.CreateTaskRequest, response *pb.CreateTaskResponse) error {
	panic("implement me")
}

func (g Task) UpdateTaskMeta(ctx context.Context, request *pb.UpdateTaskMetaRequest, response *pb.UpdateTaskMetaResponse) error {
	panic("implement me")
}

func (g Task) DeleteTask(ctx context.Context, request *pb.DeleteTaskRequest, response *pb.DeleteTaskResponse) error {
	panic("implement me")
}

func (g Task) UserInTask(ctx context.Context, request *pb.UserInTaskRequest, response *pb.UserInTaskResponse) error {
	panic("implement me")
}

func (g Task) GetGroup(ctx context.Context, request *pb.GetGroupRequest, response *pb.GetGroupResponse) error {
	db := g.SqlDb

	sql1 := fmt.Sprintf("SELECT name FROM `task` WHERE `id` = %d", request.Id)
	log.Println(sql1)
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
