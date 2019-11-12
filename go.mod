module github.com/crazy-canux/go-devops

go 1.13

require (
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/coreos/go-iptables v0.4.3
	github.com/denisenkom/go-mssqldb v0.0.0-20191001013358-cfbb681360f0
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v1.13.1
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/go-ini/ini v1.51.0
	github.com/go-redis/redis v6.15.6+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/google/go-cmp v0.3.1 // indirect
	github.com/influxdata/influxdb v1.7.9
	github.com/kr/pretty v0.1.0 // indirect
	github.com/onsi/ginkgo v1.10.3 // indirect
	github.com/onsi/gomega v1.7.1 // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/streadway/amqp v0.0.0-20190827072141-edfb9018d271
	github.com/stretchr/testify v1.4.0 // indirect
	golang.org/x/crypto v0.0.0-20191111213947-16651526fdb4
	golang.org/x/net v0.0.0-20191109021931-daa7c04131f5
	golang.org/x/sys v0.0.0-20191110163157-d32e6e3b99c4 // indirect
	google.golang.org/appengine v1.6.5 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/ini.v1 v1.50.0 // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/yaml.v2 v2.2.5 // indirect
)

replace (
	golang.org/x/crypto v0.0.0-20191106202628-ed6320f186d4 => github.com/golang/crypto v0.0.0-20191106202628-ed6320f186d4
	golang.org/x/net v0.0.0-20191105084925-a882066a44e0 => github.com/golang/net v0.0.0-20191105084925-a882066a44e0
)
