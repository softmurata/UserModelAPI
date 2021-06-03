package models

type mongoDataType struct {
	Rid     string `bson:"rid"`
	Keyword string `bson:"keyword"`
}
