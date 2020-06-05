package impl

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/common/log"
	"github.com/sjtu-miniapp/dolphin/service/database"
	"github.com/sjtu-miniapp/dolphin/service/database/model"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	//"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sjtu-miniapp/dolphin/service/auth/pb"
)

type Auth struct {
	SqlDb     *gorm.DB
	RedisDb   *redis.Client
	AppId     string
	AppSecret string
}

type authVal struct {
	Openid string `json:"openid"`
	Sid    string `json:"sid"`
}

type User struct {
	Openid      string  `json:"openid"`
	Name        string  `json:"name"`
	Gender      *int32  `json:"gender"`
	Avatar      *string `json:"avatar"`
	SelfGroupId *int32  `json:"self_group_id"`
}

// openid, sid
// acquire for session_key
func (a Auth) OnLogin(ctx context.Context, request *pb.OnLoginRequest, response *pb.OnLoginResponse) error {
	//for test
	//openid, skey := *request.Code, ""
	//sbyte := sha256.Sum256([]byte(openid + skey))
	//sid := fmt.Sprintf("%x", sbyte)
	//
	//go func() {
	//	_ = database.Set(a.RedisDb, 1*time.Hour, sid, authVal{Openid: openid, Sid: sid})
	//}()
	//response.Openid = &openid
	//response.Sid = &sid
	//fmt.Println(response)
	//return nil
	if request.Code == nil {
		return fmt.Errorf("nil pointer")
	}
	type auth2SessionResponse struct {
		Openid     string `json:"openid"`
		SessionKey string `json:"session_key"`
		UnionId    string `json:"unionid"`
		Errcode    int    `json:"errcode"`
		Errormsg   string `json:"errmsg"`
	}
	loginUrl := "https://api.weixin.qq.com/sns/jscode2session?appid=" +
		url.QueryEscape(a.AppId) +
		"&secret=" + url.QueryEscape(a.AppSecret) +
		"&js_code=" + url.QueryEscape(*request.Code) +
		"&grant_type=authorization_code"
	httpClient := http.DefaultClient
	httpClient.Timeout = time.Second * 3
	httpResp, err := httpClient.Get(loginUrl)
	if err != nil {
		log.Info(err)
		return fmt.Errorf("request for wx api failed")
	}

	defer func() {
		if err := httpResp.Body.Close(); err != nil {
			log.Error(err)
		}
	}()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status got from wx server: %s", httpResp.Status)
	}

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}
	var p auth2SessionResponse
	err = json.Unmarshal(body, &p)
	if err != nil {
		return err
	}
	if p.Errcode != 0 {
		err = fmt.Errorf(p.Errormsg)
		return err
	}

	// generate unidirectly sid from openid and session_key, store {sid: {sessionkey: ..., openid:...}}
	// to redis and set the time to expire
	openid, skey := p.Openid, p.SessionKey
	sbyte := sha256.Sum256([]byte(openid + skey))
	sid := fmt.Sprintf("%x", sbyte)
	go func() {
		err = database.Set(a.RedisDb, 1*time.Hour, sid, authVal{Openid: openid, Sid: sid})
	}()

	response.Openid = &openid
	response.Sid = &sid
	return err
}

func (a Auth) PutUser(ctx context.Context, request *pb.PutUserRequest, response *pb.PutUserResponse) error {
	if request.Openid == nil || request.Name == nil {
		return fmt.Errorf("nil pointer")
	}
	db := a.SqlDb
	var user pb.GetUserResponse
	err := a.GetUser(ctx, &pb.GetUserRequest{
		Id: request.Openid,
	}, &user)
	if err != nil {
		if err.Error() == "record not found" {
			response.Err = nil
			newUser := model.User{
				Id:          *request.Openid,
				Name:        *request.Name,
				Gender:      request.Gender,
				Avatar:      request.Avatar,
				SelfGroupId: nil,
			}
			tx := db.Begin()
			// put user
			if err = tx.Create(&newUser).Error; err != nil {
				tx.Rollback()
				return err
			}
			group := model.Group{
				Name:      new(string),
				CreatorId: *request.Openid,
				Type:      1,
			}
			// put group
			if err = tx.Create(&group).Error; err != nil {
				tx.Rollback()
				return err
			}
			// add self_group_id
			if err = tx.Model(&newUser).Update("SelfGroupId", group.Id).Error; err != nil {
				tx.Rollback()
				return err
			}
			// created
			tx.Commit()
			return nil
		} else {
			return err
		}
	}
	oldUser := model.User{
		Id:     *request.Openid,
		Name:   *request.Name,
		Gender: request.Gender,
		Avatar: request.Avatar,
		SelfGroupId: user.SelfGroupId,
	}
	if request.Gender == nil {
		oldUser.Gender = user.Gender
	}
	if request.Avatar == nil {
		oldUser.Avatar = user.Avatar
	}
	if err = db.Save(&oldUser).Error; err != nil {
		return err
	}
	response.Err = new(int32)
	*response.Err = 1
	return err
}

func (a Auth) GetUser(ctx context.Context, request *pb.GetUserRequest, response *pb.GetUserResponse) error {
	if request.Id == nil {
		return fmt.Errorf("nil pointer")
	}
	db := a.SqlDb
	var user model.User
	user.Id = *request.Id
	if err := db.Find(&user).Error; err != nil {
		// perhaps "record not found"
		return err
	}
	response.Name= &user.Name
	response.SelfGroupId = user.SelfGroupId
	response.Gender = user.Gender
	response.Avatar = user.Avatar
	return nil
}

func (a Auth) CheckAuth(ctx context.Context, request *pb.CheckAuthRequest, response *pb.CheckAuthResponse) error {
	if request.Openid == nil || request.Sid == nil {
		return fmt.Errorf("nil pointer")
	}
	var val authVal
	err := database.Get(a.RedisDb, *request.Sid, &val)
	response.Ok = new(bool)
	if err != nil {
		*response.Ok = false
		return err
	}
	if val.Openid != *request.Openid {
		*response.Ok = false
		return fmt.Errorf("user id and session id don't match")
	}
	*response.Ok = true
	return nil
}
