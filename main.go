package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	//"errors"

	"github.com/ethereum/go-ethereum/ethclient"
	//"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	//"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/rockiecn/interact/storage"
	//"github.com/rockiecn/interact/root"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rockiecn/interact/store"
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

const (
	HOST = "http://localhost:8545"
)

//
func main() {

	var cmd string

	fmt.Println("intput cmd, 1 to retrieve data, 2 to store data, 3 to deploy contract")
	fmt.Println("intput cmd, 4 to query balance, 5 to transfer value")

	fmt.Scanf("%s", &cmd)

	switch cmd {
	case "1":
		CallRetrieve()

	case "2":
		fmt.Println("input data:")
		var idata int64

		fmt.Scanf("%d", &idata)
		fmt.Printf("your input: %d\n", idata)
		data := big.NewInt(idata)
		CallStore(data)

	case "3":
		fmt.Println("deploying")
		CallDeploy()

	case "4":
		fmt.Println("query balance of account0")
		CallQuery()

	case "5":
		fmt.Println("transfer value")
		CallTransferTo()

	}

}

//
func CallRetrieve() error {
	//cli, err := ethclient.Dial("http://localhost:8545")
	//cli, err := ethclient.Dial("../myGeth/data/geth.ipc")
	fmt.Println("HOST: ", HOST)
	cli, err := getClient(HOST)
	if err != nil {
		fmt.Println("failed to dial geth", err)
		return err
	}
	defer cli.Close()

	//
	storageInstance, err := storage.NewStorage(common.HexToAddress("0x1714888Ede3a57b72781ff876dcb491cA3b4f744"), cli)
	if err != nil {
		fmt.Println("NewStorage err: ", err)
	} else {
		fmt.Println("NewStorage success: ", storageInstance)
	}

	//
	res, err := storageInstance.Retrieve(nil)
	if err != nil {
		fmt.Println("retrieve err:", err)
		return err
	}

	fmt.Println("result is", res)
	return err
}

//
func CallStore(n *big.Int) error {
	fmt.Println("HOST: ", HOST)
	cli, err := getClient(HOST)
	if err != nil {
		fmt.Println("failed to dial geth", err)
		return err
	}
	defer cli.Close()

	/*
		//
		sk, err := crypto.HexToECDSA("cb61e1519b560d994e4361b34c181656d916beb68513cff06c37eb7d258bf93d")
		if err != nil {
			fmt.Println("HexToECDSA err: ", err)
		} else {
			fmt.Println("get sk: ", sk)
		}
		auth := bind.NewKeyedTransactor(sk)
		if err != nil {
			fmt.Println("auth err:", err)
		} else {
			fmt.Printf("auth success: %v\n", auth)
		}
	*/

	hexSk := "cb61e1519b560d994e4361b34c181656d916beb68513cff06c37eb7d258bf93d"
	auth, err := makeAuth(hexSk, nil, nil, big.NewInt(1000), 3000000)
	if err != nil {
		return err
	}

	//
	storageInstance, err := storage.NewStorage(common.HexToAddress("0x1714888Ede3a57b72781ff876dcb491cA3b4f744"), cli)
	if err != nil {
		fmt.Println("NewStorage err: ", err)
		return err
	} else {
		fmt.Println("NewStorage success: ", storageInstance)
	}

	//
	fmt.Printf("n = %s\n", n.String())

	tx, err := storageInstance.Store(auth, n)
	if err != nil {
		fmt.Println("tx failed :", err)
		return err
	}

	fmt.Println("tx:", tx)

	return err
}

//
func CallDeploy() (common.Address, error) {
	var storageAddr common.Address

	fmt.Println("HOST: ", HOST)
	client, err := getClient(HOST)
	//client, err := ethclient.Dial(HOST)
	if err != nil {
		fmt.Println("failed to dial geth", err)
		return storageAddr, err
	}
	defer client.Close()

	//
	sk, err := crypto.HexToECDSA("cb61e1519b560d994e4361b34c181656d916beb68513cff06c37eb7d258bf93d")
	if err != nil {
		fmt.Println("HexToECDSA err: ", err)
	} else {
		fmt.Println("get sk: ", sk)
	}

	//
	publicKey := sk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	//
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	//
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)
	auth := bind.NewKeyedTransactor(sk)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	//auth.GasLimit = uint64(300000) // in units
	auth.GasLimit = uint64(230000) // in units
	auth.GasPrice = gasPrice

	fmt.Printf("auth success: %v\n", auth)

	/*
		storageAddr, _, _, err = storage.DeployStorage(auth, cli)
		if err != nil {
			log.Println("deployStoragedErr:", err)
			return storageAddr, err
		}
		log.Println("storageAddr:", storageAddr.String())
		return storageAddr, nil
	*/

	storeAddr, tx, ins, err := store.DeployStore(auth, client, "1.0")
	if err != nil {
		log.Println("deployStoreErr:", err)
		log.Println("storeAddr:", storeAddr)
		log.Println("tx:", tx)
		return storeAddr, err
	} else {
		log.Println("tx:", tx)
		log.Println("ins:", ins)

	}
	txCID := tx.ChainId()
	log.Println("txCID:", txCID)

	log.Println("storeAddr:", storeAddr.String())
	return storeAddr, nil

}

//QueryBalance(account string)
func CallQuery() (*big.Int, error) {
	//var addr common.Address
	//addr = common.HexToAddress("0x9e0153496067c20943724b79515472195a7aedaa")
	//0x9e0153496067c20943724b79515472195a7aedaa
	//0xd6071743390681c792cef53bedfef72a5a0cd8ef
	//0x2dc957d527ddb25af35d3f2593f289f48843d4dc
	log.Println("querying balance of 0xd6071743390681c792cef53bedfef72a5a0cd8ef")
	ret, err := QueryBalance("0xd6071743390681c792cef53bedfef72a5a0cd8ef")
	if err != nil {
		log.Println("query balance error:", err)
	}

	//log.Println("query balance success, balance:", ret)
	return ret, err

}

//TransferTo(value *big.Int, addr, eth string)
func CallTransferTo() error {
	log.Println("transfering from account0 to account1")
	value := new(big.Int)
	value.SetString("100000000000000000000", 10) //100 eth
	err := TransferTo(value, "0xd6071743390681c792cef53bedfef72a5a0cd8ef", "http://localhost:8545")
	if err != nil {
		log.Println("transfer error:", err)
	}
	return err
}

//
func getClient(endPoint string) (*ethclient.Client, error) {
	rpcClient, err := rpc.Dial(endPoint)
	if err != nil {
		log.Println("rpc.Dial err:", err)
		return nil, err
	}

	conn := ethclient.NewClient(rpcClient)
	return conn, nil
}

//MakeAuth make the transactOpts to call contract
func makeAuth(hexSk string, moneyToContract, nonce, gasPrice *big.Int, gasLimit uint64) (*bind.TransactOpts, error) {
	auth := &bind.TransactOpts{}
	sk, err := crypto.HexToECDSA(hexSk)
	if err != nil {
		log.Println("HexToECDSA err: ", err)
		return auth, err
	}

	auth = bind.NewKeyedTransactor(sk)
	auth.GasPrice = gasPrice
	auth.Value = moneyToContract //放进合约里的钱
	auth.Nonce = nonce
	auth.GasLimit = gasLimit
	return auth, nil
}

//
func QueryBalance(account string) (*big.Int, error) {
	var result string

	client, err := rpc.Dial(HOST)
	if err != nil {
		log.Println("rpc.dial err:", err)
		return big.NewInt(0), err
	}

	retryCount := 0
	for {
		retryCount++

		err = client.Call(&result, "eth_getBalance", account, "latest")

		if err != nil {
			if retryCount > 3 {
				return big.NewInt(0), err
			}
			time.Sleep(1000)
			continue
		} else {
			log.Println("call getbalance success: result = ", result)
		}
		//balance := utils.HexToBigInt(result)

		// hex to bitInt
		trimResult := strings.TrimPrefix(result, "0x")
		balance := new(big.Int)
		balance.SetString(trimResult, 16)

		log.Println("in queryBalance")
		log.Println("result:", result)
		log.Println("balance:", balance)
		return balance, nil

		//log.Println("balance:", result)
		//return nil, nil
	}
}

//
func TransferTo(value *big.Int, addr, eth string) error {
	client, err := ethclient.Dial(eth)
	if err != nil {
		fmt.Println("rpc.Dial err", err)
		log.Fatal(err)
	}

	//account0 cb61e1519b560d994e4361b34c181656d916beb68513cff06c37eb7d258bf93d
	//account2 920ffe90f05741f3b27aeec8a843f870d51b2a2d30d65afcb7390c0851af39f3
	privateKey, err := crypto.HexToECDSA("cb61e1519b560d994e4361b34c181656d916beb68513cff06c37eb7d258bf93d")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("from addr:", fromAddress.String())

	gasLimit := uint64(23000)           // in units
	gasPrice := big.NewInt(30000000000) // in wei (30 gwei)

	toAddress := common.HexToAddress(addr[2:])
	log.Println("toAddress: ", toAddress)

	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println("client.NetworkID error,use the default chainID")
		chainID = big.NewInt(666)
	}
	log.Println("networkID: ", networkID)

	log.Println("constructing and sending tx..")

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
		//continue
	}

	gasPrice, err = client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
		//continue
	}

	// construct tx
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	// sign tx
	//log.Println("chainID: ", chainID)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Println("sign transaction failed")
		return err
		//continue
	}

	// send tx
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Println("send transcation failed:", err)
		return err
		//continue
	} else {
		log.Println("send transaction success")
	}

	/*
		log.Println("quering balance ...")

		//balance, _ := QueryBalance(addr, eth)
		balance, err := QueryBalance(addr)
		if err != nil {
			log.Println("query balance error")
		} else {
			log.Println("query balance success")
		}

		log.Println("balance：", balance)
		log.Println("value", value)

		log.Println(addr, "'s Balance now:", balance.String())
	*/

	// wait 200 seconds?
	/*
		fmt.Println(addr, "'s Balance now:", balance.String(), ", waiting for transfer success")
		t := 20 * (qCount + 1)
		time.Sleep(time.Duration(t) * time.Second)
	*/

	log.Println("transfer ", value.String(), "wei to", addr, "complete")
	return nil
}
