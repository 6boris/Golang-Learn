package v3

import "testing"

func TestStartReceive(t *testing.T) {
	StartReceive("test", "test-channel", "nsq.kyle.link:4152")
}
