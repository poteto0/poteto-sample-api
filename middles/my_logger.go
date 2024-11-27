package middles

import (
	"github.com/poteto0/poteto"
	"github.com/poteto0/poteto/middleware"
	"github.com/sirupsen/logrus"
)

func MyHTTPLoggerFunc(ctx poteto.Context, rlv middleware.RequestLoggerValues) error {
	log := logrus.New()
	if rlv.Error == nil {
		log.WithFields(logrus.Fields{
			"method":    rlv.Method,
			"routePath": rlv.RoutePath,
			"realIP":    rlv.RealIP,
			"status":    rlv.Status,
		}).Info("request")
	} else {
		log.WithFields(logrus.Fields{
			"method":    rlv.Method,
			"routePath": rlv.RoutePath,
			"realIP":    rlv.RealIP,
			"status":    rlv.Status,
		}).Error("request")
	}
	return nil
}
