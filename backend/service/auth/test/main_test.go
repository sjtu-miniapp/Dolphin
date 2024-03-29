package main

import (
	"context"
	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/sjtu-miniapp/dolphin/service/auth/pb"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)


var auth pb.AuthService
var seededRand *rand.Rand

func setup() {
	seededRand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"127.0.0.1:2379"}
	})
	service := micro.NewService(
		micro.Name("go.micro.cli.auth"),
		micro.Registry(reg),
		micro.Flags(
			&cli.StringFlag{
				Name:  "test.timeout",
				Usage: "dumbass go-micro",
			},
		),
	)
	service.Init()
	auth = pb.NewAuthService("go.micro.srv.auth", service.Client())

}

func TestOnLoginAndCheckAuth(t *testing.T) {
	code := string(rand.Uint64())
	resp, err := auth.OnLogin(context.TODO(), &pb.OnLoginRequest{
		Code: &code,
	})

	if err != nil {
		t.Fatal(err)
		return
	}
	if resp == nil {
		t.Fatalf("no response")
	}

	r, err := auth.CheckAuth(context.TODO(), &pb.CheckAuthRequest{
		Openid: resp.Openid,
		Sid:    resp.Sid,
	})

	if err != nil {
		t.Fatal(err)
	}
	if !*r.Ok {
		t.Errorf("check auth fail")
	}
	tmp := *resp.Openid + "123"
	r, err = auth.CheckAuth(context.TODO(), &pb.CheckAuthRequest{
		Openid: &tmp,
		Sid:    resp.Sid,
	})

	if r != nil && *r.Ok {
		t.Errorf("check auth fail")
	}

}

func TestPutUserAndGetUser(t *testing.T) {
	randStr := strconv.Itoa(rand.Intn(10000))
	resp, err := auth.OnLogin(context.TODO(), &pb.OnLoginRequest{
		Code: &randStr,
	})
	assert.Empty(t, err)
	assert.NotEmpty(t, resp)
	rs, err := auth.GetUser(context.Background(), &pb.GetUserRequest{
		Id: resp.Openid,
	})
	assert.Empty(t, rs)
	gender := int32(1)
	avatar := "ugly boy"
	_, err = auth.PutUser(context.Background(), &pb.PutUserRequest{
		Openid: resp.Openid,
		Name:   &randStr,
		Gender: &gender,
		Avatar: &avatar,
	})
	assert.Empty(t, err, err)
	rs, err = auth.GetUser(context.Background(), &pb.GetUserRequest{
		Id: resp.Openid,
	})
	assert.Empty(t, err, err)
	assert.Equal(t, *rs.Name, randStr)
	assert.NotEmpty(t, rs.SelfGroupId)
}


func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}