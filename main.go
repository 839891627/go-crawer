package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/persist"
	"crawler/zhenai/parser"
)

func main() {

	//simpleEngine := engine.SimpleEngine{}
	//simpleEngine.Run(engine.Request{
	//	Url:        "https://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	itemChan, err := persist.ItemSaver()
	if err != nil {
		panic(err)
	}
	e := &engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		//Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{
		//Url:        "https://www.zhenai.com/zhenghun",
		//ParserFunc: parser.ParseCityList,
		Url:        "http://www.zhenai.com/zhenghun/beijing",
		ParserFunc: parser.ParseCity,
	})
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/beijing",
	//	ParserFunc: parser.ParseCity,
	//})
}
