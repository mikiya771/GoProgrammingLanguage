package params

import (
	"net/url"
	"testing"
)

func TestUnpack(t *testing.T) {
	type Result struct {
		Zip string `http:"1" valdator:"zip"`
	}
	testTable := []struct {
		UrlValues url.Values
		expected  Result
		valid     bool
	}{
		{
			url.Values{
				"1": []string{"123-4567"},
				"2": []string{"123-4567"},
			},
			Result{
				"123-4567",
			},
			true,
		},
		{
			url.Values{
				"1": []string{"123-4567"},
				"2": []string{"123-4567"},
			},
			Result{
				"123-4567",
			},
			false,
		},
	}
	for _, tt := range testTable {
		var res Result
		err := unpack(tt.UrlValues, &res)
		if err != nil && tt.valid {
			t.Errorf("error:%v", err)
		}
		if err == nil && res.Zip != tt.expected.Zip {
			t.Errorf("error:%v != %v", res.Zip, tt.expected.Zip)
		}

	}
}
