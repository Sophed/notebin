package mongodb

import (
	"context"
	"notebin/config"
	"notebin/data"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (s *StorageMongo) AddNote(n *data.Note) error {
	coll := client.Database(config.Get().Mongo.Database).Collection(config.Get().Mongo.CollNotes)
	_, err := coll.InsertOne(context.TODO(), n)
	return err
}

func (s *StorageMongo) RemoveNote(n *data.Note) error {
	coll := client.Database(config.Get().Mongo.Database).Collection(config.Get().Mongo.CollNotes)
	_, err := coll.DeleteOne(context.TODO(), bson.D{{
		Key:   "_id",
		Value: n.ID,
	}})
	return err
}

func (s *StorageMongo) FindNote(id string) (*data.Note, error) {
	coll := client.Database(config.Get().Mongo.Database).Collection(config.Get().Mongo.CollNotes)
	var result data.Note
	err := coll.FindOne(context.TODO(), bson.D{{
		Key:   "_id",
		Value: id,
	}}).Decode(&result)
	return &result, err
}

func (s *StorageMongo) GetNotes(u *data.User) (*[]data.Note, error) {
	coll := client.Database(config.Get().Mongo.Database).Collection(config.Get().Mongo.CollNotes)
	filter := bson.D{{
		Key:   "owner",
		Value: u.ID,
	}}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	var results []data.Note
	err = cursor.All(context.TODO(), &results)
	return &results, nil
}
