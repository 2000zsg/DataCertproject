package blockchain

import (
	"bytes"
	"encoding/gob"
	"time"
)

/**
 *  区块结构体的定义
 */
type Block struct {
	Height    int64  //区块高度
	TimeStamp int64  //时间戳
	Hash      []byte //区块的hash
	Data      []byte // 数据
	PrevHash  []byte //上一个区块的Hash
	Version   string //版本号
	Nonce     int64  //随机数，用于pow工作量证明算法计算
}
func CreateGenesisBlock() Block  {
	block:=NewBlock(0,[]byte{},[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
	return block
}
/**
 * 新建一个区块实例，并返回该区块
 */
func NewBlock(height int64, data []byte, prevHash []byte) (Block) {
	//1、构建一个block实例，用于生成区块
	block := Block{
		Height:    height,
		TimeStamp: time.Now().Unix(),
		Data:      data,
		PrevHash:  prevHash,
		Version:   "0x01",
	}
	pow := NewPOW(block)
	blockHash, nonce := pow.Run()
	block.Nonce = nonce
	block.Hash = blockHash
		//heightBytes, _ := util.IntToBytes(block.Height)
		//timeBytes, _ := util.IntToBytes(block.TimeStamp)
		//versionBytes := util.StringToBytes(block.Version)
		//nonceBytes, _ := util.IntToBytes(block.Nonce)
		////bytes.Join函数，用于[]byte的拼接
		//blockBytes := bytes.Join([][]byte{
		//	heightBytes,
		//	timeBytes,
		//	data,
		//	prevHash,
		//	versionBytes,
		//	nonceBytes,
		//}, []byte{})
		//
		////4、设置第7个字段hash
		//block.Hash = util.SHA256Hash(blockBytes)

		return block
}
//序列化
func (bk Block)Serialze()([]byte,error){
	buff:=new(bytes.Buffer)
	err:=gob.NewEncoder(buff).Encode(bk)
	if  err!=nil {
		return nil,err
	}
	return buff.Bytes(),nil
}
//反序列化
func Deserialze(data []byte)(*Block,error)  {
	var block Block
	err := gob.NewDecoder(bytes.NewReader(data)).Decode(&block)
		if err != nil {
			return nil,err
		}
	return &block,nil
}