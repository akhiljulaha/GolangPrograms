package auth



// here calling teh private from the public function 
func extractSession()string {
	return "loggedin"
}
func GetSession() string{
	// return "loggeedin"
	return extractSession()
}
