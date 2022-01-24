// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

var _ = (*receiptMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (r Receipt) MarshalJSON() ([]byte, error) {
	type Receipt struct {
		BlockHash         common.Hash     `json:"blockHash"`
		BlockNumber       hexutil.Uint64  `json:"blockNumber"`
		ContractAddress   *common.Address `json:"contractAddress"`
		CumulativeGasUsed hexutil.Uint64  `json:"cumulativeGasUsed"`
		EffectiveGasPrice hexutil.Uint64  `json:"effectiveGasPrice"`
		From              common.Address  `json:"from"`
		GasUsed           hexutil.Uint64  `json:"gasUsed"`
		Logs              []*Log          `json:"logs"`
		LogsBloom         types.Bloom     `json:"logsBloom"`
		Root              hexutil.Bytes   `json:"root"`
		Status            hexutil.Uint64  `json:"status"`
		To                *common.Address `json:"to"`
		TransactionHash   common.Hash     `json:"transactionHash"`
		TransactionIndex  hexutil.Uint64  `json:"transactionIndex"`
		TxExecErrorMsg    *string         `json:"txExecErrorMsg"`
		Type              hexutil.Uint    `json:"type"`
	}
	var enc Receipt
	enc.BlockHash = r.BlockHash
	enc.BlockNumber = hexutil.Uint64(r.BlockNumber)
	enc.ContractAddress = r.ContractAddress
	enc.CumulativeGasUsed = hexutil.Uint64(r.CumulativeGasUsed)
	enc.EffectiveGasPrice = hexutil.Uint64(r.EffectiveGasPrice)
	enc.From = r.From
	enc.GasUsed = hexutil.Uint64(r.GasUsed)
	enc.Logs = r.Logs
	enc.LogsBloom = r.LogsBloom
	enc.Root = r.Root
	enc.Status = hexutil.Uint64(r.Status)
	enc.To = r.To
	enc.TransactionHash = r.TransactionHash
	enc.TransactionIndex = hexutil.Uint64(r.TransactionIndex)
	enc.TxExecErrorMsg = r.TxExecErrorMsg
	enc.Type = hexutil.Uint(r.Type)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (r *Receipt) UnmarshalJSON(input []byte) error {
	type Receipt struct {
		BlockHash         *common.Hash    `json:"blockHash"`
		BlockNumber       *hexutil.Uint64 `json:"blockNumber"`
		ContractAddress   *common.Address `json:"contractAddress"`
		CumulativeGasUsed *hexutil.Uint64 `json:"cumulativeGasUsed"`
		EffectiveGasPrice *hexutil.Uint64 `json:"effectiveGasPrice"`
		From              *common.Address `json:"from"`
		GasUsed           *hexutil.Uint64 `json:"gasUsed"`
		Logs              []*Log          `json:"logs"`
		LogsBloom         *types.Bloom    `json:"logsBloom"`
		Root              *hexutil.Bytes  `json:"root"`
		Status            *hexutil.Uint64 `json:"status"`
		To                *common.Address `json:"to"`
		TransactionHash   *common.Hash    `json:"transactionHash"`
		TransactionIndex  *hexutil.Uint64 `json:"transactionIndex"`
		TxExecErrorMsg    *string         `json:"txExecErrorMsg"`
		Type              *hexutil.Uint   `json:"type"`
	}
	var dec Receipt
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.BlockHash != nil {
		r.BlockHash = *dec.BlockHash
	}
	if dec.BlockNumber != nil {
		r.BlockNumber = uint64(*dec.BlockNumber)
	}
	if dec.ContractAddress != nil {
		r.ContractAddress = dec.ContractAddress
	}
	if dec.CumulativeGasUsed != nil {
		r.CumulativeGasUsed = uint64(*dec.CumulativeGasUsed)
	}
	if dec.EffectiveGasPrice != nil {
		r.EffectiveGasPrice = uint64(*dec.EffectiveGasPrice)
	}
	if dec.From != nil {
		r.From = *dec.From
	}
	if dec.GasUsed != nil {
		r.GasUsed = uint64(*dec.GasUsed)
	}
	if dec.Logs != nil {
		r.Logs = dec.Logs
	}
	if dec.LogsBloom != nil {
		r.LogsBloom = *dec.LogsBloom
	}
	if dec.Root != nil {
		r.Root = *dec.Root
	}
	if dec.Status != nil {
		r.Status = uint64(*dec.Status)
	}
	if dec.To != nil {
		r.To = dec.To
	}
	if dec.TransactionHash != nil {
		r.TransactionHash = *dec.TransactionHash
	}
	if dec.TransactionIndex != nil {
		r.TransactionIndex = uint64(*dec.TransactionIndex)
	}
	if dec.TxExecErrorMsg != nil {
		r.TxExecErrorMsg = dec.TxExecErrorMsg
	}
	if dec.Type != nil {
		r.Type = uint(*dec.Type)
	}
	return nil
}
