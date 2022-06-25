module github.com/LyricTian/gin-admin/v6

go 1.13

require (
	github.com/LyricTian/captcha v1.1.0
	github.com/LyricTian/gzip v0.1.1
	github.com/LyricTian/queue v1.2.0
	github.com/LyricTian/structs v1.1.1
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/bwmarrin/snowflake v0.3.0
	github.com/casbin/casbin/v2 v2.4.1
	github.com/denisenkom/go-mssqldb v0.0.0-20200428022330-06a60b6afbbc // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fatih/camelcase v1.0.0 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/spec v0.20.3 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/go-redis/redis_rate v6.5.0+incompatible
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/gops v0.3.8
	github.com/google/uuid v1.1.1
	github.com/google/wire v0.4.0
	github.com/jinzhu/gorm v1.9.12
	github.com/jinzhu/now v1.1.1 // indirect
	github.com/json-iterator/go v1.1.9
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/klauspost/compress v1.10.5 // indirect
	github.com/koding/multiconfig v0.0.0-20171124222453-69c27309b2d7
	github.com/lib/pq v1.5.2 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/onsi/ginkgo v1.12.0 // indirect
	github.com/onsi/gomega v1.10.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.6.0
	github.com/stretchr/testify v1.7.0
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.7.0
	github.com/tidwall/buntdb v1.1.2
	github.com/tidwall/gjson v1.6.0 // indirect
	github.com/tidwall/pretty v1.0.1 // indirect
	github.com/urfave/cli/v2 v2.10.3
	go.mongodb.org/mongo-driver v1.5.1
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // indirect
	golang.org/x/net v0.0.0-20210716203947-853a461950ff // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1
	golang.org/x/tools v0.1.5 // indirect
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37 => github.com/golang/crypto v0.0.0-20200510223506-06a226fb4e37
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 => github.com/golang/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/net v0.0.0-20190503192946-f4e77d36d62c => github.com/golang/net v0.0.0-20190503192946-f4e77d36d62c
	golang.org/x/net v0.0.0-20190611141213-3f473d35a33a => github.com/golang/net v0.0.0-20190611141213-3f473d35a33a
	golang.org/x/net v0.0.0-20200506145744-7e3656a0809f => github.com/golang/net v0.0.0-20200506145744-7e3656a0809f
	golang.org/x/net v0.0.0-20200625001655-4c5254603344 => github.com/golang/net v0.0.0-20200625001655-4c5254603344
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a => github.com/golang/sync v0.0.0-20200317015054-43a5402ce75a
	golang.org/x/sys v0.0.0-20171017063910-8dbc5d05d6ed => github.com/golang/sys v0.0.0-20171017063910-8dbc5d05d6ed
	golang.org/x/sys v0.0.0-20181228144115-9a3f9b0469bb => github.com/golang/sys v0.0.0-20181228144115-9a3f9b0469bb
	golang.org/x/sys v0.0.0-20190610200419-93c9922d18ae => github.com/golang/sys v0.0.0-20190610200419-93c9922d18ae
	golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a => github.com/golang/sys v0.0.0-20181228144115-9a3f9b0469bb
	golang.org/x/sys v0.0.0-20191120155948-bd437916bb0e => github.com/golang/sys v0.0.0-20191120155948-bd437916bb0e
	golang.org/x/sys v0.0.0-20200116001909-b77594299b42 => github.com/golang/sys v0.0.0-20200116001909-b77594299b42
	golang.org/x/sys v0.0.0-20200509044756-6aff5f38e54f => github.com/golang/sys v0.0.0-20200509044756-6aff5f38e54f
	golang.org/x/text v0.3.3 => github.com/golang/text v0.3.3
	golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1 => github.com/golang/time v0.0.0-20200416051211-89c76fbcd5d1
	golang.org/x/tools v0.0.0-20190606050223-4d9ae51c2468 => github.com/golang/tools v0.0.0-20190606050223-4d9ae51c2468
	golang.org/x/tools v0.0.0-20190611222205-d73e1c7e250b => github.com/golang/tools v0.0.0-20190611222205-d73e1c7e250b
	golang.org/x/tools v0.0.0-20200511202723-1762287ae9dd => github.com/golang/tools v0.0.0-20200511202723-1762287ae9dd
	golang.org/x/tools v0.0.0-20200626171337-aa94e735be7f => github.com/golang/tools v0.0.0-20200626171337-aa94e735be7f
)
