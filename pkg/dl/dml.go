package dl

const CardDML = `INSERT INTO cards
    (pan, pin, balance, cvv, holderName, validity, client_id)
VALUES (11111, 1111, 1111, 111, 'NAME11', 1111, 1),
       (22222, 2222, 2222, 222, 'NAME22', 2222, 2),
       (33333, 3333, 3333, 333, 'NAME33', 3333, 1)
ON CONFLICT DO NOTHING;`
