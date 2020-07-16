package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/jparrill/pada-apps/pkg/utils"
)

const (
	cpu_icon = ""
)

func getLoadAvg(field int64) (float64, int) {
	var rc int
	var dummy float64 = 99999.99999
	loadAvgFile := "/proc/loadavg"
	if field > 5 {
		//fmt.Printf("%s please select the correct field [1, 2 or 3]\n", utils.Erro("Error:"))
		rc = 2
		return dummy, rc
	}
	contents, err := ioutil.ReadFile(loadAvgFile)
	if err != nil {
		//fmt.Printf("%s Reaging %s file: \t%s", utils.Erro("Error:"), loadAvgFile, err)
		rc = 1
		return dummy, rc
	}

	fields := strings.Split(string(contents), " ")
	loadAvg, err := strconv.ParseFloat(fields[field-1], 0)
	if err != nil {
		//fmt.Printf("%s String to Float conversion:\t%s\n", utils.Erro("Error:"), err)
		rc = 1
		return dummy, rc
	}
	return loadAvg, rc
}

func getCPUs() (float64, int) {
	var rc int
	var dummy float64 = 99999.99999
	nproc, err := exec.Command("nproc").CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		rc = 2
	}

	ncpus, err := strconv.ParseFloat(strings.TrimSpace(string(nproc)), 2)
	if err != nil {
		//fmt.Printf("Error on String to Float conversion: %s\n", err)
		rc = 1
		return dummy, rc
	}

	return ncpus, rc
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage : %s <field>\n", os.Args[0])
		fmt.Printf("1: 1 min, 2: 5 min, 3: 15 min")
		os.Exit(0)
	}

	field := os.Args[1]
	i, _ := strconv.ParseInt(field, 10, 64)

	loadAvg, _ := getLoadAvg(i)
	ncpus, _ := getCPUs()
	cpuUsage := (float64(100) * loadAvg) / ncpus

	if int(cpuUsage) >= 0 && int(cpuUsage) < 50 {
		fmt.Printf("%s %s"+utils.White("%%"), utils.Info(cpu_icon), utils.White(strconv.FormatFloat(cpuUsage, 'f', 0, 64)))
	} else if int(cpuUsage) >= 50 && int(cpuUsage) <= 70 {
		fmt.Printf("%s %s", utils.Info(cpu_icon), utils.Warn(strconv.FormatFloat(cpuUsage, 'f', 0, 64)))
	} else if int(cpuUsage) > 70 && int(cpuUsage) <= 100 {
		fmt.Printf("%s %s", utils.Info(cpu_icon), utils.Erro(strconv.FormatFloat(cpuUsage, 'f', 0, 64)))
	} else {
		fmt.Printf("%s %s", utils.Info(cpu_icon), utils.Erro("ERR"))
	}
}
