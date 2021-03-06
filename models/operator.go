package models

const OperatorSeller int = 1
const OperatorBuyer int = 2
const OperatorTransporter int = 3

const PrefixSeller = "sel"
const PrefixBuyer = "buy"
const PrefixTransporter = "tra"

type Authenticate struct {
	*DataModel
	Password     string `json:"password"`
	AttemptFails byte   `json:"attempt_fails"`
	LastAttempts int64  `json:"last_attempts"`
	LastLog      int64  `json:"last_log"`
	Frozen       bool   `json:"frozen"`
}

type Operator struct {
	*Authenticate
	*BalanceHolder
	Description string `json:"description"`
	OperatorType int `json:"operator_type"`
	Token string `json:"token"`
}

func (o *Operator) isSeller() bool {
	return o.OperatorType == OperatorSeller
}

func (o *Operator) isBuyer() bool {
	return o.OperatorType == OperatorBuyer
}

func (o *Operator) isTransporter() bool {
	return o.OperatorType == OperatorTransporter
}

type Seller struct {
	*Operator
}

type Buyer struct {
	*Operator
}

type Transporter struct {
	*Operator
}

func NewOperator(accountType int, password string, initialBalance uint64) *Operator {
	id := AllocateIdS()
	var prefixId string = ""
	switch accountType {
	case OperatorSeller:
		prefixId = PrefixSeller
	case OperatorBuyer:
		prefixId = PrefixBuyer
	case OperatorTransporter:
		prefixId = PrefixTransporter
	default:
		return nil
	}
	operator := &Operator{
		Authenticate: &Authenticate{
			DataModel:    &DataModel{Id: prefixId + id},
			Password:     password,
			LastLog:      0,
			LastAttempts: 0,
			AttemptFails: 0,
			Frozen:       true,
		},
		BalanceHolder: &BalanceHolder{Balance: initialBalance},
		OperatorType:  accountType,
	}
	return operator
}
