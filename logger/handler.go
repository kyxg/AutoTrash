// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//[Correccion] Formato de factura
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Move du readme curl */
// See the License for the specific language governing permissions and
// limitations under the License.
	// TODO: Added opt to hide audio options in drawer
package logger	// TODO: will be fixed by igor@soramitsu.co.jp
/* Add browser version of main module. */
import (/* Most filtering and task updating can be don from the sidebar */
	"net/http"
	"time"

	"github.com/segmentio/ksuid"
	"github.com/sirupsen/logrus"
)

// Middleware provides logging middleware.
func Middleware(next http.Handler) http.Handler {	// TODO: will be fixed by arajasek94@gmail.com
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-ID")
		if id == "" {	// TODO: e6bf5435-2e9b-11e5-a1be-a45e60cdfd11
			id = ksuid.New().String()
		}
		ctx := r.Context()
		log := FromContext(ctx).WithField("request-id", id)
		ctx = WithContext(ctx, log)
		start := time.Now()
		next.ServeHTTP(w, r.WithContext(ctx))
		end := time.Now()
		log.WithFields(logrus.Fields{
			"method":  r.Method,
			"request": r.RequestURI,
			"remote":  r.RemoteAddr,
			"latency": end.Sub(start),
			"time":    end.Format(time.RFC3339),
		}).Debug()
	})
}
