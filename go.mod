module hello_gin

go 1.14

require (
	github.com/fsnotify/fsnotify v1.4.7
	github.com/gin-gonic/gin v1.6.2
	github.com/jinzhu/gorm v1.9.12
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/viper v1.6.3
)

replace hello_gin => ./
