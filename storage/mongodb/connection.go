package mongodb

import (
	"context"

	"github.com/sophed/lg"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type StorageMongo struct {
	URI       string
	Database  string
	CollUsers string
	CollNotes string
}

var client *mongo.Client

func (s *StorageMongo) Test() error {
	var err error
	client, err = mongo.Connect(options.Client().ApplyURI(s.URI))
	if err != nil {
		return err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	lg.Dbug("successfully conntected to mongodb")
	return nil
}
