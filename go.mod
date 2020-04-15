module gitlab.intelligentb.com/examples/stringdemo/stringdemoservice

go 1.13

replace gitlab.intelligentb.com/devops/bplus => ../bplus

replace gitlab.intelligentb.com/examples/stringdemo/stringdemoapi => ../stringdemoapi

require (
	github.com/DATA-DOG/godog v0.7.13
	gitlab.intelligentb.com/devops/bplus v0.0.0-00010101000000-000000000000
	gitlab.intelligentb.com/examples/stringdemo/stringdemoapi v0.0.0-00010101000000-000000000000
)
