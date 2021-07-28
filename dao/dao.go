package dao

import (
	"context"
	"github.com/afagundes/mongo-generic-dao/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type DAO struct {
	Database   string
	Collection string
}

var client *mongo.Client

// Connect Conecta ao MongoDB
func (u *DAO) Connect() {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(config.MongoUrlConnection))
	checkErr(err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	checkErr(err)
}

// Disconnect Termina a conexão
func (u *DAO) Disconnect() {
	err := client.Disconnect(context.Background())
	checkErr(err)
}

// GetAll Recupera todos os documentos da coleção
func (u *DAO) GetAll(results interface{}) {
	collection := getCollection(u)
	cur, err := collection.Find(context.Background(), bson.D{})
	checkErr(err)

	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			log.Println(err)
		}
	}(cur, context.Background())

	if err := cur.All(context.Background(), results); err != nil {
		panic(err)
	}
}

// GetById Recupera um documento por ID
func (u *DAO) GetById(id string, doc interface{}) {
	collection := getCollection(u)

	objID, _ := primitive.ObjectIDFromHex(id)
	err := collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(doc)
	checkErr(err)
}

// Insert Insere um novo documento
func (u *DAO) Insert(doc interface{}) interface{} {
	collection := getCollection(u)
	insertResult, err := collection.InsertOne(context.Background(), doc, options.InsertOne())
	checkErr(err)

	return insertResult.InsertedID
}

// Update Atualiza um documento
func (u *DAO) Update(id string, doc interface{}) *mongo.UpdateResult {
	collection := getCollection(u)

	objID, _ := primitive.ObjectIDFromHex(id)

	updateResult, err := collection.UpdateOne(
		context.Background(),
		bson.M{"_id": objID},
		bson.D{
			{"$set", doc},
		})

	checkErr(err)

	return updateResult
}

// DeleteById Deleta um documento
func (u *DAO) DeleteById(id string) int64 {
	collection := getCollection(u)

	objId, _ := primitive.ObjectIDFromHex(id)
	deleteResult, err := collection.DeleteOne(context.Background(), bson.M{"_id": objId})
	checkErr(err)

	return deleteResult.DeletedCount
}

// Função utilitária. Retorna uma struct do tipo Collection
func getCollection(u *DAO) *mongo.Collection {
	return client.Database(u.Database).Collection(u.Collection)
}

// Função utilitária. Verifica se há erros e printa no console caso haja
func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
