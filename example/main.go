package main

import (
	"fmt"

	"github.com/dbadoy/signature/openchain"
)

func main() {
	client, err := openchain.New(openchain.DefaultConfig())
	if err != nil {
		panic(err)
	}

	// [transfer(address,uint256)] <nil>
	fmt.Println(client.Signature("0xa9059cbb"))
	fmt.Println(client.Signature("a9059cbb"))
}
