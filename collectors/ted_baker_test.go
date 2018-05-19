package collectors

import "testing"

func TestTedBakerCollector(t *testing.T) {
	tedBakerCollector := NewTedBakerCollector()

	urls := tedBakerCollector.Process()
	if len(urls) == 0 {
		t.Error("Collector did not return any URLs")
	}
}