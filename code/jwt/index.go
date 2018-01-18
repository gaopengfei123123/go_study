package main

import (
	"fmt"
	"encoding/json"
	"encoding/base64"
	"crypto/sha256"
	"encoding/hex"
)

// Header 消息头部
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"JWT"`
}

// PayLoad 负载
type PayLoad struct {
	Sub string `json:"sub"`
	Name string `json:"name"`
	Admin bool `json:"admin"`
}

// JWT 完整的本体
type JWT struct {
	Header	`json:"header"`
	PayLoad	`json:"payload"`
	Signature string `json:"signature"`
}

func main() {
	fmt.Println("this will check jwt")
	jwt := JWT{}
	jwt.Header = Header{"HS256","JWT"}
	jwt.PayLoad = PayLoad{"1234567890","John Doe",true}
	result := jwt.encode("secret")
	fmt.Println(result)
}

func (jwt *JWT) encode(salt string) string {
	header, err := json.Marshal(jwt.Header)
	checkError(err)
	headerString := base64.StdEncoding.EncodeToString(header)
	payload, err := json.Marshal(jwt.PayLoad)
	payloadString := base64.StdEncoding.EncodeToString(payload)
	checkError(err)
	
	format := headerString + "." + payloadString
	formatB := []byte(format)
	saltB := []byte(salt)
	hash := sha256.New()
	hash.Write(formatB)

	md := hash.Sum(saltB)
	signature := hex.EncodeToString(md)

	return format + "." + signature
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}