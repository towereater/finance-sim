module mainframe/dossier

go 1.24.5

replace mainframe-lib/common => ../../mainframe-lib/common

replace mainframe-lib/dossier => ../../mainframe-lib/dossier

replace mainframe-lib/user => ../../mainframe-lib/user

replace mainframe-lib/account => ../../mainframe-lib/account

replace mainframe-lib/checking-account => ../../mainframe-lib/checking-account

replace mainframe-lib/security => ../../mainframe-lib/security

replace mainframe-lib/xchanger => ../../mainframe-lib/xchanger

require (
	go.mongodb.org/mongo-driver v1.17.4
	mainframe-lib/account v0.0.0-00010101000000-000000000000
	mainframe-lib/checking-account v0.0.0-00010101000000-000000000000
	mainframe-lib/common v0.0.0-00010101000000-000000000000
	mainframe-lib/dossier v0.0.0-00010101000000-000000000000
	mainframe-lib/security v0.0.0-00010101000000-000000000000
	mainframe-lib/user v0.0.0-00010101000000-000000000000
	mainframe-lib/xchanger v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang/snappy v1.0.0 // indirect
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/montanaflynn/stats v0.7.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78 // indirect
	golang.org/x/crypto v0.41.0 // indirect
	golang.org/x/sync v0.16.0 // indirect
	golang.org/x/text v0.28.0 // indirect
)
