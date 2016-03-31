package main

import (
	"fmt"

	"github.com/mattermost/platform/model"
)

func main() {

	uriScheme := "https://"

	type Credentials struct {
		Login  string
		Team   string
		Pass   string
		Server string
		NoTLS  bool
	}

	cred := Credentials{}

	var myinfo *model.Result
	var appErr *model.AppError

	cred.Login = "REPLACE"
	cred.Pass = "REPLACE"
	cred.Team = "Kainos"
	cred.Server = "chatops.kainos.io/"
	cred.NoTLS = true

	client := model.NewClient(uriScheme + cred.Server)
	myinfo, appErr = client.LoginByEmail(cred.Team, cred.Login, cred.Pass)

	if appErr != nil {
		fmt.Println("Here")
		fmt.Println(appErr)
	}

	fmt.Println(myinfo)
}
