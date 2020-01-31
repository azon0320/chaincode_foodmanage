package models

import (
	"crypto/sha256"
	"fmt"
)

const PrefixToken = "tkn"

type TokenMap struct {
	Token string `json:"token"`
	AccountId string `json:"account_id"`
	CreateTime int64 `json:"create_time"`
}

func GenerateTokenWithTime(stamp int64) string{
	return fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprint(stamp))))
}

func NewTokenMap(token string, id string) *TokenMap{
	return &TokenMap{
		Token:		  PrefixToken + token,
		//Token:      PrefixToken + string(([]rune(token))[:9]),
		AccountId:  id,
		CreateTime: CurrentTimeMillis(),
	}
}

func NewTokenMapWithTime(token string, id string, nanos int64) *TokenMap{
	return &TokenMap{
		//Token:      PrefixToken + string(([]rune(token))[:9]),
		Token:		  PrefixToken + token,
		AccountId:  id,
		CreateTime: CurrentTimeMillis(nanos),
	}
}