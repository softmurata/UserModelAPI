package models

type dataType struct {
	Rid      string
	Keyword  string
	Location string
}

func NewDataType(rid string, keyword string, loc string) *dataType {
	return &dataType{
		Rid:      rid,
		Keyword:  keyword,
		Location: loc,
	}
}
