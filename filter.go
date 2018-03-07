package main

import (
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/DaveAppleton/parityclient"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/net/context"
)

func makeFilter(client *parityclient.Client, topic common.Hash) (fID *big.Int, err error) {
	filter := ethereum.FilterQuery{}
	filter.Addresses = []common.Address{common.HexToAddress("0xcc4ef9eeaf656ac1a2ab886743e98e97e090ed38")}
	filter.Topics = [][]common.Hash{[]common.Hash{topic}}

	filter.FromBlock = big.NewInt(0)
	ctx := context.Background()
	fID, err = client.NewFilter(ctx, filter)
	return
}

func filterEvents(client *parityclient.Client) {
	askQuestionTopic := common.BytesToHash(crypto.Keccak256([]byte("Romulus(address,uint256,string)")))
	fmt.Printf("Hash = 0x%x", askQuestionTopic)
	fID, err := makeFilter(client, askQuestionTopic)
	if err != nil {
		fmt.Println("shutting down due to ", err)
		log.Fatal(err)
	}
	fmt.Println("Filter ID ", fID)
	ctx := context.Background()
	for {

		logEntries, err := client.FilterChanges(ctx, fID)
		if err != nil {
			log.Fatal("FilterChanges: ", err)
		}
		if len(logEntries) > 0 {

			for _, logEnt := range logEntries {
				if logEnt.Topics[0] != askQuestionTopic {
					continue
				}
				src := common.HexToAddress(logEnt.Topics[1].Hex())
				nonce := new(big.Int).SetBytes(logEnt.Data[0:31])
				question := string(logEnt.Data[96:])
				fmt.Printf("Question number %d from 0x%x : %s\n", nonce, src.Hex(), question)
			}
		} else {
			time.Sleep(time.Second * 5)
		}
	}

}
