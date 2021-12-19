module learn.go

go 1.17

require (
	github.com/shopspring/decimal v1.3.1
	github.com/spf13/cobra v1.3.0
	github.com/wxywizard/go_bmi v0.0.0-00010101000000-000000000000
//github.com/wxywizard/go_bmi v0.0.1
)

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

replace (
	//二次开发的工具包重新指定fork
	//github.com/spf13/cobra => github.com/armstrongli/cobra v1.2.0
	//本地
	github.com/wxywizard/go_bmi => ./staging/src/github.com/wxywizard/go_bmi
)
