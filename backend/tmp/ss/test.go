package main

import (
	"github.com/sjtu-miniapp/dolphin/service/database"
	"github.com/sjtu-miniapp/dolphin/service/database/model"
	"github.com/sjtu-miniapp/dolphin/service/task/pb"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func GetTaskContent() error {
	coll := g.MongoDb.Collection(collection)
	var content model.TaskContent
	err := coll.FindOne(context.TODO(),
		bson.M{"task_id": *request.Id, "version": *request.Version}).Decode(&content)
	response.Content = &content.Content
	now := time.Now().String()
	response.CreatedAt = &now
	// TODO
	//response.Diff = content.
	response.Modifier = content.Modifier
	response.UpdatedAt = &now
	return err
}

func createContent(g Task, taskId int32) error {
	coll := g.MongoDb.Collection(collection)
	_, err := coll.InsertOne(context.TODO(),
		bson.M{"task_id": taskId, "version": "1", "content": "",
			"modifier": bson.A{}, "updated_at": "", "created_at": "", "diff": ""})
	return err
}

//func (g Task) update
func (g Task) UpdateTaskContent(ctx context.Context, request *pb.UpdateTaskContentRequest, response *pb.UpdateTaskContentResponse) error {
	coll := g.MongoDb.Collection(collection)
	_, err := coll.UpdateOne(context.TODO(),
		bson.M{"task_id": *request.Id, "version": *request.Version},
		bson.M{"$set": bson.M{"content": *request.Content}})
	return err
}


func main() {
	redisdb, _ := database.InitRedis("121.199.33.44", "610878")
	database.Set(redisdb, 0, "test", authVal{Openid: "test", Sid: "test"})

}
