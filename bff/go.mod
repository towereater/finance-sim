module bff

go 1.24.5

replace mainframe-lib/user => ../mainframe-lib/user

replace mainframe-lib/security => ../mainframe-lib/security

replace mainframe-lib/common => ../mainframe-lib/common

// replace mainframe/account => ../mainframe/account

require (
	github.com/golang-jwt/jwt/v5 v5.3.0
	mainframe-lib/common v0.0.0-00010101000000-000000000000
	mainframe-lib/security v0.0.0-00010101000000-000000000000
	mainframe-lib/user v0.0.0-00010101000000-000000000000
// mainframe/account v0.0.0-00010101000000-000000000000
)

require github.com/kelseyhightower/envconfig v1.4.0 // indirect
