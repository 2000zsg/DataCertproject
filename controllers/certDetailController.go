package controllers

import (
	"DataCertproject/blockchan"
	"DataCertproject/models"
	"DataCertproject/util"
	"github.com/astaxie/beego"
	"strings"
)

type CerDetailController struct {
	beego.Controller
}

func (c *CerDetailController) Get() {
	certId := c.GetString("cert_id")
	//fmt.Println("要查询的认证id")
	block, err := blockchain.CHAIN.QueryBlockByCertId([]byte(certId))
	if err != nil {
		c.Ctx.WriteString("链上数据查询失败！")
		return
	}
	if block == nil {
		c.Ctx.WriteString("抱歉，未查到链上数据，请重试！")
		return
	}
	//certId = hex.EncodeToString(block.Data)
	certRecord, err := models.DeserializeRecord(block.Data)
	certRecord.CertHashSth = string(certRecord.CertHash)
	certRecord.CerIDStr = strings.ToUpper(string(certRecord.CertId))
	certRecord.CertTimeFormat = util.TimeFormat(certRecord.CertTime, 0, util.TIME_FORMAT_THREE)
	if err != nil {
		return
	}
	c.Data["CertRecord"] = certRecord
	c.Data["Phone"] = certRecord.Phone
	c.TplName = "cert_detail.html"

}
