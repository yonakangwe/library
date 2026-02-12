package main

import (
	"library/package/log"
	"library/services/database"
	webserver "library/webserver"
	"os"
	"os/signal"
	"syscall"

	"library/config"
)

func init() {
	path := config.LoggerPath()
	log.Infoln(path)
	log.SetOptions(
		log.Development(),
		log.WithCaller(true),
		log.WithLogDirs(path),
	)
}
func main() {

	database.Connect()
	defer database.Close()
	go webserver.StartWebserver()

	defer os.Exit(0)

	stop := make(chan os.Signal, 1)


	
	signal.Notify(
		stop,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)
	<-stop
	log.Infoln("auth webserver is shutting down .... 👋 !")
	go func() {
		<-stop
	}()
}
