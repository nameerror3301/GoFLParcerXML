package pkg

import "testing"

func TestRequest(t *testing.T) {
	t.Run("check-Request-valid", func(t *testing.T) {
		urlSocks, urlFl := "https://hidemy.name/ru/proxy-list/?type=5#list", "https://www.fl.ru/rss/all.xml"
		body, err := Request(urlSocks, urlFl)

		if body == nil && err != nil {
			t.Error("SocksParce() - Did not pass the test with the correct data")
		}
	})

	t.Run("check-Request-invalid", func(t *testing.T) {
		urlSocks, urlFl := "", ""
		body, err := Request(urlSocks, urlFl)

		if body == nil && err == nil {
			t.Error("SocksParce() - Didn't pass the test with the wrong data")
		}
	})
}

func TestGetXmlItem(t *testing.T) {
	t.Run("check-GetXmlItem-valid", func(t *testing.T) {
		urlSocks, urlFl := "https://hidemy.name/ru/proxy-list/?type=5#list", "https://www.fl.ru/rss/all.xml"
		_, _, _, _, _, err := GetXmlItem(urlSocks, urlFl)

		if err != nil {
			t.Error("GetXmlItem() - Did not pass the test with the correct data")
		}
	})

	t.Run("check-GetXmlItem-invalid", func(t *testing.T) {
		urlSocks, urlFl := "", ""
		_, _, _, _, _, err := GetXmlItem(urlSocks, urlFl)

		if err == nil {
			t.Error("GetXmlItem() - Didn't pass the test with the wrong data")
		}
	})
}
