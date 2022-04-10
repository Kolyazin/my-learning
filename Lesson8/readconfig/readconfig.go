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

type UrlUrl url.URL

func (uu *UrlUrl) Decode(value string) error {
	zz, err := url.Parse(value)
	*uu = UrlUrl(*zz)
	return err
}

type ValidConfig struct {
	Port int
	Db_url *UrlUrl
	Jaeger_url *UrlUrl
	Sentry_url *UrlUrl
	Kafka_broker *UrlUrl
	Some_app_id string
	Some_app_key string
}


var (
	RawConf RawConfig
	ValidConf = ValidConfig{0, new(UrlUrl), new(UrlUrl), new(UrlUrl), new(UrlUrl), "",""}
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
	var err error
	flag.Parse()

	if flag.NFlag() == 0 {
		//если флагов нет, то беру из переменных окружения
		err = envconfig.Process("", &ValidConf)
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
//		ValidConf = ValidConfig{}
		//валидирую url
		ValidConf.Port = RawConf.Port
		db_url, err := url.Parse(RawConf.Db_url)
		if err != nil {
			fmt.Println(err.Error())
		}
		*ValidConf.Db_url = UrlUrl(*db_url)
		jaeger_url, err := url.Parse(RawConf.Jaeger_url)
		if err != nil {
			fmt.Println(err.Error())
		}
		*ValidConf.Jaeger_url = UrlUrl(*jaeger_url)
		sentry_url, err := url.Parse(RawConf.Sentry_url)
		if err != nil {
			fmt.Println(err.Error())
		}
		*ValidConf.Sentry_url = UrlUrl(*sentry_url)
		kafka_broker, err := url.Parse(RawConf.Kafka_broker)
		if err != nil {
			fmt.Println(err.Error())
		}
		*ValidConf.Kafka_broker = UrlUrl(*kafka_broker)
		ValidConf.Some_app_id = RawConf.Some_app_id
		ValidConf.Some_app_key = RawConf.Some_app_key
	}
	return ValidConf
}