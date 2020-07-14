package utils

import "fmt"

var (
	Info = Green
	Warn = Yellow
	Erro = Red
	NoEx = Strikthro
)

var (
	Black     = Color("\033[1;30m%s\033[0m")
	Red       = Color("\033[1;31m%s\033[0m")
	Green     = Color("\033[1;32m%s\033[0m")
	Yellow    = Color("\033[1;33m%s\033[0m")
	Purple    = Color("\033[1;34m%s\033[0m")
	Magenta   = Color("\033[1;35m%s\033[0m")
	Teal      = Color("\033[1;36m%s\033[0m")
	White     = Color("\033[1;37m%s\033[0m")
	Grey      = Color("\033[1;2m%s\033[0m")
	HiBlack   = Color("\033[1;90m%s\033[0m")
	HiRed     = Color("\033[1;91m%s\033[0m")
	HiGreen   = Color("\033[1;92m%s\033[0m")
	HiYellow  = Color("\033[1;93m%s\033[0m")
	HiPurple  = Color("\033[1;94m%s\033[0m")
	HiMagenta = Color("\033[1;95m%s\033[0m")
	HiTeal    = Color("\033[1;96m%s\033[0m")
	HiWhite   = Color("\033[1;97m%s\033[0m")
	Strikthro = Color("\033[1;9m%s\033[0m")
)

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}
