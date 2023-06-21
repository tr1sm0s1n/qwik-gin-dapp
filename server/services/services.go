package services

import (
	"context"
	"log"
	"math/big"

	"github.com/DEMYSTIF/qwik-gin-dapp/server/interfaces"
	"github.com/DEMYSTIF/qwik-gin-dapp/server/lib"
	"github.com/DEMYSTIF/qwik-gin-dapp/server/middlewares"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
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
	trx, err := instance.Issue(auth, big.NewInt(int64(newCertificate.ID)), newCertificate.Name, newCertificate.Course, newCertificate.Grade, newCertificate.Date)

	return trx, err
}

func FetchService(instance *lib.Cert, id int64) (interfaces.ReturnCertificate, error) {
	opts := bind.CallOpts{}
	certID := big.NewInt(id)

	result, err := instance.Certificates(&opts, certID)
	return result, err
}
