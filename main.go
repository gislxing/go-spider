package main

import (
	"zbs/go-spider/engine"
	"zbs/go-spider/zhengai"
)

const startUrl = `http://www.zhenai.com/zhenghun`

func main() {
	engine.Run(engine.Request{
		Url:       startUrl,
		ParseFunc: zhengai.ParseCityList,
	})
}
