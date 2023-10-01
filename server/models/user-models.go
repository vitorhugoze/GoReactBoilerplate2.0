package models

type Identifier interface {
	GetIdentifier() string
}

type User struct {
	User_id   string `bson:"_id"`
	User_name string `bson:"user_name"`
	User_mail string `bson:"user_mail"`
	User_pass string `bson:"user_pass"`
}

func (user *User) GetIdentifier() string {

	if user.User_name != "" {
		return user.User_name
	} else {
		return user.User_mail
	}
}
