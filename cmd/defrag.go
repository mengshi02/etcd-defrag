package main

import (
	"fmt"
	"github.com/mengshi02/etcd-defrag"
)

func main() {
	err := etcd_defrag.Clean(
		[]string{
			"https://",
		},
		"./certificate/cert.",
		"./certificate/cert.",
		"./certificate/cert.",
	)
	if err != nil {
		fmt.Printf("clean failure: %s\n", err.Error())
	}
	fmt.Println("clean success")
}
