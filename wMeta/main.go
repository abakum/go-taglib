package main

import (
	"fmt"
	"os"

	"github.com/abakum/go-taglib"
)

func main() {
	err := taglib.WriteTags(os.Args[1], map[string][]string{
		// Multi-valued tags allowed
		taglib.AlbumArtist: {"David Bynre", "Brian Eno"},
		taglib.Album:       {"My Life in the Bush of Ghosts"},
		taglib.TrackNumber: {"1"},

		// Non-standard allowed too
		"ALBUMARTIST_CREDIT": {"Brian Eno & David Bynre"},
	}, 0)
	fmt.Println(err)
}
