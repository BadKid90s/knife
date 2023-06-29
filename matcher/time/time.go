package time

import (
	"errors"
	"log"
	"strings"
	"time"
)

func After(datetime time.Time, patterns string) bool {
	if len(patterns) == 0 {
		return true
	}
	// 提取时区信息
	zone, err := parseZone(patterns)
	if err != nil {
		log.Fatalln("无法解析时区信息:", err)
	}
	// 解析字符串中的时间和时区
	t := parseDatetime(patterns)

	// 将时间转换到指定的时区
	actualTime := settingZone(t, zone)
	anticipateTime := settingZone(datetime, zone)

	return actualTime.After(anticipateTime)
}

func Before(datetime time.Time, patterns string) bool {
	if len(patterns) == 0 {
		return true
	}
	// 提取时区信息
	zone, err := parseZone(patterns)
	if err != nil {
		log.Fatalln("无法解析时区信息:", err)
	}
	// 解析字符串中的时间和时区
	t := parseDatetime(patterns)

	// 将时间转换到指定的时区
	actualTime := settingZone(t, zone)
	anticipateTime := settingZone(datetime, zone)

	return actualTime.Before(anticipateTime)
}

func Between(datetime time.Time, patterns string) bool {
	if len(patterns) == 0 {
		return true
	}

	pats := strings.Split(patterns, ",")

	datetime1 := parseTime(pats[0])
	datetime2 := parseTime(pats[1])

	return datetime1.After(datetime) && datetime2.Before(datetime)
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