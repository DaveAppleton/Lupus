package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"

	"github.com/DaveAppleton/etherUtils"

	"github.com/DaveAppleton/ether_go/ethKeys"

	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var askQuestionTopic common.Hash
var codeQuestionTopic common.Hash

// GetEthClient connect to the local parity node - assume localhost:8545
func GetEthClient(endPoint string) (client *ethclient.Client, err error) {
	if client != nil {
		return client, nil
	}
	client, err = ethclient.Dial(endPoint)
	return
}

func askQuestionType1(question string) (answer string, err error) {
	query := viper.GetString("Q1") + url.QueryEscape(strings.TrimSpace(question))
	//fmt.Println(query)
	resp, err := http.Get(query)
	if err != nil {
		return
	}
	var answerBytes []byte
	answerBytes, err = ioutil.ReadAll(resp.Body)
	answer = string(answerBytes)
	return
}

// Q2Reply data
type Q2Reply struct {
	Interpretation string
	Result         string
	Failure        string
}

func askQuestionType2(question string) (q2repl Q2Reply, err error) {
	query := viper.GetString("Q2") + url.QueryEscape(strings.TrimSpace(question))
	//fmt.Println(query)
	resp, err := http.Get(query)
	if err != nil {
		return
	}
	var answerBytes []byte
	answerBytes, err = ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(answerBytes, &q2repl)
	return
}

func askQuestionType3(question string) (answer string, err error) {
	query := viper.GetString("Q3") + url.QueryEscape(strings.TrimSpace(question))
	//fmt.Println(query)
	resp, err := http.Get(query)
	if err != nil {
		return
	}
	var answerBytes []byte
	answerBytes, err = ioutil.ReadAll(resp.Body)
	answer = string(answerBytes)
	return
}

func askQuestion(question string) (answer string, err error) {
	answer, err = askQuestionType1(question)
	if strings.Compare(answer, "Wolfram|Alpha did not understand your input") != 0 {
		return
	}
	Q2, err := askQuestionType2(question)
	if err != nil {
		return
	}
	if len(Q2.Failure) != 0 {
		answer = "Could not get answer"
		return
	}
	answer = Q2.Result
	return
}

func keyTx(key *ethKeys.AccountKey) *bind.TransactOpts {
	return bind.NewKeyedTransactor(key.GetKey())
}

type ctxReply struct {
	Src       common.Address
	Nonce     *big.Int
	Question  string
	Status    string
	LastBlock uint64
}

func checkTx(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var reply ctxReply
	txHashStr := r.FormValue("hash")

	hash := common.HexToHash(txHashStr)
	client, err := GetEthClient(viper.GetString("IPC_HOST"))
	if err != nil {
		log.Fatal("GetClient: ", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	rct, err := client.TransactionReceipt(ctx, hash)
	if err != nil {
		w.WriteHeader(418)
		log.Println("tx Receipt ", txHashStr, err)
		return
	}
	for _, log := range rct.Logs {
		if log.Topics[0] == askQuestionTopic || log.Topics[0] == codeQuestionTopic {
			if len(log.Topics) > 1 && len(log.Data) > 96 {

				reply.Src = common.HexToAddress(log.Topics[1].Hex())
				reply.Nonce = new(big.Int).SetBytes(log.Data[0:31])
				qBytes := log.Data[96:]

				n := bytes.IndexByte(qBytes, 0)
				reply.Question = string(qBytes[:n])
				reply.Status = "OK"
				reply.LastBlock = log.BlockNumber
				json.NewEncoder(w).Encode(reply)
				return
			}
		}
	}
	w.WriteHeader(418)
}

func checkAnswer(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}
	wolfAddress := common.HexToAddress(viper.GetString("CONTRACT_ADDRESS"))
	client, err := GetEthClient(viper.GetString("IPC_HOST"))
	if err != nil {
		log.Fatal("GetClient: ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	addrStr := r.FormValue("address")
	nonceStr := r.FormValue("nonce")
	blkNoStr := r.FormValue("block")

	addrHash := common.HexToHash(addrStr)
	bigNonce, ok := new(big.Int).SetString(nonceStr, 10)
	if !ok {
		w.WriteHeader(404)
		return
	}

	filter := ethereum.FilterQuery{}
	filter.Addresses = make([]common.Address, 0)
	filter.Addresses = append(filter.Addresses, wolfAddress)

	answerQuestionTopic := common.BytesToHash(crypto.Keccak256([]byte("Remus(address,uint256,string)")))
	filter.Topics = [][]common.Hash{[]common.Hash{answerQuestionTopic, addrHash}}
	filter.FromBlock, _ = etherUtils.StrToDecimals(blkNoStr, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	logEntries, err := client.FilterLogs(ctx, filter)
	if err != nil {
		log.Fatal("FilterLogs: ", err)
	}

	for _, logEnt := range logEntries {
		fmt.Println(".")
		if logEnt.Topics[0] != answerQuestionTopic {
			continue
		}
		src := common.HexToAddress(logEnt.Topics[1].Hex())
		if src != common.HexToAddress(addrStr) {
			continue
		}
		nonce := new(big.Int).SetBytes(logEnt.Data[0:31])
		if nonce.Cmp(bigNonce) != 0 {
			continue
		}
		qBytes := logEnt.Data[96:]
		n := bytes.IndexByte(qBytes, 0)
		answer := string(qBytes)
		if n > 0 {
			answer = string(qBytes[:n])
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(answer)

		return
	}
	w.WriteHeader(418)
}

func ethFilter() {
	payment := ethKeys.NewKey("keys/wolf")
	err := payment.RestoreOrCreate()
	fmt.Println("Transactor at ", payment.PublicKeyAsHexString())
	answerTx := keyTx(payment)
	wolfAddress := common.HexToAddress(viper.GetString("CONTRACT_ADDRESS"))
	client, err := GetEthClient(viper.GetString("IPC_HOST")) //"/Users/daveappleton/Library/Ethereum/geth.ipc") //
	if err != nil {
		log.Fatal("GetClient: ", err)
	}
	fmt.Println("Connected")

	contract, err := NewDogsOfRome(wolfAddress, client)
	if err != nil {
		fmt.Println("Error instantiating contract ", err)
		log.Fatal("Error instantiating contract ", err)
	}
	lastBlock, err := contract.LastQBlock(nil)
	if err != nil {
		fmt.Println("Error getting last block ", err)
		log.Fatal("Error getting last block ", err)
	}

	askQuestionTopic = common.BytesToHash(crypto.Keccak256([]byte("Romulus(address,uint256,string)")))
	codeQuestionTopic = common.BytesToHash(crypto.Keccak256([]byte("Scooby(address,uint256,string)")))
	fmt.Println("Filter is ", askQuestionTopic.Hex())
	filter := ethereum.FilterQuery{}
	filter.Addresses = make([]common.Address, 0)
	filter.Addresses = append(filter.Addresses, wolfAddress)

	filter.Topics = [][]common.Hash{[]common.Hash{askQuestionTopic, codeQuestionTopic}}

	fmt.Println("Starting at block number ", lastBlock)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	bal, err := client.BalanceAt(ctx, payment.PublicKey(), nil)
	if err != nil {
		fmt.Println("Check bal ", err)
		os.Exit(1)
	}
	if bal.Cmp(big.NewInt(10000000000000000)) < 0 {
		fmt.Println("WARNING LOW BALANCE", etherUtils.EtherToStr(bal))
		os.Exit(1)
	}
	fmt.Println("balance = ", etherUtils.EtherToStr(bal))

	for {
		filter.FromBlock = lastBlock
		ctx, cancel = context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		logEntries, err := client.FilterLogs(ctx, filter)
		if err != nil {
			log.Fatal("FilterLogs: ", err)
		}

		for _, logEnt := range logEntries {
			fmt.Println(".")
			mode := 0
			if logEnt.Topics[0] == askQuestionTopic {
				mode = 0
			} else if logEnt.Topics[0] == codeQuestionTopic {
				mode = 1
			} else {
				continue
			}

			src := common.HexToAddress(logEnt.Topics[1].Hex())
			nonce := new(big.Int).SetBytes(logEnt.Data[0:31])
			qBytes := logEnt.Data[96:]
			n := bytes.IndexByte(qBytes, 0)
			question := string(qBytes)
			if n > 0 {
				question = string(qBytes[:n])
			}
			answer := ""
			if mode == 0 {
				fmt.Println("Q) ", question)
				answer, err = askQuestion(question)
				if err != nil {
					fmt.Println(err)
					continue
				}
			} else {
				fmt.Println("C) ", question)
				answer, err = askQuestionType3(question)
				if err != nil {
					fmt.Println(err)
					continue
				}

			}
			fmt.Println("A) ", answer)
			answering := true
			if answering {
				tx, err := contract.LetRomulusReply(answerTx, src, nonce, answer, big.NewInt(int64(logEnt.BlockNumber)))
				if err != nil {
					fmt.Println(err)
					log.Println(err)
				} else {
					fmt.Println(tx.Hash().Hex())
					log.Println(tx.Hash().Hex())
				}
			}
			lastBlock = big.NewInt(int64(logEnt.BlockNumber + 1))

			// amount := new(big.Int).SetBytes(logEnt.Data)
			// src := common.HexToAddress(logEnt.Topics[1].Hex())
			// dest := common.HexToAddress(logEnt.Topics[2].Hex())
			time.Sleep(10 * time.Second)
		}
	}
}
