package main

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/bogem/id3v2"
)

// WriteTags maps an array of args to the file path as ID3v2 tags
func WriteTags(path string, tags [][]string) {
	file, err := id3v2.Open(path, id3v2.Options{Parse: true})
	if err != nil {
		log.Fatal("Error while opening file: ", err)
	}
	defer file.Close()

	for _, tag := range tags {
		switch tag[0] {
		case "title":
			file.SetTitle(tag[1])

		case "artist":
			file.SetArtist(tag[1])

		case "lyrics":
			lyrics := id3v2.UnsynchronisedLyricsFrame{
				Encoding: id3v2.EncodingUTF8,
				Lyrics:   tag[1],
			}
			file.AddUnsynchronisedLyricsFrame(lyrics)

		case "cover":
			pic, err := ioutil.ReadFile(tag[1])
			if err != nil {
				log.Fatal("Error while reading image: ", err)
			}
			cover := id3v2.PictureFrame{
				Encoding:    id3v2.EncodingUTF8,
				MimeType:    "image/" + filepath.Ext(tag[1]),
				PictureType: id3v2.PTFrontCover,
				Description: "Front cover",
				Picture:     pic,
			}
			file.AddAttachedPicture(cover)

		case "album":
			file.SetAlbum(tag[1])

		default:
		}
	}

	if err = file.Save(); err != nil {
		log.Fatal("Error while saving tags: ", err)
	}
}

// func ReadTags(path string) {
// 	file, err := id3v2.Open(path, id3v2.Options{Parse: true})
// 	if err != nil {
// 		log.Fatal("Error while opening file: ", err)
// 	}
// 	defer file.Close()

// 	metadata:= make(map[string]string)
// 	metadata["title"] = file.Title()
// 	metadata["artist"] = file.Artist()
// 	metadata["album"]  = file.Album()
// 	metadata["lyrics"] = string(file.GetFrames(file.CommonID("APIC"))[0].(id3v2.PictureFrame).Picture)

// }
