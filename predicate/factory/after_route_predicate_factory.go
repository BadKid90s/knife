package factory

import (
	"errors"
	"gateway/definition"
	"gateway/predicate"
	"gateway/web"
	"log"
	"strings"
	"time"
)

type AfterRoutePredicateFactory struct{}

func (f *AfterRoutePredicateFactory) Apply(definition *definition.PredicateDefinition) (predicate.Predicate[*web.ServerWebExchange], error) {
	config, err := f.parseConfig(definition)
	if err != nil {
		return nil, err
	}
	//return nil
	return f.apply(config), nil

}
func (f *AfterRoutePredicateFactory) parseConfig(definition *definition.PredicateDefinition) (*AfterPredicateConfig, error) {

	val := definition.Args["pattern"]
	t := parseTime(val)
	return &AfterPredicateConfig{
		time: t,
	}, nil
}

func (f *AfterRoutePredicateFactory) apply(config *AfterPredicateConfig) predicate.Predicate[*web.ServerWebExchange] {
	return &AfterPredicate[*web.ServerWebExchange]{
		time: config.time,
	}
}

type AfterPredicate[T any] struct {
	predicate.DefaultPredicate[T]
	time time.Time
}

func (p *AfterPredicate[T]) Apply(T) bool {
	return time.Now().After(p.time)
}

type AfterPredicateConfig struct {
	time time.Time
}

func (c *AfterPredicateConfig) ToString() string {
	return ""
}

func parseTime(datetimeStr string) time.Time {
	// 提取时区信息
	zone, err := parseZone(datetimeStr)
	if err != nil {
		log.Fatalln("无法解析时区信息:", err)
	}
	// 解析字符串中的时间和时区
	t := parseDatetime(datetimeStr)

	// 将时间转换到指定的时区
	return settingZone(t, zone)
}

func parseDatetime(dateStr string) time.Time {
	// 查找第一个方括号的索引
	startIndex := strings.Index(dateStr, "[")
	str := dateStr[0:startIndex]
	datetime, err := time.Parse("2006-01-02T15:04:05.999Z07:00", str)
	if err != nil {
		log.Fatalln("无法解析时间:", err)
	}
	return datetime
}

func parseZone(dateStr string) (*time.Location, error) {
	// 查找第一个方括号的索引
	startIndex := strings.Index(dateStr, "[")
	if startIndex == -1 {
		return nil, errors.New("未找到开始方括号")
	}

	// 查找第二个方括号的索引
	endIndex := strings.Index(dateStr, "]")
	if endIndex == -1 {
		return nil, errors.New("未找到结束方括号")
	}
	// 截取方括号之间的内容
	return time.LoadLocation(dateStr[startIndex+1 : endIndex])
}

func settingZone(datetime time.Time, zone *time.Location) time.Time {
	return datetime.In(zone)
}
