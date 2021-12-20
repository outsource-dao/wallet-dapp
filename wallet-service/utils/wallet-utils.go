package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"outsource-dao/wallet-dapp/proto"
	"outsource-dao/wallet-dapp/types"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	AUTH_URL      = GoDotEnvVariable("AUTH_URL")
	GRPC_URL      = GoDotEnvVariable("GRPC_URL")
	CONTRACT_ADDR = GoDotEnvVariable("MNGR_CONTRACT_ADDR")
	ABI           = GoDotEnvVariable("MNGR_CONTRACT_ABI")
	TKN_ADDR      = GoDotEnvVariable("TKN_CONTRACT_ADDR")
	TKN_ABI       = GoDotEnvVariable("TKN_CONTRACT_ABI")
)

// GenToken -> Generates auth token from auth service
func GenToken(clientID string) (types.TokenType, error) {
	url := fmt.Sprintf("%s/getToken/%s", AUTH_URL, clientID)
	resp, err := http.Get(url)
	if err != nil {
		return types.TokenType{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		var token types.TokenType
		err := json.Unmarshal(body, &token)
		if err != nil {
			return types.TokenType{}, err
		}
		return token, nil
	}
	return types.TokenType{}, errors.New("Something Went Wrong")
}

// VerifyToken -> Verifies uth token from auth service
func VerifyToken(token string) bool {
	url := fmt.Sprintf("%s/verifyToken/%s", AUTH_URL, token)
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return true
	}
	return false
}

// CreateTransferTokenTx -> Sends transaction parameters to create unsigned transaction
func CreateTransferTokenTx(authToken string, txParams types.TxParams, netID uint32) (string, error) {
	conn, err := grpc.Dial(GRPC_URL, grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	defer conn.Close()
	client := proto.NewProtoEthServiceClient(conn)
	header := metadata.New(map[string]string{"authorization": authToken})
	ctxHeader := metadata.NewOutgoingContext(context.Background(), header)
	// Creating unsigned `transfer` request
	p, err := json.Marshal(txParams.Params)
	if err != nil {
		return "", err
	}
	// transfer() call to ERC-20 token contract
	c := &proto.CallRequest{
		Networkid:   netID,
		Fn:          "transfer",
		Params:      string(p),
		Abi:         TKN_ABI,
		Address:     TKN_ADDR,
		FromAddress: txParams.From,
		Value:       txParams.Value,
	}
	unsignedTx, err := client.ContractCall(ctxHeader, c)
	if err != nil {
		return "", err
	}
	return unsignedTx.Result, nil
}

// GetTokenBalance -> Get balance of given address for Token
func GetTokenBalance(authToken string, txParams types.TxParams, netID uint32) (string, error) {
	conn, err := grpc.Dial(GRPC_URL, grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	defer conn.Close()
	client := proto.NewProtoEthServiceClient(conn)
	header := metadata.New(map[string]string{"authorization": authToken})
	ctxHeader := metadata.NewOutgoingContext(context.Background(), header)
	p, err := json.Marshal(txParams.Params)
	if err != nil {
		return "", err
	}
	// balanceOf() call to manager contract
	c := &proto.CallRequest{
		Networkid:   netID,
		Fn:          "balanceOf",
		Params:      string(p),
		Abi:         ABI,
		Address:     CONTRACT_ADDR,
		FromAddress: txParams.From,
	}
	fmt.Println(c)
	BalanceResp, err := client.ContractCall(ctxHeader, c)
	if err != nil {
		return "", err
	}
	return BalanceResp.Result, nil
}

// CreateSwapTokenTx -> Sends transaction parameters to create unsigned transaction
func CreateSwapTokenTx(authToken string, txParams types.TxParams, netID uint32) (string, error) {
	conn, err := grpc.Dial(GRPC_URL, grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	defer conn.Close()
	client := proto.NewProtoEthServiceClient(conn)
	header := metadata.New(map[string]string{"authorization": authToken})
	ctxHeader := metadata.NewOutgoingContext(context.Background(), header)
	// Creating unsigned `transfer` request
	p, err := json.Marshal(txParams.Params)
	if err != nil {
		return "", err
	}
	// swap() call to manager contract
	c := &proto.CallRequest{
		Networkid:   netID,
		Fn:          "swap",
		Params:      string(p),
		Abi:         ABI,
		Address:     CONTRACT_ADDR,
		FromAddress: txParams.From,
		Value:       txParams.Value,
	}
	unsignedTx, err := client.ContractCall(ctxHeader, c)
	if err != nil {
		return "", err
	}
	return unsignedTx.Result, nil
}

// CreateApproveTx -> Creates a transaction to approve manager to spend on behaf of the user
func CreateApproveTx(authToken string, txParams types.TxParams, netID uint32) (string, error) {
	conn, err := grpc.Dial(GRPC_URL, grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	defer conn.Close()
	client := proto.NewProtoEthServiceClient(conn)
	header := metadata.New(map[string]string{"authorization": authToken})
	ctxHeader := metadata.NewOutgoingContext(context.Background(), header)
	// Creating unsigned `approve` request
	p, err := json.Marshal(txParams.Params)
	if err != nil {
		return "", err
	}
	// approve() call to ERC-20 token contract
	c := &proto.CallRequest{
		Networkid:   netID,
		Fn:          "approve",
		Params:      string(p),
		Abi:         TKN_ABI,
		Address:     TKN_ADDR,
		FromAddress: txParams.From,
		Value:       txParams.Value,
	}
	unsignedTx, err := client.ContractCall(ctxHeader, c)
	if err != nil {
		return "", err
	}
	return unsignedTx.Result, nil
}
