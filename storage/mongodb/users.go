package mongodb

import (
	"context"
	"notebin/config"
	"notebin/data"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (s *StorageMongo) AddUser(u *data.User) error {
	coll := client.Database(config.Get().Mongo.Database).Collection(config.Get().Mongo.CollUsers)
	_, err := coll.InsertOne(context.TODO(), u)
	return err
}

func (s *StorageMongo) RemoveUser(u *data.User) error {
	coll := client.Database(config.Get().Mongo.Database).Collection(config.Get().Mongo.CollUsers)
	_, err := coll.DeleteOne(context.TODO(), bson.D{{
		Key:   "_id",
		Value: u.ID,
	}})
	return err
}

func (s *StorageMongo) FindUser(id string) (*data.User, error) {
	coll := client.Database(config.Get().Mongo.Database).Collection(config.Get().Mongo.CollUsers)
	var result data.User
	err := coll.FindOne(context.TODO(), bson.D{{
		Key:   "_id",
		Value: id,
	}}).Decode(&result)
	return &result, err
}
