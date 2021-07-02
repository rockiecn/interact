package main

import (
	"fmt"
	"log"
	"math/big"
	//"strings"
	//"time"
	"github.com/ethereum/go-ethereum/ethclient"
	//"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/rockiecn/interact/storage"
)

//
func main() {

	var cmd string

	fmt.Println("intput cmd, 1 to retrieve data, 2 to store data.")
	fmt.Scanf("%s",&cmd)

	switch cmd {
	case "1":
		CallRetrieve()

	case "2": 
		fmt.Println("input data:")
		var idata int64

		fmt.Scanf("%d",&idata)
		fmt.Printf("your input: %d\n",idata)
		data := big.NewInt(idata)
		CallStore(data)
		
	case "3":
		fmt.Println("deploying")
		CallDeploy()
	}
	
}

//
func CallRetrieve() error {
    //cli, err := ethclient.Dial("http://localhost:8545")
	//cli, err := ethclient.Dial("../myGeth/data/geth.ipc")
    cli ,err := getClient("http://localhost:8545")
	if err != nil {
			fmt.Println("failed to dial geth", err)
			return err
	}
	defer cli.Close()
	
    //
    storageInstance, err := storage.NewStorage(common.HexToAddress("0x1714888Ede3a57b72781ff876dcb491cA3b4f744"), cli)
	if err != nil {
		fmt.Println("NewStorage err: ",err)
	} else {
		fmt.Println("NewStorage success: ",storageInstance)
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
	cli ,err := getClient("http://localhost:8545")
	if err != nil {
			fmt.Println("failed to dial geth", err)
			return err
	}
	defer cli.Close()

	//
    sk, err :=  crypto.HexToECDSA("cb61e1519b560d994e4361b34c181656d916beb68513cff06c37eb7d258bf93d")
    if err != nil {
	    fmt.Println("HexToECDSA err: ",err)
    } else {
		fmt.Println("get sk: ",sk)
	}
    auth := bind.NewKeyedTransactor(sk)
	if err != nil {
		fmt.Println("auth err:", err)
	} else {
		fmt.Printf("auth success: %v\n",auth)
	}

    //
    storageInstance, err := storage.NewStorage(common.HexToAddress("0x1714888Ede3a57b72781ff876dcb491cA3b4f744"), cli)
	if err != nil {
		fmt.Println("NewStorage err: ",err)
		return err
	} else {
		fmt.Println("NewStorage success: ",storageInstance)
	}
	
    //
    fmt.Printf("n = %s\n",n.String())
	
	tx, err:= storageInstance.Store(auth,n)
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

    //cli, err := ethclient.Dial("http://192.168.105.141:8545")

	cli, err := getClient("https://127.0.0.1:8545")
	if err != nil {
			fmt.Println("failed to dial geth", err)
			return storageAddr, err
	}
	defer cli.Close()

    sk, err :=  crypto.HexToECDSA("cb61e1519b560d994e4361b34c181656d916beb68513cff06c37eb7d258bf93d")
    if err != nil {
	    fmt.Println("HexToECDSA err: ",err)
    } else {
		fmt.Println("get sk: ",sk)
	}

    auth := bind.NewKeyedTransactor(sk)
	if err != nil {
		fmt.Println("auth err:", err)
	} else {
		fmt.Printf("auth success: %v\n",auth)
	}

	storageAddr, _, _, err = storage.DeployStorage(auth, cli)
	if err != nil {
		log.Println("deployStoragedErr:", err)
		return storageAddr, err
	}
	log.Println("storageAddr:", storageAddr.String())
	return storageAddr, nil

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
