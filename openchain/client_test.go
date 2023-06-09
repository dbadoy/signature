package openchain

import (
	"testing"
	"time"
)

func TestDefaultConfig(t *testing.T) {
	c, err := New(nil)
	if err != nil {
		t.Fatal(err)
	}

	if c.cfg.Timeout != 0 {
		t.Fatalf("TestDefaultConfig: invalid DefaultConfig.Timeout want %d got %d", 0, c.cfg.Timeout)
	}
}

func TestSignature(t *testing.T) {
	testset := []struct {
		id      string
		success bool
	}{
		{
			id:      "0xa9059cbb",
			success: true,
		},
		{
			id:      "a9059cbb",
			success: true,
		},
		{
			id:      "",
			success: false,
		},
		{
			id:      "0x",
			success: false,
		},
		{
			id:      "hello",
			success: false,
		},
	}

	c, err := New(DefaultConfig())
	if err != nil {
		t.Fatal(err)
	}

	for _, task := range testset {
		_, err := c.Signature(task.id)
		if err != nil {
			if task.success {
				t.Fatal("TestSignature: got failed want success")
			}
			continue
		}

		if err == nil && !task.success {
			t.Fatal("TestSignature: got success want failed")
		}
	}
}

func TestSignatureWithBytes(t *testing.T) {
	testset := []struct {
		id      []byte
		success bool
	}{
		{
			id:      []byte{169, 5, 156, 187}, // a9059cbb
			success: true,
		},
		{
			id:      nil,
			success: false,
		},
		{
			id:      []byte{1, 2, 3, 4, 5, 6},
			success: false,
		},
	}

	c, err := New(DefaultConfig())
	if err != nil {
		t.Fatal(err)
	}

	for _, task := range testset {
		_, err := c.SignatureWithBytes(task.id)
		if err != nil {
			if task.success {
				t.Fatal("TestSignatureWithBytes: got failed want success")
			}
			continue
		}

		if err == nil && !task.success {
			t.Fatal("TestSignatureWithBytes: got success want failed")
		}
	}
}

func TestSignatureNotFound(t *testing.T) {
	c, err := New(DefaultConfig())
	if err != nil {
		t.Fatal(err)
	}

	if _, err := c.Signature("0x00000000"); err == nil {
		t.Fatal("TestSignatureNotFound: want fail got succeed")
	}
}

func TestTimeout(t *testing.T) {
	testset := []struct {
		timeout time.Duration
		success bool
	}{
		{0, true},
		{time.Millisecond, false /* timeout */},
		{0, true /* it should fail, but now it succeeds because the version is fixed. */},
	}

	for _, task := range testset {
		c, err := New(&Config{Timeout: task.timeout})
		if err != nil {
			t.Fatal(err)
		}

		_, err = c.Signature("0xa9059cbb")
		if err != nil {
			if task.success {
				t.Fatal("TestTimeout: got failed want success")
			}
			continue
		}

		if err == nil && !task.success {
			t.Fatal("TestTimeout: got success want failed")
		}
	}
}
