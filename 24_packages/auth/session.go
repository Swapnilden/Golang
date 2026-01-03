package auth

func extractSession() string { // private to package
	return "loggedin"
}

func GetSession() string {
	return extractSession()
}
