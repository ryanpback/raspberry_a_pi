package models

import (
	"database/sql"
)

// DBConn this is here so that we can hydrate all the web handlers with the same
// DB connection that we are using in the main package
var DBConn *sql.DB

// Payload is how all requests are wrapped
type Payload map[string]interface{}
