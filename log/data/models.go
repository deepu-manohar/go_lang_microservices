package data

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LogRepo struct {
	client *mongo.Client
}

type LogEntry struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string    `bson:"name" json:"name"`
	Data      string    `bson:"data" json:"data"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

func New(mongo *mongo.Client) LogRepo {
	return LogRepo{
		client: mongo,
	}
}

func (m *LogRepo) Insert(entry LogEntry) error {
	collection := m.client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), LogEntry{
		Name:      entry.Name,
		Data:      entry.Data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Print("failed to insert to mongo ", err)
		return err
	}
	return nil
}

func (m *LogRepo) GetAll() ([]*LogEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	collection := m.client.Database("logs").Collection("logs")
	defer cancel()
	ops := options.Find()
	ops.SetSort(bson.D{{"created_at", -1}})
	cursor, err := collection.Find(context.TODO(), bson.D{}, ops)
	if err != nil {
		log.Println("failed to fetch all records ", err)
		return nil, err
	}
	defer cursor.Close(ctx)
	var logs []*LogEntry
	for cursor.Next(ctx) {
		var item LogEntry
		err := cursor.Decode(&item)
		if err == nil {
			logs = append(logs, &item)
		} else {
			log.Println("failed to decode item ", err)
		}
	}
	return logs, nil
}

func (m *LogRepo) GetOne(id string) (*LogEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	collection := m.client.Database("logs").Collection("logs")

	docId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("error generating id ", err)
		return nil, err
	}
	var entry LogEntry
	err = collection.FindOne(ctx, bson.M{"_id": docId}).Decode(&entry)
	if err != nil {
		log.Println("error generating id ", err)
		return nil, err
	}
	return &entry, nil
}

func (m *LogRepo) DropCollection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	collection := m.client.Database("logs").Collection("logs")
	if err := collection.Drop(ctx); err != nil {
		return err
	}
	return nil
}

func (m *LogRepo) Update(logEntry LogEntry) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	collection := m.client.Database("logs").Collection("logs")
	result, err := collection.UpdateOne(ctx, bson.M{"_id": logEntry.ID},
		bson.D{
			{"$set", bson.D{
				{"name", logEntry.Name},
				{"data", logEntry.Data},
				{"updated_at", time.Now()},
			}},
		},
	)
	if err != nil {
		log.Println("failed to update record ", err)
		return nil, err
	}
	return result, nil
}
