package twilio

import (
	"github.com/hetiansu5/urlquery"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type SmsResponseStruct struct {
	ToCountry     string
	ToState       string
	SmsMessageSid string
	NumMedia      string
	ToCity        string
	FromZip       string
	SmsSid        string
	FromState     string
	SmsStatus     string
	FromCity      string
	Body          string
	FromCountry   string
	To            string
	ToZip         string
	NumSegments   string
	MessageSid    string
	AccountSid    string
	From          string
	ApiVersion    string
}

func SmsWebhookHandler(w http.ResponseWriter, r *http.Request) {
	logger := log.New(os.Stdout, "RoL", log.LstdFlags|log.Lshortfile)
	var smsResponse SmsResponseStruct

	if r.Method == "POST" {
		rawBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger.Printf("There was an error getting the request body: %v", err)
		}

		urlquery.Unmarshal(rawBody, &smsResponse)
		ProcessMessage(smsResponse, logger)

	}
}

func ProcessMessage(response SmsResponseStruct, logger *log.Logger) {
	logger.Print(response.Body)
}
