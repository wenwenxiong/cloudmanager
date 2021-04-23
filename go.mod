module cloudmanager

require (
	github.com/go-redis/redis/v8 v8.8.2
	github.com/gogf/gf v1.15.6
	github.com/gookit/color v1.4.2
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/viper v1.7.1
	github.com/wenwenxiong/cloudmanager v0.0.0-00010101000000-000000000000
	gorm.io/gorm v1.21.8
)

replace github.com/wenwenxiong/cloudmanager => ../cloudmanager

go 1.11
