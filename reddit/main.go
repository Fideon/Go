package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://reddit.com/r/golang.json", nil)
	if err != nil {
		log.Println("NOT OK")
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "Fideon bot")

	response, err := client.Do(req)

	if err != nil {
		log.Println("NOT OK")
		log.Fatalln(err)
	}

	if response.StatusCode != http.StatusOK {
		log.Println("NOT OK")
		log.Fatal(response.Status)
	}
	_, err = io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}
}

/*
package main

import (
 "io"
 "log"
 "net/http"
 "os"
)

func main() {
 client := &http.Client{}

 req, err := http.NewRequest("GET", "http://reddit.com/r/golang.json", nil)
 if err != nil {
  log.Fatal(err)
 }

 req.Header.Set("User-Agent", "your bot 0.1")

 response, err := client.Do(req)

 if err != nil {
  log.Fatalln(err)
 }

 if response.StatusCode != http.StatusOK {
  log.Fatal(response.Status)
 }

 _, err = io.Copy(os.Stdout, response.Body)
 if err != nil {
  log.Fatal(err)
 }

}*/
