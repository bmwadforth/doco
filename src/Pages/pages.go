package Pages

import (
	"fmt"
	"log"
)

func (page *Page) SetMargin(margin Margin) {
	page.Margin = margin
}

func (page *Page) Write(content string) uint {
	bytesWritten, err := page.Body.WriteString(content)
	if err != nil {
		page.Errors = append(page.Errors, err)
		log.Print(fmt.Errorf("%v", err))
	}
	return uint(bytesWritten)
}
