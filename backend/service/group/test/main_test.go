package main

import (
	"context"
	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	pb2 "github.com/sjtu-miniapp/dolphin/service/auth/pb"
	"github.com/sjtu-miniapp/dolphin/service/group/pb"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)


var group pb.GroupService
var auth pb2.AuthService
var seededRand *rand.Rand

func setup() {
	seededRand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"127.0.0.1:2379"}
	})
	service := micro.NewService(
		micro.Name("go.micro.cli.group"),
		micro.Registry(reg),
		micro.Flags(
			&cli.StringFlag{
				Name:  "test.timeout",
				Usage: "dumbass go-micro",
			},
		),
	)
	service.Init()
	group = pb.NewGroupService("go.micro.srv.group", service.Client())

	service2 := micro.NewService(
		micro.Name("go.micro.cli.auth"),
		micro.Registry(reg),
	)
	service2.Init()
	auth = pb2.NewAuthService("go.micro.srv.auth", service2.Client())
}

func TestGroup(t *testing.T) {
	ctx := context.Background()
	openid := strconv.Itoa(rand.Intn(10000))
	openid2 := strconv.Itoa(rand.Intn(10000))
	_, err := auth.PutUser(ctx, &pb2.PutUserRequest{
		Openid:               openid,
		Name:                 openid,
		Gender:               1,
		Avatar:               "ugly girl",
	})
	assert.Empty(t, err)
	_, err = auth.PutUser(ctx, &pb2.PutUserRequest{
		Openid:               openid2,
		Name:                 openid2,
		Gender:               0,
		Avatar:               "ugly boy",
	})
	assert.Empty(t, err)
	rsp, err := auth.GetUser(ctx, &pb2.GetUserRequest{
		Id:                   openid,
	})
	assert.Empty(t, err)
	assert.NotEmpty(t, rsp.SelfGroupId)
	groupid := rsp.SelfGroupId
	rsp2, err := group.GetGroup(ctx, &pb.GetGroupRequest{
		Id:                   groupid,
	})
	assert.Empty(t, err)
	assert.NotEmpty(t, rsp2)
	assert.Equal(t, rsp2.Type, uint32(1))
	assert.Equal(t, rsp2.CreatorId, openid)
	_, err = group.AddUser(ctx, &pb.AddUserRequest{
		GroupId:              groupid,
		UserIds:              []string{strconv.Itoa(rand.Intn(10000)), openid2},
	})
	assert.NotEmpty(t, err)
	_, err = group.DeleteGroup(ctx, &pb.DeleteGroupRequest{
		Id:                   groupid,
	})
	assert.NotEmpty(t, err)
	rsp3, err := group.CreateGroup(ctx, &pb.CreateGroupRequest{
		Name:                 "tfboys",
		CreatorId:            openid,
		Type:                 0,
	})
	assert.Empty(t, err)
	assert.NotEmpty(t, rsp3.Id)
	rsp4, err := group.GetGroup(ctx, &pb.GetGroupRequest{
		Id:                   rsp3.Id,
	})
	assert.Empty(t, err)
	assert.Equal(t, rsp4.Name, "tfboys")
	assert.Equal(t, rsp4.CreatorId, openid)
	assert.Equal(t, len(rsp4.Users), 1)
	_, err = group.AddUser(ctx, &pb.AddUserRequest{
		GroupId:              rsp3.Id,
		UserIds:              []string{strconv.Itoa(rand.Intn(10000)), openid2},
	})
	assert.NotEmpty(t, err)
	_, err = group.AddUser(ctx, &pb.AddUserRequest{
		GroupId:              rsp3.Id,
		UserIds:              []string{openid2},
	})
	assert.Empty(t, err)
	rsp4, err = group.GetGroup(ctx, &pb.GetGroupRequest{
		Id:                   rsp3.Id,
	})
	assert.Empty(t, err)
	assert.Equal(t, len(rsp4.Users), 2)
	assert.Equal(t, true, rsp4.Users[1].Id == openid2 || rsp4.Users[0].Id == openid2)
	_, err = group.UpdateGroup(ctx, &pb.UpdateGroupRequest{
		Id:                   rsp3.Id,
		Name:                 "tfgirls",
	})
	assert.Empty(t, err)
	_, err = group.UpdateGroup(ctx, &pb.UpdateGroupRequest{
		Id:                   rsp3.Id,
		Name:                 "",
	})
	assert.NotEmpty(t, err)
	rsp4, err = group.GetGroup(ctx, &pb.GetGroupRequest{
		Id:                   rsp3.Id,
	})
	assert.Empty(t, err)
	assert.Equal(t, rsp4.Name, "tfgirls")
	rsp5, err := group.GetGroupByUserId(context.TODO(), &pb.GetGroupByUserIdRequest{
		UserId:    openid,
	})
	assert.Empty(t, err)
	assert.Equal(t, len(rsp5.Groups), 1)
	assert.Equal(t, rsp5.Groups[0].Name, "tfgirls")
	rsp6, err := group.UserInGroup(ctx, &pb.UserInGroupRequest{
		UserId:               openid2,
		GroupId:              rsp3.Id,
	})
	assert.Empty(t, err)
	assert.Equal(t, rsp6.Ok, true)
	rsp6, err = group.UserInGroup(ctx, &pb.UserInGroupRequest{
		UserId:               "wangjunkai",
		GroupId:              rsp3.Id,
	})
	assert.Empty(t, err)
	assert.Equal(t, rsp6.Ok, false)
	_, err = group.DeleteGroup(ctx, &pb.DeleteGroupRequest{
		Id:                   rsp3.Id,
	})
	assert.Empty(t, err)
	rsp5, err = group.GetGroupByUserId(context.TODO(), &pb.GetGroupByUserIdRequest{
		UserId:    openid2,
	})
	assert.Empty(t, err)
	assert.Equal(t, len(rsp5.Groups), 0)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}



