package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/julienschmidt/httprouter"
	"github.com/sebastianbroekhoven/go-emailverification-api/controllers"
)

func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile)

	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// Settings
	ListenHost := flag.String("host", "127.0.0.1", "Set the server host")
	ListenPort := flag.String("port", "3003", "Set the server port")

	// Read flags
	flag.Usage = func() {
		fmt.Println("\nUSAGE :")
		flag.PrintDefaults()
	}
	flag.Parse()

	// Loggins
	log.Println("---------------------------------------------")
	log.Println("  Get Email Verification API written in Go.  ")
	log.Println("----------------------------------------")
	log.Println("    Listening: http://" + *ListenHost + ":" + *ListenPort + "    ")
	log.Println("---------------------------------------------")

	// Instantiate a new router
	r := httprouter.New()

	evc := controllers.NewEmailVerifyController()
	r.GET("/v1/email/verify/:emailaddress", evc.EmailVerify)

	// Fire up the server
	http.ListenAndServe(*ListenHost+":"+*ListenPort, r)
}
