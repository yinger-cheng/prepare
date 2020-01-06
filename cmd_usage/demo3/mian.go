package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	err error
	output []byte
}

func main(){
	var (
		ctx context.Context
		cancleFunc context.CancelFunc
		cmd *exec.Cmd
		resultChan chan *result
		res *result
	)
	resultChan = make (chan *result,1000)

	ctx,cancleFunc = context.WithCancel(context.TODO())
	go func(){
		var (
			output []byte
			err error
		)
		cmd = exec.CommandContext(ctx,"C:\\cygwin64\\bin\\bash.exe","-c","sleep 2;echo hello")
		//select { case <- ctx.Done()}
		output,err = cmd.CombinedOutput()

		resultChan <- &result{
			err:    err,
			output: output,
		}
	}()
	time.Sleep(1* time.Second)
	cancleFunc()

	res = <- resultChan

	fmt.Println(res.err,string(res.output))

}
