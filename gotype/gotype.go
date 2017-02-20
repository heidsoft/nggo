package main

import (
	"time"
	"fmt"
)

func main() {
	fmt.Println(int64(time.Millisecond))
	fmt.Println(int64(time.Nanosecond))
	timestamp :=  makeTimestamp()
	fmt.Println(timestamp)

	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	// Note that there is no `UnixMillis`, so to get the
	// milliseconds since epoch you'll need to manually
	// divide from nanoseconds.
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// You can also convert integer seconds or nanoseconds
	// since the epoch into the corresponding `time`.
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}


func makeTimestamp() int64 {
	return time.Now().UnixNano() / (int64(time.Second)/int64(time.Nanosecond))
}