package option

import (
	"testing"
)

type temp struct {
	A string `json:"a"`
	B string `json:"b"`
	C bool   `json:"c"`
}

func TestEncodeQueryParam(t *testing.T) {
	testset := []struct {
		given  temp
		expect string
	}{
		{
			given:  temp{"", "", false},
			expect: "",
		},
		{
			given:  temp{"", "world", false},
			expect: "b=world",
		},
		{
			given:  temp{"", "world", true},
			expect: "b=world&c=true",
		},
		{
			given:  temp{"hello", "world", false},
			expect: "a=hello&b=world",
		},
		{
			given:  temp{"hello", "", true},
			expect: "a=hello&c=true",
		},
		{
			given:  temp{"hello", "0", false},
			expect: "a=hello&b=0",
		},
	}

	for _, task := range testset {
		queryParam := EncodeQueryParam(task.given)
		if task.expect != queryParam {
			t.Fatalf("TestEncodeQueryParam: got %v want %v", queryParam, task.expect)
		}
	}
}
