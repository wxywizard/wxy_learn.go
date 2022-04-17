module learn.go

go 1.17

require (
	//github.com/wxywizard/go_bmi v0.0.1
	//github.com/wxywizard/go_example v0.0.4
	github.com/armstrongli/go-bmi v0.0.1
	github.com/shopspring/decimal v1.3.1
	github.com/spf13/cobra v1.3.0
	github.com/wxywizard/go_bmi v0.0.0-00010101000000-000000000000
	github.com/wxywizard/go_example v0.0.4
)

require (
	github.com/xuri/excelize/v2 v2.5.0
	google.golang.org/grpc v1.43.0
	gorm.io/driver/mysql v1.3.2
	gorm.io/gorm v1.23.3
)

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/richardlehane/mscfb v1.0.3 // indirect
	github.com/richardlehane/msoleps v1.0.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/xuri/efp v0.0.0-20210322160811-ab561f5b45e3 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d // indirect
	golang.org/x/sys v0.0.0-20211205182925-97ca703d548d // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

//二次开发的工具包重新指定fork
//github.com/spf13/cobra => github.com/armstrongli/cobra v1.2.0
//本地
replace (
	github.com/armstrongli/go-bmi => ./staging/src/github.com/armstrongli/go-bmi
	github.com/wxywizard/go_bmi => ./staging/src/github.com/wxywizard/go_bmi
)
