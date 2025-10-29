package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	url := "http://google.co.jp/idjsaodajsdoi"

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if res.StatusCode != http.StatusOK {
		fmt.Println("statuscode=:", res.StatusCode)
	}
	fmt.Println(res)
	loadEnv()
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("ERROR!!!:", err)
	}
	message := os.Getenv("SAMPLE_MESSAGE")

	fmt.Println(message)
}
