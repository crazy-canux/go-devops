package tickstack

import (
	"testing"
)

func TestQuery(t *testing.T) {
	col, val, err := Query("http://10.103.239.60:8086", "sandboxav", "select * from tail order by time desc limit 10")
	if err != nil {
		t.Error("query failed.")
	} else {
		t.Log(col)
		t.Log(val)
	}
}