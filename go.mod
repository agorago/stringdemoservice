module github.com/agorago/stringdemoservice

go 1.13

replace github.com/agorago/wego => ../wego

replace github.com/agorago/stringdemoapi => ../stringdemoapi

require (
	github.com/DATA-DOG/godog v0.7.13
	github.com/agorago/stringdemoapi v0.0.0-00010101000000-000000000000
	github.com/agorago/wego v0.0.0-00010101000000-000000000000
)
