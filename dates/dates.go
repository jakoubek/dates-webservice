package dates

import (
	"github.com/jakoubek/dates-webservice/l10n"
	"strconv"
	"time"
)

type DateCoreConfig func(core *DateCore)

type DateCore struct {
	dateObject time.Time
	language string
	userFormat string
	ResultString string
}

func WithLanguage(language string) DateCoreConfig {
	if language != "de" {
		language = ""
	}
	return func(dc *DateCore) {
		dc.language = language
	}
}

func WithFormat(formatString string) DateCoreConfig {
	return func(dc *DateCore) {
		dc.userFormat = formatString
	}
}

func NewDateCore(opts ...DateCoreConfig) *DateCore {
	dc := DateCore{
		dateObject: time.Now(),
	}
	for _, opt := range opts {
		opt(&dc)
	}
	return &dc
}

func (dc *DateCore) LastOfMonth() string {
	dc.dateObject = dc.dateObject.AddDate(0, 0, -dc.dateObject.Day()).AddDate(0, 1, 0)
	dc.ResultString = dc.getLocalization(dc.dateObject.Format(dc.getFormat("2006-01-02")))
	return dc.ResultString
}

func (dc *DateCore) NextYear() string {
	dc.dateObject = dc.dateObject.AddDate(1, 0, 0)
	dc.ResultString = dc.dateObject.Format("2006")
	return dc.ResultString
}

func (dc *DateCore) LastYear() string {
	dc.dateObject = dc.dateObject.AddDate(-1, 0, 0)
	dc.ResultString = dc.dateObject.Format("2006")
	return dc.ResultString
}

func (dc *DateCore) ThisYear() string {
	dc.ResultString = dc.dateObject.Format("2006")
	return dc.ResultString
}

func (dc *DateCore) NextMonth() string {
	dc.dateObject = dc.dateObject.AddDate(0, 1, 0)
	dc.ResultString = dc.getLocalization(dc.dateObject.Format(dc.getFormat("January 2006")))
	return dc.ResultString
}

func (dc *DateCore) LastMonth() string {
	dc.dateObject = dc.dateObject.AddDate(0, -1, 0)
	dc.ResultString = dc.getLocalization(dc.dateObject.Format(dc.getFormat("January 2006")))
	return dc.ResultString
}

func (dc *DateCore) ThisMonth() string {
	dc.ResultString = dc.getLocalization(dc.dateObject.Format(dc.getFormat("January 2006")))
	return dc.ResultString
}

func (dc *DateCore) Today() string {
	dc.ResultString = dc.getLocalization(dc.dateObject.Format(dc.getFormat("2006-01-02")))
	return dc.ResultString
}

func (dc *DateCore) Tomorrow() string {
	dc.dateObject = dc.dateObject.AddDate(0, 0, 1)
	dc.ResultString = dc.getLocalization(dc.dateObject.Format(dc.getFormat("2006-01-02")))
	return dc.ResultString
}

func (dc *DateCore) Yesterday() string {
	dc.dateObject = dc.dateObject.AddDate(0, 0, -1)
	dc.ResultString = dc.getLocalization(dc.dateObject.Format(dc.getFormat("2006-01-02")))
	return dc.ResultString
}

func (dc *DateCore) Weeknumber() string {
	_, wknr := dc.dateObject.ISOWeek()
	dc.ResultString = strconv.Itoa(wknr)
	return dc.ResultString
}

func (dc *DateCore) getFormat(formatString string) string {
	if dc.userFormat != "" {
		return dc.userFormat
	} else {
		return formatString
	}
}

func (dc *DateCore) getLocalization(inputString string) string {
	switch dc.language {
	case "de":
		inputString = (l10n.NewDeTranslation()).Translate(inputString)
	}
	return inputString
}