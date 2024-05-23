package Session

import (
	"regexp"
)

func VerifySesssion(_Session Session) bool {
	if VerifySessionOs(_Session.Os) && VerifySessionIp(_Session.IPAddress) {
		return true
	}

	return false
}

func VerifySessionIp(SessionIp string) bool {
	Regex := regexp.MustCompile(`[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}`)
	return Regex.Match([]byte(SessionIp))
}

func VerifySessionOs(SessionOs string) bool {
	var SupportedOs []string = []string{"windows", "linux"}

	for i := 0; i < len(SupportedOs); i++ {
		if SupportedOs[i] == SessionOs {
			return true
		}
	}

	return false
}
