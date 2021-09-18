package dao

import (
	"cloudcost-demo/model"
)

func CheckAuth(accesskey, secretkey string) (err error) {
	var auth model.Auth
	db.Where("access_key = ? AND secret_key = ? AND enabled = true", accesskey, secretkey).First(&auth)
	if err != nil {
		panic(err)
	}

	return nil
}
