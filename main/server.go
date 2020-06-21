package main

import (
	"fmt"
	"sody.com/chat/config"
)

func main() {
	fmt.Println(config.GetEndPoint())
}
