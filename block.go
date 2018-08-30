package main

import (
	"time"
)

type Block struct {
	Version int64  //版本
	PrevBlockHash []byte  //前区块哈希值
	Hash []byte //当前区块的哈希
	MerkelRoot []byte //梅克尔跟
	TimeStamp int64  //时间截
	Bits int64  //难度值
	Nonce int64 //随机值

	//交易信息
	Data []byte
}

func NewBlock(data string,prevBlockHash []byte) *Block{
	var block = Block{
		Version:1,
		PrevBlockHash:prevBlockHash,
		//Hash TODO
		MerkelRoot:[]byte{},
		TimeStamp:time.Now().Unix(),
		Bits:targetBits,
		Nonce:0,
		Data:[]byte(data)}
	//block.SetHash()
	pow:=NewProofOfWork(&block)
	nonce,hash :=pow.Run()
	block.Nonce = nonce
	block.Hash = hash
	return &block
}
/*
func (block *Block)SetHash(){
	tmp:=[][]byte{
		IntToByte(block.Version),
		block.PrevBlockHash,
		block.MerkelRoot,
		IntToByte(block.TimeStamp),
		IntToByte(block.Bits),
		IntToByte(block.Nonce),
		block.Data}
	data:=bytes.Join(tmp,[]byte{})
	hash:=sha256.Sum256(data)
	block.Hash = hash[:]
}
*/

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block",[]byte{})
}