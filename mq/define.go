package mq

import "CloudWebOfGin/common"

type TransferData struct {
	FileHash string
	CurLocation string
	DestLocation string
	DestStoreType common.StoreType
}