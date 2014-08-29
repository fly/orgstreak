package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type contribTime time.Time

type ContribData struct {
	date          contribTime
	contributions int
}

type Contributions []ContribData

type GHError struct {
}

const contribDateFmt = "2006-01-02"

func (cTime *contribTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	t, err := time.Parse(contribDateFmt, s)
	if err != nil {
		return err
	}
	*cTime = (contribTime)(t)
	return nil
}

func getContributions(user string, userData chan Contributions) {
	url := "https://github.com/users/" + user + "/contributions"
	var contribData Contributions

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if data != nil {
		err = json.Unmarshal(data, &contribData)
	}

	userData <- contribData
	return
}

func main() {
}
