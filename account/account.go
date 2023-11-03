package account

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type IAccount interface {
	Address() common.Address
	SignTx(tx *types.Transaction, chainID *big.Int) (*types.Transaction, error)
	Nonce(client *ethclient.Client) int64
}

type Account struct {
	private    string
	addr       common.Address
	rawPrivate *ecdsa.PrivateKey
}

func NewAccount(private string) (*Account, error) {
	account := &Account{
		private: private,
	}

	var err error
	account.rawPrivate, err = crypto.HexToECDSA(private)
	if err != nil {
		return nil, err
	}

	// Deduce public address
	account.addr = crypto.PubkeyToAddress(account.rawPrivate.PublicKey)

	return account, nil
}

func (a *Account) Address() common.Address {
	return a.addr
}

func (a *Account) SignTx(tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	if tx.GasFeeCap().Cmp(tx.GasTipCap()) == 0 {
		return types.SignTx(tx, types.NewEIP155Signer(chainID), a.rawPrivate)
	}

	return types.SignTx(tx, types.NewLondonSigner(chainID), a.rawPrivate)
}

func (a *Account) Nonce(client *ethclient.Client) int64 {
	nonce, err := client.PendingNonceAt(context.Background(), a.addr)
	if err != nil {
		return -1
	}

	return int64(nonce)
}
