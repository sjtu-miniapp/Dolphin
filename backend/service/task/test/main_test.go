package main

import (
	"context"
	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	pb3 "github.com/sjtu-miniapp/dolphin/service/auth/pb"
	pb2 "github.com/sjtu-miniapp/dolphin/service/group/pb"
	"github.com/sjtu-miniapp/dolphin/service/task/pb"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)


var task pb.TaskService
var group pb2.GroupService
var auth pb3.AuthService
var seededRand *rand.Rand

func setup() {
	seededRand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"127.0.0.1:2379"}
	})

	service := micro.NewService(
		micro.Name("go.micro.cli.task"),
		micro.Registry(reg),
		micro.Flags(
			&cli.StringFlag{
				Name:  "test.timeout",
				Usage: "dumbass go-micro",
			},
		),
	)
	service.Init()
	task = pb.NewTaskService("go.micro.srv.task", service.Client())

	service2 := micro.NewService(
		micro.Name("go.micro.cli.group"),
		micro.Registry(reg),
	)
	service2.Init()
	group = pb2.NewGroupService("go.micro.srv.group", service.Client())

	service3 := micro.NewService(
		micro.Name("go.micro.cli.auth"),
		micro.Registry(reg),
	)
	service3.Init()
	auth = pb3.NewAuthService("go.micro.srv.auth", service2.Client())
}

func TestTask(t *testing.T) {
	ctx := context.Background()
	openid := strconv.Itoa(rand.Intn(10000))
	openid2 := strconv.Itoa(rand.Intn(10000))
	openid3 := strconv.Itoa(rand.Intn(10000))
	_, err := auth.PutUser(ctx, &pb3.PutUserRequest{
		Openid:               openid,
		Name:                 openid,
		Gender:               1,
		Avatar:               "ugly girl",
	})
	assert.Empty(t, err)
	_, err = auth.PutUser(ctx, &pb3.PutUserRequest{
		Openid:               openid2,
		Name:                 openid2,
		Gender:               0,
		Avatar:               "ugly boy",
	})
	assert.Empty(t, err)
	_, err = auth.PutUser(ctx, &pb3.PutUserRequest{
		Openid:               openid3,
		Name:                 openid3,
		Gender:               0,
		Avatar:               "ugly boy",
	})
	assert.Empty(t, err)
	rsp3, err := group.CreateGroup(ctx, &pb2.CreateGroupRequest{
		Name:                 "tfboys",
		CreatorId:            openid,
		Type:                 0,
	})
	assert.Empty(t, err)
	_, err = group.AddUser(ctx, &pb2.AddUserRequest{
		GroupId:              rsp3.Id,
		UserIds:              []string{openid2},
	})
	assert.Empty(t, err)
	rsp4, err := group.CreateGroup(ctx, &pb2.CreateGroupRequest{
		Name:                 "tfboys",
		CreatorId:            openid,
		Type:                 0,
	})
	assert.Empty(t, err)
	_, err = group.AddUser(ctx, &pb2.AddUserRequest{
		GroupId:              rsp3.Id,
		UserIds:              []string{openid2, openid3},
	})
	assert.Empty(t, err)
	gid1, gid2 := rsp3.Id, rsp4.Id
	// group task
	_, err = task.CreateTask(ctx, &pb.CreateTaskRequest{
		GroupId:              int32(gid1),
		UserIds:              []string{openid3},
		Name:                 "t1",
		Type:                 0,
		LeaderId:             "",
		StartDate:            "",
		EndDate:              "",
		Description:          "",
		PublisherId:          openid,
		Readonly:             false,
	})
	assert.NotEmpty(t, err)
	rsp7, err := task.CreateTask(ctx, &pb.CreateTaskRequest{
		GroupId:              int32(gid1),
		UserIds:              []string{openid, openid2},
		Name:                 "t2",
		Type:                 0,
		LeaderId:             "",
		StartDate:            "",
		EndDate:              "",
		Description:          "",
		PublisherId:          openid,
		Readonly:             true,
	})
	assert.Empty(t, err)
	task1 := rsp7.Id
	rsp8, err := task.CreateTask(ctx, &pb.CreateTaskRequest{
		GroupId:              int32(gid2),
		UserIds:              []string{openid3, openid2},
		Name:                 "t3",
		Type:                 0,
		LeaderId:             openid3,
		StartDate:            "2020-05-20",
		EndDate:              "2020-05-21",
		Description:          "",
		PublisherId:          openid,
		Readonly:             false,
	})
	assert.Empty(t, err)
	task2 := rsp8.Id
	rsp5, err := task.GetTaskMetaByGroupId(ctx, &pb.GetTaskMetaByGroupIdRequest{
		GroupId:              gid1,
	})
	assert.Empty(t, err)
	assert.Equal(t, 1, len(rsp5.Metas))
	assert.Equal(t, rsp5.Metas[0].Name, "t1")
	rsp6, err := task.UserInTask(ctx, &pb.UserInTaskRequest{
		UserId:               openid3,
		TaskId:               task1,
	})
	assert.Empty(t, err)
	assert.False(t, rsp6.Ok)
	rsp6, err = task.UserInTask(ctx, &pb.UserInTaskRequest{
		UserId:               openid3,
		TaskId:               task2,
	})
	assert.Empty(t, err)
	assert.True(t, rsp6.Ok)
	rsp9, err := task.GetTaskPeolple(ctx, &pb.GetTaskPeopleRequset{
		Id:                   task2,
	})
	assert.Empty(t, err)
	assert.Equal(t, len(rsp9.Workers), 3)
	assert.False(t, rsp9.Workers[1].Done)
	_, err = task.DeleteTask(ctx, &pb.DeleteTaskRequest{
		Id:                   task1,
	})
	assert.Empty(t, err)
	_, err = task.GetTaskMeta(ctx, &pb.GetTaskMetaRequest{
		Id:                   task1,
	})
	assert.NotEmpty(t, err)
	_, err = task.UpdateTaskMeta(ctx, &pb.UpdateTaskMetaRequest{
		Id:                   task2,
		Name:                 "",
		StartDate:            "2020-05-30",
		EndDate:              "2020-05-20",
		Readonly:             false,
		Description:          "",
	})
	assert.NotEmpty(t, err)
	_, err = task.UpdateTaskMeta(ctx, &pb.UpdateTaskMetaRequest{
		Id:                   task2,
		Name:                 "modified",
		StartDate:            "2020-05-21",
		EndDate:              "2020-05-21",
		Readonly:             false,
		Description:          "nothing",
	})
	assert.Empty(t, err)
	_, err = task.UpdateTaskMeta(ctx, &pb.UpdateTaskMetaRequest{
		Id:                   task2,
		Name:                 "don't modifiy",
		StartDate:            "",
		EndDate:              "2020-05-21",
		Readonly:             false,
		Description:          "really nothing",
	})
	assert.NotEmpty(t, err)
	rsp11, err := task.GetTaskMeta(ctx, &pb.GetTaskMetaRequest{
		Id:                   task2,
	})
	assert.NotEmpty(t, err)
	assert.Equal(t, rsp11.Meta.Name, "modified")
	// TODO: done and donetime api
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}
