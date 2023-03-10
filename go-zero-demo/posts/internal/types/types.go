// Code generated by goctl. DO NOT EDIT.
package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Request struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
}

type Response struct {
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Comment     []Comment          `json:"comment,omitempty" bson:"comment,omitempty" default:"[]"`
}

type Post struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Comment     []Comment          `json:"comment,omitempty" bson:"comment,omitempty" default:"[]"`
}

type (
	NewPost struct {
		ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
		Title       string    `json:"title,omitempty" bson:"title,omitempty"`
		Description string    `json:"description,omitempty" bson:"description,omitempty"`
		Comment     []Comment `json:"comment,omitempty" bson:"comment,omitempty" default:"[]"`
	}
	Comment struct {
		ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
		Content string             `json:"content,omitempty" bson:"content,omitempty"`
	}

)

