package data

import (
	"online-teaching/internal/conf"
	"online-teaching/internal/ent"

	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewCourseRepo)

// Data .
type Data struct {
	Client *ent.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	client, err := ent.Open("mysql", "test_rw:54rltyi5BCdcm06wu22A0brvvzU5uDgB@tcp(bjfk-d13.yz02:15101)/online-teaching?parseTime=true&loc=Asia%2FShanghai&charset=utf8mb4")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	return &Data{
		Client: client,
	}, cleanup, nil
}
