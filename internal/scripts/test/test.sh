curl  -XPOST -d'{"s":"hello, world"}' localhost:5000/stringdemo/count
curl  -XPOST -H "SecureToken: passpass" -d'{"s":"hello, world"}' localhost:5000/stringdemo/uppercase
curl  -XGET -d'{}' -H "Arg1: 29" -H "arg2: 33" localhost:5000/stringdemo/add-numbers
curl  -XGET -d'{}' "localhost:5000/stringdemo/add-numbers?Arg1=29&Arg2=33"
curl  -XGET -d'{}' localhost:5000/stringdemo/add-numbers-path/29/33
