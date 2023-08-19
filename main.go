package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/pedroalbanese/readline"
)

const pageWidth = 80

var (
	t     = time.Now()
	year  = flag.Int("year", t.Year(), "year")
	months = [12]string{
		"Jan", "Feb", "Mar", "Apr", "May", "Jun",
		"Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
	}
)

func main() {
	flag.Parse()

	if flag.Arg(0) == "." {
		t := time.Now()
		year, month, day := t.Date()

		calendars := []string{
			getCalendar(year, int(month)-1),
			getCalendar(year, int(month)),
			getCalendar(year, int(month)+1),
		}

		lines := make([]string, 9)
		for i := range lines {
			if i == 0 {
				lines[i] = strings.Repeat(" ", 32) + fmt.Sprintf("%d", year)
				continue
			}
			lines[i] = strings.Join([]string{
				" " + getMonthRows(calendars[0], i-1),
				" " + getMonthRows(calendars[1], i-1),
				" " + getMonthRows(calendars[2], i-1),
			}, " ")
		}

		rl, err := readline.New("")
		if err != nil {
			fmt.Println("Error initializing readline:", err)
			return
		}
		defer rl.Close()

		dayStr := fmt.Sprintf("%2d", day)
		dayUnderlined := underline(dayStr)

		lines[2] = strings.Replace(lines[2], dayStr, dayUnderlined, 1)

		for _, line := range lines {
			fmt.Fprintln(rl.Stdout(), line)
		}
	} else if flag.Arg(0) == "+" {
		printFullCal(*year)
	} else {
		year, month, day := time.Now().Date()
		printCalWithUnderline(year, int(month), day)
	}
}

func underline(text string) string {
	return "\033[4m" + text + "\033[0m"
}

func getCalendar(year, month int) string {
	t := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	monthName := t.Month().String()
	cal := fmt.Sprintf("%11s          \nSu Mo Tu We Th Fr Sa \n", monthName[:3])

	lastDay := getDaysInMonth(year, month)
	startWeekday := int(t.Weekday())

	weekday := startWeekday
	week := 0

	cal += strings.Repeat("   ", weekday)

	for day := 1; day <= lastDay; day++ {
		dayStr := fmt.Sprintf("%2d", day)

		if year == time.Now().Year() && month == int(time.Now().Month()) && day == time.Now().Day() {
			// Add the underlined day
			dayStr = "\033[4m" + dayStr + "\033[0m"
		}

		cal += dayStr + " "
		weekday = (weekday + 1) % 7

		if weekday == 0 && day < lastDay {
			cal += "\n"
			week++
		}
	}

	if weekday != 0 {
		cal += strings.Repeat("   ", 7-weekday)
		cal += "\n"
	}

	return cal
}

func getDaysInMonth(year, month int) int {
	return time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day()
}

func getMonthRows(calendar string, row int) string {
	rows := strings.Split(calendar, "\n")
	if len(rows) > row {
		rowStr := rows[row]
		if len(rowStr) < 21 {
			rowStr += strings.Repeat(" ", 21-len(rowStr))
		}
		return rowStr
	}
	return strings.Repeat(" ", 21)
}

func printFullCal(year int) {
	thisDate := time.Date(year, 1, 1, 1, 1, 1, 1, time.UTC)
	var (
		dayArr                  [12][7][6]int // month, weekday, week
		month, lastMonth        time.Month
		weekInMonth, dayInMonth int
	)
	for thisDate.Year() == year {
		if month = thisDate.Month(); month != lastMonth {
			weekInMonth = 0
			dayInMonth = 1
		}
		weekday := thisDate.Weekday()
		if weekday == 0 && dayInMonth > 1 {
			weekInMonth++
		}
		dayArr[int(month)-1][weekday][weekInMonth] = thisDate.Day()
		lastMonth = month
		dayInMonth++
		thisDate = thisDate.Add(time.Hour * 24)
	}
	fmt.Println(strings.Repeat(" ", 32) + fmt.Sprintf("%d", year))

	for qtr := 0; qtr < 4; qtr++ {
		for monthInQtr := 0; monthInQtr < 3; monthInQtr++ { // Month names
			fmt.Printf("         %s           ", months[qtr*3+monthInQtr])
		}
		fmt.Println()
		for monthInQtr := 0; monthInQtr < 3; monthInQtr++ { // Day names
			days := [7]string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
			for _, day := range days {
				fmt.Printf(" %s", day[:2])
			}
			fmt.Printf("  ")
		}
		fmt.Println()
		for weekInMonth = 0; weekInMonth < 6; weekInMonth++ {
			for monthInQtr := 0; monthInQtr < 3; monthInQtr++ {
				for day := 0; day < 7; day++ {
					if dayArr[qtr*3+monthInQtr][day][weekInMonth] == 0 {
						fmt.Printf("   ")
					} else {
						fmt.Printf("%3d", dayArr[qtr*3+monthInQtr][day][weekInMonth])
					}
				}
				fmt.Printf("  ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func printCalWithUnderline(year, month, currentDay int) {
	monthName := time.Month(month).String()[:3]
	fmt.Printf("       %s %d\n", monthName, year)
	fmt.Println(" Su Mo Tu We Th Fr Sa")

	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	startWeekday := int(startDate.Weekday())
	daysInMonth := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day()

	day := 1
	lines := []string{}
	line := ""
	for week := 0; day <= daysInMonth; week++ {
		line = ""
		for weekday := 0; weekday < 7; weekday++ {
			if week == 0 && weekday < startWeekday {
				line += "   "
			} else if day == currentDay {
				line += " " + underline(fmt.Sprintf("%2d", day))
				day++
			} else if day <= daysInMonth {
				line += fmt.Sprintf(" %2d", day)
				day++
			}
		}
		lines = append(lines, line)
	}

	rl, err := readline.New("")
	if err != nil {
		fmt.Println("Error initializing readline:", err)
		return
	}
	defer rl.Close()

	for _, line := range lines {
		fmt.Fprintln(rl.Stdout(), line)
	}

	fmt.Println()
}
