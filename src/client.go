package main

import (
    "fmt"
    "log"
    "context"

    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/common"
)

func main() {
    client, err := ethclient.Dial("http://localhost:8545")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("we have a connection")

    // Hardcoded account check ganache output to change the account number!
    account := common.HexToAddress("0xbD53Ef4e1d736776F4458E9B43f6781CA43596cc")
    balance, err := client.BalanceAt(context.Background(), account, nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(balance)
}
