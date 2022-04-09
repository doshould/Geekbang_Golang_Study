// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"hello/internal/data/ent/migrate"

	"hello/internal/data/ent/movie"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Movie is the client for interacting with the Movie builders.
	Movie *MovieClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Movie = NewMovieClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Movie:  NewMovieClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Movie:  NewMovieClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Movie.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Movie.Use(hooks...)
}

// MovieClient is a client for the Movie schema.
type MovieClient struct {
	config
}

// NewMovieClient returns a client for the Movie from the given config.
func NewMovieClient(c config) *MovieClient {
	return &MovieClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `movie.Hooks(f(g(h())))`.
func (c *MovieClient) Use(hooks ...Hook) {
	c.hooks.Movie = append(c.hooks.Movie, hooks...)
}

// Create returns a create builder for Movie.
func (c *MovieClient) Create() *MovieCreate {
	mutation := newMovieMutation(c.config, OpCreate)
	return &MovieCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Movie entities.
func (c *MovieClient) CreateBulk(builders ...*MovieCreate) *MovieCreateBulk {
	return &MovieCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Movie.
func (c *MovieClient) Update() *MovieUpdate {
	mutation := newMovieMutation(c.config, OpUpdate)
	return &MovieUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MovieClient) UpdateOne(m *Movie) *MovieUpdateOne {
	mutation := newMovieMutation(c.config, OpUpdateOne, withMovie(m))
	return &MovieUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MovieClient) UpdateOneID(id int64) *MovieUpdateOne {
	mutation := newMovieMutation(c.config, OpUpdateOne, withMovieID(id))
	return &MovieUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Movie.
func (c *MovieClient) Delete() *MovieDelete {
	mutation := newMovieMutation(c.config, OpDelete)
	return &MovieDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MovieClient) DeleteOne(m *Movie) *MovieDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MovieClient) DeleteOneID(id int64) *MovieDeleteOne {
	builder := c.Delete().Where(movie.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MovieDeleteOne{builder}
}

// Query returns a query builder for Movie.
func (c *MovieClient) Query() *MovieQuery {
	return &MovieQuery{
		config: c.config,
	}
}

// Get returns a Movie entity by its id.
func (c *MovieClient) Get(ctx context.Context, id int64) (*Movie, error) {
	return c.Query().Where(movie.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MovieClient) GetX(ctx context.Context, id int64) *Movie {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MovieClient) Hooks() []Hook {
	return c.hooks.Movie
}
