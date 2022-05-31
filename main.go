package main

import(
	"fmt"
	"os"
	"runtime/debug"
)

type MyWrapErr struct{
	message string
	trace string
}

func NewError (msg string) error{
	return &MyWrapErr{
		message: msg,
		trace: string(debug.Stack()),
	}
} 

func (e *MyWrapErr) Error() string {
	return fmt.Sprintf("--ошибка: %s\n--трассировка:\n%s", e.message, e.trace)
}

func main(){
	file, err := os.Open("./nofile")
	defer func(){
		if rec:=recover(); rec!=nil{
			fmt.Println("Восстановлено", NewError("file not found"))
		}
	}()
	if err != nil{
		panic(err)
	}
	defer file.Close()
}