package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Amazon struct {
	ID      primitive.ObjectID `json:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Seasons string             `json:"seasons,omitempty"`
	Genre   string             `json:"genre,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}
