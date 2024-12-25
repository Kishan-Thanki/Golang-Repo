package auth

func extractSession() string {
	return "loggedIN"
}

func GetSession() string {
	return extractSession()
}
