package main

import (
	"fmt"
	"github.com/mengshi02/etcd-defrag"
)

func main() {
	err := etcd_defrag.Run(
		[]string{
			"https://127.0.0.1:4001",
		},
		"./certificate/service.cert",
		"./certificate/service.key",
		"./certificate/ca.cert",
	)
	if err != nil {
		fmt.Printf("clean failure: %s\n", err.Error())
	}
	fmt.Println("clean success")
}
