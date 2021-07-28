package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Usuario struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nome string `bson:"nome" json:"nome,omitempty"`
	Idade int `bson:"idade" json:"idade,omitempty"`
	Bio string `bson:"bio" json:"bio,omitempty"`
	Foto string `bson:"foto" json:"foto,omitempty"`
}
