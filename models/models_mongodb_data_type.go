package models

type mongoDataType struct {
	Rid     string `bson:"rid"`
	Keyword string `bson:"keyword"`
}

func NewMongoDataType(rid string, keyword string) *mongoDataType {
	return &mongoDataType{
		Rid:     rid,
		Keyword: keyword,
	}
}
