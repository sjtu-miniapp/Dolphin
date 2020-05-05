package impl

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"github.com/sjtu-miniapp/dolphin/service/database"
	//"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sjtu-miniapp/dolphin/service/auth/pb"
	//"io/ioutil"
	//"net/http"
	//"net/url"
	"time"
)

type Auth struct {
	SqlDb     *sql.DB
	RedisDb   *redis.Client
	AppId     string
	AppSecret string
}

type authVal struct {
	Openid string `json:"openid"`
	Sid    string `json:"sid"`
}

type User struct {
	Openid      string `json:"openid"`
	Name        string `json:"name"`
	Gender      int32   `json:"gender"`
	Avatar      string `json:"avatar"`
	SelfGroupId sql.NullInt32    `json:"self_group_id"`
}

// openid, sid
// acquire for session_key
func (a Auth) OnLogin(ctx context.Context, request *pb.OnLoginRequest, response *pb.OnLoginResponse) error {
	// for test
	openid, skey := request.Code, ""
	sbyte := sha256.Sum256([]byte(openid + skey))
	sid := fmt.Sprintf("%x", sbyte)

	go func() {
		database.Set(a.RedisDb, 1*time.Hour, sid, authVal{Openid: openid, Sid: sid})
	}()
	response.Openid = openid
	response.Sid = sid
	return nil

	//
	//type auth2SessionResponse struct {
	//	Openid     string `json:"openid"`
	//	SessionKey string `json:"session_key"`
	//	UnionId    string `json:"unionid"`
	//	Errcode    int    `json:"errcode"`
	//	Errormsg   string `json:"errmsg"`
	//}
	//loginUrl := "https://api.weixin.qq.com/sns/jscode2session?appid=" +
	//	url.QueryEscape(a.AppId) +
	//	"&secret=" + url.QueryEscape(a.AppSecret) +
	//	"&js_code=" + url.QueryEscape(request.Code) +
	//	"&grant_type=authorization_code"
	//httpClient := http.DefaultClient
	//httpClient.Timeout = time.Second * 3
	//httpResp, err := httpClient.Get(loginUrl)
	//if err != nil {
	//	return err
	//}
	//
	//defer httpResp.Body.Close()
	//if httpResp.StatusCode != http.StatusOK {
	//	return fmt.Errorf("http.Status got from wx server: %s", httpResp.Status)
	//}
	//
	//body, err := ioutil.ReadAll(httpResp.Body)
	//if err != nil {
	//	return err
	//}
	//var p auth2SessionResponse
	//err = json.Unmarshal(body, &p)
	//if err != nil {
	//	return err
	//}
	//if p.Errcode != 0 {
	//	err = fmt.Errorf(p.Errormsg)
	//	return err
	//}
	//
	//// generate unidirectly sid from openid and session_key, store {sid: {sessionkey: ..., openid:...}}
	//// to redis and set the time to expire
	//openid, skey := p.Openid, p.SessionKey
	//sbyte := sha256.Sum256([]byte(openid + skey))
	//sid := fmt.Sprintf("%x", sbyte)
	//go func() {
	//	err = database.Set(a.RedisDb, 1*time.Hour, sid, authVal{Openid: openid, Sid: sid})
	//}()
	//response.Openid = openid
	//response.Sid = sid
	//return err
}

func (a Auth) PutUser(ctx context.Context, request *pb.PutUserRequest, response *pb.PutUserResponse) error {
	db := a.SqlDb
	var user pb.GetUserResponse
	err := a.GetUser(ctx, &pb.GetUserRequest{
		Id: request.Openid,
	}, &user)
	if err != nil {
		if err.Error() == "not found" {
			response.Err = 0
			rows, err := db.ExecContext(ctx, "INSERT INTO `user`(`id`, `avatar`, `gender`, `name`) VALUES(?, ?, ?, ?)",
				request.Openid, request.Avatar, request.Gender, request.Name)
			if err != nil {
				return err
			}
			n, _ := rows.RowsAffected()
			if n == 1 {
				rows, err = db.ExecContext(ctx, "INSERT INTO `group` (`creator_id`, `type`) VALUES(?, 1)", request.Openid)
				if err != nil {
					return err
				}
				gid, err := rows.LastInsertId()
				if err != nil {
					return err
				}
				_, err = db.ExecContext(ctx, "UPDATE `user` SET `self_group_id` = ? WHERE `id` = ?", gid, request.Openid)
				if err != nil {
					return err
				}
				// created
				response.Err = 0
				return nil
			} else {
				return fmt.Errorf("not inserted")
			}
		} else {
			return err
		}
	}
	_, err = db.ExecContext(ctx, "UPDATE `user` SET `avatar` = ?, `gender` = ?, `name` = ? WHERE `id` = ?",
		request.Avatar, request.Gender, request.Name, request.Openid)
	if err == nil {
		response.Err = 1
	}
	return err
}

func (a Auth) GetUser(ctx context.Context, request *pb.GetUserRequest, response *pb.GetUserResponse) error {
	db := a.SqlDb
	rows, err := db.QueryContext(ctx, "SELECT `name`, `gender`, `self_group_id`, `avatar` FROM `user` WHERE `id` = ?", request.Id)
	if err != nil {
		return err
	}
	var sgid sql.NullInt64
	if rows.Next() {
		err = rows.Scan(&response.Name, &response.Gender, &sgid, &response.Avatar)
		if sgid.Valid {
			response.SelfGroupId = uint32(sgid.Int64)
		}
	} else {
		return fmt.Errorf("not found")
	}
	return err
}

func (a Auth) CheckAuth(ctx context.Context, request *pb.CheckAuthRequest, response *pb.CheckAuthResponse) error {
	var val authVal
	err := database.Get(a.RedisDb, request.Sid, &val)
	if err != nil {
		response.Ok = false
		return err
	}
	if val.Openid != request.Openid {
		response.Ok = false
		return fmt.Errorf("user id and session id don't match")
	}
	response.Ok = true
	return nil
}
