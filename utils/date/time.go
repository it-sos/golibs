/*
   Copyright (c) [2021] IT.SOS
   kn is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package date

import (
	"strconv"
	"time"
)

// NowYmd 获取当前年月日，时分秒为 0
func NowYmd() time.Time {
	now := time.Now()
	year, month, day := now.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, now.Location())
}

// TimeDuration 距今持续时间
func TimeDuration(timeAnchor time.Time) string {
	t := time.Now().Sub(timeAnchor)
	s := ""
	day := int(t.Hours() / 24)
	if t.Seconds() < 60 {
		s = strconv.Itoa(int(t.Seconds())) + "秒"
	} else if t.Minutes() < 60 {
		s = strconv.Itoa(int(t.Minutes())) + "分钟"
	} else if t.Hours() < 24 {
		s = strconv.Itoa(int(t.Hours())) + "小时"
	} else if day < 7 {
		s = strconv.Itoa(day) + "天"
	} else if day < 30 {
		s = strconv.Itoa(day/7) + "周"
	} else if day < 365 {
		s = strconv.Itoa(day/30) + "月"
	} else {
		s = strconv.Itoa(day/365) + "年"
	}
	return s + "前"
}
