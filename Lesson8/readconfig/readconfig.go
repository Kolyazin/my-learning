package readconfig

import (
	"fmt"
	"flag"
	"github.com/kelseyhightower/envconfig"
	"net/url"
)

type RawConfig struct {
	Port int
	Db_url string
	Jaeger_url string
	Sentry_url string
	Kafka_broker string
	Some_app_id string
	Some_app_key string
}

type ValidConfig struct {
	Port int
	Db_url_parsed *url.URL
	Jaeger_url_parsed *url.URL
	Sentry_url_parsed *url.URL
	Kafka_broker_parsed *url.URL
	Some_app_id string
	Some_app_key string
}

var (
	RawConf RawConfig
	ValidConf ValidConfig
)


func init() {
	flag.IntVar(&RawConf.Port, "port", 8080, "port")                                                                                //: 8080
	flag.StringVar(&RawConf.Db_url, "db_url", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "db_url") //: postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable
	flag.StringVar(&RawConf.Jaeger_url, "jaeger_url", "http://jaeger:16686", "jaeger")                                              //: http://jaeger:16686
	flag.StringVar(&RawConf.Sentry_url, "sentry_url", "http://sentry:9000", "sentry_url")                                           //: http://sentry:9000
	flag.StringVar(&RawConf.Kafka_broker, "kafka_broker", "kafka:9092", "kafka_broker")                                             //: kafka:9092
	flag.StringVar(&RawConf.Some_app_id, "some_app_id", "DefaultId", "some_app_id")                                                    //: testid
	flag.StringVar(&RawConf.Some_app_key, "some_app_key", "DefaultKey", "some_app_key")                                                //: testkey
}

func ReadConfig() ValidConfig {
	flag.Parse()

	if flag.NFlag() == 0 {
//если флагов нет, то беру из переменных окружения
		err := envconfig.Process("", &RawConf)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
//валидирую url
	ValidConf.Port = RawConf.Port
	ValidConf.Db_url_parsed, _ = url.Parse(RawConf.Db_url)
	ValidConf.Jaeger_url_parsed, _ = url.Parse(RawConf.Jaeger_url)
	ValidConf.Sentry_url_parsed, _ = url.Parse(RawConf.Sentry_url)
	ValidConf.Kafka_broker_parsed, _ = url.Parse(RawConf.Kafka_broker)
	ValidConf.Some_app_id = RawConf.Some_app_id
	ValidConf.Some_app_key = RawConf.Some_app_key

	return ValidConf
}