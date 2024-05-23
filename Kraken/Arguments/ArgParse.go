package Arguments

import (
	"os"
	"strings"
)

var ServerInfo map[string]string = map[string]string{
	"EngineServer": "0.0.0.0:8080",
}

func ParseAndSet() {
	if len(ParseArguments(os.Args)["lhost"]) > 0 {
		ServerInfo["EngineServer"] = ParseArguments(os.Args)["lhost"]
	}
}

func ParseArguments(Args []string) map[string]string {
	var ParsedArguments map[string]string = map[string]string{}

	for i := 1; i < len(Args); i++ {
		if strings.Contains(Args[i], "=") {
			ParsedArguments[strings.Split(Args[i], "=")[0]] = strings.Split(Args[i], "=")[1]
		} else {
			ParsedArguments[Args[i]] = "True"
		}
	}

	return ParsedArguments
}
