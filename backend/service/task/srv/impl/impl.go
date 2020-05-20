package impl

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sjtu-miniapp/dolphin/service/task/pb"
	"time"
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
	} else {
		return fmt.Errorf("no task found")
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
		response.Workers = append(response.Workers, user)
	}
	return nil
}

func (g Task) GetTaskMetaByGroupId(ctx context.Context, request *pb.GetTaskMetaByGroupIdRequest, response *pb.GetTaskMetaByGroupIdResponse) error {
	db := g.SqlDb
	rows, err := db.QueryContext(ctx, "SELECT `type`, `publisher_id`, `leader_id`, `readonly`, `description`, `start_date`, `end_date`, `name`, `done` FROM `task` WHERE `group_id` = ?", request.GroupId)
	if err != nil {
		return err
	}
	for rows.Next() {
		meta := new(pb.TaskMeta)
		var sd sql.NullString
		var ed sql.NullString
		err = rows.Scan(&meta.Type, &meta.PublisherId, &meta.LeaderId, &meta.Readonly,
			&meta.Description, &sd, &ed, &meta.Name, &meta.Done)
		if err != nil {
			return err
		}
		meta.GroupId = request.GroupId
		if sd.Valid {
			meta.StartDate = sd.String
			if !ed.Valid {
				return fmt.Errorf("wrong date format")
			}
			meta.EndDate = ed.String
			t1, err := time.Parse("2000-01-01", sd.String)
			if err != nil {
				return err
			}
			t2, err := time.Parse("2000-01-01", sd.String)
			if err != nil {
				return err
			}
			if t1.After(t2) {
				return fmt.Errorf("start date after end date")
			}
		} else if ed.Valid {
			return fmt.Errorf("wrong date format")
		}
		response.Metas = append(response.Metas, meta)
	}
	return nil
}
// including add user_task
func (g Task) CreateTask(ctx context.Context, request *pb.CreateTaskRequest, response *pb.CreateTaskResponse) error {
	db := g.SqlDb
	result, err := db.ExecContext(ctx, "INSERT INTO `task`(`type`, `publisher_id`, `readonly`, `description`, `name`, `group_id`) VALUES(?, ?, ?, ?, ?, ?)",
		request.Type, request.PublisherId, request.Readonly, request.Description, request.Name, request.GroupId)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	response.Id = uint32(id)
	errChan := make(chan error)
	go func() {
		defer func() {
			errChan <- err
		}()
		if len(request.UserIds) == 0 {
			err = fmt.Errorf("at least one user should be added")
			return
		}
		sql1 := fmt.Sprintf("INSERT INTO `user_task`(`user_id`, `task_id`) VALUES('%s', %d)", request.UserIds[0], id)
		for _, v := range request.UserIds[1:] {
			sql1 = sql1 + fmt.Sprintf(", ('%s', %d)", v, id)
		}
		_, err = db.ExecContext(ctx, sql1)
	}()

	go func() {
		defer func() {
			errChan <- err
		}()
		var str string
		if request.StartDate != "" {
			str := fmt.Sprintf("`start_date` = '%s', `end_date` = '%s'", request.StartDate, request.EndDate)
			if request.LeaderId != "" {
				str += fmt.Sprintf(", `leader_id` = '%s'", request.LeaderId)
			}
		} else if request.LeaderId != "" {
			str = fmt.Sprintf("`leader_id` = '%s'", request.LeaderId)
		} else {
			err = nil
			return
		}
		sql1 := fmt.Sprintf("UPDATE `task` SET %s WHERE `id` = %d", str, id)
		_, err = db.ExecContext(ctx, sql1)
	}()
	if <-errChan != nil {
		return err
	}
	if <-errChan != nil {
		return err
	}
	return nil
	// TODO: ADD CONTENT
}

func (g Task) UpdateTaskMeta(ctx context.Context, request *pb.UpdateTaskMetaRequest, response *pb.UpdateTaskMetaResponse) error {
	db := g.SqlDb
	str := ""
	if request.StartDate != "" {
		str = fmt.Sprintf(", `start_date` = '%s', `end_date` = '%s'", request.StartDate, request.EndDate)
	}
	_, err := db.ExecContext(ctx, "UPDATE `task` SET `description` = ?, `readonly` = ?, `name` = ?"+ str + " WHERE `id` = ?",
		request.Description, request.Readonly, request.Name, request.Id)
	return err
}

func (g Task) DeleteTask(ctx context.Context, request *pb.DeleteTaskRequest, response *pb.DeleteTaskResponse) error {
	db := g.SqlDb
	_, err := db.ExecContext(ctx, "DELETE FROM `task` WHERE `id` = ?", request.Id)
	return err
}

func (g Task) UserInTask(ctx context.Context, request *pb.UserInTaskRequest, response *pb.UserInTaskResponse) error {
	db := g.SqlDb
	rows, err := db.QueryContext(ctx, "SELECT * FROM `task_user` WHERE `task_id` = ? and `user_id` = ?", request.TaskId, request.UserId)
	if err != nil {
		response.Ok = false
		return err
	}
	response.Ok = rows.Next()
	return nil
}
