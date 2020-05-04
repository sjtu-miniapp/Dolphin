package impl

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sjtu-miniapp/dolphin/service/task/pb"
	"log"
)


type Task struct {
	SqlDb *sql.DB
}

func (g Task) GetGroup(ctx context.Context, request *pb.GetGroupRequest, response *pb.GetGroupResponse) error {
	db := g.SqlDb

	sql1 := fmt.Sprintf("SELECT name FROM `task` WHERE `id` = %d", request.Id)
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
