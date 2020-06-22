package main

import (
	"crypto/sha256"
	"time"
)

//1. 定义区块结构
type Block struct {
	//1.版本号
	Version uint64
	//2. 前区块hash
	PrevHash []byte
	//3.Merkel根
	MerkelRoot []byte
	//4.时间戳
	TimeStamp uint64
	//5.难度值
	Difficulty uint64
	//a. 当前区块hash 正常比特币区块中没有当前区块的哈希
	Hash []byte
	//b. 数据
	Data []byte
}

//2.创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		Version:    00,
		PrevHash:   prevBlockHash,
		MerkelRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0,
		Hash:       []byte{}, //先空，后面计算hash  //TODO
		Data:       []byte(data),
	}
	block.SetHash()
	return &block
}

//辅助函数 将uint64 转[]byte
func Uint64ToByte(num uint64) []byte {
	return []byte{}
}

//3.生成hash
func (block *Block) SetHash() {
	var blockInfo []byte
	//1. 拼装数据
	blockInfo = append(blockInfo, Uint64ToByte(block.Version)...)
	blockInfo = append(blockInfo, block.PrevHash...)
	blockInfo = append(blockInfo, block.MerkelRoot...)
	blockInfo = append(blockInfo, Uint64ToByte(block.TimeStamp)...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Difficulty)...)
	blockInfo = append(blockInfo, block.Hash...)
	blockInfo = append(blockInfo, block.Data...)

	//2.sha256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}
