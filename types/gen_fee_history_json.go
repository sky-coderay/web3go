// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*feeHistoryMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (f FeeHistory) MarshalJSON() ([]byte, error) {
	type FeeHistory struct {
		OldestBlock      *hexutil.Big           `json:"oldestBlock"`
		Reward           blockRewardsMarshaling `json:"reward,omitempty"`
		BaseFee          []*hexutil.Big         `json:"baseFeePerGas,omitempty"`
		GasUsedRatio     []float64              `json:"gasUsedRatio"`
		BlobBaseFee      []*hexutil.Big         `json:"baseFeePerBlobGas,omitempty"`
		BlobGasUsedRatio []float64              `json:"blobGasUsedRatio,omitempty"`
	}
	var enc FeeHistory
	enc.OldestBlock = (*hexutil.Big)(f.OldestBlock)
	enc.Reward = f.Reward
	if f.BaseFee != nil {
		enc.BaseFee = make([]*hexutil.Big, len(f.BaseFee))
		for k, v := range f.BaseFee {
			enc.BaseFee[k] = (*hexutil.Big)(v)
		}
	}
	enc.GasUsedRatio = f.GasUsedRatio
	if f.BlobBaseFee != nil {
		enc.BlobBaseFee = make([]*hexutil.Big, len(f.BlobBaseFee))
		for k, v := range f.BlobBaseFee {
			enc.BlobBaseFee[k] = (*hexutil.Big)(v)
		}
	}
	enc.BlobGasUsedRatio = f.BlobGasUsedRatio
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (f *FeeHistory) UnmarshalJSON(input []byte) error {
	type FeeHistory struct {
		OldestBlock      *hexutil.Big            `json:"oldestBlock"`
		Reward           *blockRewardsMarshaling `json:"reward,omitempty"`
		BaseFee          []*hexutil.Big          `json:"baseFeePerGas,omitempty"`
		GasUsedRatio     []float64               `json:"gasUsedRatio"`
		BlobBaseFee      []*hexutil.Big          `json:"baseFeePerBlobGas,omitempty"`
		BlobGasUsedRatio []float64               `json:"blobGasUsedRatio,omitempty"`
	}
	var dec FeeHistory
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.OldestBlock != nil {
		f.OldestBlock = (*big.Int)(dec.OldestBlock)
	}
	if dec.Reward != nil {
		f.Reward = *dec.Reward
	}
	if dec.BaseFee != nil {
		f.BaseFee = make([]*big.Int, len(dec.BaseFee))
		for k, v := range dec.BaseFee {
			f.BaseFee[k] = (*big.Int)(v)
		}
	}
	if dec.GasUsedRatio != nil {
		f.GasUsedRatio = dec.GasUsedRatio
	}
	if dec.BlobBaseFee != nil {
		f.BlobBaseFee = make([]*big.Int, len(dec.BlobBaseFee))
		for k, v := range dec.BlobBaseFee {
			f.BlobBaseFee[k] = (*big.Int)(v)
		}
	}
	if dec.BlobGasUsedRatio != nil {
		f.BlobGasUsedRatio = dec.BlobGasUsedRatio
	}
	return nil
}
