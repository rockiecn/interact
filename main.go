package main

import (
	"fmt"
	"math/big"

	//"github.com/ethereum/go-ethereum/accounts/abi"

	//"github.com/ethereum/go-ethereum/core/types"

	"github.com/rockiecn/interact/callcontact"
)

var (
	// 交易发起方keystore文件地址
	fromKeyStoreFile = "./data/keystore/UTC--2021-07-01T07-05-14.130901482Z--9e0153496067c20943724b79515472195a7aedaa"

	// keystore文件对应的密码
	password = "123123"

	// 交易接收方地址
	toAddress = "0x1714888Ede3a57b72781ff876dcb491cA3b4f744"

	// http服务地址, 例:http://localhost:8545
	httpUrl = "http://localhost:8545"

	chainID = big.NewInt(1337)
)

//
func main() {

	var cmd string

	fmt.Println("intput cmd, 1 to retrieve data, 2 to store data, 3 to deploy contract")
	fmt.Println("intput cmd, 4 to query balance, 5 to transfer value")

	fmt.Scanf("%s", &cmd)

	switch cmd {
	case "1":
		callcontact.CallRetrieve()

	case "2":
		fmt.Println("input data:")
		var idata int64

		fmt.Scanf("%d", &idata)
		fmt.Printf("your input: %d\n", idata)
		data := big.NewInt(idata)
		callcontact.CallStore(data)

	case "3":
		fmt.Println("deploying")
		callcontact.CallDeploy()

	case "4":
		fmt.Println("query balance of account0")
		callcontact.CallQuery()

	case "5":
		fmt.Println("transfer value")
		callcontact.CallTransferTo()
	}
}
