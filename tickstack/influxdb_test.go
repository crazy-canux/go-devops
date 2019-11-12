package tickstack

import (
	"testing"
)

func TestQuery(t *testing.T) {
	col, val, err := Query("http://127.0.0.1:8086", "sandboxav", "select * from tail order by time desc limit 10")
	if err != nil {
		t.Error("query failed.")
	} else {
		t.Log(col)
		t.Log(val)
	}
}
