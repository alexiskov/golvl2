package main

import(
	"fmt"
	"os"
	//"io"
)

func main(){
	file, err := os.Open("./nofile")
	defer func(){
		if rec:=recover(); rec!=nil{
			fmt.Println("Восстановлено", rec)
		}
	}()
	if err != nil{
		panic(err)
	}
	defer file.Close()
}