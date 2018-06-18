package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	crypto "github.com/tendermint/go-crypto"
	//"github.com/cosmos/cosmos-sdk/wire"
	"github.com/cosmos/cosmos-sdk/x/auth"
	//"github.com/cosmos/cosmos-sdk/x/stake"
)

/*
var _ auth.Account = (*AppAccount)(nil)

// Custom extensions for this application.  This is just an example of
// extending auth.BaseAccount with custom fields.
//
// This is compatible with the stock auth.AccountStore, since
// auth.AccountStore uses the flexible go-amino library.
type AppAccount struct {
	auth.BaseAccount
	Name string `json:"name"`
}

// nolint
func (acc AppAccount) GetName() string      { return acc.Name }
func (acc *AppAccount) SetName(name string) { acc.Name = name }

// Get the AccountDecoder function for the custom AppAccount
func GetAccountDecoder(cdc *wire.Codec) auth.AccountDecoder {
	return func(accBytes []byte) (res auth.Account, err error) {
		if len(accBytes) == 0 {
			return nil, sdk.ErrTxDecode("accBytes are empty")
		}
		acct := new(AppAccount)
		err = cdc.UnmarshalBinaryBare(accBytes, &acct)
		if err != nil {
			panic(err)
		}
		return acct, err
	}
}
*/
//___________________________________________________________________________________

type GenTx struct {
	Address sdk.Address   `json:"address"`
	PubKey  crypto.PubKey `json:"pub_key"`
}

// State to Unmarshal
type GenesisState struct {
	Accounts []GenesisAccount `json:"accounts"`
	//StakeData stake.GenesisState `json:"stake"`
}

// GenesisAccount doesn't need pubkey or sequence
type GenesisAccount struct {
	//Name    string      `json:"name"`
	Address sdk.Address   `json:"address"`
	Coins   sdk.Coins     `json:"coins"`
	PubKey  crypto.PubKey `json:"pub_key"` //add in pub key so I can send coins?
}

func NewGenesisAccount(aa *auth.BaseAccount) GenesisAccount {
	return GenesisAccount{
		//Name:    aa.Name,
		Address: aa.Address,
		Coins:   aa.Coins.Sort(),
		PubKey:  aa.PubKey, // add in pub key
	}
}

// convert GenesisAccount to AppAccount
func (ga *GenesisAccount) ToAppAccount() (acc *auth.BaseAccount, err error) {
	return &auth.BaseAccount{
		Address: ga.Address,
		Coins:   ga.Coins.Sort(),
		PubKey:  ga.PubKey, // add in pub key
	}, nil
}