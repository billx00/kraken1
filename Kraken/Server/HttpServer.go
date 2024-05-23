package Server

import (
	"fmt"
	"framework/kraken/Gostyle/Fore"
	"framework/kraken/Gostyle/Log"
	"framework/kraken/Kraken/Arguments"
	"framework/kraken/Kraken/Session"
	"framework/kraken/Utilities"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

var Hanshakes map[string]string = map[string]string{
	"buff": "\x0F\x0F\xFF",
}

func StartServerAsync() {
	http.HandleFunc("/", Handle)

	go func() {
		if err := http.ListenAndServe(Arguments.ServerInfo["EngineServer"], nil); err != nil {
			Log.Error(err.Error())
			os.Exit(0)
		}
	}()
}

func Handle(w http.ResponseWriter, r *http.Request) {
	notify := w.(http.CloseNotifier).CloseNotify()

	var close_conn bool = false
	go func() {
		<-notify

		close_conn = true
	}()

	var oname string = r.URL.Query().Get("oname")
	var uname string = r.URL.Query().Get("uname")
	var hname string = r.URL.Query().Get("hname")

	var raddr string = strings.Split(r.RemoteAddr, ":")[0]

	var reqtm int64 = time.Now().UTC().Unix()

	var sesid string = Utilities.GetMD5Hash(uname + "@" + hname + "#" + raddr)

	var _session Session.Session = Session.Session{
		SessionId: sesid,
		Username:  uname,
		Hostname:  hname,
		Os:        oname,
		Job:       Session.NoJob,
		IPAddress: raddr,
		Seen:      reqtm,
	}

	var isSessionNew bool = Session.GetSessionIndex(_session.SessionId) == -1

	if !Session.VerifySesssion(_session) {
		return
	}

	if isSessionNew {
		Session.Sessions = append(Session.Sessions, _session)

		fmt.Fprint(w, Session.NoJob)

		return
	}

	Session.Sessions[Session.GetSessionIndex(_session.SessionId)].Seen = reqtm

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return
	}

	defer r.Body.Close()

	for Session.SelectedSession != _session.SessionId {
		if close_conn {
			return
		}

		Session.Sessions[Session.GetSessionIndex(_session.SessionId)].Seen = time.Now().UTC().Unix()
		time.Sleep(350 * time.Millisecond)
	}

	if string(body) != Hanshakes["buff"] && Session.SelectedSession == _session.SessionId {
		fmt.Printf("%s\n", Fore.Colorize(Fore.GREEN, string(body)))

		Session.Sessions[Session.GetSessionIndex(_session.SessionId)].Job = Session.NoJob

		Session.SessionStopWait()
	}

	fmt.Fprintf(w, "%s", Session.Sessions[Session.GetSessionIndex(_session.SessionId)].Job)
}
