package astruct

import "fmt"

type authenticationInfo struct {
	username string
	password string
}


func (authI authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf(
		"Authorization : Basic %s:%s",
		authI.username,
		authI.password,
	)
}

func assignS(){

	auth := authenticationInfo{
		username: "user",
		password: "pass",
	}

	fmt.Println(authenticationInfo.getBasicAuth(auth))

}