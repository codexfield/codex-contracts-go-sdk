package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"

	"github.com/codexfield/codex-contracts-go-sdk/account"
	"github.com/codexfield/codex-contracts-go-sdk/contracts/codexam"
)

const (
	PrivateKey         = ""
	TestnetRPC         = "https://data-seed-prebsc-1-s1.binance.org:8545/"
	AccountManagerAddr = "0xae5c57a7285602830aEA302f56e8Cf647a82F022"
)

func init() {
	log.SetDefault(log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stderr, log.LevelInfo, true)))
}

func main() {
	// Prepare account
	acc, err := account.NewAccount(PrivateKey)
	if err != nil {
		log.Crit("prepare account failed", "err", err)
	}

	// Prepare account manager
	client, err := ethclient.Dial(TestnetRPC)
	if err != nil {
		log.Crit("dial rpc failed", "err", err)
	}
	codexAM, err := codexam.NewICodexAM(common.HexToAddress(AccountManagerAddr), client)
	if err != nil {
		log.Crit("new codex account manager instance failed", "err", err)
	}

	// Register
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Crit("get gas price failed", "err", err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Crit("get chainID failed", "err", err)
	}
	tx, err := codexAM.Register(&bind.TransactOpts{
		From:  acc.Address(),
		Nonce: big.NewInt(acc.Nonce(client)),
		Signer: func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
			return acc.SignTx(transaction, chainID)
		},
		GasPrice: gasPrice,
		GasLimit: uint64(500000),
	}, acc.Address(), "darbo", "", "I am Darbo from CodexField", "CodexField", "SG", "https://codexfield.com", []string{})
	if err != nil {
		log.Crit("register failed", "err", err)
	}
	err = waitTxReceipt(client, tx.Hash())
	if err != nil {
		log.Error("wait register receipt failed", "err", err)
	}

	// Get Account ID
	accountID, err := codexAM.GetAccountId(&bind.CallOpts{}, acc.Address())
	if err != nil {
		log.Crit("get account id failed", "err", err)
	}

	log.Info("Register and get account id successfully", "account", acc.Address(), "id", accountID.Int64())
}

func waitTxReceipt(client *ethclient.Client, txHash common.Hash) error {
	for i := 0; i < 6; i++ {
		time.Sleep(10 * time.Second)
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err != nil {
			return fmt.Errorf("get transaction: %s receipt failed: %v", txHash.String(), err)
		}
		if receipt != nil {
			if receipt.Status == types.ReceiptStatusSuccessful {
				return nil
			} else if receipt.Status == types.ReceiptStatusFailed {
				return fmt.Errorf("transaction failed: %s", txHash.String())
			}
		}
	}
	return fmt.Errorf("get transaction: %s receipt timeout", txHash.String())
}
