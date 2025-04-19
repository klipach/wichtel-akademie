// package main

// import (
// 	"encoding/json"
// 	"fmt"
//     "os"
//     "io"
//     "strconv"
//     "net/http"
// )

// type photo struct {
//     PhotoId int `json:"PhotoId"`
// }

//  func main() {
//     photosFile, err := os.Open("photos.json")
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
//     fmt.Println(photos)
//     for i := 0; i < len(photos); i++ {
//         // fmt.Println("PhotoId: ", photos[i].PhotoId)
//         downloadPhoto(photos[i].PhotoId)
//     }
//  }


//  func downloadPhoto(photoId int) {
//     out, err := os.Create(strconv.Itoa(photoId) + ".jpg")
//     if err != nil {
//         fmt.Println("Error creating file:", err)
//         return
//     }
//     defer out.Close()

//     requestURL := fmt.Sprintf("https://apiwichtelakademie.nemborn.com//Media/GetImageByIdAndType/%d/NemData/1500", photoId)
//     req, err := http.NewRequest(http.MethodGet, requestURL, nil)
// 	if err != nil {
// 		fmt.Printf("client: could not create request: %s\n", err)
// 		os.Exit(1)
// 	}
//     req.Header.Set("accept", "application/json")
//     req.Header.Set("content-type", "application/json")
//     req.Header.Set("user-agent", "Wichtel Eltern/2.984 (dk.assemble.wichtelakademie.parents; build:1; iOS 17.5.1) Alamofire/4.9.1")
//     req.Header.Set("accept-language", "en-DE;q=1.0, uk-DE;q=0.9, ru-DE;q=0.8")
//     req.Header.Set("assemble-metadata", `{"language":"en","v":"2.984","app":"parent","platform":"ios","token":"2LSJwl5mhmTR9HJDHUYhFA%3D%3D"}`)

//     resp, err := http.DefaultClient.Do(req)
//     if err != nil {
//         fmt.Println("Error downloading image:", err)
//         return
//     }
//     defer resp.Body.Close()

//     // Check server response
//     if resp.StatusCode != http.StatusOK {
//         fmt.Println("Bad status:", resp.Status)
//         return
//     }

//     // Writer the body to file
//     _, err = io.Copy(out, resp.Body)
//     if err != nil {
//         fmt.Println("Error saving image:", err)
//         return
//     }

//     fmt.Println("Image downloaded successfully")
//  }
