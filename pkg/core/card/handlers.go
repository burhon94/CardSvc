package card

import (
	"context"
	"github.com/burhon94/CardSvc/pkg/dl"
	"log"
	"net/http"
)

func (c *Card) HandleGetCards(ctx context.Context) (cards []StructCard, err error) {
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
		var card StructCard
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

func (c *Card) HandleMyCards(ctx context.Context, request *http.Request) (cards []StructCard, err error) {
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
		var card StructCard
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

func (c *Card) HandleMyCard(ctx context.Context, request *http.Request, idCard int64) (card StructCard, err error) {
	id := 1
	rows, err := c.pool.Query(ctx, dl.GetMyCard, id, idCard)
	if err != nil {
		log.Print(err)
		return card, err
	}
	defer func() {
		if innerErr := rows.Close; innerErr != nil {
			return
		}
	}()
	for rows.Next() {
		err = rows.Scan(&card.Pan, &card.Pin, &card.Balance, &card.Cvv, &card.HolderName, &card.Validity, &card.ClientId)
		if err != nil {
			log.Print(err)
			return card, err
		}

	}
	if rows.Err() != nil {
		log.Print(rows.Err())
		return card, rows.Err()
	}

	return card, nil
}