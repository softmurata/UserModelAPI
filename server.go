package main

import (
	"fmt"

	"github.com/softmurata/UserModelAPI/models"
)

func main() {

	user := models.NewUser("murata", 24)
	data := models.NewDataType("0", "1", "Tokyo")
	fmt.Println(user.Name)
	fmt.Println(data.Keyword, data.Location)
}

/*

create go file except for go.mod
create github repository
push github
go get githubrepo


*/
