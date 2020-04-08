package dl

const CardDML = `INSERT INTO cards
    (pan, pin, balance, cvv, holderName, validity, client_id)
VALUES (11111, 1111, 1111, 111, 'NAME11', 1111, 1),
       (22222, 2222, 2222, 222, 'NAME22', 2222, 2),
       (33333, 3333, 3333, 333, 'NAME33', 3333, 1)
ON CONFLICT DO NOTHING;`

const GetCards = `SELECT pan, pin, balance, cvv, holderName, validity, client_id FROM cards;`

const GetMyCards = `SELECT pan, pin, balance, cvv, holderName, validity, client_id FROM cards WHERE client_id = $1;`

const GetMyCard = `SELECT pan, pin, balance, cvv, holderName, validity, client_id FROM cards WHERE pan = $2 and client_id = $1;`
