package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"sync"
	"time"

	"github.com/DEMYSTIF/qwik-gin-dapp/server/helpers"
	"github.com/DEMYSTIF/qwik-gin-dapp/server/lib"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress, trx, err := deployContract(client)
	if err != nil {
		log.Fatal(err)
	}

	wg.Add(1)
	go checkStatus(client, trx.Hash(), &wg, ch)

	select {
	case <-ch:
		fmt.Printf("Contract Address: \x1b[32m%s\x1b[0m\n", contractAddress.String())
		fmt.Println("-----------------")
		fmt.Printf("Transaction Hash: \x1b[32m%s\x1b[0m\n", trx.Hash())
		fmt.Println("-----------------")
		helpers.UpdateEnv(contractAddress.String())
	case <-time.After(30 * time.Second):
		fmt.Println("\x1b[31mTimeout reached!!\x1b[0m")
	}

	wg.Wait()
}

func deployContract(client *ethclient.Client) (common.Address, *types.Transaction, error) {
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasLimit := uint64(927000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	contract, transaction, _, err := lib.DeployCert(auth, client)
	return contract, transaction, err
}

func checkStatus(client *ethclient.Client, tHash common.Hash, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	trxReceipt, err := client.TransactionReceipt(context.Background(), tHash)
	if err != nil {
		log.Fatal(err)
	}

	if trxReceipt.Status == 1 {
		ch <- 1
		return
	}

	time.Sleep(1 * time.Second)
	checkStatus(client, tHash, wg, ch)
}
