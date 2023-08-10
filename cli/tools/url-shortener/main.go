package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Data = map[string]string

var wg sync.WaitGroup

func listData(urlFrom, urlTo string) {
	http.HandleFunc("/"+urlFrom, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, urlTo, http.StatusSeeOther)
	})
}

func getData(d Data) (string, string, error) {
	// get input
	urlFromPtr := new(string)
	fmt.Print("Enter URL from: ")
	fmt.Scanf("%s\n", urlFromPtr)

	urlToPtr := new(string)
	fmt.Print("Enter URL to: ")
	fmt.Scanf("%s\n", urlToPtr)
	fmt.Println()

	// add data
	_, ok := d[*urlFromPtr]
	if !ok {
		d[*urlFromPtr] = *urlToPtr
		return *urlFromPtr, *urlToPtr, nil
	}

	return "", "", fmt.Errorf("URL is already taken")
}

func main() {
	var data Data = map[string]string{}

	wg.Add(1)

	// start server
	go func() {
		defer wg.Done()
		err := http.ListenAndServe(":4200", nil)

		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("Server started successfully")
		}
	}()

	// send data
	for {
        inputPtr := new(int)

        fmt.Printf("0: exit\n1: add item\n2: view all items\n")
        fmt.Scanf("%d\n", inputPtr)

        switch *inputPtr {
        case 0:
            panic("Program exited")
        case 1:
            urlFrom, urlTo, err := getData(data)
            if err != nil { return }

            data[urlFrom] = urlTo
            listData(urlFrom, urlTo)
        case 2:
            fmt.Println(data)
        default: 
            continue
        }
	}
}
