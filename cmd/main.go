package main

import (
	"flag"
	"fmt"
	"log"
	msBeer "msBeer/pkg"
	"msBeer/pkg/beers"
	"msBeer/pkg/log/logrus"
	"msBeer/pkg/server"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error loading .env file, error:", err)
	}
	var (
		hostName, _     = os.Hostname()
		defaultServerID = fmt.Sprintf("%s-%s", os.Getenv("API_NAME"), hostName)
		defaultHost     = os.Getenv("API_SERVER_HOST")
		defaultPort, _  = strconv.Atoi(os.Getenv("API_SERVER_PORT"))
	)

	fmt.Println("hostname is:", hostName)
	fmt.Println("defaultServerID:", defaultServerID)
	fmt.Println("defaultHost:", defaultHost)
	fmt.Println("defaultPort:", defaultPort)

	host := flag.String("host", defaultHost, "define host of the server")
	port := flag.Int("port", defaultPort, "define port of the server")
	serverID := flag.String("server-id", defaultServerID, "define server identifier")
	flag.Parse()

	logger := logrus.NewLogger()

	//simulation of repository
	var repo msBeer.Repository
	beerServices := beers.NewService(repo, logger)

	httpAddr := fmt.Sprintf("%s:%d", *host, *port)

	s := server.New(
		*serverID,
		beerServices,
	)

	fmt.Println("The server is on tap now:", httpAddr)
	log.Fatal(http.ListenAndServe(httpAddr, s.Router()))
}
