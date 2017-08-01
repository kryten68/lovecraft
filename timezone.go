package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	p := fmt.Println

	inputedTime := "2017-07-31 10:00:00.545"
	p("Timestamp provided:", inputedTime)

	timezone, err := time.LoadLocation("Europe/London")
	if err != nil {
		log.Println("Err is", err)
	} else {
		p("Timezone: ", timezone)
	}

	layout := "2006-01-02 15:04:05.000"
	p("layout provided:", layout)

	tt, err := time.ParseInLocation(layout, inputedTime, timezone)
	if err != nil {
		log.Println("Err is", err)
	} else {
		p("\nParsed time is: ", tt.String())
		p("\nUnix time is:\t", tt.Unix())
		p("ES time is:\t", tt.Unix()*1000)
	}

}
