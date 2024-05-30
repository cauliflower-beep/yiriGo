package main

import "time"

/*
	时间处理很复杂, 会包含很多主观意义上的错误假设
	因此在处理时间时, 使用time包会有助于以更安全、更准确的方式处理这些不正确的假设
*/

// isActiveBad
//  @Description: 表达瞬时时间-Bad
//  @param now
//  @param start
//  @param stop
//  @return bool
func isActiveBad(now, start, stop int) bool {
	return start <= now && now < stop
}

// isActiveGood
//  @Description: 使用time.time表达瞬时时间-Good
//  @param now
//  @param start
//  @param stop
//  @return bool
func isActiveGood(now, start, stop time.Time) bool {
	return (start.Before(now) || start.Equal(now)) && now.Before(stop)
}

// pollBad
//  @Description:表达时间段-Bad
//  @param delay 10是毫秒还是秒? 不清楚
func pollBad(delay int) {
	for {
		// ...
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}

// pollGood
//  @Description: 表达时间段-Good
//  @param delay 10*time.Second
func pollGood(delay time.Duration) {
	for {
		// ...
		time.Sleep(delay)
	}
}

/*
	若要在某个时间瞬间加上24小时,使用的添加时间的方法,取决于添加时间的意图
	1.想要下一个日历日的同一个时间点：time.AddDate()
	2.想保证某一时刻比前一时刻晚24小时：time.Add
*/
