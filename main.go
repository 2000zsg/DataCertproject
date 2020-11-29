package main

import (
	blockchain "DataCertproject/blockchan"
	"DataCertproject/qkl_mysql"
	_ "DataCertproject/routers"
	"github.com/astaxie/beego"
)

func main() {
	//fmt.Println("====")
	blockchain.NewBlockChain()

	qkl_mysql.Qkl()
	//fmt.Println(bc)
	//fmt.Printf("最新的Hash值:%x\n", bc.LastHash)
	//block,err :=bc.SavaData([]byte("这里储存的是上链的数据信息"))
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	////fmt.Printf("区块:%x\n", block)
	//fmt.Printf("区块的高度:%d\n", block.Height)
	//fmt.Printf("区块的PrevHash:%x\n", block.PrevHash)
	//fmt.Printf("区块的Hash值：%x\n",block.Hash)
	//fmt.Printf("区块的nonce值：%d\n",block.Nonce)
	//return
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")

	beego.Run()
}
