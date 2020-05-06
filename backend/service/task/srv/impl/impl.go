package impl

import (
	"context"
	"database/sql"
	"github.com/sjtu-miniapp/dolphin/service/task/pb"
)


type Task struct {
	SqlDb *sql.DB
}


func (g Task) GetTaskMeta(ctx context.Context, request *pb.GetTaskMetaRequest, response *pb.GetTaskMetaResponse) error {
	db := g.SqlDb

	rows, err := db.QueryContext(ctx, "SELECT `name`, `type`, `done`, `group_id`, `publisher_id`, `leader_id`, `start_date`, `end_date`, `readonly`, `description` FROM `task` WHERE `id`=?", request.Id)
	if err != nil {
		return err
	}
	var sd sql.NullString
	var ed sql.NullString
	if rows.Next() {
		err = rows.Scan(&response.Meta.Name, &response.Meta.Type, &response.Meta.Done, &response.Meta.GroupId, &response.Meta.PublisherId,
			&response.Meta.LeaderId, &sd, &ed, &response.Meta.Readonly, &response.Meta.Description,
			)
		if sd.Valid {
			response.Meta.StartDate = sd.String
		}
		if ed.Valid {
			response.Meta.EndDate = ed.String
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (g Task) GetTaskPeolple(ctx context.Context, request *pb.GetTaskPeopleRequset, response *pb.GetTaskPeopleResponse) error {
	db := g.SqlDb
	rows, err := db.QueryContext(ctx, "SELECT `user_id`, `name`, `done`, `done_time` FROM `task_user` JOIN `user` ON `user_id`=`id` WHERE `task_id`=?", request.Id)
	if err != nil {
		return err
	}
	for rows.Next() {
		user := new(pb.GetTaskPeopleResponse_User)
		var dt sql.NullString
		err = rows.Scan(&user.Id, &user.Name, &user.Done, &dt)
		if err != nil {
			return err
		}
		if dt.Valid {
			user.DoneTime = dt.String
		}
	}
	return nil
}

func (g Task) GetTaskMetaByGroupId(ctx context.Context, request *pb.GetTaskMetaByGroupIdRequest, response *pb.GetTaskMetaByGroupIdResponse) error {
	db := g.SqlDb
	request
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


