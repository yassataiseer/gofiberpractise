module example.com/config

go 1.19

require (
	example.com/models v0.0.0
	gorm.io/driver/mysql v1.3.6
	gorm.io/gorm v1.23.8
)

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
)

replace example.com/models => ../models
