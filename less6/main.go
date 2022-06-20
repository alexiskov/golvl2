package main

import(
	"os"
	"sync"
	"time"
	"runtime"
	"runtime/trace"
)

var(
	wrk = make(chan struct{}, runtime.NumCPU())
	mutex sync.Mutex
	iter int
)

func main(){
	trace.Start(os.Stderr)
	defer trace.Stop()
	
	for	i:=0;i<10;i++{
		go func(){
			mutex.Lock()
				iter+=1
			mutex.Unlock()
			wrk<-struct{}{}
			time.Sleep(10*time.Millisecond)
		}()
		if i%3==0{
			runtime.Gosched()
		}
		<-wrk
	}
}