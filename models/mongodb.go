package models

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserSettingModel struct {
	Collection *mongo.Collection
}

func (us UserSettingModel) CreateDefaultSettings(userID int64) (interface{}, error) {
	defaultSettings := &QRSettings{}
	defaultSettings.SetDefault(userID)

	result, err := us.Collection.InsertOne(
		context.TODO(),
		defaultSettings,
	)
	return result.InsertedID, err
}

//func toDoc(v interface{}) (doc *bson.Document, err error) {
//	data, err := bson.Marshal(v)
//	if err != nil {
//		return
//	}
//
//	err = bson.Unmarsh
//	al(data, &doc)
//	return
//}

//func (us UserSettingModel) tmp() {
//
//	db, err := database.Open()
//	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
//
//}
