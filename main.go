package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"
)

var fname string = getEnv("APP_T_FILE_NAME", "tmp")
var port string = getEnv("APP_PORT", "8080")

func main() {
	log.Println(fname)
	if _, err := os.Stat(fname); os.IsNotExist(err) {
		f, err := os.Create(fname)

		if err != nil {
			log.Fatal("Cannot create file", err)
		}
		f.Close()
	}

	go write_loop(fname)
	http.HandleFunc("/", handler)
	fmt.Println("Server started at port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func write_loop(name string) {
	key := "HOSTNAME"
	if runtime.GOOS == "windows" {
		key = "COMPUTERNAME"
	}
	val, ok := os.LookupEnv(key)

	log.Println(val)
	if !ok {
		log.Println("No " + key + " env var set")
	}

	f, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal("Cannot open file", err)
	}

	defer f.Close()
	for i := 1; ; i++ {

		s := strconv.Itoa(i)
		_, err2 := f.WriteString(val + " " + s + "\n")
		if err2 != nil {
			log.Fatal(err2)
		}
		time.Sleep(1 * time.Second)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK\n")
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
