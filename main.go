package main

import (
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("urls")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	viper.WatchConfig()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		code := strings.TrimPrefix(r.RequestURI, "/")

		if url := viper.GetString(code); url != "" {
			w.Header().Set("Location", url)
			w.WriteHeader(http.StatusPermanentRedirect)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	http.ListenAndServe(":5050", nil)
}
