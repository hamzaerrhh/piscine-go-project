package helper

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ToHexa(str string) string {
	n, err := strconv.ParseInt(str, 16, 64)
	if err != nil {
		fmt.Println("error in conv")
		os.Exit(0)
	}
	return strconv.FormatInt(n, 10)
}

func ToBinary(str string) string {
	n, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		fmt.Println("error in converte binary")
		os.Exit(0)
	}
	return strconv.FormatInt(n, 10)
}

// up,low,cap
func Up(str string) string {
	return strings.ToUpper(str)
}

func Low(str string) string {
	return strings.ToLower(str)
}

func Cap(str string) string {
	fmt.Println("str", str)
	
	return strings.ToUpper(str[:1]) + strings.ToLower(str[1:])
}

// up,low,cap
func UpN(strs ...string) string {
	return "up with number"
}

func LowN(strs ...string) string {
	return "low with number"
}

func CapN(strs ...string) string {
	return "cap with number"
}

// check after a
func Acheck(str string) bool {
	// check if the string began with voyel [a, e, i, o, u, h]

	return true
}
