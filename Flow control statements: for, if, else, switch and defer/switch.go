package main

import (
	"fmt"
	"runtime"
)

//switch文はif-else文を簡略化したもの
//switch-caseは上から下に評価し、合致したらそこでbreakされる。
//defaultはswitch trueに相当。
func main() {
	fmt.Println("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}