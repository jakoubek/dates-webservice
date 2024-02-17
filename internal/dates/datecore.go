package dates

import (
	"fmt"
	"github.com/golang-module/carbon/v2"
	"log"
	"strconv"
	"time"

	"github.com/jakoubek/dates-webservice/l10n"
)

const defaultDateFormat string = "2006-01-02"

type DateCoreConfig func(core *DateCore)

type DateCore struct {
	useCarbon    bool
	carbonObject carbon.Carbon
	dateObject   time.Time
	dateFormat   string
	language     string
	userFormat   string
	resultString string
}

func NewDateCore(opts ...DateCoreConfig) *DateCore {

	carbon.SetDefault(carbon.Default{
		Layout:       carbon.DateTimeLayout,
		Timezone:     carbon.Local,
		WeekStartsAt: carbon.Monday,
		Locale:       "en",
	})

	dc := DateCore{
		carbonObject: carbon.Now(),
		dateObject:   time.Now().UTC(),
		dateFormat:   defaultDateFormat,
	}
	for _, opt := range opts {
		opt(&dc)
	}
	return &dc
}

func WithUserFormat(format string) DateCoreConfig {
	return func(dc *DateCore) {
		dc.userFormat = format
	}
}

func WithLanguage(language string) DateCoreConfig {
	return func(dc *DateCore) {
		dc.language = language
		dc.carbonObject.SetLocale(language)
	}
}

func WithDayFormat() DateCoreConfig {
	return func(dc *DateCore) {
		dc.dateFormat = "2006-01-02"
	}
}

func WithMonthFormat() DateCoreConfig {
	return func(dc *DateCore) {
		dc.useCarbon = true
		dc.dateFormat = "January 2006"
	}
}

func WithYearFormat() DateCoreConfig {
	return func(dc *DateCore) {
		dc.dateFormat = "2006"
	}
}

func WithDatetimeFormat() DateCoreConfig {
	return func(dc *DateCore) {
		dc.dateFormat = time.RFC3339
	}
}

func WithDayAdd(numberOfDays int) DateCoreConfig {
	return func(dc *DateCore) {
		dc.carbonObject = dc.carbonObject.AddDays(numberOfDays)
		dc.dateObject = dc.dateObject.AddDate(0, 0, numberOfDays)
	}
}

func WithMonthAdd(numberOfMonths int) DateCoreConfig {
	return func(dc *DateCore) {
		dc.carbonObject = dc.carbonObject.AddMonths(numberOfMonths)
		dc.dateObject = dc.dateObject.AddDate(0, numberOfMonths, 0)
	}
}

func WithYearAdd(numberOfYears int) DateCoreConfig {
	return func(dc *DateCore) {
		dc.carbonObject = dc.carbonObject.AddYears(numberOfYears)
		dc.dateObject = dc.dateObject.AddDate(numberOfYears, 0, 0)
	}
}

func WithLastOfMonth() DateCoreConfig {
	return func(dc *DateCore) {
		dc.dateObject = dc.dateObject.
			AddDate(0, 0, -dc.dateObject.Day()+1).
			AddDate(0, 1, 0).
			AddDate(0, 0, -1)
	}
}

func WithWeeknumber() DateCoreConfig {
	return func(dc *DateCore) {
		_, wknr := dc.dateObject.ISOWeek()
		dc.resultString = strconv.Itoa(wknr)
	}
}

func WithTimestamp(asMilliseconds bool) DateCoreConfig {
	return func(dc *DateCore) {
		var timestamp int64
		if asMilliseconds {
			timestamp = time.Now().UnixMilli()
		} else {
			timestamp = time.Now().Unix()
		}
		dc.resultString = strconv.FormatInt(timestamp, 10)
	}
}

func (dc *DateCore) ResultString() string {
	if dc.resultString == "" {
		if dc.useCarbon {
			log.Println("Format: ", dc.getFormat())
			dc.resultString = fmt.Sprintf("%s %d", dc.carbonObject.ToMonthString(), dc.carbonObject.Year())
		} else {
			dc.resultString = dc.dateObject.Format(dc.getFormat())
		}
	}

	return dc.resultString
}

func (dc *DateCore) getFormat() string {
	if dc.userFormat != "" {
		return dc.userFormat
	} else {
		return dc.dateFormat
	}
}

func (dc *DateCore) getLocalization(inputString string) string {
	switch dc.language {
	case "de":
		inputString = (l10n.NewDeTranslation()).Translate(inputString)
	}
	return inputString
}
