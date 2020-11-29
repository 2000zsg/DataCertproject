package models

import (
	"bytes"
	"encoding/gob"
)

type CertRecord struct {
	CertHash   []byte
	CertHashSth string
	CertId     []byte
	CerIDStr string
	CertName   string
	Phone      string
	AuthorCard string
	FileName   string
	FileSize   int64
	CertTime   int64
	CertTimeFormat string
}

func (c CertRecord) SerializeRecord() ([]byte, error) {
	buff := new(bytes.Buffer)
	err := gob.NewEncoder(buff).Encode(c)
	return buff.Bytes(), err
}
func DeserializeRecord(data []byte) (*CertRecord, error) {
	var certRecord *CertRecord
	err := gob.NewDecoder(bytes.NewReader(data)).Decode(&certRecord)
	//fmt.Println("============================")
	//fmt.Println(certRecord)
	//fmt.Println(data)
	return certRecord, err
}
