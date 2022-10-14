package sqlite

// import (
// 	"context"
// 	"database/sql"
// 	"database/sql/driver"
// 	"time"
// )

// func (c *Client) Begin() (*sql.Tx, error) {
// 	if c.db == nil {
// 		return nil, ErrNotConnected
// 	}
// 	return c.db.Begin()
// }

// func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
// 	if c.db == nil {
// 		return nil, ErrNotConnected
// 	}
// 	return c.db.BeginTx(ctx, opts)
// }

// func (c *Client) Close() {
// 	if c.db == nil {
// 		return
// 	}
// 	c.db.Close()
// }

// func (c *Client) Conn(ctx context.Context) (*sql.Conn, error) {
// 	if c.db == nil {
// 		return nil, ErrNotConnected
// 	}
// 	return c.db.Conn(ctx)
// }

// func (c *Client) Driver() driver.Driver {
// 	if c.db == nil {
// 		return nil
// 	}
// 	return c.db.Driver()
// }

// func (c *Client) Exec(query string, args ...any) (sql.Result, error) {
// 	if c.db == nil {
// 		return nil, ErrNotConnected
// 	}
// 	return c.db.Exec(query, args)
// }

// func (c *Client) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
// 	if c.db == nil {
// 		return nil, ErrNotConnected
// 	}
// 	return c.db.ExecContext(ctx, query, args)
// }

// func (c *Client) Ping() error {
// 	if c.db == nil {
// 		return ErrNotConnected
// 	}
// 	return c.db.Ping()
// }

// func (c *Client) PingContext(ctx context.Context) error {
// 	if c.db == nil {
// 		return ErrNotConnected
// 	}
// 	return c.db.PingContext(ctx)
// }

// func (c *Client) Prepare(query string) (*sql.Stmt, error) {
// 	if c.db == nil {
// 		return nil, ErrNotConnected
// 	}
// 	return c.db.Prepare(query)
// }

// func (c *Client) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
// 	if c.db == nil {
// 		return nil, ErrNotConnected
// 	}
// 	return c.db.PrepareContext(ctx, query)
// }

// func (c *Client) Query(query string, args ...any) (*sql.Rows, error) {
// 	if c.db == nil {
// 		return nil, ErrNotConnected
// 	}
// 	return c.db.Query(query, args)
// }

// func (c *Client) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
// 	if c.db == nil {
// 		return nil, ErrNotConnected
// 	}
// 	return c.db.QueryContext(ctx, query, args)
// }

// func (c *Client) QueryRow(query string, args ...any) *sql.Row {
// 	if c.db == nil {
// 		return nil
// 	}
// 	return c.db.QueryRow(query, args)
// }

// func (c *Client) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
// 	if c.db == nil {
// 		return nil
// 	}
// 	return c.db.QueryRowContext(ctx, query, args)
// }

// func (c *Client) SetConnMaxIdleTime(d time.Duration) {
// 	if c.db == nil {
// 		return
// 	}
// 	c.db.SetConnMaxIdleTime(d)
// }

// func (c *Client) SetConnMaxLifetime(d time.Duration) {
// 	if c.db == nil {
// 		return
// 	}
// 	c.db.SetConnMaxLifetime(d)
// }

// func (c *Client) SetMaxIdleConns(n int) {
// 	if c.db == nil {
// 		return
// 	}
// 	c.db.SetMaxIdleConns(n)
// }

// func (c *Client) SetMaxOpenConns(n int) {
// 	if c.db == nil {
// 		return
// 	}
// 	c.db.SetMaxOpenConns(n)
// }

// func (c *Client) Stats() sql.DBStats {
// 	if c.db == nil {
// 		return sql.DBStats{}
// 	}
// 	return c.db.Stats()
// }
