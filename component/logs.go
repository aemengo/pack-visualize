package component

import (
	"fmt"
	"github.com/rivo/tview"
	"time"
)

var detectText = fmt.Sprintf(`%s
[%[2]s] 6 of 6 buildpacks participating
[%[2]s] paketo-buildpacks/ca-certificates 2.1.0
[%[2]s] paketo-buildpacks/go-dist 0.3.1
[%[2]s] paketo-buildpacks/go-build 0.3.0
[%[2]s] paketo-buildpacks/procfile 4.1.0
[%[2]s] paketo-buildpacks/environment-variables 3.0.0
[%[2]s] paketo-buildpacks/image-labels 3.0.0
%[3]s
`, cyan.Sprint("===> DETECTING"), cyan.Sprint("detector"), boldWhite.Sprint("---"))

var analyzeText = fmt.Sprintf(`%s
[%[2]s] Previous image with name "index.docker.io/anthony/my-sample-golang:0.1" not found
`, cyan.Sprint("===> ANALYZING"), cyan.Sprint("analyzer"))

var restoreText = fmt.Sprintf(`%s
`, cyan.Sprint("===> RESTORING"))


var buildText1 = fmt.Sprintf(`%s
[%[2]s] Paketo Go Distribution Buildpack 0.3.1
[%[2]s]   Resolving Go version
[%[2]s]     Candidate version sources (in priority order):
[%[2]s]       <unknown> -> ""
[%[2]s] 
[%[2]s]     Selected Go version (using <unknown>): 1.15.9
`, cyan.Sprint("===> BUILDING"), cyan.Sprint("builder"))

var buildText2 = fmt.Sprintf(`[%[1]s] 
[%[1]s]   Executing build process
[%[1]s]     Installing Go 1.15.9
[%[1]s]       Completed in 24.928s
`, cyan.Sprint("builder"))

var buildText3 = fmt.Sprintf(`[%[1]s]
[%[1]s] Paketo Go Build Buildpack 0.3.0
[%[1]s]   Executing build process
[%[1]s]     Running 'go build -o /layers/paketo-buildpacks_go-build/targets/bin -buildmode pie .'
[%[1]s]       Completed in 5.635s
[%[1]s]
`, cyan.Sprint("builder"))

var failedText = boldRed.Sprint(`
BUILD FAILED`)

type Logs struct {
	tv *tview.TextView
}

func NewLogs() *Logs {
	tv := tview.NewTextView()
	tv.SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft).
		SetText(tview.TranslateANSI(detectText)).
		SetBorder(true).
		SetTitle(tview.TranslateANSI(bold.Sprint("| logs |"))).
		SetTitleAlign(tview.AlignLeft).
		SetBackgroundColor(backgroundColor)
	return &Logs{tv: tv}
}

func (l *Logs) View() *tview.TextView {
	return l.tv
}

func (l *Logs) Strobe() {
	l.tv.SetText(tview.TranslateANSI(detectText))

	time.Sleep(500 * time.Millisecond)
	l.tv.SetText(tview.TranslateANSI(detectText + analyzeText))

	time.Sleep(4 * time.Second)
	l.tv.SetText(tview.TranslateANSI(detectText + analyzeText +  restoreText))

	time.Sleep(time.Second)
	l.tv.SetText(tview.TranslateANSI(detectText + analyzeText +  restoreText + buildText1))

	time.Sleep(2 * time.Second)
	l.tv.SetText(tview.TranslateANSI(detectText + analyzeText +  restoreText + buildText1 + buildText2))

	time.Sleep(time.Second)
	l.tv.SetText(tview.TranslateANSI(detectText + analyzeText +  restoreText + buildText1 + buildText2 + buildText3))

	time.Sleep(4 * time.Second)
	l.tv.SetText(tview.TranslateANSI(detectText + analyzeText +  restoreText + buildText1 + buildText2 + buildText3 + failedText))
}