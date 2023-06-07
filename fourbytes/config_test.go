package fourbytes

import "testing"

func TestConfigValidate(t *testing.T) {
	tests := []struct {
		cfg     *Config
		success bool
	}{
		{DefaultConfig(), true},
	}

	for _, task := range tests {
		err := task.cfg.validate()
		if err != nil {
			if task.success {
				t.Fatal("TestConfigValidate: got failed want success")
			}
			continue
		}

		if err == nil && !task.success {
			t.Fatal("TestConfigValidate: got success want failed")
		}
	}
}
