package Utilities

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var StringEmpty string

func StrConcat(s1 string, s2 string) string {
	return fmt.Sprintf("%s%s", s1, s2)
}

func MakeStr(c string, l int) string {
	var r string
	for i := 0; i < l; i++ {
		r += c
	}
	return r
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GetInput() string {
	cin_reader := bufio.NewReader(os.Stdin)
	user_input, err := cin_reader.ReadString('\n')

	if err != nil {
		return ""
	}

	return user_input[0 : len(user_input)-1]
}

func CmdParse(cmdLine string) []string {
	re := regexp.MustCompile(`"([^"\\]*(?:\\.[^"\\]*)*)"|'([^'\\]*(?:\\.[^'\\]*)*)'|(\S+)`)

	args := re.FindAllString(cmdLine, -1)

	for i := 0; i < len(args); i++ {
		args[i] = strings.Trim(regexp.MustCompile(`\\(.)`).ReplaceAllString(args[i], "$1"), " ")

		if strings.HasPrefix(args[i], "\"") && strings.HasSuffix(args[i], "\"") ||
			strings.HasPrefix(args[i], "'") && strings.HasSuffix(args[i], "'") {
			args[i] = args[i][1 : len(args[i])-1]
		}
	}

	return args
}
