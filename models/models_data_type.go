package models

type dataType struct {
	Rid     string
	Keyword string
}

func NewDataType(rid string, keyword string) *dataType {
	return &dataType{
		Rid:     rid,
		Keyword: keyword,
	}
}
