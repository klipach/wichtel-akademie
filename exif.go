// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"os"

//     // exif "github.com/dsoprea/go-exif/v3"
//     // exifcommon "github.com/dsoprea/go-exif/v3/common"
//     jis "github.com/dsoprea/go-jpeg-image-structure/v2"
// )

// type photo struct {
//     PhotoId      int    `json:"PhotoId"`
//     FileName     string `json:"FileName"`
//     Description  string `json:"Description"`
//     Date         string `json:"Date"`
// }

// func main() {
// 	photosFile, err := os.Open("photos.json")
//     if err != nil {
//         fmt.Println(err)
//     }
//     defer photosFile.Close()

//     photosBytes, err := io.ReadAll(photosFile)
//     if err != nil {
//         fmt.Println(err)
//     }

//     var photos []photo

//     err = json.Unmarshal(photosBytes, &photos)
//     if err != nil {
//         fmt.Println(err)
//     }
//     // fmt.Println(photos)
//     for _, p := range photos {
//         ec, err := jis.NewJpegMediaParser().ParseFile(fmt.Sprint("photos/", p.PhotoId))
//         if err != nil {
//             fmt.Println(err)
//         }
//         fmt.Println(ec.Exif())
//         panic("stop")
//     }
// }
