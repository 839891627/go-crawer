package parser

import (
	"crawler/engine"
	"regexp"
	"strings"
	"strconv"
	"crawler/model"
)

var profileRe = regexp.MustCompile(`<div[\s.]*class="des f-cl"\s*[^>]*>([^<]*)`)
var idRe = regexp.MustCompile(`http://album.zhenai.com/u/(\d+)`)

func ParseProfile(content []byte, name, url string) engine.ParseResult {
	match := profileRe.FindSubmatch(content)

	profile := model.Profile{}
	profile.Name = name

	// 可能没找到，那么需要跳过
	if len(match) < 2 {
		return engine.ParseResult{}
	}

	fields := strings.Split(string(match[1]), "|")
	age, _ := strconv.Atoi(strings.TrimSuffix(strings.Trim(fields[1], " "), "岁"))
	profile.Age = age

	height, _ := strconv.Atoi(strings.TrimSuffix(strings.Trim(fields[4], " "), "cm"))
	profile.Height = height
	profile.City = string(fields[0])
	profile.Education = string(fields[2])
	profile.Income = string(fields[5])

	id := idRe.FindStringSubmatch(url)[1]

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      id,
				Payload: profile,
			},
		},
	}

	return result
}
