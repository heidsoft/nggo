package main

import (
	"fmt"
	"time"
	"os"
	"encoding/json"
)

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04"))
	return []byte(stamp), nil
}

func main() {
	fmt.Println(time.Now().UTC().Format("2006-01-02T15:04:05-0700"))

	t, err := time.Parse(time.RFC3339Nano, "2013-06-05T14:10:43.678Z")


	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	fmt.Println(time.Now().Format("2006-01-02 15:04"))

	var jsontime JSONTime
	jsontime = time.Now()
	fmt.Println(jsontime)

	type MyUser struct {
		ID       int64     `json:"id"`
		Name     string    `json:"name"`
		LastSeen time.Time `json:"lastSeen"`
	}

	_ = json.NewEncoder(os.Stdout).Encode(
		&MyUser{1, "Ken", time.Now()},
	)
}