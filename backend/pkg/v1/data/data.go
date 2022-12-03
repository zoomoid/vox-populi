package data

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/google/wire"
	conf "github.com/zoomoid/vox-populi/backend/pkg/config"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent"
)

var Repositories = wire.NewSet(
	NewData,
	NewPollRepo,
	NewVoteRepo,
	// NewVoteTemplateRepo,
	NewReactionRepo,
	// NewReactionTemplateRepo,
)

type Data struct {
	db *ent.Client
}

func NewData(c *conf.Configuration) (*Data, func(), error) {
	driver, err := sql.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		return nil, nil, err
	}
	sqlDriver := dialect.DebugWithContext(driver, func(ctx context.Context, a ...any) {})

	client := ent.NewClient(ent.Driver(sqlDriver))

	if err := client.Schema.Create(context.TODO()); err != nil {
		return nil, nil, err
	}

	d := &Data{
		db: client,
	}

	cleanup := func() {
		if err = d.db.Close(); err != nil {
			panic(err)
		}
	}

	return d, cleanup, nil
}

func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
