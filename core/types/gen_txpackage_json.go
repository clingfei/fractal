// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"github.com/fractal-platform/fractal/common"
	"github.com/fractal-platform/fractal/common/hexutil"
	"math/big"
)

var _ = (*pkgDataMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (p pkgData) MarshalJSON() ([]byte, error) {
	type pkgData struct {
		Packer        common.Address `json:"packer"`
		PackNonce     hexutil.Uint64 `json:"nonce"`
		Transactions  []*Transaction `json:"transactions"`
		BlockFullHash common.Hash    `json:"blockFullHash"`
		R             *hexutil.Big   `json:"r"`
		S             *hexutil.Big   `json:"s"`
		V             *hexutil.Big   `json:"v"`
		Hash          *common.Hash   `json:"hash" rlp:"-"`
	}
	var enc pkgData
	enc.Packer = p.Packer
	enc.PackNonce = hexutil.Uint64(p.PackNonce)
	enc.Transactions = p.Transactions
	enc.BlockFullHash = p.BlockFullHash
	enc.R = (*hexutil.Big)(p.R)
	enc.S = (*hexutil.Big)(p.S)
	enc.V = (*hexutil.Big)(p.V)
	enc.Hash = p.Hash
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (p *pkgData) UnmarshalJSON(input []byte) error {
	type pkgData struct {
		Packer        *common.Address `json:"packer"`
		PackNonce     *hexutil.Uint64 `json:"nonce"`
		Transactions  []*Transaction  `json:"transactions"`
		BlockFullHash *common.Hash    `json:"blockFullHash"`
		R             *hexutil.Big    `json:"r"`
		S             *hexutil.Big    `json:"s"`
		V             *hexutil.Big    `json:"v"`
		Hash          *common.Hash    `json:"hash" rlp:"-"`
	}
	var dec pkgData
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Packer != nil {
		p.Packer = *dec.Packer
	}
	if dec.PackNonce != nil {
		p.PackNonce = uint64(*dec.PackNonce)
	}
	if dec.Transactions != nil {
		p.Transactions = dec.Transactions
	}
	if dec.BlockFullHash != nil {
		p.BlockFullHash = *dec.BlockFullHash
	}
	if dec.R != nil {
		p.R = (*big.Int)(dec.R)
	}
	if dec.S != nil {
		p.S = (*big.Int)(dec.S)
	}
	if dec.V != nil {
		p.V = (*big.Int)(dec.V)
	}
	if dec.Hash != nil {
		p.Hash = dec.Hash
	}
	return nil
}