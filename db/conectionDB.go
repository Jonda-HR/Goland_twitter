package db

import (
	"context"
	"fmt"

	"github.com/Jonda-HR/Goland_twitter/v2/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN *mongo.Client
var DatabaseName string

func ConectionDB(ctx context.Context) error {
	user := ctx.Value(models.Key("user")).(string)
	passwd := ctx.Value(models.Key("password")).(string)
	host := ctx.Value(models.Key("host")).(string)
	connStr := fmt.Sprintf("mongodb+srv://%s:%s@%s?retryWrites=true&w=majority", user, passwd, host)

	var clientOption = options.Client().ApplyURI(connStr)

	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("SUCCESSFUL CONNECTION")

	MongoCN = client
	DatabaseName = ctx.Value(models.Key("database")).(string)

	return nil
}

func IsConnected() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	return err == nil
}