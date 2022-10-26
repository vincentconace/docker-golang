package task

import (
	"context"
	"time"

	"github.com/vincentconace/golang-docker/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Text      string             `bson:"text"`
	Completed bool               `bson:"completed"`
}

var ctx context.Context

func SaveTask(t *Task) (Task, error) {
	taskB := bson.M{
		"created_at": t.CreatedAt,
		"updated_at": t.UpdatedAt,
		"text":       t.Text,
		"completed":  t.Completed,
	}

	result, err := db.Collection.InsertOne(ctx, taskB)
	if err != nil {
		return Task{}, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	t.ID = objID

	return *t, nil
}
