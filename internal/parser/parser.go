// Package parser parser.go provides funcs to set values in url.Values
package parser

import (
	"bytes"
	"net/url"
	"unsafe"
)

// AddTitle add values from 'FindTitle' field
// to query string
func AddTitle(query *string, find []byte) {
	var buf bytes.Buffer
	buf.Write([]byte("NAME:("))
	rangeByByte(find, byte(','), func(start, end int) {
		if start == end {
			return
		}
		src := find[start:end]
		buf.Write(src)
		buf.Write([]byte(" OR "))
	})
	findBytes := buf.Bytes()

	lastOrIdx := bytes.LastIndex(findBytes, []byte(" OR "))
	if lastOrIdx != -1 {
		findBytes = findBytes[:lastOrIdx]
	}
	findBytes = append(findBytes, ')')

	findStr := unsafe.String(unsafe.SliceData(findBytes), len(findBytes))

	if findStr == "NAME:()" {
		*query = ""
		return
	}

	*query = findStr
}

// AddSalary add values from 'Salary' field
// to query string
func AddSalary(query *string, salary []byte) {
	trimSpaceBytes(&salary)
	salaryStr := unsafe.String(unsafe.SliceData(salary), len(salary))
	*query = salaryStr
}

// AddExp add values from 'ExpRange' field
// to query url.Values
func AddExp(query *url.Values, exp []byte) {
	rangeByByte(exp, byte(','), func(start, end int) {
		if start == end {
			return
		}
		src := exp[start:end]
		srcStr := unsafe.String(unsafe.SliceData(src), len(src))
		query.Add("experience", srcStr)
	})
}

// AddSchedule add values from 'Schedule' field
func AddSchedule(query *string, schedule []byte) {
	trimSpaceBytes(&schedule)
	scheduleStr := unsafe.String(unsafe.SliceData(schedule), len(schedule))
	*query = scheduleStr
}
