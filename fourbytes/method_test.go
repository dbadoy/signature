package fourbytes

import (
	"context"
	"testing"
)

func TestMethodSignature(t *testing.T) {
	testset := []struct {
		opts    MethodSigV1Options
		success bool
	}{
		{
			opts:    MethodSigV1Options{}, // Missing required
			success: false,
		},
		{
			opts: MethodSigV1Options{
				HexSignature: "0xa9059cbb",
			},
			success: true,
		},
		{
			opts: MethodSigV1Options{
				HexSignature: "a9059cbb",
			},
			success: true,
		},
		{
			opts: MethodSigV1Options{
				HexSignature: "0x",
			},
			success: false,
		},
		{
			opts: MethodSigV1Options{
				HexSignature: "0x12",
			},
			success: false,
		},
	}

	client, err := New("", 0)
	if err != nil {
		t.Fatal(err)
	}

	for _, task := range testset {
		_, statusCode, err := client.MethodSignatureV1(context.Background(), task.opts)
		if err != nil || statusCode != 200 {
			if task.success {
				t.Fatal("TestMethodSignature: got failed want success")
			}
			continue
		}

		if (err == nil && statusCode == 200) && !task.success {
			t.Fatal("TestMethodSignature: got success want failed")
		}
	}
}
