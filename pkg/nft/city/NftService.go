package city

import (
	"composable-nfts-service/logging"
	citynft "composable-nfts-service/pkg/nft/city/abi"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
)

type RolledEvent struct {
	requestId big.Int
	roller    common.Address
}

type RandomResultEvent struct {
	requestId big.Int
	result    big.Int
}

var (
	rolledEventSig       = crypto.Keccak256Hash([]byte("Rolled(uint256,address)")).Hex()
	randomResultEventSig = crypto.Keccak256Hash([]byte("RandomResult(uint256,uint256)")).Hex()
)

func eventListener() {
	client, err := ethclient.DialContext(context.Background(), "https://rpc-testnet.morphl2.io")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		logging.Error(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(citynft.CitynftMetaData.ABI)))
	if err != nil {
		logging.Error(err)
	}

	for _, vLog := range logs {
		fmt.Println(vLog.BlockHash.Hex())
		fmt.Println(vLog.BlockNumber)
		fmt.Println(vLog.TxHash.Hex())

		if len(vLog.Topics) > 0 {
			switch vLog.Topics[0].Hex() {
			case rolledEventSig:
				var event RolledEvent
				err := contractAbi.UnpackIntoInterface(&event, "Rolled", vLog.Data)
				if err != nil {
					logging.Error("Error unpacking rolled event: ", err)
					continue
				}
				fmt.Println("Rolled event: ", event.requestId, hexutil.Encode(event.roller.Bytes()))

			case randomResultEventSig:
				var event RandomResultEvent
				err := contractAbi.UnpackIntoInterface(&event, "RandomResult", vLog.Data)
				if err != nil {
					logging.Error("Error unpacking randomResult event: ", err)
					continue
				}
				fmt.Println("RandomResult event: ", event.requestId, event.result)
			}
		}
	}

}
