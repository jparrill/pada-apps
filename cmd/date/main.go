package main

import (
	"fmt"
	"time"

	"github.com/jparrill/pada-apps/pkg/utils"
)

var (
	date_icon = ""
	separator = "|"
	time_icon = ""
)

func main() {
	date := time.Now()
	fmt.Printf("%s %s %s %s %s", utils.Green(date_icon), utils.White(date.Format("Mon 02 Jan 2006")), utils.Grey(separator), utils.Green(time_icon), utils.White(date.Format("15:04")))
}
