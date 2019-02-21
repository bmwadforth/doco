package Doco

import (
	"fmt"
	"log"
)

func (page *DocoPage) SetMargin(margin DocoMargin) {
	page.Margin = margin
}

func (page *DocoPage) Write(content string) uint {
	bytesWritten, err := page.Body.WriteString(content)
	if err != nil {
		page.Errors = append(page.Errors, err)
		log.Print(fmt.Errorf("%v", err))
	}
	return uint(bytesWritten)
}
