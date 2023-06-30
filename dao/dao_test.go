package dao

import (
	"testing"

	"github.com/dylenfu/gmx-spider/config"
)

// go test -v -count=1 github.com/dylenfu/gmx-spider/dao -run TestNewDao
func TestNewDao(t *testing.T) {
	cfg := &config.DBConfig{
		URL:      "localhost:3306",
		User:     "root",
		Password: "111111",
		Scheme:   "gmx",
		Debug:    true,
	}
	dao := NewGmxDao(cfg)
	t.Log(dao.cfg)
}
