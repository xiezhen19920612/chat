package config

import "fmt"

const Ip string = "0.0.0.0"
const Port int = 3333

func GetEndPoint() string {
	return fmt.Sprintf("%s:%d", Ip, Port)
}
