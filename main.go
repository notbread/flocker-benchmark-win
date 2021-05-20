package main

import (
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {
	var fname, ok = os.LookupEnv("APP_T_FILE_NAME")
	if !ok {
		log.Fatal("No APP_T_FILE_NAME env var set")
	}
	log.Println(fname)
	if _, err := os.Stat(fname); os.IsNotExist(err) {
		f, err := os.Create(fname)

		if err != nil {
			log.Fatal("Cannot create file", err)
		}
		f.Close()
	}

	write_loop(fname)

	log.Println("done")
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

	f, err := os.OpenFile(name, os.O_RDWR, 0666)

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
