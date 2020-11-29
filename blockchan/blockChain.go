package blockchain

import (
	"DataCertproject/models"
	"errors"
	"github.com/boltdb/bolt"
	"math/big"
)
var BUCKET_NAME = "blocks"
var LAST_KEY  ="lasthash"
var CHANINDB = "chain.db"

type BlockChain struct {
	LastHash []byte//最新区块的hash
	BoltDb *bolt.DB
}

var CHAIN  BlockChain
//
func (bc BlockChain)QueryAllbolks() []*Block  {
	bloks:=make([]*Block,0)
	db:=bc.BoltDb
	db.View(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte(BUCKET_NAME))
		if bucket== nil{
			panic("查询数据出错")
		}
		enchKey := bc.LastHash
		preHashBig := new(big.Int)
		zeroBig :=big.NewInt(0)
		for{
			enchBlockBytes:=bucket.Get(enchKey)
			enchBlock,_:=Deserialze(enchBlockBytes)
			bloks =append(bloks,enchBlock)
			preHashBig=preHashBig.SetBytes(enchBlock.PrevHash)
			if preHashBig.Cmp(zeroBig)==0 {
				break
			}
			enchKey=enchBlock.PrevHash
		}

		return nil
	})
	return bloks
}
func NewBlockChain() BlockChain {
	db,err:=bolt.Open(CHANINDB,0600,nil)
	if err!=nil{
		panic(err.Error())
	}
	var bl BlockChain
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(BUCKET_NAME))
			if err != nil {
				panic(err.Error())
			}
		}
		lastHash := bucket.Get([]byte(LAST_KEY))
		if len(lastHash)==0{
			genesis:=CreateGenesisBlock()
			bl=BlockChain{
				LastHash: genesis.Hash,
				BoltDb:  db,
			}
			genesisBytes,_:=genesis.Serialze()
			bucket.Put(genesis.Hash,genesisBytes)
			bucket.Put([]byte(LAST_KEY),genesis.Hash)
		}else {
			lastHash := bucket.Get([]byte(LAST_KEY))
			lastBlockBytes:=bucket.Get(lastHash)
			lastBlock,err:=Deserialze(lastBlockBytes)
			if err != nil {
				panic("读取区块链数据失败")
			}
			bl =BlockChain{
				LastHash: lastBlock.Hash,
				BoltDb:   db,
			}
		}
		return nil
	})
	CHAIN = bl
	return bl
}
func(bc BlockChain)QueryBlockHeight(height int64)*Block{
		if height < 0{
			return nil
		}
		var block *Block
		db := bc.BoltDb
		db.View(func(tx *bolt.Tx) error {
			bucket:=tx.Bucket([]byte(BUCKET_NAME))
			if bucket==nil {
			panic("查询数据失败")
			}
			hashkey:=bc.LastHash
			for  {
				lastBlockbytes:=bucket.Get(hashkey)
				enchBlock,_:=Deserialze(lastBlockbytes)
				if enchBlock.Height<height{
					break
				}
				if enchBlock.Height == height{
					block = enchBlock
					break
				}
				hashkey=enchBlock.Hash
			}
		return nil
		})
	return block
}

func (bc BlockChain)QueryBlockByCertId(cert_id []byte)(*Block, error ){
  var block *Block
  db:=bc.BoltDb
	var err error
  db.View(func(tx *bolt.Tx) error {
	   bucket:=tx.Bucket([]byte(BUCKET_NAME))
	   if bucket == nil{
	   	err =errors.New("查询区块数据遇到错误！")
		return err
	   }
			enchHash:=bucket.Get([]byte(LAST_KEY))
			eacBig:=new(big.Int)
			zero:=big.NewInt(0)
			var certRecord *models.CertRecord
	   for  {
		   enchBlockBytes:= bucket.Get(enchHash)
		   enchBolck,_:= Deserialze(enchBlockBytes)
		   certRecord ,err = models.DeserializeRecord(enchBolck.Data)
		 //  fmt.Println(err)
		 //  fmt.Println(enchBolck.Data)
		 //  fmt.Println(certRecord)
		   if string(certRecord.CertId) == string(cert_id) {
				block = enchBolck
				break
		   }
			eacBig=eacBig.SetBytes(enchBolck.Hash)
			if eacBig.Cmp(zero)==0{
			break
			}
			enchHash = enchBolck.PrevHash
	   }
	   return nil
  })
  return block ,err
}
func (bc BlockChain)SavaData(data []byte) (Block,error) {
	db:=bc.BoltDb
	var e  error
	var lastBlock  *Block
	db.View(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte(BUCKET_NAME))
		if bucket ==nil{
		e = errors.New("boltdb未创建，请重试!")
		return e
		}
		lastBlockBytes:=bucket.Get(bc.LastHash)
		lastBlock,_=Deserialze(lastBlockBytes)
		return nil
	})
	newBlock:=NewBlock(lastBlock.Height+1,data,lastBlock.Hash)
	db.Update(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte(BUCKET_NAME))
		newBlockBytes,_:=newBlock.Serialze()
		bucket.Put(newBlock.Hash,newBlockBytes)
		bucket.Put([]byte(LAST_KEY),newBlock.Hash)
		bc.LastHash =newBlock.Hash
		return nil
	})
	return newBlock,e
}
