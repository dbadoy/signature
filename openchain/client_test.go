package openchain

import (
	"net/http"
	"testing"
	"time"
)

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

	c, err := New("", 0)
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

	c, err := New("", 0)
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
	c, err := New("", 0)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := c.Signature("0x00000000"); err == nil {
		t.Fatal("TestSignatureNotFound: want fail got succeed")
	}
}

func TestInvalidVersion(t *testing.T) {
	c := Client{
		version: "unknown/",
		caller:  &http.Client{},
	}

	if _, err := c.Signature("0xa9059cbb"); err == nil {
		t.Fatal("TestInvalidVersion: version (unknown/) must be fail")
	}
}

func TestTimeout(t *testing.T) {
	testset := []struct {
		version string
		timeout time.Duration
		success bool
	}{
		{Version, 0, true},
		{Version, time.Millisecond, false /* timeout */},
		{"unknown/", 0, true /* it should fail, but now it succeeds because the version is fixed. */},
	}

	for _, task := range testset {
		c, err := New(task.version, task.timeout)
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
