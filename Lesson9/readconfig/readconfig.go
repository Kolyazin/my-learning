package readconfig

import (
	"encoding/json"
	"flag"
	"github.com/kelseyhightower/envconfig"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
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

// функция читает конфигурацию из флагов, если флаги не переданы, то берет из конфигурационного файла
// переданного в качестве аргумента, если ничего не передается, то читает из переменных окружения
func ReadConfig() (Config, error) {
	var err error
	flag.Parse()

	if flag.NFlag() == 0 {
		if flag.NArg() == 0 {
			//если флагов и аргументов нет, то беру из переменных окружения
			err = envconfig.Process("", &Conf)
			if err != nil {
				return Conf, err
			}
		} else {
			// если есть конфигурационный файл в качестве аргумента
			f, err := os.Open(flag.Arg(0))
			if err != nil {
				return Conf, err
			}

			defer func() {
				f.Close()
			}()

			content, err := ioutil.ReadAll(f)
			if err != nil {
				return Conf, err
			}
			if strings.HasSuffix(flag.Arg(0), ".yaml") {
				err = yaml.Unmarshal(content, &Conf)
				if err != nil {
					return Conf, err
				}
			} else {
				err = json.Unmarshal(content, &Conf)
				if err != nil {
					return Conf, err
				}
			}
		}
	}
//валидирую url
	Conf.Db_url_parsed, err = url.Parse(Conf.Db_url)
	if err != nil {
		return Conf, err
	}
	Conf.Jaeger_url_parsed, err = url.Parse(Conf.Jaeger_url)
	if err != nil {
		return Conf, err
	}
	Conf.Sentry_url_parsed, err = url.Parse(Conf.Sentry_url)
	if err != nil {
		return Conf, err
	}
	Conf.Kafka_broker_parsed, err = url.Parse(Conf.Kafka_broker)
	if err != nil {
		return Conf, err
	}

	return Conf, nil
}