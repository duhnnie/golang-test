module hello

go 1.15

replace example.com/greetings => ../greetings

replace example.com/math => ../math

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	example.com/math v0.0.0-00010101000000-000000000000
	rsc.io/quote v1.5.2 // indirect
)
