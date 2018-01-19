package main

import (
	"fmt"
	"encoding/json"
	"encoding/base64"
	"crypto/sha256"
	"crypto/hmac"
	"encoding/hex"
	"strings"
)

// SALT 密钥
const SALT = "SECRET"

// Header 消息头部
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
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

	if jwt.decode(result) {
		fmt.Println("data verify success")
	} else {
		fmt.Println("something was wrong")
	}
}

func (jwt *JWT) encode(salt string) string {
	header, err := json.Marshal(jwt.Header)
	checkError(err)
	headerString := base64.StdEncoding.EncodeToString(header)
	payload, err := json.Marshal(jwt.PayLoad)
	payloadString := base64.StdEncoding.EncodeToString(payload)
	checkError(err)
	
	format := headerString + "." + payloadString
    signature := getHmacCode(format)

	return format + "." + signature
}

func getHmacCode(s string) string {
    h := hmac.New(sha256.New, []byte("secret"))
	h.Write([]byte(s))
	key := h.Sum(nil)
    return hex.EncodeToString(key)
}


func (jwt *JWT) decode( code string) bool {

	arr := strings.Split(code,".")
	if len(arr) != 3 {
		return false
	}

	// 验证签名是否正确
	format := arr[0] + "." + arr[1]
	signature := getHmacCode(format)
	if signature != arr[2] {
		return false
	}


	header, err := base64.StdEncoding.DecodeString(arr[0])
	checkError(err)
	payload, err := base64.StdEncoding.DecodeString(arr[1])
	checkError(err)

	
	json.Unmarshal(header, &jwt.Header)
	json.Unmarshal(payload,&jwt.PayLoad)

	return true
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}