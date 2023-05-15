package fourbytes

import (
	"context"
	"testing"
)

func TestEventSignature(t *testing.T) {
	testset := []struct {
		opts    EventSigV1Options
		success bool
	}{
		{
			opts:    EventSigV1Options{}, // Missing required
			success: false,
		},
		{
			opts: EventSigV1Options{
				HexSignature: "0x15aac4af776447c09d895192c86bab463c38b92191f3ba3f7b8831723c548d6e",
			},
			success: true,
		},
		{
			opts: EventSigV1Options{
				HexSignature: "15aac4af776447c09d895192c86bab463c38b92191f3ba3f7b8831723c548d6e",
			},
			success: true,
		},
		{
			opts: EventSigV1Options{
				HexSignature: "0x",
			},
			success: false,
		},
		{
			opts: EventSigV1Options{
				HexSignature: "0x12",
			},
			success: false,
		},
	}

	client, err := New(DefaultConfig())
	if err != nil {
		t.Fatal(err)
	}

	for _, task := range testset {
		_, statusCode, err := client.EventSignatureV1(context.Background(), task.opts)
		if err != nil || statusCode != 200 {
			if task.success {
				t.Fatal("TestEventSignature: got failed want success")
			}
			continue
		}

		if (err == nil && statusCode == 200) && !task.success {
			t.Fatal("TestEventSignature: got success want failed")
		}
	}
}
