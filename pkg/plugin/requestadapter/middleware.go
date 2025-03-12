package requestadapter

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	Mapping map[string]string `json:"mapping"`
}

// NewRequestAdapter creates a new instance of RequestAdapter
func NewRequestAdapter(config Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// remove request body json
			bytes, err := ioutil.ReadAll(r.Body)
			log.Info(string(bytes))
			if err == nil {
				var result map[string]string
				if err := json.Unmarshal(bytes, &result); err != nil {
					result = make(map[string]string)
				}

				log.WithField("result", result).Info("decoded value")

				form := url.Values{}

				// add form-urlencode
				for jsonField, formField := range config.Mapping {
					form.Set(formField, result[jsonField])
				}

				formData := form.Encode()
				r.Body = ioutil.NopCloser(strings.NewReader(formData))
				r.ContentLength = int64(len(formData))
			}

			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			next.ServeHTTP(w, r)
		})
	}
}
