package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tr1sm0s1n/qwik-gin-dapp/server/controllers"
	"github.com/tr1sm0s1n/qwik-gin-dapp/server/lib"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	contract := os.Getenv("CONTRACT_ADDRESS")
	printContract := fmt.Sprintf("Contract: %s", contract)
	fmt.Println(printContract)
	contractAddress := common.HexToAddress(contract)
	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		log.Fatal(err)
	}

	cert := lib.NewCert()
	instance := cert.Instance(client, contractAddress)

	router := gin.Default()
	router.POST("/issue", func(ctx *gin.Context) {
		controllers.IssueController(ctx, client, instance)
	})
	router.GET(("/fetch"), func(ctx *gin.Context) {
		controllers.FetchController(ctx, instance)
	})
	router.GET("/info", func(ctx *gin.Context) {
		controllers.InfoController(ctx, client)
	})
	router.Run("localhost:8080")
}
