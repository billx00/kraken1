package Log

import (
	"fmt"
	"framework/kraken/Gostyle/Fore"
)

func Error(str string) {
	fmt.Printf("["+Fore.RED+"Erro"+Fore.RESET+"] %s", str)
}

func Info(str string) {
	fmt.Printf("["+Fore.GREEN+"Info"+Fore.RESET+"] %s", str)
}

func Warning(str string) {
	fmt.Printf("["+Fore.YELLOW+"Warn"+Fore.RESET+"] %s", str)
}

func Debug(str string) {
	fmt.Printf("["+Fore.BLUE+"Debg"+Fore.RESET+"] %s", str)
}
