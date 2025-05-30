package blocks

import (
	"context"
	"fmt"

	"github.com/blockfrost/blockfrost-go"
)

func GetPoolBlocksForEpoch(poolId string, blockfrostClient blockfrost.APIClient) int {

	epoch, err := blockfrostClient.EpochLatest(context.TODO())

	if err != nil {
		fmt.Println(err)
	}

	blocks, err := blockfrostClient.EpochBlockDistributionByPool(context.TODO(), epoch.Epoch, poolId, blockfrost.APIQueryParams{})

	if err != nil {
		fmt.Println(err)
	}

	return len(blocks)

}
