package testhelpers

// TruncateUsers will truncate the users table
func TruncateUsers() {
	_, err := testConfig.DBConn.Exec("TRUNCATE TABLE users RESTART IDENTITY")

	if err != nil {
		panic(err.Error())
	}
}
