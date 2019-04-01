package clicolorer

//go:generate counterfeiter -o ./fake.go --fake-name Fake ./ CliColorer

import (
	"github.com/gookit/color"
)

type CliColorer interface {
	// silently disables coloring
	Disable()

	// attention colors
	Attention(
		format string,
		values ...interface{},
	) string

	// errors colors
	Error(
		format string,
		values ...interface{},
	) string

	// info colors
	Info(
		format string,
		values ...interface{},
	) string

	// success colors
	Success(
		format string,
		values ...interface{},
	) string
}

func New() CliColorer {
	attentionCliColorer := color.Style{color.FgYellow, color.OpBold}
	errorCliColorer := color.Style{color.FgRed, color.OpBold}
	infoCliColorer := color.Style{color.FgCyan, color.OpBold}
	successCliColorer := color.Style{color.FgGreen, color.OpBold}

	return &cliColorer{
		attentionCliColorer: &attentionCliColorer,
		errorCliColorer:     &errorCliColorer,
		infoCliColorer:      &infoCliColorer,
		successCliColorer:   &successCliColorer,
	}
}

type cliColorer struct {
	attentionCliColorer *color.Style
	errorCliColorer     *color.Style
	infoCliColorer      *color.Style
	successCliColorer   *color.Style
}

func (this *cliColorer) Disable() {
	colorDefault := color.Style{color.FgDefault}
	this.attentionCliColorer = &colorDefault
	this.errorCliColorer = &colorDefault
	this.infoCliColorer = &colorDefault
	this.successCliColorer = &colorDefault
}

func (this cliColorer) Attention(
	format string,
	values ...interface{},
) string {
	return this.attentionCliColorer.Sprintf(format, values...)
}

func (this cliColorer) Error(
	format string,
	values ...interface{},
) string {
	return this.errorCliColorer.Sprintf(format, values...)
}

func (this cliColorer) Info(
	format string,
	values ...interface{},
) string {
	return this.infoCliColorer.Sprintf(format, values...)
}

func (this cliColorer) Success(
	format string,
	values ...interface{},
) string {
	return this.successCliColorer.Sprintf(format, values...)
}
