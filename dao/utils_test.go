package dao

import "testing"

// go test -v -count=1 github.com/dylenfu/gmx-spider/dao -run TestNewBig
func TestNewBig(t *testing.T) {
	data := int64(12)
	value := NewBigIntFromInt(data)
	t.Log(value)
}
