package main

import (
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	factory "github.com/arteptiber/techchart/contracts/factory"
	pool "github.com/arteptiber/techchart/contracts/pool"
	token "github.com/arteptiber/techchart/contracts/token"
)

var tokens map[string]common.Address = map[string]common.Address{
	"WETH": common.HexToAddress("0xc778417e063141139fce010982780140aa0cd5ab"),
	"UNI":  common.HexToAddress("0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"),
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Panic(err)
	}
}

func main() {
	if providerURL, exists := os.LookupEnv("PROVIDER_URL"); exists {
		connect(providerURL)
	}
}

func connect(providerURL string) {
	client, err := ethclient.Dial(providerURL)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984")
	instance, err := factory.NewFactory(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fee := big.NewInt(10000)

	poolAddress, err := instance.GetPool(nil, tokens["WETH"], tokens["UNI"], fee)
	if err != nil {
		log.Fatal(err)
	}

	poolInstance, err := pool.NewPool(poolAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	_ = poolInstance

	token0Instance, err := token.NewToken(tokens["WETH"], client)
	if err != nil {
		log.Fatal(err)
	}

	token1Instance, err := token.NewToken(tokens["UNI"], client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(token0Instance.BalanceOf(nil, poolAddress))
	fmt.Println(token1Instance.BalanceOf(nil, poolAddress))
}
