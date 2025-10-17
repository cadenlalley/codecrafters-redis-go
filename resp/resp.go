package resp

import "fmt"

const (
	BULKSTRING   = "$%d\r\n%s\r\n"
	SIMPLESTRING = "+%s\r\n"
)

func GetOKResponse() string {
	return fmt.Sprintf(SIMPLESTRING, "OK")
}
