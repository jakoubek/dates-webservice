package l10n

import "strings"

type De struct {

}

func NewDeTranslation() *De {
	return &De{}
}

func (de *De) Translate(inputString string) string {
	inputString = de.longMonthNames(inputString)
	inputString = de.shortMonthNames(inputString)
	inputString = de.longDayNames(inputString)
	return inputString
}

func (de *De) longMonthNames(inputString string) string {
	inputString = strings.ReplaceAll(inputString, "January", "Januar")
	inputString = strings.ReplaceAll(inputString, "February", "Februar")
	inputString = strings.ReplaceAll(inputString, "March", "März")
	//inputString = strings.ReplaceAll(inputString, "April", "April")
	inputString = strings.ReplaceAll(inputString, "May", "Mai")
	inputString = strings.ReplaceAll(inputString, "June", "Juni")
	inputString = strings.ReplaceAll(inputString, "July", "Juli")
	//inputString = strings.ReplaceAll(inputString, "August", "August")
	//inputString = strings.ReplaceAll(inputString, "September", "September")
	inputString = strings.ReplaceAll(inputString, "October", "Oktober")
	//inputString = strings.ReplaceAll(inputString, "November", "November")
	inputString = strings.ReplaceAll(inputString, "December", "Dezember")
	return inputString
}

func (de *De) shortMonthNames(inputString string) string {
	inputString = strings.ReplaceAll(inputString, "Jan", "Jan")
	inputString = strings.ReplaceAll(inputString, "Feb", "Feb")
	inputString = strings.ReplaceAll(inputString, "Mar", "März")
	//inputString = strings.ReplaceAll(inputString, "Apr", "Apr")
	inputString = strings.ReplaceAll(inputString, "May", "Mai")
	inputString = strings.ReplaceAll(inputString, "Jun", "Jun")
	inputString = strings.ReplaceAll(inputString, "Jul", "Jul")
	//inputString = strings.ReplaceAll(inputString, "Aug", "Aug")
	//inputString = strings.ReplaceAll(inputString, "Sep", "Sep")
	inputString = strings.ReplaceAll(inputString, "Oct", "Okt")
	//inputString = strings.ReplaceAll(inputString, "Nov", "Nov")
	inputString = strings.ReplaceAll(inputString, "Dec", "Dez")
	return inputString
}

func (de *De) longDayNames(inputString string) string {
	inputString = strings.ReplaceAll(inputString, "Monday", "Montag")
	inputString = strings.ReplaceAll(inputString, "Tuesday", "Dienstag")
	inputString = strings.ReplaceAll(inputString, "Wednesday", "Mittwoch")
	inputString = strings.ReplaceAll(inputString, "Thursday", "Donnerstag")
	inputString = strings.ReplaceAll(inputString, "Friday", "Freitag")
	inputString = strings.ReplaceAll(inputString, "Saturday", "Samstag")
	inputString = strings.ReplaceAll(inputString, "Sunday", "Sonntag")
	return inputString
}
