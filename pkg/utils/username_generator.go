package utils

import "strings"

func UserIDGenerator(email, userName string) string {

	emailName := strings.Split(email, "@")

	return emailName[0] + userName
}
