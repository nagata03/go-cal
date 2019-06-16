package main

import (
	"fmt"
	"strings"
	"time"
)

func printCal(today time.Time, wday int, days int) {
	fmt.Printf("     %v %d\n", today.Month(), today.Year()) // 月と年を表示
	fmt.Println("Su Mo Tu We Th Fr Sa")                     // 曜日を表示
	fmt.Print(strings.Repeat("   ", int(wday)))             // 1日までのスペースを表示

	// 日付を表示
	for day := 1; day <= days; day++ {
		fmt.Printf("%2d ", day)
		wday++
		if wday == 7 {
			fmt.Println("")
			wday = 0
		}
	}
	fmt.Println("")
}

func main() {
	t := time.Now()                                                                             // 現在の日付・時刻を取得 tはtime.Time型
	wday := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local).Weekday()                 // 今月初日の曜日を取得
	days := time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, time.Local).AddDate(0, 0, -1).Day() // 今月最終日の日付（=今月の日数）を取得

	printCal(t, int(wday), days)
}
