package main

import (
	"context"
	"fmt"
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
	int0 := int32(0)
	int1 := int32(1)
	avatar0 := "ugly girl"
	avatar1 := "ugly boy"
	_, err := auth.PutUser(ctx, &pb3.PutUserRequest{
		Openid:               &openid,
		Name:                 &openid,
		Gender:               &int1,
		Avatar:               &avatar0,
	})
	assert.Empty(t, err)
	_, err = auth.PutUser(ctx, &pb3.PutUserRequest{
		Openid:               &openid2,
		Name:                 &openid2,
		Gender:               &int0,
		Avatar:               &avatar1,
	})
	assert.Empty(t, err)
	_, err = auth.PutUser(ctx, &pb3.PutUserRequest{
		Openid:               &openid3,
		Name:                 &openid3,
		Gender:               &int0,
		Avatar:               &avatar1,
	})
	assert.Empty(t, err)
	name1 := "tfboys"
	name2 := "tfgirls"
	rsp3, err := group.CreateGroup(ctx, &pb2.CreateGroupRequest{
		Name:                 &name1,
		CreatorId:            &openid,
		Type:                 &int0,
	})
	assert.Empty(t, err)
	_, err = group.AddUser(ctx, &pb2.AddUserRequest{
		GroupId:              rsp3.Id,
		UserIds:              []string{openid2},
	})
	assert.Empty(t, err)
	rsp4, err := group.CreateGroup(ctx, &pb2.CreateGroupRequest{
		Name:                 &name2,
		CreatorId:            &openid,
		Type:                 &int0,
	})
	assert.Empty(t, err)
	_, err = group.AddUser(ctx, &pb2.AddUserRequest{
		GroupId:              rsp4.Id,
		UserIds:              []string{openid2, openid3},
	})
	assert.Empty(t, err)
	gid1, gid2 := rsp3.Id, rsp4.Id
	// group task
	name3 := "t2"
	name4 := "t3"
	true_ := true
	false_ := true
	d1 := "2020-05-20T00:00:00"
	d2 := "2020-05-21T00:00:00"
	d6 := "2020-05-23T00:00:00"
	rsp7, err := task.CreateTask(ctx, &pb.CreateTaskRequest{
		GroupId:              gid1,
		UserIds:              []string{openid, openid2},
		Name:                 &name3,
		Type:                 &int0,
		LeaderId:             nil,
		StartDate:            &d1,
		EndDate:             &d6,
		Description:          nil,
		PublisherId:          &openid,
		Readonly:             &true_,
	})
	assert.Empty(t, err)
	task1 := rsp7.Id

	rsp8, err := task.CreateTask(ctx, &pb.CreateTaskRequest{
		GroupId:              gid2,
		UserIds:              []string{openid3, openid2},
		Name:                 &name4,
		Type:                 &int0,
		LeaderId:             &openid3,
		StartDate:            &d1,
		EndDate:              &d2,
		Description:          nil,
		PublisherId:          &openid,
		Readonly:             &false_,
	})
	assert.Empty(t, err)
	task2 := *rsp8.Id
	rsp5, err := task.GetTaskMetaByGroupId(ctx, &pb.GetTaskMetaByGroupIdRequest{
		GroupId:              gid1,
	})
	assert.Empty(t, err)
	assert.Equal(t, 1, len(rsp5.Metas))
	assert.Equal(t, *rsp5.Metas[0].Name, name3)
	assert.Equal(t, *rsp5.Metas[0].Id, *task1)
	rsp13, err := task.GetTaskMetaByUserId(ctx, &pb.GetTaskMetaByUserIdRequest{
		UserId:               &openid2,
	})
	assert.Empty(t, err)
	assert.Equal(t, 2, len(rsp13.Metas))
	t1, err := time.Parse("2006-01-02T15:04:05", *rsp13.Metas[0].EndDate)
	t2, err := time.Parse("2006-01-02T15:04:05", *rsp13.Metas[1].EndDate)
	fmt.Println( *rsp13.Metas[0].EndDate)
	fmt.Println(t2)
	assert.True(t, t1.Before(t2))
	rsp12, err := group.GetGroup(ctx, &pb2.GetGroupRequest{
		Id:                   gid1,
	})
	assert.Empty(t, err)
	assert.Equal(t, int32(1), *rsp12.TaskNum)
	assert.NotEmpty(t, *rsp12.UpdatedAt)
	rsp6, err := task.UserInTask(ctx, &pb.UserInTaskRequest{
		UserId:               &openid3,
		TaskId:               task1,
	})
	assert.Empty(t, err)
	assert.False(t, *rsp6.Ok)
	rsp6, err = task.UserInTask(ctx, &pb.UserInTaskRequest{
		UserId:               &openid3,
		TaskId:               &task2,
	})
	assert.Empty(t, err)
	assert.True(t, *rsp6.Ok)
	rsp9, err := task.GetTaskPeolple(ctx, &pb.GetTaskPeopleRequset{
		Id:                   &task2,
	})
	assert.Empty(t, err)
	assert.Equal(t, len(rsp9.Workers), 2)
	assert.False(t, *rsp9.Workers[1].Done)
	_, err = task.DeleteTask(ctx, &pb.DeleteTaskRequest{
		Id:                   task1,
	})
	assert.Empty(t, err)
	_, err = task.GetTaskMeta(ctx, &pb.GetTaskMetaRequest{
		Id:                   task1,
	})
	assert.NotEmpty(t, err)
	name5 := ""
	d3 := "2020-05-30T00:00:00"
	des1 := ""
	_, err = task.UpdateTaskMeta(ctx, &pb.UpdateTaskMetaRequest{
		Id:                   &task2,
		Name:                 &name5,
		StartDate:            &d3,
		EndDate:              &d1,
		Readonly:             &false_,
		Description:          &des1,
	})
	assert.NotEmpty(t, err)
	name6 := "modified"
	d4 := "2020-05-21T00:00:00"
	des2 := "nothing"
	_, err = task.UpdateTaskMeta(ctx, &pb.UpdateTaskMetaRequest{
		Id:                   &task2,
		Name:                 &name6,
		StartDate:            &d4,
		EndDate:              &d4,
		Readonly:             &false_,
		Description:          &des2,
	})
	assert.Empty(t, err)
	name7 := "don't modified"
	d5 := ""
	des3 := "really nothing"
	_, err = task.UpdateTaskMeta(ctx, &pb.UpdateTaskMetaRequest{
		Id:                   &task2,
		Name:                 &name7,
		StartDate:            &d5,
		EndDate:              &d4,
		Readonly:             &false_,
		Description:          &des3,
	})
	assert.NotEmpty(t, err)
	rsp11, err := task.GetTaskMeta(ctx, &pb.GetTaskMetaRequest{
		Id:                   &task2,
	})
	assert.Empty(t, err)
	assert.NotEmpty(t, rsp11.Meta)
	assert.Equal(t, *rsp11.Meta.Name, name6)

	// TODO content api
	// TODO: done and donetime api
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}
