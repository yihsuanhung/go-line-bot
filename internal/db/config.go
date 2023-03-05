package db

import (
	"fmt"
	"strconv"
)

const (
	Host     = "localhost"
	Port     = 27017
	Username = ""
	Password = ""
	Database = "line_bot"
)

func GetMongoURI() string {
	uri := "mongodb://"
	if Username != "" && Password != "" {
		uri += Username + ":" + Password + "@"
	}
	uri += Host + ":" + strconv.Itoa(Port)

	fmt.Println("URI", uri)
	return uri
}
