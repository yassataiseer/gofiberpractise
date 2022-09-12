module example.com

go 1.19

require github.com/gofiber/fiber/v2 v2.37.1

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	gorm.io/driver/mysql v1.3.6 // indirect
	gorm.io/gorm v1.23.8 // indirect
)

require (
	example.com/handlers v0.0.0
	example.com/logic v0.0.0
	example.com/models v0.0.0
	example.com/config v0.0.0	
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/klauspost/compress v1.15.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.40.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20220227234510-4e6760a101f9 // indirect
)

replace example.com/handlers => ./handlers
replace example.com/config => ./config
replace example.com/logic => ./logic
replace example.com/models => ./models