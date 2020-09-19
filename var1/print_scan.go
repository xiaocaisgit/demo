package main

import (
	"fmt"
)

func main() {
	/*jidu := 1
	switch jidu {
	case 1 :
		fmt.Println("第一季度")
	case 2:
		fmt.Println("第二季度")
	case 3:
		fmt.Println("第三季度")
	case 4:
		fmt.Println("第四季度")
	default:
		fmt.Println("重新输入")
	}
	fmt.Println("main...over...")*/

	year := 2020
	month := 2
	day := 0
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		day = 31
	case 4, 6, 9, 11:
		day = 30
	case 2:
		if year%4 == 0 {
			day = 28
		} else {
			day = 29
		}
	default:
		fmt.Println("输入有误")

	}
	fmt.Printf("%d年%d有%d天", year, month, day)

}

/*
{{ define "__subject" }}
[
{{ .Status | toUpper }}
{{ if eq .Status "firing" }}:{{ .Alerts.Firing | len }}{{ end }}
]
{{ .GroupLabels.SortedPairs.Values | join " " }}

{{ if gt (len .CommonLabels) (len .GroupLabels) }}({{ with .CommonLabels.Remove .GroupLabels.Names }}{{ .Values | join " " }}{{ end }}){{ end }}
{{ end }}
*/
