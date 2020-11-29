package blockchain

import (
	"DataCertproject/util"
	"bytes"
	"crypto/sha256"
	"math/big"
)

const DIFFFICULTY  = 20
type Proofofwork struct {
	Target *big.Int
	Block Block
}

func NewPOW(block Block) Proofofwork {
	target:=big.NewInt(1)
	target.Lsh(target,255-DIFFFICULTY)
	pow := Proofofwork{
		Target:target,
		Block:block,
	}
	return pow
}
func (p Proofofwork) Run() ([]byte,int64) {
	var nonce  int64
	//var bigblock * big.Int
	bigblock := new(big.Int)
	var block256Hash []byte
	for{
		block :=p.Block
		heightBytes,_:=util.IntToBytes(block.Height)
		TimeBytes,_:=util.IntToBytes(block.TimeStamp)
		versionBytes :=util.StringToBytes(block.Version)
		nonceBytes,_ :=util.IntToBytes(nonce)
		blockBytes:=bytes.Join([][]byte{
			heightBytes,
			TimeBytes,
			versionBytes,
			nonceBytes,
			block.Data,
			block.PrevHash,
		},[]byte{})
		SHA256Hash := sha256.New()
		SHA256Hash.Write(blockBytes)
		block256Hash = SHA256Hash.Sum(nil)
		//fmt.Printf("挖矿中，当前尝试nonce值：%d\n",nonce)
		bigblock = bigblock.SetBytes(block256Hash)
		//fmt.Printf("目标值：%x\n",p.Target)
		//fmt.Printf("hash值：%x\n",bigblock)
		if p.Target.Cmp(bigblock)==1 {
			break
		}
		nonce++
	}

	return block256Hash, nonce
}