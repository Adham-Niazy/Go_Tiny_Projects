package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Here we defined our new type of output
type fileWriter struct {}

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// 1st method to parse the response ||

	// The read function can't resize slices so you should declare it with huge amount 
	// and them pass it to read function
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)

	// Then we will use the byte slice that created through `read` method to write it to seperate file
	// os.WriteFile("index.html", bs, 0666)

	// 2nd method to parse the response || with built in Stdout
	// io.Copy(os.Stdout, resp.Body)

	// 3rd method to parse the response || with our custom type
	fw := fileWriter{}
	io.Copy(fw, resp.Body)
}

// We are trying to make our type implement the write interface
func (fileWriter) Write(bs []byte) (int, error) {
	err := os.WriteFile("index.html", bs, 0666)
	if err != nil {
		fmt.Println("Error: ", err)
		return 0, err
	} else {
		return len(bs), nil
	}
}