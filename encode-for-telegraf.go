package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	//"calls-generated,runid=123456,uuid=6562-HHHHG-HVGHHF-665556 value=130 {{timestamp}}"
	title := "calls-generated"
	runid := "123456"
	uuid := "7627632-jhbhvbb-27282-vhghgvhg"
	tags := "runid=" + runid + "," + "uuid=" + uuid
	value := "value=" + "130"
	timeNow := time.Now()
	timeUnix := strconv.FormatInt(timeNow.UnixNano(), 10)

	for i := 0; i < 20; i++ {
		encodeForLineProtocol(title, tags, value, timeUnix)
	}
}

func encodeForLineProtocol(title string, tags string, value string, timeUnix string) {
	lineProtocolTemplate := "{{measurement}},{{tags}} {{value}} {{timestamp}}"
	i := strings.Replace(lineProtocolTemplate, "{{measurement}}", title, 1)
	i = strings.Replace(i, "{{tags}}", tags, 1)
	i = strings.Replace(i, "{{value}}", value, 1)
	i = strings.Replace(i, "{{timestamp}}", timeUnix, 1)

	sendToDispatcher(i)
}

func sendToDispatcher(i string) {
	fmt.Println("from dispatch:", i)
}
