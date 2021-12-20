package types

// TokenType -> Type for the auth token
type TokenType struct {
	Token string `json:"token"`
}

// AuthHeaderType -> Type for bearer token auth header
type AuthHeaderType struct {
	Token string `header:"Authorization"`
}

// TxParams -> Type for Transaction Parameters
type TxParams struct {
	From   string
	Value  uint64
	Params []Params
	To     string
	Fn     string
}

// Params for TxParams
type Params struct {
	InternalType string `json:"internalType"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Value        string `json:"value"`
}

// ReqBody is Unsigned Transaction Request
type ReqBody struct {
	From        string `json:"from"`
	ToAddress   string `json:"toAddr,omitempty"`
	TokenAmount string `json:"tokenAmount,omitempty"`
	Sender      string `json:"sender,omitempty"`
	Pubkey      string `json:"pubkey,omitempty"`
	Spender     string `json:"spender,omitempty"`
}
