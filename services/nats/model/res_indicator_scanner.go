package model

import "github.com/InsuranceTech/shared/scanner"

type ResIndicatorScanner struct {
	ResultCode int                   `json:"result_code"`
	Results    []*scanner.ResultItem `json:"results"`
}
