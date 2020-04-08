package card

import (
	"context"
	"github.com/burhon94/CardSvc/pkg/dl"
	"log"
	"net/http"
)

func (c *Card) HandleGetCards(ctx context.Context) (cards []Struct, err error) {
	rows, err := c.pool.Query(ctx, dl.GetCards)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer func() {
		if innerErr := rows.Close; innerErr != nil {
			return
		}
	}()
	for rows.Next() {
		var card Struct
		err = rows.Scan(&card.Pan, &card.Pin, &card.Balance, &card.Cvv, &card.HolderName, &card.Validity, &card.ClientId)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		cards = append(cards, card)
	}
	if rows.Err() != nil {
		log.Print(rows.Err())
		return nil, rows.Err()
	}

	return cards, nil
}

func (c *Card) HandleMyCards(ctx context.Context, request *http.Request) (cards []Struct, err error) {
	id := 2
	rows, err := c.pool.Query(ctx, dl.GetMyCards, id)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer func() {
		if innerErr := rows.Close; innerErr != nil {
			return
		}
	}()
	for rows.Next() {
		var card Struct
		err = rows.Scan(&card.Pan, &card.Pin, &card.Balance, &card.Cvv, &card.HolderName, &card.Validity, &card.ClientId)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		cards = append(cards, card)
	}
	if rows.Err() != nil {
		log.Print(rows.Err())
		return nil, rows.Err()
	}
	return cards, nil
}