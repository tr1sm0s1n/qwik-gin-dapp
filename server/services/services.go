package services

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tr1sm0s1n/qwik-gin-dapp/server/helpers"
	"github.com/tr1sm0s1n/qwik-gin-dapp/server/interfaces"
	"github.com/tr1sm0s1n/qwik-gin-dapp/server/lib"
	"github.com/tr1sm0s1n/qwik-gin-dapp/server/middlewares"
)

func InfoService(client *ethclient.Client) (*big.Int, *big.Int, uint64) {
	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	latestBlock, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return networkID, chainID, latestBlock
}

func IssueService(client *ethclient.Client, instance *lib.Cert, newCertificate interfaces.InputCertificate) (*types.Transaction, error) {
	auth := middlewares.AuthGenerator(client)
	intID, err := helpers.ParseInt(newCertificate.ID)
	if err != nil {
		log.Fatal(err)
	}

	trx, err := instance.Issue(auth, big.NewInt(int64(intID)), newCertificate.Name, newCertificate.Course, newCertificate.Grade, newCertificate.Date)

	return trx, err
}

func FetchService(instance *lib.Cert, id int64) (interfaces.ReturnCertificate, error) {
	opts := bind.CallOpts{}
	certID := big.NewInt(id)

	result, err := instance.Certificates(&opts, certID)
	return result, err
}
