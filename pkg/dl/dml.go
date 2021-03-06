package dl

const CardDML = `INSERT INTO cards
    (pan, pin, balance, cvv, holderName, validity, client_id, lock)
VALUES (11111, 1111, 1111, 111, 'NAME11', 1111, 1, false),
       (22222, 2222, 2222, 222, 'NAME22', 2222, 2, false),
       (33333, 3333, 3333, 333, 'NAME33', 3333, 1, false)
ON CONFLICT DO NOTHING;`

const GetCards = `SELECT pan, pin, balance, cvv, holderName, validity, client_id, lock FROM cards;`

const GetMyCards = `SELECT pan, pin, balance, cvv, holderName, validity, client_id, lock FROM cards WHERE client_id = $1;`

const GetMyCard = `SELECT pan, pin, balance, cvv, holderName, validity, client_id, lock FROM cards WHERE pan = $2 and client_id = $1;`

const GetPan = `select pan from cards order by pan desc;`

const CreateCard = `INSERT INTO cards
    (pan, pin, balance, cvv, holderName, validity, client_id)
VALUES ($1, $2, $3, $4, $5, $6, $7);`

const CheckPan = `SELECT pan FROM cards WHERE pan = $1;`

const LockCard = `UPDATE cards SET lock = TRUE WHERE pan = $1;`

const UnlockCard = `UPDATE cards SET lock = FALSE WHERE pan = $1;`

const LockMyCard = `UPDATE cards SET lock = TRUE WHERE id = $1;`

const UnlockMyCard = `UPDATE cards SET lock = FALSE WHERE id = $1;`
