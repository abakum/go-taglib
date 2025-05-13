package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/abakum/go-taglib"
)

func main() {
	tags, err := taglib.ReadTags(os.Args[1])
	if err != nil {
		fmt.Println(os.Args, err)
		return
	}
	for k, v := range tags {
		fmt.Println(k, strings.Join(v, "/"))
	}
}
