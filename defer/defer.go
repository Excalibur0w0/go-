package main 

import "fmt"
import "os"

func main() {
	f, err := os.Open("file")
	if err != nil {
		panic(err)
	}
	fmt.Println(f)
	
	// defer f.close()
}