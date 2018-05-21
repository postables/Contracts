package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	bindings "./bindings/ERC20UTXO"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	key := os.Getenv("ETH_KEY")
	pass := os.Getenv("ETH_PASS")
	connURL := os.Getenv("CONN_URL")
	if key == "" {
		log.Fatal("invalid key provided")
	}
	if pass == "" {
		log.Fatal("invalid password provided")
	}
	if connURL == "" {
		log.Fatal("invlaid connection URL specified")
	}
	client, err := ethclient.Dial(connURL)
	if err != nil {
		log.Fatal("erorr connecting to ethclient ", err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(key), pass)
	if err != nil {
		log.Fatal("erorr creating authneticator ", err)
	}

	contractAddress, tx, contract, err := bindings.DeployERC20UTXO(auth, client)
	if err != nil {
		log.Fatal("error deploying contract ", err)
	}
	fmt.Println(contractAddress, tx, contract)
}
