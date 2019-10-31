package places

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

const (
	// dd.mm.yyyy
	dateRegExp = `^([0-2][0-9]|(3)[0-1])(\.)(((0)[0-9])|((1)[0-2]))(\.)\d{4}$`
	urlFormat  = "https://618.by/api/v2/route/schedule?id_route=%d&date=%s"
)

type response struct {
	Schedule []scheduleItem `json:"schedule"`
}

type scheduleItem struct {
	Count int    `json:"count"`
	Time  string `json:"time"`
}

func ValidateDate(date *string) error {
	if date == nil || *date == "" {
		return fmt.Errorf("Error: The date is empty.")
	}
	matched, err := regexp.MatchString(dateRegExp, *date)
	if err != nil {
		return err
	}
	if !matched {
		return fmt.Errorf("Error: Date '%s' doesn't match dd.mm.yyyy format.", *date)
	}

	return nil
}

// CheckPlaces sends GET requests to the server and check received JSON
func CheckPlaces(date string, route uint, interval uint) error {
	client := http.Client{
		Timeout: time.Second * 5,
	}

	url := fmt.Sprintf(urlFormat, route, date)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	// initialize buffer for sound (notifications)
	beepBuffer, err := getBeepBuffer()
	if err != nil {
		return err
	}

	log.Println("Start check places...")
	log.Println("URL: ", url)

	for {
		log.Println(time.Now())
		res, err := client.Do(req)
		if err != nil {
			log.Println(err)
			continue
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			continue
		}
		defer res.Body.Close()

		resp := response{}
		if err := json.Unmarshal(body, &resp); err != nil {
			log.Println(err)
			continue
		}

		if err := processResponse(resp, beepBuffer); err != nil {
			log.Println(err)
			continue
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func processResponse(resp response, beepBuffer *beep.Buffer) error {
	for _, s := range resp.Schedule {
		if s.Count != 0 {
			log.Printf("Empty place: time: %s, amount: %d.\n", s.Time, s.Count)

			shot := beepBuffer.Streamer(0, beepBuffer.Len())
			speaker.Play(shot)
		}
	}
	return nil
}

func getBeepBuffer() (*beep.Buffer, error) {
	f, err := os.Open("../files/beep.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	if err := speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10)); err != nil {
		return nil, err
	}

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	if err := streamer.Close(); err != nil {
		return nil, err
	}

	return buffer, err
}
