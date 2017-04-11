// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	luis "github.com/kkdai/luis"
	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client
var luisAction *LuisAction
var allIntents *luis.IntentListResponse
var currentUtterance string

var apiURL string = "http://107.167.183.27:3000/api/v1/tf-image/"

//TFResponse :
type TFResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.ImageMessage:
				HandleImage(message, event.ReplyToken)

			case *linebot.TextMessage:
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Hi, Just upload animal photo to me.")).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}

//HandleImage :
func HandleImage(message *linebot.ImageMessage, replyToken string) error {
	content, err := bot.GetMessageContent(message.ID).Do()
	if err != nil {
		log.Println("Get msg err:", err)
		return err
	}
	defer content.Content.Close()
	log.Printf("Got file: %s", content.ContentType)

	file, err := ioutil.TempFile("/tmp", "")
	if err != nil {
		log.Println("Create tmp error:", err)
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, content.Content)
	if err != nil {
		log.Println("Copy tmp error:", err)
		return err
	}
	log.Printf("Saved to %s", file.Name())

	repBody, err := PredictContent(file.Name())
	if err != nil {
		log.Println("Save file err:", err)
		return err
	}

	//unmarshall result
	var tfRet TFResponse
	if err := json.Unmarshal(repBody, &tfRet); err != nil {
		log.Print(err)
	}

	if _, err = bot.ReplyMessage(replyToken, linebot.NewTextMessage(tfRet.Message)).Do(); err != nil {
		log.Print(err)
		return err
	}
	return nil
}

// PredictContent :
func PredictContent(filename string) ([]byte, error) {

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fw, err := w.CreateFormFile("upload", filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	fh, err := os.Open(filename)
	if err != nil {
		log.Println("error opening file")
		return nil, err
	}

	if _, err = io.Copy(fw, fh); err != nil {
		log.Println(err)
		return nil, err
	}

	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	if err = w.Close(); err != nil {
		log.Println(err)
		return nil, err
	}

	log.Printf("Total file length: %d \n", b.Len())
	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", apiURL, &b)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())
	log.Println("data content type:", w.FormDataContentType())

	// Submit the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Check the response
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
		log.Println(err)
	}

	defer res.Body.Close()
	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	log.Println("Successed: ", string(respBody))
	return respBody, nil
}
