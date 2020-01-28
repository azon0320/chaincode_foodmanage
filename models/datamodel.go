package models

import (
	"strconv"
	"time"
)

/*
 * 非常基本的 数据模型，包括一个 ID 作为键
 */
type DataModel struct {
	Id string `json:"id"`
}

type BalanceHolder struct {
	Balance uint64 `json:"balance"`
}

type Credentials struct {
	AccountId string `json:"account_id"`
	Password  string `json:"password"`
}

// 毫秒级对象ID
func allocateId() int64 {
	return time.Now().UnixNano() / 1e6
}
func allocateIdS() string { return strconv.Itoa(int(allocateId())) }

func CurrentTimeMillis() int64 {
	return allocateId()
}
