package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func deployOld() {
	const oldABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"DataStored\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"compute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getResult\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"
	const oldBin = "0x608060405234801561001057600080fd5b5060056000819055506003600181905550600060028190555061018b806100386000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80631a43c3381461003b578063de29278914610045575b600080fd5b610043610063565b005b61004d6100b4565b60405161005a91906100d7565b60405180910390f35b6001546000546100739190610121565b6002819055507f9455957c3b77d1d4ed071e2b469dd77e37fc5dfd3b4d44dc8a997cc97c7b3d496002546040516100aa91906100d7565b60405180910390a1565b6000600254905090565b6000819050919050565b6100d1816100be565b82525050565b60006020820190506100ec60008301846100c8565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061012c826100be565b9150610137836100be565b925082820190508082111561014f5761014e6100f2565b5b9291505056fea26469706673582212201f3b057b9f2a5c82890f074082290d895d835063fa71a2bb96895d4b612558cb64736f6c63430008120033"

	byteCode := common.Hex2Bytes(oldBin[2:])
	oldabi, _ := abi.JSON(strings.NewReader(oldABI))
	input, _ := oldabi.Pack("")
	byteCode = append(byteCode, input...)

	client, error := ethclient.Dial("http://localhost:8545")
	if error != nil {
		fmt.Println("Client error")
	}

	chainID := big.NewInt(32382)

	privateKey, err := crypto.HexToECDSA("09a0b04d809cb0dc46006e766db9f892bd2143c78cf8ca3e8df4d8e19e980407")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0)
	gasLimit := uint64(30000000)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	var accessList types.AccessList
	var stkhldrs []common.Address
	stkhldrs = append(stkhldrs, common.HexToAddress("0x0c8eb26280FC04ee9417ad3c3359bD4Eb28A9246"))
	stkhldrs = append(stkhldrs, common.HexToAddress("0x5415F5ce710ce1512B9b286CB52Ab498f91b34e4"))

	scupdateTx := types.SmartContractUpdateTx{ChainID: chainID, ProposalNumber: 0, VotesNeededToWin: 1, Nonce: nonce, GasPrice: gasPrice, Gas: gasLimit, To: nil, Value: value, Data: byteCode, AccessList: accessList, Stakeholders: stkhldrs}

	signedTx, err := types.SignTx(types.NewTx(&scupdateTx), types.NewEIP2930Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
	fmt.Println()

	time.Sleep(30 * time.Second)

	receipt, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(receipt.Status)
	fmt.Println(receipt.ContractAddress)

}

func approve(priv string, address string) {
	const newABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"DataStored\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"compute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	const newBin = "0x608060405234801561001057600080fd5b50610158806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80631a43c33814610030575b600080fd5b61003861003a565b005b60015460005461004a91906100c4565b6002819055507f9455957c3b77d1d4ed071e2b469dd77e37fc5dfd3b4d44dc8a997cc97c7b3d496002546040516100819190610107565b60405180910390a1565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006100cf8261008b565b91506100da8361008b565b92508282039050818111156100f2576100f1610095565b5b92915050565b6101018161008b565b82525050565b600060208201905061011c60008301846100f8565b9291505056fea26469706673582212208836426f631b8645ee912649c3484a565de11204f36f2364ada2ea23e40676e564736f6c63430008120033"

	byteCode := common.Hex2Bytes(newBin[2:])
	newabi, _ := abi.JSON(strings.NewReader(newABI))
	input, _ := newabi.Pack("")
	byteCode = append(byteCode, input...)

	client, error := ethclient.Dial("http://localhost:8545")
	if error != nil {
		fmt.Println("Client error")
	}

	chainID := big.NewInt(32382)

	privateKey, err := crypto.HexToECDSA(priv)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	toAddress := common.HexToAddress(address)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0)
	gasLimit := uint64(30000000)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	var accessList types.AccessList
	var stkhldrs []common.Address
	stkhldrs = append(stkhldrs, common.HexToAddress("0x0c8eb26280FC04ee9417ad3c3359bD4Eb28A9246"))
	stkhldrs = append(stkhldrs, common.HexToAddress("0x5415F5ce710ce1512B9b286CB52Ab498f91b34e4"))

	scupdateTx := types.ApproveProposalTx{ChainID: chainID, ProposalNumber: 1, VotesNeededToWin: 2, Nonce: nonce, GasPrice: gasPrice, Gas: gasLimit, To: &toAddress, Value: value, Data: byteCode, AccessList: accessList, Stakeholders: stkhldrs}

	signedTx, err := types.SignTx(types.NewTx(&scupdateTx), types.NewEIP2930Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
	fmt.Println()

	time.Sleep(30 * time.Second)

	receipt, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(receipt.Status)
	fmt.Println(receipt.ContractAddress)

}

func callComputeInOld(address string) {

	const oldABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"DataStored\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"compute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getResult\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"
	const oldBin = "0x608060405234801561001057600080fd5b5060056000819055506003600181905550600060028190555061018b806100386000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80631a43c3381461003b578063de29278914610045575b600080fd5b610043610063565b005b61004d6100b4565b60405161005a91906100d7565b60405180910390f35b6001546000546100739190610121565b6002819055507f9455957c3b77d1d4ed071e2b469dd77e37fc5dfd3b4d44dc8a997cc97c7b3d496002546040516100aa91906100d7565b60405180910390a1565b6000600254905090565b6000819050919050565b6100d1816100be565b82525050565b60006020820190506100ec60008301846100c8565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061012c826100be565b9150610137836100be565b925082820190508082111561014f5761014e6100f2565b5b9291505056fea26469706673582212201f3b057b9f2a5c82890f074082290d895d835063fa71a2bb96895d4b612558cb64736f6c63430008120033"

	oldabi, _ := abi.JSON(strings.NewReader(oldABI))
	bytesData, _ := oldabi.Pack("compute")

	client, error := ethclient.Dial("http://localhost:8545")
	if error != nil {
		fmt.Println("Client error")
	}

	chainID := big.NewInt(32382)

	privateKey, err := crypto.HexToECDSA("09a0b04d809cb0dc46006e766db9f892bd2143c78cf8ca3e8df4d8e19e980407")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	toAddress := common.HexToAddress(address)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0)
	gasLimit := uint64(30000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	var accessList types.AccessList

	accessListTx := types.AccessListTx{ChainID: chainID, Nonce: nonce, GasPrice: gasPrice, Gas: gasLimit, To: &toAddress, Value: value, Data: bytesData, AccessList: accessList}

	signedTx, err := types.SignTx(types.NewTx(&accessListTx), types.NewEIP2930Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
	fmt.Println()

	time.Sleep(30 * time.Second)

	receipt, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(receipt.Status)
	fmt.Println(receipt.ContractAddress)
}

func proposeNew(privKey string, address string) {
	const newABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"DataStored\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"compute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	const newBin = "0x608060405234801561001057600080fd5b50610158806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80631a43c33814610030575b600080fd5b61003861003a565b005b60015460005461004a91906100c4565b6002819055507f9455957c3b77d1d4ed071e2b469dd77e37fc5dfd3b4d44dc8a997cc97c7b3d496002546040516100819190610107565b60405180910390a1565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006100cf8261008b565b91506100da8361008b565b92508282039050818111156100f2576100f1610095565b5b92915050565b6101018161008b565b82525050565b600060208201905061011c60008301846100f8565b9291505056fea26469706673582212208836426f631b8645ee912649c3484a565de11204f36f2364ada2ea23e40676e564736f6c63430008120033"

	byteCode := common.Hex2Bytes(newBin[2:])
	newabi, _ := abi.JSON(strings.NewReader(newABI))
	input, _ := newabi.Pack("")
	byteCode = append(byteCode, input...)

	client, error := ethclient.Dial("http://localhost:8545")
	if error != nil {
		fmt.Println("Client error")
	}

	chainID := big.NewInt(32382)

	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	toAddress := common.HexToAddress(address)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0)
	gasLimit := uint64(30000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	var accessList types.AccessList
	var stkhldrs []common.Address
	stkhldrs = append(stkhldrs, common.HexToAddress("0x0c8eb26280FC04ee9417ad3c3359bD4Eb28A9246"))
	stkhldrs = append(stkhldrs, common.HexToAddress("0x5415F5ce710ce1512B9b286CB52Ab498f91b34e4"))

	scupdateTx := types.SmartContractUpdateTx{ChainID: chainID, Nonce: nonce, ProposalNumber: 1, GasPrice: gasPrice, Gas: gasLimit, To: &toAddress, Value: value, Data: byteCode, AccessList: accessList, Stakeholders: stkhldrs, VotesNeededToWin: 2}

	signedTx, err := types.SignTx(types.NewTx(&scupdateTx), types.NewEIP2930Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
	fmt.Println()

	time.Sleep(30 * time.Second)

	receipt, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(receipt.Status)
	fmt.Println(receipt.ContractAddress)
}

func deployNew(privKey string, address string) {
	const newABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"DataStored\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"compute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	const newBin = "0x608060405234801561001057600080fd5b50610158806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80631a43c33814610030575b600080fd5b61003861003a565b005b60015460005461004a91906100c4565b6002819055507f9455957c3b77d1d4ed071e2b469dd77e37fc5dfd3b4d44dc8a997cc97c7b3d496002546040516100819190610107565b60405180910390a1565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006100cf8261008b565b91506100da8361008b565b92508282039050818111156100f2576100f1610095565b5b92915050565b6101018161008b565b82525050565b600060208201905061011c60008301846100f8565b9291505056fea26469706673582212208836426f631b8645ee912649c3484a565de11204f36f2364ada2ea23e40676e564736f6c63430008120033"

	byteCode := common.Hex2Bytes(newBin[2:])
	newabi, _ := abi.JSON(strings.NewReader(newABI))
	input, _ := newabi.Pack("")
	byteCode = append(byteCode, input...)

	client, error := ethclient.Dial("http://localhost:8545")
	if error != nil {
		fmt.Println("Client error")
	}

	chainID := big.NewInt(32382)

	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	toAddress := common.HexToAddress(address)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0)
	gasLimit := uint64(30000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	var accessList types.AccessList
	var stkhldrs []common.Address
	stkhldrs = append(stkhldrs, common.HexToAddress("0x0c8eb26280FC04ee9417ad3c3359bD4Eb28A9246"))
	stkhldrs = append(stkhldrs, common.HexToAddress("0x5415F5ce710ce1512B9b286CB52Ab498f91b34e4"))

	scupdateTx := types.SmartContractUpdateTx{ChainID: chainID, Nonce: nonce, ProposalNumber: 1, GasPrice: gasPrice, Gas: gasLimit, To: &toAddress, Value: value, Data: byteCode, AccessList: accessList, Stakeholders: stkhldrs, VotesNeededToWin: 2}

	signedTx, err := types.SignTx(types.NewTx(&scupdateTx), types.NewEIP2930Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
	fmt.Println()

	time.Sleep(30 * time.Second)

	receipt, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(receipt.Status)
	fmt.Println(receipt.ContractAddress)
}

func callComputeInNew(address string) {

	const newABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"DataStored\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"compute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	const newBin = "0x608060405234801561001057600080fd5b50610158806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80631a43c33814610030575b600080fd5b61003861003a565b005b60015460005461004a91906100c4565b6002819055507f9455957c3b77d1d4ed071e2b469dd77e37fc5dfd3b4d44dc8a997cc97c7b3d496002546040516100819190610107565b60405180910390a1565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006100cf8261008b565b91506100da8361008b565b92508282039050818111156100f2576100f1610095565b5b92915050565b6101018161008b565b82525050565b600060208201905061011c60008301846100f8565b9291505056fea26469706673582212208836426f631b8645ee912649c3484a565de11204f36f2364ada2ea23e40676e564736f6c63430008120033"

	newabi, _ := abi.JSON(strings.NewReader(newABI))
	bytesData, _ := newabi.Pack("compute")

	client, error := ethclient.Dial("http://localhost:8545")
	if error != nil {
		fmt.Println("Client error")
	}

	chainID := big.NewInt(32382)

	privateKey, err := crypto.HexToECDSA("09a0b04d809cb0dc46006e766db9f892bd2143c78cf8ca3e8df4d8e19e980407")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	toAddress := common.HexToAddress(address)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0)
	gasLimit := uint64(30000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	var accessList types.AccessList

	accessListTx := types.AccessListTx{ChainID: chainID, Nonce: nonce, GasPrice: gasPrice, Gas: gasLimit, To: &toAddress, Value: value, Data: bytesData, AccessList: accessList}

	signedTx, err := types.SignTx(types.NewTx(&accessListTx), types.NewEIP2930Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
	fmt.Println()

	time.Sleep(30 * time.Second)

	receipt, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(receipt.Status)
	fmt.Println(receipt.ContractAddress)
}

func getCodeHash(address string) {

	client, error := ethclient.Dial("http://localhost:8545")
	if error != nil {
		fmt.Println("Client error")
	}

	codeHash, err := client.CodeAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(common.BytesToHash(codeHash))
}

func getStakeholders(address string) {

	client, error := ethclient.Dial("http://localhost:8545")
	if error != nil {
		fmt.Println("Client error")
	}

	stakeholders, err := client.StakeholdersAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		fmt.Println(err)
	}

	for _, addr := range stakeholders {

		fmt.Println(addr)
	}
}

func getProposalNumber(address string) {

	client, error := ethclient.Dial("http://localhost:8545")
	if error != nil {
		fmt.Println("Client error")
	}

	proposalNumber, err := client.ProposalNumberAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("Proposal Number:")
	fmt.Println(proposalNumber)
}

func getExecutionState(address string) {

	client, error := ethclient.Dial("http://localhost:8545")
	if error != nil {
		fmt.Println("Client error")
	}

	executionState, err := client.ExecutionStateAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		fmt.Println(err)
	}

	if executionState == types.ProposalRejected || executionState == types.ChangesApplied {
		fmt.Println("Executable")
	} else {
		fmt.Println("Not Executable")
	}

}

func getVotesNeededToWin(address string) {

	client, error := ethclient.Dial("http://localhost:8545")
	if error != nil {
		fmt.Println("Client error")
	}

	votesNeededToWin, err := client.VotesNeededToWinAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(votesNeededToWin)
}

func checkStorage(address string) {
	client, error := ethclient.Dial("http://localhost:8545")
	if error != nil {
		fmt.Println("Client error")
	}

	storage, err := client.StorageAt(context.Background(), common.HexToAddress(address), common.BigToHash(big.NewInt(2)), nil)
	if err != nil {
		fmt.Println("Client error")
	}
	fmt.Println(storage)
}

func main() {

	//deployOld()
	//getCodeHash("0xcbC258788f3d47Ad29B9971D7CC318003B0CEACb")
	//getStakeholders("0xcbC258788f3d47Ad29B9971D7CC318003B0CEACb")
	//getExecutionState("0xcbC258788f3d47Ad29B9971D7CC318003B0CEACb")
	//callComputeInOld("0xcbC258788f3d47Ad29B9971D7CC318003B0CEACb")
	//checkStorage("0xcbC258788f3d47Ad29B9971D7CC318003B0CEACb")
	//proposeNew("09a0b04d809cb0dc46006e766db9f892bd2143c78cf8ca3e8df4d8e19e980407", "0xcbC258788f3d47Ad29B9971D7CC318003B0CEACb")
	//getExecutionState("0xcbC258788f3d47Ad29B9971D7CC318003B0CEACb")
	//callComputeInOld("0xcbC258788f3d47Ad29B9971D7CC318003B0CEACb")
	//approve("09a0b04d809cb0dc46006e766db9f892bd2143c78cf8ca3e8df4d8e19e980407", "0xcbC258788f3d47Ad29B9971D7CC318003B0CEACb")
	//getExecutionState("0xcbC258788f3d47Ad29B9971D7CC318003B0CEACb")
	//deployNew("09a0b04d809cb0dc46006e766db9f892bd2143c78cf8ca3e8df4d8e19e980407", "0xcbC258788f3d47Ad29B9971D7CC318003B0CEACb")
	//getCodeHash("0xcbC258788f3d47Ad29B9971D7CC318003B0CEACb")
	//callComputeInNew("0xcbC258788f3d47Ad29B9971D7CC318003B0CEACb")
	//checkStorage("0xcbC258788f3d47Ad29B9971D7CC318003B0CEACb")

}
