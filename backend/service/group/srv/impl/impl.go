package impl

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sjtu-miniapp/dolphin/service/group/pb"
	"log"
)


type Group struct {
	SqlDb *sql.DB
}

func (g Group) GetGroup(ctx context.Context, request *pb.GetGroupRequest, response *pb.GetGroupResponse) error {
	db := g.SqlDb

	sql1 := fmt.Sprintf("SELECT name FROM `group` WHERE `id` = %d", request.Id)
	log.Println(sql1)
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
