package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/tr1sm0s1n/qwik-gin-dapp/server/helpers"
	"github.com/tr1sm0s1n/qwik-gin-dapp/server/lib"
	"github.com/tr1sm0s1n/qwik-gin-dapp/server/middlewares"
)

func init() {
	if _, ok := os.LookupEnv("DOCKER"); !ok {
		if err := godotenv.Load(); err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	done := make(chan bool)

	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		log.Fatal(err)
	}

	contractAddress, trx, err := deployContract(client)
	if err != nil {
		log.Fatal(err)
	}

	wg.Add(2)
	go checkStatus(client, trx.Hash(), &wg, ch)
	go helpers.Loading(&wg, done)

	select {
	case <-ch:
		fmt.Printf("1\rTransaction has been committed!!")
		fmt.Println("\n--------------------------------")
		fmt.Printf("Contract Address: \033[32m%s\033[0m\n", contractAddress.String())
		fmt.Println("-----------------")
		fmt.Printf("Transaction Hash: \033[32m%s\033[0m\n", trx.Hash())
		fmt.Println("-----------------")
		helpers.UpdateEnv(contractAddress.String())
		done <- true
	case <-time.After(5 * time.Second):
		fmt.Println("1\r\033[31mSorry, timeout has reached!!\033[0m")
		done <- true
		ch <- 1
	}

	wg.Wait()
}

func deployContract(client *ethclient.Client) (common.Address, *types.Transaction, error) {
	auth := middlewares.AuthGenerator(client)
	result, err := bind.LinkAndDeploy(
		&bind.DeploymentParams{
			Contracts: []*bind.MetaData{&lib.CertMetaData},
		},
		bind.DefaultDeployer(auth, client))
	if err != nil {
		return common.Address{}, nil, err
	}
	return result.Addresses[lib.CertMetaData.ID], result.Txs[lib.CertMetaData.ID], err
}

func checkStatus(client *ethclient.Client, tHash common.Hash, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for {
		select {
		case <-ch:
			fmt.Println("\033[31mUnable to fetch the transaction receipt.\033[0m\033[?25h")
			return
		default:
			time.Sleep(1 * time.Second)
			trxReceipt, err := client.TransactionReceipt(context.Background(), tHash)
			if err != nil {
				continue
			}

			if trxReceipt.Status == types.ReceiptStatusSuccessful {
				ch <- 1
				return
			}
		}
	}
}
