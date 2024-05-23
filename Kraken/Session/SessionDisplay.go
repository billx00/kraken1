package Session

import (
	"fmt"
	"framework/kraken/Gostyle/Fore"
	"framework/kraken/Utilities"
	"time"
)

func IsSessionActive(_session Session) bool {
	return _session.Seen+7 >= time.Now().UTC().Unix()
}

func SesssionsCommand() {
	if len(Sessions) == 0 {
		fmt.Printf("%s", "No Active Sessions.")
		return
	}

	var sessionid_info string = "Session ID"
	var ipaddr_info string = "IP Address"
	var os_info string = "Os"
	var user_info string = "User"
	var status_info string = "Status"

	var longest_sessionid int = len(sessionid_info)
	var longest_ipaddr int = len(ipaddr_info)
	var longest_os int = len(os_info)
	var longest_user int = len(user_info)
	var longest_status int = len(status_info)

	for i := 0; i < len(Sessions); i++ {
		if longest_sessionid < len(Sessions[i].SessionId) {
			longest_sessionid = len(Sessions[i].SessionId)
		}

		if longest_ipaddr < len(Sessions[i].IPAddress) {
			longest_ipaddr = len(Sessions[i].IPAddress)
		}

		if longest_os < len(Sessions[i].Os) {
			longest_os = len(Sessions[i].Os)
		}

		if longest_user < len(Sessions[i].Username+"@"+Sessions[i].Hostname) {
			longest_user = len(Sessions[i].Username + "@" + Sessions[i].Hostname)
		}

		if !IsSessionActive(Sessions[i]) {
			longest_status = 8
		}
	}

	fmt.Printf("%s\n", Utilities.MakeStr("-", 1+len(Utilities.MakeStr("-", longest_sessionid)+" "+Utilities.MakeStr("-", longest_ipaddr)+" "+Utilities.MakeStr("-", longest_os)+" "+Utilities.MakeStr("-", longest_user)+" "+Utilities.MakeStr("-", longest_status)+" ")))
	fmt.Printf("|%s|\n",
		sessionid_info+(Utilities.MakeStr(" ", 1+longest_sessionid-(len(sessionid_info))))+
			ipaddr_info+(Utilities.MakeStr(" ", 1+longest_ipaddr-len(ipaddr_info)))+
			os_info+(Utilities.MakeStr(" ", 1+longest_os-len(os_info)))+
			user_info+(Utilities.MakeStr(" ", 1+longest_user-len(user_info)))+
			status_info+(Utilities.MakeStr(" ", longest_status-len(status_info))))
	fmt.Printf("|%s|\n", Utilities.MakeStr("-", longest_sessionid)+" "+Utilities.MakeStr("-", longest_ipaddr)+" "+Utilities.MakeStr("-", longest_os)+" "+Utilities.MakeStr("-", longest_user)+" "+Utilities.MakeStr("-", longest_status))

	for i := 0; i < len(Sessions); i++ {
		var Active_len int = 6
		var Activeniss string = Fore.Colorize(Fore.GREEN, "Active")
		if !IsSessionActive(Sessions[i]) {
			Activeniss = Fore.Colorize(Fore.RED, "Inactive")
			Active_len = 8
		}

		fmt.Printf("|%s|\n",
			Fore.Colorize(Fore.YELLOW, Sessions[i].SessionId)+(Utilities.MakeStr(" ", longest_sessionid-len(Sessions[i].SessionId)))+" "+
				Sessions[i].IPAddress+(Utilities.MakeStr(" ", longest_ipaddr-len(Sessions[i].IPAddress)))+" "+
				Sessions[i].Os+(Utilities.MakeStr(" ", longest_os-len(Sessions[i].Os)))+" "+
				Sessions[i].Username+"@"+Sessions[i].Hostname+(Utilities.MakeStr(" ", longest_user-len(Sessions[i].Username+"@"+Sessions[i].Hostname)))+" "+
				Activeniss+(Utilities.MakeStr(" ", longest_status-Active_len)))
	}
	fmt.Printf("%s\n", Utilities.MakeStr("-", 1+len(Utilities.MakeStr("-", longest_sessionid)+" "+Utilities.MakeStr("-", longest_ipaddr)+" "+Utilities.MakeStr("-", longest_os)+" "+Utilities.MakeStr("-", longest_user)+" "+Utilities.MakeStr("-", longest_status)+" ")))

}
