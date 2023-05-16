package openchain

import (
	"context"
	"testing"
)

func TestStats(t *testing.T) {
	client, err := New(DefaultConfig())
	if err != nil {
		t.Fatal(err)
	}

	_, _, err = client.StatsV1(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}
