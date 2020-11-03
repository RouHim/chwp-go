package display

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

//Struct that represents the current display configuration
type Information struct {
	Count               int
	Resolutions         []string
	TotalResolution     string
	MaxSingleResolution string
}

func GetDisplayInfo() Information {
	displayResolutions := getDisplayResolutions()
	maxSingleResolution := getMaxSingleDisplayResolution()
	totalResolution := getTotalResolution()

	return Information{
		Resolutions:         displayResolutions,
		Count:               len(displayResolutions),
		TotalResolution:     totalResolution,
		MaxSingleResolution: maxSingleResolution,
	}
}

func getTotalResolution() (totalResolution string) {
	if isDisplayVarSet() {
		totalResolution = executeCommand(
			"(xrandr -q|sed -n 's/.*current[ ]\\([0-9]*\\) x \\([0-9]*\\),.*/\\1x\\2/p')",
		)
	} else {
		totalResolution = executeCommand(
			"(DISPLAY=:0 xrandr -q|sed -n 's/.*current[ ]\\([0-9]*\\) x \\([0-9]*\\),.*/\\1x\\2/p')",
		)
	}

	return strings.TrimSpace(totalResolution)
}

func getDisplayResolutions() (resolutions []string) {
	resolutionsString := strings.TrimSpace(
		executeDisplayCommand("xrandr | grep \\* | cut -d' ' -f4"),
	)

	if strings.Contains(resolutionsString, "\n") {
		resolutions = strings.Split(resolutionsString, "\n")
	} else {
		resolutions = append(resolutions, resolutionsString)
	}

	return resolutions
}

func getMaxSingleDisplayResolution() string {
	resolutions := getDisplayResolutions()
	maxResolution := multiplyResolution(resolutions[0])
	resolutionString := resolutions[0]

	for _, resolution := range resolutions {
		loopResolution := multiplyResolution(resolution)

		if loopResolution > maxResolution {
			maxResolution = loopResolution
			resolutionString = resolution
		}
	}

	return resolutionString
}

func multiplyResolution(resolutionString string) int {
	decomposition := strings.Split(resolutionString, "x")
	width, err := strconv.Atoi(decomposition[0])
	height, err := strconv.Atoi(decomposition[1])

	if err != nil {
		log.Fatal(err)
	}

	return width * height
}

func executeDisplayCommand(cmd string) (resolution string) {
	if isDisplayVarSet() {
		resolution = executeCommand(cmd)
	} else {
		resolution = executeCommand("DISPLAY=:0 " + cmd)
	}

	return resolution
}

func isDisplayVarSet() bool {
	return len(strings.TrimSpace(os.Getenv("DISPLAY"))) > 0
}

func executeCommand(cmd string) string {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}
