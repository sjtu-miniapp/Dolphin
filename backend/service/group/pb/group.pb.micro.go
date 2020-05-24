// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: group.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Group service

type GroupService interface {
	GetGroup(ctx context.Context, in *GetGroupRequest, opts ...client.CallOption) (*GetGroupResponse, error)
	GetGroupByUserId(ctx context.Context, in *GetGroupByUserIdRequest, opts ...client.CallOption) (*GetGroupByUserIdResponse, error)
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...client.CallOption) (*CreateGroupResponse, error)
	AddUser(ctx context.Context, in *AddUserRequest, opts ...client.CallOption) (*AddUserResponse, error)
	UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...client.CallOption) (*UpdateGroupResponse, error)
	DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...client.CallOption) (*DeleteGroupResponse, error)
	UserInGroup(ctx context.Context, in *UserInGroupRequest, opts ...client.CallOption) (*UserInGroupResponse, error)
}

type groupService struct {
	c    client.Client
	name string
}

func NewGroupService(name string, c client.Client) GroupService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "group"
	}
	return &groupService{
		c:    c,
		name: name,
	}
}

func (c *groupService) GetGroup(ctx context.Context, in *GetGroupRequest, opts ...client.CallOption) (*GetGroupResponse, error) {
	req := c.c.NewRequest(c.name, "Group.GetGroup", in)
	out := new(GetGroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) GetGroupByUserId(ctx context.Context, in *GetGroupByUserIdRequest, opts ...client.CallOption) (*GetGroupByUserIdResponse, error) {
	req := c.c.NewRequest(c.name, "Group.GetGroupByUserId", in)
	out := new(GetGroupByUserIdResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...client.CallOption) (*CreateGroupResponse, error) {
	req := c.c.NewRequest(c.name, "Group.CreateGroup", in)
	out := new(CreateGroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) AddUser(ctx context.Context, in *AddUserRequest, opts ...client.CallOption) (*AddUserResponse, error) {
	req := c.c.NewRequest(c.name, "Group.AddUser", in)
	out := new(AddUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...client.CallOption) (*UpdateGroupResponse, error) {
	req := c.c.NewRequest(c.name, "Group.UpdateGroup", in)
	out := new(UpdateGroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...client.CallOption) (*DeleteGroupResponse, error) {
	req := c.c.NewRequest(c.name, "Group.DeleteGroup", in)
	out := new(DeleteGroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) UserInGroup(ctx context.Context, in *UserInGroupRequest, opts ...client.CallOption) (*UserInGroupResponse, error) {
	req := c.c.NewRequest(c.name, "Group.UserInGroup", in)
	out := new(UserInGroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Group service

type GroupHandler interface {
	GetGroup(context.Context, *GetGroupRequest, *GetGroupResponse) error
	GetGroupByUserId(context.Context, *GetGroupByUserIdRequest, *GetGroupByUserIdResponse) error
	CreateGroup(context.Context, *CreateGroupRequest, *CreateGroupResponse) error
	AddUser(context.Context, *AddUserRequest, *AddUserResponse) error
	UpdateGroup(context.Context, *UpdateGroupRequest, *UpdateGroupResponse) error
	DeleteGroup(context.Context, *DeleteGroupRequest, *DeleteGroupResponse) error
	UserInGroup(context.Context, *UserInGroupRequest, *UserInGroupResponse) error
}

func RegisterGroupHandler(s server.Server, hdlr GroupHandler, opts ...server.HandlerOption) error {
	type group interface {
		GetGroup(ctx context.Context, in *GetGroupRequest, out *GetGroupResponse) error
		GetGroupByUserId(ctx context.Context, in *GetGroupByUserIdRequest, out *GetGroupByUserIdResponse) error
		CreateGroup(ctx context.Context, in *CreateGroupRequest, out *CreateGroupResponse) error
		AddUser(ctx context.Context, in *AddUserRequest, out *AddUserResponse) error
		UpdateGroup(ctx context.Context, in *UpdateGroupRequest, out *UpdateGroupResponse) error
		DeleteGroup(ctx context.Context, in *DeleteGroupRequest, out *DeleteGroupResponse) error
		UserInGroup(ctx context.Context, in *UserInGroupRequest, out *UserInGroupResponse) error
	}
	type Group struct {
		group
	}
	h := &groupHandler{hdlr}
	return s.Handle(s.NewHandler(&Group{h}, opts...))
}

type groupHandler struct {
	GroupHandler
}

func (h *groupHandler) GetGroup(ctx context.Context, in *GetGroupRequest, out *GetGroupResponse) error {
	return h.GroupHandler.GetGroup(ctx, in, out)
}

func (h *groupHandler) GetGroupByUserId(ctx context.Context, in *GetGroupByUserIdRequest, out *GetGroupByUserIdResponse) error {
	return h.GroupHandler.GetGroupByUserId(ctx, in, out)
}

func (h *groupHandler) CreateGroup(ctx context.Context, in *CreateGroupRequest, out *CreateGroupResponse) error {
	return h.GroupHandler.CreateGroup(ctx, in, out)
}

func (h *groupHandler) AddUser(ctx context.Context, in *AddUserRequest, out *AddUserResponse) error {
	return h.GroupHandler.AddUser(ctx, in, out)
}

func (h *groupHandler) UpdateGroup(ctx context.Context, in *UpdateGroupRequest, out *UpdateGroupResponse) error {
	return h.GroupHandler.UpdateGroup(ctx, in, out)
}

func (h *groupHandler) DeleteGroup(ctx context.Context, in *DeleteGroupRequest, out *DeleteGroupResponse) error {
	return h.GroupHandler.DeleteGroup(ctx, in, out)
}

func (h *groupHandler) UserInGroup(ctx context.Context, in *UserInGroupRequest, out *UserInGroupResponse) error {
	return h.GroupHandler.UserInGroup(ctx, in, out)
}
