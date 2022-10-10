package pkg

import (
	"testing"
)

func TestSocksParce(t *testing.T) {
	t.Run("check-SocksParce-valid", func(t *testing.T) {
		url := "https://hidemy.name/ru/proxy-list/?type=5#list"
		socks := SocksParce(url)
		if socks == "" {
			t.Error("SocksParce() - Did not pass the test with the correct data")
		}
	})

	t.Run("check-SocksParce-invalid", func(t *testing.T) {
		url := "htt1ps://hidem543y.name/ru/proxy-list/?type=5#list"
		socks := SocksParce(url)
		if socks != "" {
			t.Error("SocksParce() - Didn't pass the test with the wrong data")
		}
	})
}
