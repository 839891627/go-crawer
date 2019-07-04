package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {

	//simpleEngine := engine.SimpleEngine{}
	//simpleEngine.Run(engine.Request{
	//	Url:        "https://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	e := &engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		//Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/beijing",
	//	ParserFunc: parser.ParseCity,
	//})
}
