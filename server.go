package main

import (
	"fmt"

	"github.com/softmurata/UserModelAPI/models"
)

func main() {

	user := models.NewUser("murata", 24)
	var data models.dataType
	fmt.Println(user.Name)
}

/*

create go file except for go.mod
create github repository
push github
go get githubrepo


*/