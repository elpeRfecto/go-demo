package readandwrite

import (
	"fmt"
	"os"
)

func ReadFile(filename string)  {
	
}

func WriteFile(filename, name string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}

	_, err = file.WriteString(name)
	if err != nil {
		file.Close()
		fmt.Println(err)
	}
	file.Close()
}
