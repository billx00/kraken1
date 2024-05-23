package Session

import (
	"framework/kraken/Kraken/Public"
	"framework/kraken/Utilities"
	"time"
)

type Session struct {
	SessionId string
	IPAddress string
	Username  string
	Hostname  string
	Os        string
	Job       string
	Alias     string
	Seen      int64
}

var NoJob string = "None"

var SelectedSession string
var Sessions []Session = []Session{}

func SessionStopWait() {
	SelectedSession = Utilities.StringEmpty
	Public.ContinueShell = true
}

func SessionStartWait(Session string) {
	SelectedSession = Session
	Public.ContinueShell = false
}

func SessionWait() {
	for !Public.ContinueShell {
		time.Sleep(time.Millisecond * 350)
	}
}

func GetSessionIndex(ia string) int {
	for i := 0; i < len(Sessions); i++ {
		if Sessions[i].SessionId == ia || Sessions[i].Alias == ia {
			return i
		}
	}
	return -1
}
