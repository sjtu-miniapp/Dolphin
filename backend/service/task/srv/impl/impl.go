package impl

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sjtu-miniapp/dolphin/service/database/model"
	"github.com/sjtu-miniapp/dolphin/service/task/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Task struct {
	SqlDb   *gorm.DB
	MongoDb *mongo.Database
}

const collection = "task"

func time2string(time time.Time) string {
	s := fmt.Sprintf("%d-%d-%d", time.Year(),
		time.Month(), time.Day())
	return s
}

func string2time(str string) (time.Time, error) {
	return time.Parse("2006-01-02", str)
}

func (g Task) GetTaskMeta(ctx context.Context, request *pb.GetTaskMetaRequest, response *pb.GetTaskMetaResponse) error {
	if request.Id == nil {
		return fmt.Errorf("nil pointer")
	}
	db := g.SqlDb
	task := model.Task{
		Id: *request.Id,
	}
	if err := db.Find(&task).Error; err != nil {
		return err
	}
	response.Meta = &pb.TaskMeta{
		Name:        task.Name,
		Type:        &task.Type,
		Done:        &task.Done,
		GroupId:     &task.GroupId,
		PublisherId: &task.PublisherId,
		LeaderId:    task.LeaderId,
		StartDate:   nil,
		EndDate:     nil,
		Readonly:    &task.Readonly,
		Description: task.Desciption,
	}
	if task.EndDate != nil && task.StartDate != nil {
		ed := time2string(*task.EndDate)
		sd := time2string(*task.StartDate)
		response.Meta.EndDate = &ed
		response.Meta.StartDate = &sd
	}

	return nil
}

func (g Task) GetTaskPeolple(ctx context.Context, request *pb.GetTaskPeopleRequset, response *pb.GetTaskPeopleResponse) error {
	if request.Id == nil {
		return fmt.Errorf("nil pointer")
	}
	db := g.SqlDb
	task := model.Task{
		Id: *request.Id,
	}
	if err := db.Preload("Users").Find(&task).Error; err != nil {
		return err
	}
	for _, v := range task.Users {
		user := pb.GetTaskPeopleResponse_User{
			Id:       &v.Id,
			Name:     &v.Name,
			Done:     nil,
			DoneTime: nil,
			Avatar:   v.Avatar,
			Gender:   v.Gender,
		}
		userTask := model.UserTask{
			UserId: v.Id,
			TaskId: *request.Id,
		}
		if err := db.Find(&userTask).Error; err != nil {
			return err
		}
		user.Done = &userTask.Done
		if userTask.DoneTime != nil {
			s := time2string(*userTask.DoneTime)
			user.DoneTime = &s
		}
		response.Workers = append(response.Workers, &user)
	}
	return nil
}

func (g Task) GetTaskMetaByGroupId(ctx context.Context, request *pb.GetTaskMetaByGroupIdRequest, response *pb.GetTaskMetaByGroupIdResponse) error {
	if request.GroupId == nil {
		return fmt.Errorf("nil pointer")
	}
	db := g.SqlDb
	group := model.Group{
		Id: *request.GroupId,
	}

	if err := db.Preload("Tasks").Find(&group).Error; err != nil {
		return err
	}
	for _, v := range group.Tasks {
		task := pb.TaskMeta{
			Name:        v.Name,
			Type:        &v.Type,
			Done:        &v.Done,
			GroupId:     &v.GroupId,
			PublisherId: &v.PublisherId,
			LeaderId:    v.LeaderId,
			StartDate:   nil,
			EndDate:     nil,
			Readonly:    &v.Readonly,
			Description: v.Desciption,
		}
		if v.StartDate != nil {
			s := time2string(*v.StartDate)
			task.StartDate = &s
		}
		if v.EndDate != nil {
			s := time2string(*v.EndDate)
			task.EndDate = &s
		}
		response.Metas = append(response.Metas, &task)
	}
	return nil
}

// including add user_task
func (g Task) CreateTask(ctx context.Context, request *pb.CreateTaskRequest, response *pb.CreateTaskResponse) error {
	if request.GroupId == nil || request.PublisherId == nil ||
		request.Type == nil || request.Readonly == nil {
		return fmt.Errorf("nil pointer")
	}
	db := g.SqlDb
	if len(request.UserIds) == 0 {
		return fmt.Errorf("at least one user should be added")
	}
	task := model.Task{
		GroupId:     *request.GroupId,
		Name:        request.Name,
		PublisherId: *request.PublisherId,
		LeaderId:    request.LeaderId,
		StartDate:   nil,
		EndDate:     nil,
		Readonly:    *request.Readonly,
		Type:        *request.Type,
		Desciption:  request.Description,
		Done:        false,
		Users:       nil,
	}
	if (request.EndDate != nil) != (request.StartDate != nil) {
		return fmt.Errorf("wrong date format")
	} else if request.EndDate != nil {
		ed, err := string2time(*request.EndDate)
		if err != nil {
			return fmt.Errorf("wrong date format: %v", err)
		}
		sd, err := string2time(*request.StartDate)
		if err != nil {
			return fmt.Errorf("wrong date format: %v", err)
		}
		if sd.After(ed) {
			return fmt.Errorf("start date after end date")
		}
		task.EndDate = &ed
		task.StartDate = &sd
	}

	for _, v := range request.UserIds {
		user := model.User{
			Id: v,
		}
		if err := db.First(&user).Error; err != nil {
			return err
		}
		task.Users = append(task.Users, &user)
	}
	if err := db.Save(&task).Error; err != nil {
		return err
	}
	response.Id = &task.Id
	go func() {
		createContent(task.Id)
	}()
	return nil
}

func (g Task) UpdateTaskMeta(ctx context.Context, request *pb.UpdateTaskMetaRequest, response *pb.UpdateTaskMetaResponse) error {
	if request.Id == nil {
		return fmt.Errorf("nil pointer")
	}
	db := g.SqlDb
	task := model.Task{
		Id: *request.Id,
	}
	if err := db.First(&task).Error; err != nil {
		return err
	}

	if (request.EndDate != nil) != (request.StartDate != nil) {
		return fmt.Errorf("wrong date format")
	} else if request.EndDate != nil {
		ed, err := string2time(*request.EndDate)
		if err != nil {
			return fmt.Errorf("wrong date format: %v", err)
		}
		sd, err := string2time(*request.StartDate)
		if err != nil {
			return fmt.Errorf("wrong date format: %v", err)
		}
		if sd.After(ed) {
			return fmt.Errorf("start date after end date")
		}
		task.EndDate = &ed
		task.StartDate = &sd
	}
	if request.Description != nil {
		task.Desciption = request.Description
	}
	if request.Readonly != nil {
		task.Readonly = *request.Readonly
	}
	if request.Name != nil {
		task.Name = request.Name
	}
	if err := db.Save(&task).Error; err != nil {
		return err
	}
	return nil
}

func (g Task) DeleteTask(ctx context.Context, request *pb.DeleteTaskRequest, response *pb.DeleteTaskResponse) error {
	if request.Id == nil {
		return fmt.Errorf("nil pointer")
	}
	db := g.SqlDb
	err := db.Delete(&model.Task{
		Id: *request.Id,
	}).Error
	return err
}

func (g Task) UserInTask(ctx context.Context, request *pb.UserInTaskRequest, response *pb.UserInTaskResponse) error {
	if request.UserId == nil || request.TaskId == nil {
		return fmt.Errorf("nil pointer")
	}
	db := g.SqlDb
	err := db.First(&model.UserTask{
		UserId: *request.UserId,
		TaskId: *request.TaskId,
	}).Error
	b := false
	if err != nil {
		if err.Error() != "record not found" {
			return err
		}
		response.Ok = &b
		return nil
	}
	b = true
	response.Ok = &b
	return nil
}


// TODO
// TODO: TEST
func (g Task) GetTaskContent(ctx context.Context, request *pb.GetTaskContentRequest, response *pb.GetTaskContentResponse) error {
	coll := g.MongoDb.Collection(collection)

}

func createContent(taskId int32) {
	a, err := collection.InsertOne(context.Background(), bson.E{"hah", "al"})

}

//func (g Task) update
// commit update
