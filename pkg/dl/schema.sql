CREATE TABLE IF NOT EXISTS cards
(
    id         BIGSERIAL PRIMARY KEY,
    pan        INTEGER NOT NULL UNIQUE,
    pin        INTEGER NOT NULL,
    balance    INTEGER NOT NULL,
    cvv        INTEGER NOT NULL,
    holderName TEXT    NOT NULL,
    validity   INTEGER NOT NULL,
    client_id  INTEGER NOT NULL,
    lock       bool DEFAULT FALSE
);

INSERT INTO cards
    (pan, pin, balance, cvv, holderName, validity, client_id, lock)
VALUES (11111, 1111, 1111, 111, 'NAME11', 1111, 1, false),
       (22222, 2222, 2222, 222, 'NAME22', 2222, 2, false),
       (33333, 3333, 3333, 333, 'NAME33', 3333, 1, false)
ON CONFLICT DO NOTHING;

SELECT pan, pin, balance, cvv, holderName, validity, client_id
FROM cards;

SELECT pan, pin, balance, cvv, holderName, validity, client_id
FROM cards
WHERE pan = 11111
  and client_id = 1;

select pan
from cards
order by pan desc;

update cards
set lock = true
where pan = ?;