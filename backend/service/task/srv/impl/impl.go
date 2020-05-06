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
	db := g.SqlDb

	rows, err := db.QueryContext(ctx, )
	if err != nil {
		return err
	}
	if rows.Next() {
		err = rows.Scan(&response.Name)
		if err != nil {
			return err
		}
	}

	panic("implement me")
}

func (g Task) GetTaskPeolple(ctx context.Context, requset *pb.GetTaskPeopleRequset, response *pb.GetTaskPeopleResponse) error {
	db := g.SqlDb
	panic("implement me")
}

func (g Task) GetTaskMetaByGroupId(ctx context.Context, request *pb.GetTaskMetaByGroupIdRequest, response *pb.GetTaskMetaByGroupIdResponse) error {
	db := g.SqlDb
	panic("implement me")
}

func (g Task) CreateTask(ctx context.Context, request *pb.CreateTaskRequest, response *pb.CreateTaskResponse) error {
	db := g.SqlDb
	panic("implement me")
}

func (g Task) UpdateTaskMeta(ctx context.Context, request *pb.UpdateTaskMetaRequest, response *pb.UpdateTaskMetaResponse) error {
	db := g.SqlDb
	panic("implement me")
}

func (g Task) DeleteTask(ctx context.Context, request *pb.DeleteTaskRequest, response *pb.DeleteTaskResponse) error {
	db := g.SqlDb
	panic("implement me")
}

func (g Task) UserInTask(ctx context.Context, request *pb.UserInTaskRequest, response *pb.UserInTaskResponse) error {
	db := g.SqlDb
	panic("implement me")
}


