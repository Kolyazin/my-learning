package readconfig

import (
	"fmt"
	"flag"
	"github.com/kelseyhightower/envconfig"
	"net/url"
)

type Config struct {
	Port int
	Db_url string
	Db_url_parsed *url.URL
	Jaeger_url string
	Jaeger_url_parsed *url.URL
	Sentry_url string
	Sentry_url_parsed *url.URL
	Kafka_broker string
	Kafka_broker_parsed *url.URL
	Some_app_id string
	Some_app_key string
}

var (
	Conf Config
)


func init() {
	flag.IntVar(&Conf.Port, "port", 8080, "port")                                                                                //: 8080
	flag.StringVar(&Conf.Db_url, "db_url", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "db_url") //: postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable
	flag.StringVar(&Conf.Jaeger_url, "jaeger_url", "http://jaeger:16686", "jaeger")                                              //: http://jaeger:16686
	flag.StringVar(&Conf.Sentry_url, "sentry_url", "http://sentry:9000", "sentry_url")                                           //: http://sentry:9000
	flag.StringVar(&Conf.Kafka_broker, "kafka_broker", "kafka:9092", "kafka_broker")                                             //: kafka:9092
	flag.StringVar(&Conf.Some_app_id, "some_app_id", "DefaultId", "some_app_id")                                                    //: testid
	flag.StringVar(&Conf.Some_app_key, "some_app_key", "DefaultKey", "some_app_key")                                                //: testkey
}

func ReadConfig() Config {
	flag.Parse()

	if flag.NFlag() == 0 {
//если флагов нет, то беру из переменных окружения
		err := envconfig.Process("", &Conf)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
//валидирую url
	Conf.Db_url_parsed, _ = url.Parse(Conf.Db_url)
	Conf.Jaeger_url_parsed, _ = url.Parse(Conf.Jaeger_url)
	Conf.Sentry_url_parsed, _ = url.Parse(Conf.Sentry_url)
	Conf.Kafka_broker_parsed, _ = url.Parse(Conf.Kafka_broker)

	return Conf
}