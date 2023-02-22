package postgres

import (
	"fmt"
	"github.com/gozelle/gorm"
	"github.com/gozelle/gorm/logger"
)

func NewDB(dsn string, opts ...gorm.Option) (*gorm.DB, error) {
	useOpts := make([]gorm.Option, 0)
	if len(opts) == 0 {
		useOpts = append(useOpts, &gorm.Config{
			Logger: logger.Default,
		})
	} else {
		useOpts = append(useOpts, opts...)
	}
	db, err := gorm.Open(Open(dsn), useOpts...)
	if err != nil {
		err = fmt.Errorf("connect postgres error: %s", err)
		return nil, err
	}
	sqldb, err := db.DB()
	if err != nil {
		err = fmt.Errorf("get postgres sql.DB error: %s", err)
		return nil, err
	}
	err = sqldb.Ping()
	if err != nil {
		err = fmt.Errorf("ping postgres error:%s", err)
		return nil, err
	}
	return db, nil
}
