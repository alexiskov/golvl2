package outthdd

import(
	"os"
	"io"
	"fmt"
)

func ToHDD(str string){
	file, err := os.Opent("./globaldict")
	if err != nil{
		file, err = osCreate("./globaldict")
		if err != nil {
			fmt.Println(err)
		}
	}
	defer file.Close()
	file.WriteString(str)
}