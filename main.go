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

var fname string = getEnv("APP_T_FILE_NAME", "logfile.log")
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
	http.HandleFunc("/cpu", cpuLoadHandler)
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

func cpuLoadHandler(w http.ResponseWriter, r *http.Request) {
	RunCPULoad(2, 5, 100)
	fmt.Fprintf(w, "CPU LOAD\n")
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// RunCPULoad run CPU load in specify cores count and percentage
func RunCPULoad(coresCount int, timeSeconds int, percentage int) {
	runtime.GOMAXPROCS(coresCount)

	// second     ,s  * 1
	// millisecond,ms * 1000
	// microsecond,μs * 1000 * 1000
	// nanosecond ,ns * 1000 * 1000 * 1000
	// every loop : run + sleep = 1 unit
	// 1 unit = 100 ms may be the best
	unitHundresOfMicrosecond := 1000
	runMicrosecond := unitHundresOfMicrosecond * percentage
	sleepMicrosecond := unitHundresOfMicrosecond*100 - runMicrosecond
	for i := 0; i < coresCount; i++ {
		go func() {
			runtime.LockOSThread()
			// endless loop
			for {
				begin := time.Now()
				for {
					// run 100%
					if time.Now().Sub(begin) > time.Duration(runMicrosecond)*time.Microsecond {
						break
					}
				}
				// sleep
				time.Sleep(time.Duration(sleepMicrosecond) * time.Microsecond)
			}
		}()
	}
	// how long
	time.Sleep(time.Duration(timeSeconds) * time.Second)
}
