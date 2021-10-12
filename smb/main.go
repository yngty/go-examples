package main

import (
	"fmt"

	"github.com/mvo5/libsmbclient-go"
)

func authHandle(serverName, shareName string) (domain, username, password string) {

	return "WORKGROUP", "test@123", "1"
}
func main() {
	client := libsmbclient.New()
	client.SetAuthCallback(authHandle)
	dh, err := client.Opendir("smb://127.0.0.1/123")
	if err != nil {
		panic(err)
	}
	defer dh.Closedir()
	for {
		dirent, err := dh.Readdir()
		if err != nil {
			break
		}
		fmt.Println(dirent)
	}
}
