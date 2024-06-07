package main

import (
	"fmt"
	dingo "github.com/DIN-center/din-sc/apps/din-go/pkg/dinregistry"
)

func main() {
	fmt.Println(dingo.GetDINRegistry())
}
