package tron

import (
	"context"
	"log"

	api "github.com/tron-us/go-btfs-common/protos/protocol/api"
	"google.golang.org/grpc"
)

type TronClient struct {
	client api.WalletClient
}

func NewTronClient() *TronClient {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Tron node: %v", err)
	}
	client := api.NewWalletClient(conn)
	return &TronClient{client: client}
}

func (c *TronClient) GetBalance(address string) (int64, error) {
	addr := &api.AccountAddress{Address: address}
	acc, err := c.client.GetAccount(context.Background(), addr)
	if err != nil {
		return 0, err
	}
	return acc.Balance, nil
}

func (c *TronClient) GetTransactions(address string) ([]*api.Transaction, error) {
	addr := &api.AccountAddress{Address: address}
	txs, err := c.client.ListTransactions(context.Background(), addr)
	if err != nil {
		return nil, err
	}
	return txs.Transactions, nil
}
