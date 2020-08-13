/*
 * Copyright (c) 2013 IBM Corp.
 *
 * All rights reserved. This program and the accompanying materials
 * are made available under the terms of the Eclipse Public License v1.0
 * which accompanies this distribution, and is available at
 * http://www.eclipse.org/legal/epl-v10.html
 *
 * Contributors:
 *    Seth Hoenig
 *    Allan Stockdill-Mander
 *    Mike Robertson
 */

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/eclipse/paho.mqtt.golang"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {
	// 打印程序运行结果, 感觉干扰视线就把 mqtt.DEBUG个注释掉
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	// 这里使用的本地地址
	opts := mqtt.NewClientOptions().AddBroker("tcp://0.0.0.0:1883").SetClientID("gotrivial")
	opts.SetKeepAlive(2 * time.Second)

	// 这里需要注入一个client收到消息后对消息处理的方法
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	// 启动一个链接
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 订阅一个topic
	if token := c.Subscribe("go-mqtt/sample", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// 向topic发布消息, SetDefaultPublishHandler里收消息
	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("this is msg #%d!", i)
		token := c.Publish("go-mqtt/sample", 0, false, text)
		token.Wait()
	}

	time.Sleep(6 * time.Second)

	if token := c.Unsubscribe("go-mqtt/sample"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	c.Disconnect(250)

	time.Sleep(1 * time.Second)
}
