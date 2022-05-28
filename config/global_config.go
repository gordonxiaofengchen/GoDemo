package config

import(
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
	"time"
)

var projectPath string

func init(){
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if !strings.HasSuffix(wd, string(os.PathSeparator) + "main") {
		projectPath = wd
	}else{
		projectPath = wd[:strings.Index(wd, string(os.PathSeparator) + "main")]
	}

	viper.SetConfigFile(projectPath + string(os.PathSeparator) + "resources" + string(os.PathSeparator) + "config.yml")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func ProjectPath() string {
	return projectPath
}

func Get(key string) interface{}{
	return viper.Get(key)
}

func GetString(key string) string{
	return viper.GetString(key)
}

func GetBool(key string) bool{
	return viper.GetBool(key)
}

func GetInt(key string) int{
	return viper.GetInt(key)
}

func GetStringSlice(key string) []string{
	return viper.GetStringSlice(key)
}

func GetDuration(key string) time.Duration{
	return viper.GetDuration(key)
}

func Has(key string) bool {
	return viper.InConfig(key)
}

