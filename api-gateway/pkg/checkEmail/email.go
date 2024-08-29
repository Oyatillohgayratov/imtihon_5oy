package checkemail

import "regexp"

func ValidateGmail(email string) bool {
	const gmailRegex = `^[a-zA-Z0-9._%+-]+@gmail\.com$`
	
	re := regexp.MustCompile(gmailRegex)
	
	return re.MatchString(email)
}