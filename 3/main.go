package main

import(
	"os"
	"fmt"
	"strconv"
)

func CreateFiles(val int){
		file, err := os.Create("./files/"+strconv.Itoa(val)+".testfile")
		defer func(){
				file.Close()
		}()
		if err!=nil{
			panic(err)
		}
		fmt.Println("Файл создан, имя: ", val)
}

func main(){
	fCount := 1000000
	for iter:=0; iter<fCount; iter++{
		CreateFiles(iter)
	}
}

//что бы аварийная остановка программы не выполнилась - нужно recover() запустить для той рутины, где вызвана паника (к подзаданию 4)