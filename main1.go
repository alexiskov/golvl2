package main

import(
	"log"
	"time"
	"os"
	"os/signal"
	"syscall"
)

type chanSig struct{}

var(
	workerPull = make(chan chanSig, 1000)
	i=0
)

func printMyIncrement(){
	defer func(){ <- workerPull }()
	time.Sleep(60 * time.Millisecond)
	i++
}

func signalInterceptor(s chan os.Signal){
	<-s
	time.Sleep(time.Second)
	log.Println("Завершено пользователем")
	os.Exit(0)
}

func main(){
	sig:=make(chan os.Signal, 1)
	signal.Notify(sig,syscall.SIGTERM)
	go signalInterceptor(sig)

	for i<1000{
		workerPull <- chanSig{}
		printMyIncrement()
	}
	log.Println(i)
}