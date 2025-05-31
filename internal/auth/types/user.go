package types

type User struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
	Phone    string `bson:"phone"`
}
