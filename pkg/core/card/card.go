package card

import (
	"context"
	"fmt"
	"github.com/burhon94/CardSvc/pkg/dl"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Card struct {
	pool *pgxpool.Pool
}

func NewCard(pool *pgxpool.Pool) *Card {
	return &Card{pool: pool}
}

func (c *Card) Start() {
	_, err := c.pool.Exec(context.Background(), dl.CardDDL)
	if err != nil {
		panic(fmt.Sprintf("can't init DB: %v", err))
	}

	_, err = c.pool.Exec(context.Background(), dl.CardDML)
	if err != nil {
		panic(fmt.Sprintf("can't set DB: %v", err))
	}
}