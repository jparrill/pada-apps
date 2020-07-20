package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/jparrill/pada-apps/pkg/utils"
)

const (
	tmp_icon = ""
	suffix   = "ºC"
)

type Meta struct {
	CoreAdapter CoreAdapter `json:"coretemp-isa-0000"`
}

type CoreAdapter struct {
	Core CoreTemp `json:"Package id 0"`
}

type CoreTemp struct {
	TempInput float64 `json:"temp1_input"`
}

func islmSensorsInstalled() (bool, int) {
	var rc int
	_, err := exec.LookPath("sensors")
	if err != nil {
		rc = 1
		return false, rc
	}
	return true, rc
}

func getCoreTemp() (int, int) {
	var rc int
	var CoreTempMeta Meta
	raw_temp, err := exec.Command("sensors", "-j").Output()
	if err != nil {
		log.Fatal(err)
		rc = 1
	}

	if err := json.Unmarshal([]byte(raw_temp), &CoreTempMeta); err != nil {
		rc = 1
		//panic(err)
	}

	return int(CoreTempMeta.CoreAdapter.Core.TempInput), rc
}

func main() {
	installed, _ := islmSensorsInstalled()
	if !(installed) {
		fmt.Println("Please install lm_sensors Package")
		os.Exit(0)
	}

	temp, _ := getCoreTemp()
	if int(temp) >= 30 && int(temp) < 60 {
		fmt.Printf("%s %s%s", utils.Green(tmp_icon), utils.White(temp), utils.White(suffix))
	} else if int(temp) >= 65 && int(temp) <= 80 {
		fmt.Printf("%s %s%s", utils.Green(tmp_icon), utils.Warn(temp), utils.Warn(suffix))
	} else if int(temp) > 80 && int(temp) <= 95 {
		fmt.Printf("%s %s%s", utils.Green(tmp_icon), utils.Erro(temp), utils.Erro(suffix))
	} else {
		fmt.Printf("%s %s%s", utils.Green(tmp_icon), utils.Erro("ERR"), utils.Erro(suffix))
	}

}
