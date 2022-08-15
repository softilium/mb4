package domains

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type JSDateOnly time.Time

const cL = "2006-01-02"

func (ct *JSDateOnly) String() string {
	t := time.Time(*ct)
	return fmt.Sprintf("%q", t.Format(cL))
}

func (ct *JSDateOnly) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	nt, err := time.Parse(cL, s)
	*ct = JSDateOnly(nt)
	return
}

func (ct JSDateOnly) MarshalJSON() ([]byte, error) {
	return []byte(ct.String()), nil
}

func (ct *JSDateOnly) Scan(src any) error {
	*ct = JSDateOnly(src.(time.Time))
	return nil
}

func (ct *JSDateOnly) Value() (driver.Value, error) {
	return time.Time(*ct), nil
}
