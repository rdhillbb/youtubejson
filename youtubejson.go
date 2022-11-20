//package main
package youtubejson

import "fmt"
import "encoding/json"
import "strings"

type YoutubeUrl struct {
	Ytuburl string `json:"url"`
}

func GetYubeKey(k string) string {
	keyIdx := strings.Index(k[1:], "=") + 2
	return k[keyIdx:]
}

func CreateStrk(b []byte) YoutubeUrl {
	var s YoutubeUrl
	json.Unmarshal(b, &s)
	return s

}
func GetJsonBytes(s YoutubeUrl) []byte {
	u, _ := json.Marshal(s)
	return u

}
func main() {
	//var testUrl = "https://www.youtube.com/watch?v=r8dfQja43Ik"
	var jsonYoutube YoutubeUrl
	jsonYoutube.Ytuburl = "https://www.youtube.com/watch?v=r8dfQja43Ik"
	u, _ := json.Marshal(jsonYoutube)
	fmt.Printf("\nGetJsonStrk called :\n input(%s) \n output: %s\n ",string(u),CreateStrk(u))

	b := GetJsonBytes(CreateStrk(u))

        fmt.Printf("\nGetJsonBytes called Input :%s\n ", CreateStrk(u)	)
        fmt.Printf("\nGetJsonBytes calledi Output :%s\n\n ", string(b))
	json.Unmarshal(u, &jsonYoutube)
	fmt.Printf("\nGet Youtube Key: %s\n\n",GetYubeKey(jsonYoutube.Ytuburl))

}
