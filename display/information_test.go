package display

import "os"
import "fmt"

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}

	got := GetDisplayInfo()

	fmt.Println(got)

	if len(got.TotalResolution) <= 0 {
		t.Error("TotalResolution is empty")
	}

	if len(got.MaxSingleResolution) <= 0 {
		t.Error("MaxSingleResolution is empty")
	}
}
