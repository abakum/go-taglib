package main

import (
	"fmt"
	"os"

	"github.com/abakum/go-taglib"
)

func main() {
	tags, err := taglib.ReadTags(os.Args[1])
	// check(err)

	fmt.Printf("tags: %v %v\n", tags, err) // map[string][]string
	if err != nil {
		return
	}
	fmt.Printf("AlbumArtist: %q\n", tags[taglib.AlbumArtist])
	fmt.Printf("Album: %q\n", tags[taglib.Album])
	fmt.Printf("TrackNumber: %q\n", tags[taglib.TrackNumber])
}
