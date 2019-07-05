package parser

import (
	"crawler/engine"
	"regexp"
)

var profilRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]*)</a>`)
var cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)

func ParseCity(contents []byte) engine.ParseResult {
	matches := profilRe.FindAllSubmatch(contents, -1)

	//var result engine.ParseResult
	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		url := string(m[1])
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(contents []byte) engine.ParseResult {
				return ParseProfile(contents, name, url)
			},
		})
		//result.Items = append(result.Items, "User "+name)
	}
	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
