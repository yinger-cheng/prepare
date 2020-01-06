package main

import (
	"fmt"
	"os/exec"
)

func main(){
	var(
		cmd *exec.Cmd
		output []byte
		err error
	)

	//cmd = exec.Command("C:\\cygwin64\\bin\\bash.exe","-c","echo 1;echo 2;")
	//err = cmd.Run()
	//fmt.Println(err)

	cmd = exec.Command("C:\\cygwin64\\bin\\bash.exe","-c","echo hello")
	if output,err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
}