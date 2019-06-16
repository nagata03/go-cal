package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	t := time.Now() // 現在の時刻を取得 tはtime.Time型

	// 今月初日の曜日を取得する
	firstDay := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
	wday := firstDay.Weekday()

	// 今月の最終日を取得する
	lastDay := time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, time.Local).AddDate(0, 0, -1)

	//
	// カレンダーを表示する
	//
	fmt.Printf("     %v %d\n", t.Month(), t.Year()) // 月と年を表示
	fmt.Println("Su Mo Tu We Th Fr Sa")             // 曜日を表示

	// 1日までのスペースを表示
	fmt.Print(strings.Repeat("   ", int(wday)))

	// 日付を表示
	for day := 1; day <= lastDay.Day(); day++ {
		fmt.Printf("%2d ", day)
		wday++
		if wday == 7 {
			fmt.Println("")
			wday = 0
		}
	}
}
