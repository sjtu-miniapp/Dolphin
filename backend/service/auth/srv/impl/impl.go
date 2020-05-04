package impl

import (
	"database/sql"
)

type Auth struct {
	SqlDb *sql.DB
}

func (g Auth) OnLogin(ctx context.Context, request *pb.OnLoginRequest, response *pb.OnLoginResponse) error {
	panic("implement me")
}

func (g Auth) AfterLogin(ctx context.Context, request *pb.AfterLoginRequest, response *pb.AfterLoginResponse) error {
	panic("implement me")
}

func (g Auth) GetUser(ctx context.Context, request *pb.GetUserRequest, response *pb.GetUserResponse) error {
	panic("implement me")
}

func (g Auth) CheckAuth(ctx context.Context, request *pb.CheckAuthRequest, response *pb.CheckAuthResponse) error {
	panic("implement me")
}

//func (g Auth) GetAuth(ctx context.Context, request *pb.GetAuthRequest, response *pb.GetAuthResponse) error {
//	db := g.SqlDb
//
//	sql1 := fmt.Sprintf("SELECT name FROM `auth` WHERE `id` = %d", request.Id)
//	log.Println(sql1)
//	rows, err := db.QueryContext(ctx, sql1)
//	if err != nil {
//		return err
//	}
//	if rows.Next() {
//		err = rows.Scan(&response.Name)
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}

//type AuthService struct {
//	Db        *sql.DB
//	AppId     string
//	AppSecret string
//}
//
//type auth2SessionResponse struct {
//	Openid     string `json:"openid"`
//	SessionKey string `json:"session_key"`
//	UnionId    string `json:"unionid"`
//	Errcode    int    `json:"errcode"`
//	Erromsg    string `json:"errmsg"`
//}
//
//// openid, sid
//// acquire for session_key
//func (as AuthService) OnLogin(ctx context.Context, code string) (string, string, error) {
//	loginUrl := "https://api.weixin.qq.com/sns/jscode2session?appid=" +
//		url.QueryEscape(as.AppId) +
//		"&secret=" + url.QueryEscape(as.AppSecret) +
//		"&js_code=" + url.QueryEscape(code) +
//		"&grant_type=authorization_code"
//	httpClient := http.DefaultClient
//	httpClient.Timeout = time.Second * 3
//	httpResp, err := httpClient.Get(loginUrl)
//	if err != nil {
//		return "", "", err
//	}
//	defer httpResp.Body.Close()
//	if httpResp.StatusCode != http.StatusOK {
//		return "", "", fmt.Errorf("http.Status: %s", httpResp.Status)
//	}
//	// read the payload, in this case, Jhon's info
//	body, err := ioutil.ReadAll(httpResp.Body)
//	if err != nil {
//		return "", "", err
//	}
//	var p auth2SessionResponse
//	err = json.Unmarshal(body, &p)
//	if err != nil {
//		return "", "", err
//	}
//	if p.Errcode != 0 {
//		err = fmt.Errorf(p.Erromsg)
//		return "", "", err
//	}
//	// TODO:
//	// generate unidirectly sid from openid and session_key, store {sid: {sessionkey: ..., openid:...}}
//	// to redis and set the time to expire
//	sid := ""
//	//
//
//	// return openid and sid
//	return p.Openid, sid, err
//}
//
//// err
//// prepare the user into the database
//// insert id, name, gender
//func (as AuthService) AfterLogin(ctx context.Context, id, name string, gender uint32, sid string) (int, error) {
//	if ok, err := as.CheckAuth(ctx, id, sid); !ok || err != nil {
//		return -1, fmt.Errorf("auth check failed")
//	}
//	err := db.InsertUser(ctx, as.Db, id, name, gender)
//	if err != nil {
//		return -2, err
//	}
//	return 0, nil
//}
//
//// ok
//// TODO: retrieve value of session_key and openid, check id equals openid
//func (as AuthService) CheckAuth(ctx context.Context, id, sid string) (bool, error) {
//	return true, nil
//}
//
//// name, gender
//func (as AuthService) GetUser(ctx context.Context, id string) (string, uint32, error) {
//	user, err := db.GetUser(ctx, as.Db, id)
//	if err != nil {
//		return "", 0, err
//	}
//	return user.Name, uint32(user.Gender), err
//}
