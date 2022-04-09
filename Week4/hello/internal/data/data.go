package data

import (
	_ "entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"hello/internal/conf"
	"hello/internal/data/ent"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	db *ent.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)
	client, err := ent.Open(
		c.Database.Driver,
		c.Database.Source,
	)
	//sqlDrv := dialect.DebugWithContext(drv, func(ctx context.Context, i ...interface{}) {})
	//client := ent.NewClient(ent.Driver(drv))
	if err != nil {
		return nil, nil, err
	}

	//if err := client.Schema.Create(context.Background()); err != nil {
	//	log.Errorf("failed creating schema resources: %v", err)
	//	return nil, nil, err
	//}

	d := &Data{db: client}
	cleanup := func() {
		log.Info("closing the data resources")
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}
	return d, cleanup, nil
}
