package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Hello, World!")
	})
	//logf, err := rotatelogs.New(
	//	"/log.%Y%m%d%H%M",
	//	rotatelogs.WithMaxAge(24*time.Hour),
	//	rotatelogs.WithRotationTime(time.Hour),
	//)
	//if err != nil {
	//	log.Printf("failed to create rotatelogs: %s", err)
	//	return
	//}

	// Now you must write to logf. apache-logformat library can create
	// a http.Handler that only writes the approriate logs for the request
	// to the given handle
	http.ListenAndServe(":8080", mux)
}

//func main() {
//	rl, _ := rotatelogs.New("/path/to/access_log.%Y%m%d%H%M")
//
//	log.SetOutput(rl)
//
//	/* elsewhere ... */
//	log.Printf("Hello, World!")
//}
