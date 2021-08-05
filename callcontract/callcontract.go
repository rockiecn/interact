package callstorage

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/rockiecn/interact/clientops"
	"github.com/rockiecn/interact/storage"
)

const (
	HOST = "http://localhost:8545"
)

// to do: pick contract address as param,
// and call.go is irrelavent to specific contract address
func CallRetrieve() error {
	fmt.Println("HOST: ", HOST)
	cli, err := clientops.GetClient(HOST)
	if err != nil {
		fmt.Println("failed to dial geth", err)
		return err
	}
	defer cli.Close()

	//
	storageInstance, err := storage.NewStorage(common.HexToAddress("0x52D898a941C0b3d12C19b45008c210965644c7D5"), cli)
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
	cli, err := clientops.GetClient(HOST)
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
	auth, err := clientops.MakeAuth(hexSk, nil, nil, big.NewInt(1000), 3000000)
	if err != nil {
		return err
	}

	//
	storageInstance, err := storage.NewStorage(common.HexToAddress("0x52D898a941C0b3d12C19b45008c210965644c7D5"), cli)
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
	client, err := clientops.GetClient(HOST)
	//client, err := ethclient.Dial(HOST)
	if err != nil {
		fmt.Println("failed to dial geth", err)
		return storageAddr, err
	}
	defer client.Close()

	// get sk
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
	auth, err := bind.NewKeyedTransactorWithChainID(sk, big.NewInt(1337))
	if err != nil {
		log.Panic("NewKeyedTransactorWithChainID err:", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	//auth.GasLimit = uint64(300000) // in units
	auth.GasLimit = uint64(230000) // in units
	auth.GasPrice = gasPrice

	fmt.Printf("auth success: %v\n", auth)

	storageAddr, _, _, err = storage.DeployStorage(auth, client)
	if err != nil {
		log.Println("deployStoragedErr:", err)
		return storageAddr, err
	}
	log.Println("storageAddr:", storageAddr.String())
	return storageAddr, nil

	/*
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
	*/

}

//QueryBalance(account string)
func CallQuery() (*big.Int, error) {
	//var addr common.Address
	//addr = common.HexToAddress("0x9e0153496067c20943724b79515472195a7aedaa")
	//0x9e0153496067c20943724b79515472195a7aedaa
	//0xd6071743390681c792cef53bedfef72a5a0cd8ef
	//0x2dc957d527ddb25af35d3f2593f289f48843d4dc
	log.Println("querying balance of 0xd6071743390681c792cef53bedfef72a5a0cd8ef")
	ret, err := clientops.QueryBalance("0xd6071743390681c792cef53bedfef72a5a0cd8ef")
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
	err := clientops.TransferTo(value, "0xd6071743390681c792cef53bedfef72a5a0cd8ef", "http://localhost:8545")
	if err != nil {
		log.Println("transfer error:", err)
	}
	return err
}
