package grid

import (
	"encoding/json"

	"github.com/yidane/gotest/beegoTest/models/grid/data"
	"github.com/yidane/gotest/beegoTest/models/grid/style"
)

type ReportResult struct {
	IsSuccess bool
	Data      *data.Define
	Style     *style.Grid
	Message   string
}

func (p ReportResult) Success(d data.Define, s style.Grid) ReportResult {
	p = ReportResult{
		IsSuccess: true,
		Data:      &d,
		Style:     &s,
		Message:   "",
	}

	return p
}

func (p ReportResult) Error(str string) string {
	r := &ReportResult{
		IsSuccess: false,
		Message:   str,
	}
	result, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(result)
}
