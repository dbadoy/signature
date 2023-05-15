package file

import (
	"testing"
)

func TestSignature(t *testing.T) {
	tests := []struct {
		id      string
		want    string
		succeed bool
	}{
		{
			"00000009",
			"getInitializationCodeFromContractRuntime_6CLUNS()",
			true,
		},
		{
			"000000fa",
			"getDexAggKeeperWhitelistPosition_IkFc(address)",
			true,
		},
		{
			"0xa9059cbb",
			"transfer(address,uint256)",
			true,
		},
		{
			"00000000", // Not exist methodID.
			"",
			false,
		},
		{
			"134cc3f9",
			"getTokenLockersForAccount(address)",
			false,
		},
	}

	c, err := New(0)
	if err != nil {
		t.Fatal(err)
	}

	for _, test := range tests {
		got, err := c.Signature(test.id)
		if err != nil {
			if test.succeed {
				t.Fatal(err)
			}
			continue
		}

		if test.want != got[0] {
			t.Fatalf("TestSignature: want %v got %v", test.want, got)
		}
	}
}

func TestSplitBySeperator(t *testing.T) {
	if len(splitBySeperator("getTokenLockersForAccount(address)")) != 1 {
		t.Fatalf("TestSplitBySeperator: single item case failed")
	}
	if len(splitBySeperator("getTokenLockersForAccount(address);activateCollection(uint256)")) != 2 {
		t.Fatalf("TestSplitBySeperator: couldn't handle seperator")
	}
	c, err := New(0)
	if err != nil {
		t.Fatal(err)
	}

	// getTokenLockersForAccount(address);activateCollection(uint256)
	sigs, err := c.Signature("134cc3f9")
	if err != nil || len(sigs) != 2 {
		t.Fatal("TestSplitBySeperator")
	}
}
