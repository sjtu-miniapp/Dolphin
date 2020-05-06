package impl

import (
	"context"
	"database/sql"
	"github.com/sjtu-miniapp/dolphin/service/group/pb"
)


type Group struct {
	SqlDb *sql.DB
}

func (g Group) GetGroup(ctx context.Context, request *pb.GetGroupRequest, response *pb.GetGroupResponse) error {
	db := g.SqlDb
	rows, err := db.QueryContext(ctx, sql1)
	if err != nil {
		return err
	}
	if rows.Next() {
		err = rows.Scan(&response.Name)
		if err != nil {
			return err
		}
	}
	return nil
}
