module bff

go 1.22.2

replace mainframe/user => ../mainframe/user

replace mainframe/account => ../mainframe/account

require (
	github.com/kelseyhightower/envconfig v1.4.0
	go.mongodb.org/mongo-driver v1.15.0
	gopkg.in/yaml.v3 v3.0.1
	mainframe/account v0.0.0-00010101000000-000000000000
	mainframe/user v0.0.0-00010101000000-000000000000
)
