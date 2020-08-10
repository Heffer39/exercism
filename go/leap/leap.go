// The leap package implements functionality surrounding leap years
package leap

// The IsLeapYear function reports if the passed in year is a leap year
func IsLeapYear(year int) bool {
	if year % 4 == 0 && (year % 100 != 0 || year % 400 == 0) {
		return true
	} else {
		return false
	}
}
