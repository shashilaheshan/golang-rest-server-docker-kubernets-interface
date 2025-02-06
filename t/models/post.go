package models

type Post struct {
	ID     string `bson:"_id,omitempty"`
	Name   string `json:"name",bson:"name"`
	Author string `json:"author",bson:"author"`
	Client string `json:"client",bson:"client"`
}
