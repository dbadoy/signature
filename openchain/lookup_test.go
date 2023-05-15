package openchain

import (
	"context"
	"testing"
)

func TestMethodLookup(t *testing.T) {
	testset := []struct {
		opts    LookupV1Options
		success bool
	}{
		{
			opts: LookupV1Options{
				Method: "0x00000000",
			},
			success: true, /* openchain API returns success with a null result value, even if there is no result value !! */
		},
		{
			opts:    LookupV1Options{},
			success: false,
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
			success: true,
		},
		{
			opts: LookupV1Options{
				Method: "0x",
			},
			success: false,
		},
		{
			opts: LookupV1Options{
				Method: "0x12",
			},
			success: false,
		},
	}

	client, err := New(DefaultConfig())
	if err != nil {
		t.Fatal(err)
	}

	for _, task := range testset {
		_, statusCode, err := client.LookupV1(context.Background(), task.opts)
		if err != nil || statusCode != 200 {
			if task.success {
				t.Fatal("TestMethodLookup: got failed want success")
			}
			continue
		}

		if (err == nil && statusCode == 200) && !task.success {
			t.Fatal("TestMethodLookup: got success want failed")
		}
	}
}

func TestEventLookup(t *testing.T) {
	testset := []struct {
		opts    LookupV1Options
		success bool
	}{
		{
			opts: LookupV1Options{
				Event: "0x0000000000000000000000000000000000000000000000000000000000000000",
			},
			success: true, /* openchain API returns success with a null result value, even if there is no result value !! */
		},
		{
			opts:    LookupV1Options{},
			success: false,
		},
		{
			opts: LookupV1Options{
				Event: "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
			},
			success: true,
		},
		{
			opts: LookupV1Options{
				Event: "ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
			},
			success: true,
		},
		{
			opts: LookupV1Options{
				Event: "0x",
			},
			success: false,
		},
		{
			opts: LookupV1Options{
				Event: "0x12",
			},
			success: false,
		},
	}

	client, err := New(DefaultConfig())
	if err != nil {
		t.Fatal(err)
	}

	for _, task := range testset {
		_, statusCode, err := client.LookupV1(context.Background(), task.opts)
		if err != nil || statusCode != 200 {
			if task.success {
				t.Fatal("TestEventLookup: got failed want success")
			}
			continue
		}

		if (err == nil && statusCode == 200) && !task.success {
			t.Fatal("TestEventLookup: got success want failed")
		}
	}
}
