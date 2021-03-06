package callcash

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/rockiecn/interact/cash"
	"github.com/rockiecn/interact/clientops"
)

const HOST = "http://localhost:8545"

//
func CallApplyCheque(storageAddr []byte, nonce []byte, payAmount []byte, sig []byte) error {
	fmt.Println("HOST: ", HOST)
	cli, err := clientops.GetClient(HOST)
	if err != nil {
		fmt.Println("failed to dial geth", err)
		return err
	}
	defer cli.Close()

	hexSk := "cb61e1519b560d994e4361b34c181656d916beb68513cff06c37eb7d258bf93d"
	auth, err := clientops.MakeAuth(hexSk, nil, nil, big.NewInt(1000), 3000000)
	if err != nil {
		return err
	}

	//
	cashInstance, err := cash.NewCash(common.HexToAddress("0x77AA1d64C1E85Cc4AF38046FfF5bc35e394f8eAD"), cli)
	if err != nil {
		fmt.Println("NewCash err: ", err)
		return err
	} else {
		fmt.Println("NewCash success: ", cashInstance)
	}

	//
	//fmt.Printf("n = %s\n", n.String())

	// address to receive money
	toAddress := common.HexToAddress("0xb213d01542d129806d664248a380db8b12059061")
	// transfer 1 eth to receiver
	tx, err := cashInstance.ApplyCheque(auth, toAddress, big.NewInt(1000000000000000000))
	if err != nil {
		fmt.Println("tx failed :", err)
		return err
	}

	fmt.Println("tx:", tx)

	return err
}

//
func CallDeploy() (common.Address, error) {
	var cashAddr common.Address

	fmt.Println("HOST: ", HOST)
	client, err := clientops.GetClient(HOST)
	if err != nil {
		fmt.Println("failed to dial geth", err)
		return cashAddr, err
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
	// string to bigint
	bn := new(big.Int)
	bn, ok1 := bn.SetString("100000000000000000000", 10) // deploy 100 eth
	if !ok1 {
		fmt.Println("SetString: error")
		panic("SetString error")
	}
	auth.Value = bn                // deploy 100 eth
	auth.GasLimit = uint64(230000) // in units
	auth.GasPrice = gasPrice

	fmt.Printf("auth success: %v\n", auth)

	cashAddr, _, _, err = cash.DeployCash(auth, client)
	if err != nil {
		log.Println("deployCashErr:", err)
		return cashAddr, err
	}
	log.Println("cashAddr:", cashAddr.String())
	return cashAddr, nil

}
