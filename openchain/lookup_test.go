package openchain

import (
	"context"
	"testing"
)

func TestGetEventSignature(t *testing.T) {
	testset := []struct {
		opts    LookupV1Options
		success bool
	}{
		{
			/*
				Openchain returns success with a null result value, even if there is no result value !!
			*/
			opts:    LookupV1Options{},
			success: true,
		},
		{
			opts: LookupV1Options{
				Method: "0xa9059cbb",
			},
			success: true,
		},
		{
			opts: LookupV1Options{
				Method: "a9059cbb",
			},
			success: false, // Not accept without '0x'
		},
		{
			opts: LookupV1Options{
				Method: "0x",
			},
			success: true,
		},
		{
			opts: LookupV1Options{
				Method: "0x12",
			},
			success: true,
		},
	}

	client, err := New("", 0)
	if err != nil {
		t.Fatal(err)
	}

	for _, task := range testset {
		_, statusCode, err := client.LookupV1(context.Background(), task.opts)
		if err != nil || statusCode != 200 {
			if task.success {
				t.Fatal("TestGetSignature failure: got failed want success")
			}
			continue
		}

		if (err == nil && statusCode == 200) && !task.success {
			t.Fatal("TestGetSignature failure: got success want failed")
		}
	}
}
