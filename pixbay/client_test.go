package pixbay

import (
	"fmt"
	"github.com/RouHim/chwp/cli"
	"github.com/RouHim/chwp/display"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := GetImageUrl(cli.Configuration{
		Keywords: []string{"autumn"},
		Span:     true,
	},
		display.Information{
			Count:               3,
			Resolutions:         []string{"1920x1080", "1920x1080", "2560x1440"},
			TotalResolution:     "6400x1440",
			MaxSingleResolution: "2560x1440",
		},
	)

	fmt.Println(got)

	if len(got) <= 0 {
		t.Errorf("expected the result not to be empty")
	}
}
