package CLI

import (
	"github.com/schollz/progressbar/v3"
	"os"
)

func NewProgressBar(size int) *progressbar.ProgressBar {
	return progressbar.NewOptions(
		size,
		progressbar.OptionSetWriter(os.Stdout),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetWidth(100),
		progressbar.OptionSetDescription("[cyan][SCRAPPING][reset]"),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]█[reset]",
			SaucerHead:    "[green]█[reset]",
			SaucerPadding: " ",
			BarStart:      "|",
			BarEnd:        "|",
		}))
}
