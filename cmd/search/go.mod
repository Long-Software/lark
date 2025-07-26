module github.com/Long-Software/lark/cmd/search

go 1.23.2

replace github.com/Long-Software/lark/pkg/log => ../../pkg/log

require (
	github.com/Long-Software/lark/pkg/log v0.0.0-20250719154726-ebbe9476ca44
	github.com/PuerkitoBio/goquery v1.10.3
)

require (
	github.com/andybalholm/cascadia v1.3.3 // indirect
	golang.org/x/net v0.39.0 // indirect
)
