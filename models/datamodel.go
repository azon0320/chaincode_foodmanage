package models

import (
	"strconv"
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

// 毫秒级对象ID
func AllocateId(nanos ...int64) int64 {
	return CurrentTxNanos(nanos...) / 1e6
}
func AllocateIdS(nanos ...int64) string { return strconv.Itoa(int(AllocateId(nanos...))) }

func CurrentTxNanos(nanos ...int64) int64 {
	if len(nanos) > 0 {
		return nanos[0]
	}else{
		return GetTxTimeNanos()
	}
}

func CurrentTimeMillis(nanos ...int64) int64{
	return CurrentTxNanos(nanos...) / 1e6
}
