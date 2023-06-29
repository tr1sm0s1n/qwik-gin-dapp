package controllers

import (
	"log"
	"net/http"

	"github.com/DEMYSTIF/qwik-gin-dapp/server/helpers"
	"github.com/DEMYSTIF/qwik-gin-dapp/server/interfaces"
	"github.com/DEMYSTIF/qwik-gin-dapp/server/lib"
	"github.com/DEMYSTIF/qwik-gin-dapp/server/services"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

func InfoController(ctx *gin.Context, client *ethclient.Client) {
	networkID, chainID, latestBlock := services.InfoService(client)

	ctx.IndentedJSON(http.StatusOK, gin.H{"net_version": networkID, "eth_chainId": chainID, "eth_blockNumber": latestBlock})
}

func IssueController(ctx *gin.Context, client *ethclient.Client, instance *lib.Cert) {
	var newCertificate interfaces.InputCertificate

	if err := ctx.ShouldBindJSON(&newCertificate); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"message": "Bad request"})
		return
	}

	trx, err := services.IssueService(client, instance, newCertificate)
	if err != nil {
		log.Fatal(err)
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{
		"issuedID": newCertificate.ID, "trx": trx,
	})
}

func FetchController(ctx *gin.Context, instance *lib.Cert) {
	param := ctx.Query("id")
	id, err := helpers.ParseInt(param)
	if err != nil {
		log.Fatal(err)
	}

	result, err := services.FetchService(instance, id)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"message": "Bad request"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"id": id, "name": result.Name, "course": result.Course, "grade": result.Grade, "date": result.Date,
	})
}
