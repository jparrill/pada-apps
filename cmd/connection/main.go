package main

import (
	"fmt"
	"net/http"

	"github.com/jparrill/pada-apps/pkg/utils"
)

var (
	inet_icon = ""
)

func main() {
	_, err := http.Get("http://clients3.google.com/generate_204")
	if err != nil {
		fmt.Printf("%s", utils.Erro(inet_icon))
	} else {
		fmt.Printf("%s", utils.Info(inet_icon))
	}
}
