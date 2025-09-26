module mainframe-lib/dossier

go 1.24.5

replace mainframe-lib/account => ../account

replace mainframe-lib/common => ../common

require (
	mainframe-lib/account v0.0.0-00010101000000-000000000000
	mainframe-lib/common v0.0.0-00010101000000-000000000000
)
