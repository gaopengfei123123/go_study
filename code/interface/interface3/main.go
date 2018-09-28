package main

import (
	"encoding/json"
	"log"
)

// 这里是针对出现 map[interface{}]interface{} 类型数据进行的一次转化处理示例
type respBody map[interface{}]interface{}

func main() {

	// 模拟一个从 hprose-php-server 传过来的数据
	res := map[interface{}]interface{}{
		"errorCode": 200,
		"errorMsg":  "登录成功",
		"responseData": map[interface{}]interface{}{
			"hx_password": "c427ee88c8abeeee4fcddbfbf8767025",
			"like_post":   0,
			"avatar":      "http://img2.xxx.com/user/3_100_100.png",
			"beauty_list": []int{1, 2, 3},
			"role":        []string{"admin", "emplyee", "boss"},
			"mission_status": map[interface{}]interface{}{
				"ok": 233,
			},
		},
	}

	tmp := respHandler(res)
	log.Println("tmp:", tmp)

	// log.Println("responseData:", tmp["responseData"])
	// rpData := respHandler(tmp["responseData"])

	by, err := json.Marshal(tmp)
	log.Println("output json:", string(by), err)
}

func respHandler(res interface{}) (tmp map[string]interface{}) {
	// map 需要初始化一个出来
	tmp = make(map[string]interface{})
	log.Println("input res is : ", res)
	switch res.(type) {
	case nil:
		return tmp
	case map[string]interface{}:
		return res.(map[string]interface{})
	case map[interface{}]interface{}:
		log.Println("map[interface{}]interface{} res:", res)
		for k, v := range res.(map[interface{}]interface{}) {
			log.Println("loop:", k, v)
			switch k.(type) {
			case string:
				switch v.(type) {
				case map[interface{}]interface{}:
					log.Println("map[interface{}]interface{} v:", v)
					tmp[k.(string)] = respHandler(v)
					continue
				default:
					log.Printf("default v: %v %v \n", k, v)
					tmp[k.(string)] = v
				}

			default:
				continue
			}
		}
		return tmp
	default:
		// 暂时没遇到更复杂的数据
		log.Println("unknow data:", res)
	}
	return tmp
}
