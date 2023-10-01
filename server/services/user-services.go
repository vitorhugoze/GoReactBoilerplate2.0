package services

import (
	"context"
	"errors"
	"server/models"
	"server/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
strict parameter determines if user object requires both user_name and user_pass to proceed checking
*/
func GetUser(loginUsr *models.User, strict bool) (*models.User, error) {

	var dbUsr models.User

	var username, usermail string

	if strict {
		username = loginUsr.User_name
		usermail = loginUsr.User_mail
	} else {
		username = loginUsr.GetIdentifier()
		usermail = loginUsr.GetIdentifier()
	}

	err := utils.GetMongoData().Collection("users").FindOne(context.TODO(), bson.D{
		{Key: "$or", Value: bson.A{
			bson.D{{Key: "user_name", Value: username}},
			bson.D{{Key: "user_mail", Value: usermail}},
		}}},
	).Decode(&dbUsr)

	if err != nil {
		return nil, err
	}

	if err = utils.CompareHash(loginUsr.User_pass, dbUsr.User_pass); err != nil {
		return nil, err
	}

	return &dbUsr, nil
}

func CreateUser(user *models.User) error {

	if usr, _ := GetUser(user, true); usr != nil {
		return errors.New("user already exists")
	}

	passHash, err := utils.CreateHash(user.User_pass)
	if err != nil {
		return err
	}

	res, err := utils.GetMongoData().Collection("users").InsertOne(context.TODO(), bson.D{
		{Key: "user_name", Value: user.User_name},
		{Key: "user_mail", Value: user.User_mail},
		{Key: "user_pass", Value: passHash},
	})
	if err != nil {
		return err
	}

	id, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return errors.New("error decoding insert result")
	}

	user.User_id = id.Hex()

	return nil
}

func UpdateUser(user *models.User) error {

	var update = bson.M{"$set": bson.M{}}

	if user.User_name != "" {
		update["$set"].(bson.M)["user_name"] = user.User_name
	}
	if user.User_mail != "" {
		update["$set"].(bson.M)["user_mail"] = user.User_mail
	}
	if user.User_pass != "" {
		if hashedPass, err := utils.CreateHash(user.User_pass); err != nil {
			return err
		} else {
			update["$set"].(bson.M)["user_pass"] = hashedPass
		}
	}

	oid, err := primitive.ObjectIDFromHex(user.User_id)
	if err != nil {
		return err
	}

	res, err := utils.GetMongoData().Collection("users").UpdateOne(context.TODO(), bson.M{"_id": oid}, update)

	if err != nil {
		return err
	}

	if res.ModifiedCount == 0 {
		return errors.New("no document was updated")
	}

	return nil
}

func DeleteUser(user *models.User) error {

	oid, err := primitive.ObjectIDFromHex(user.User_id)
	if err != nil {
		return err
	}

	res, err := utils.GetMongoData().Collection("users").DeleteOne(context.TODO(), bson.D{
		{Key: "_id", Value: oid},
	})

	if res.DeletedCount == 0 && err == nil {
		return errors.New("user not deleted")
	}

	return err
}
