package util

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func ParseTime(datetimeStr string) (*time.Time, error) {
	// 提取时区信息
	zone, err := parseZone(datetimeStr)
	if err != nil {
		return nil, fmt.Errorf("unable to parse time zone information: %s", err)
	}
	// 解析字符串中的时间和时区
	t, err := parseDatetime(datetimeStr)
	if err != nil {
		return nil, err
	}

	// 将时间转换到指定的时区
	zoneTime := settingZone(t, zone)
	return &zoneTime, nil
}

func parseDatetime(dateStr string) (*time.Time, error) {
	// 查找第一个方括号的索引
	startIndex := strings.Index(dateStr, "[")
	str := dateStr[0:startIndex]
	datetime, err := time.Parse("2006-01-02T15:04:05.999Z07:00", str)
	if err != nil {
		return nil, fmt.Errorf("time could not be parsed, error: %s", err)
	}
	return &datetime, nil
}

func parseZone(dateStr string) (*time.Location, error) {
	// 查找第一个方括号的索引
	startIndex := strings.Index(dateStr, "[")
	if startIndex == -1 {
		return nil, errors.New("opening bracket was not found")
	}

	// 查找第二个方括号的索引
	endIndex := strings.Index(dateStr, "]")
	if endIndex == -1 {
		return nil, errors.New("closing bracket was not found")
	}
	// 截取方括号之间的内容
	return time.LoadLocation(dateStr[startIndex+1 : endIndex])
}

func settingZone(datetime *time.Time, zone *time.Location) time.Time {
	return datetime.In(zone)
}
