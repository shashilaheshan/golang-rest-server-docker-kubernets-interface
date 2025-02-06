package main

import (
	"books-api/t/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostRepo struct {
	client      string
	mongoClient *mongo.Client
	logger      Logger
}

func NewPost(client string, mongo *mongo.Client, logger Logger) PostRepoInterface {
	return &PostRepo{client: client, mongoClient: mongo, logger: logger}
}

func (p *PostRepo) GetPosts() (*[]models.Post, error) {
	collection := p.mongoClient.Database(DB).Collection(BOOKS_COLLECTION)

	filter := bson.M{}

	options := options.Find()

	cursor, err := collection.Find(context.Background(), filter, options)

	if err != nil {
		p.logger.LogMessage(err.Error())

		return nil, err
	}

	defer cursor.Close(context.Background())

	var posts []models.Post

	for cursor.Next(context.Background()) {
		var post models.Post

		if err := cursor.Decode(&post); err != nil {
			p.logger.LogMessage(err.Error())
			continue
		}

		posts = append(posts, post)
	}
	if err := cursor.Err(); err != nil {
		p.logger.LogMessage(err.Error())
		return nil, err
	}

	return &posts, nil

}

func (p *PostRepo) AddPost(post *models.Post) (bool, error) {
	collction := p.mongoClient.Database(DB).Collection(BOOKS_COLLECTION)

	postDoc := bson.M{
		"name":   post.Name,
		"author": post.Author,
		"client": post.Client,
	}
	collction.InsertOne(context.Background(), postDoc)

	return true, nil
}

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) LogMessage(message string) {
	log.Println(message)
}
