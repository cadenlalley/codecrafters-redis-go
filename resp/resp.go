package resp

import "fmt"

const (
	BULKSTRINGFORMATTING   = "$%d\r\n%s\r\n"
	SIMPLESTRINGFORMATTING = "+%s\r\n"
	SIMPLEERRORFORMATTING  = "-ERR %s\r\n"
)

const (
	NULLBULKSTRING = "$-1\r\n"
)

func NewSimpleString(input any) string {
	return fmt.Sprintf(SIMPLESTRINGFORMATTING, input)
}

func NewBulkString(input string) string {
	return fmt.Sprintf(BULKSTRINGFORMATTING, len(input), input)
}

func NewOKResponse() string {
	return fmt.Sprintf(SIMPLESTRINGFORMATTING, "OK")
}

func NewErrorResponse(message string) string {
	return fmt.Sprintf(SIMPLEERRORFORMATTING, message)
}

func GetNullBulkString() string {
	return NULLBULKSTRING
}
