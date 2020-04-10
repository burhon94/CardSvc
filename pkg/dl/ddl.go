package dl

const CardDDL = `CREATE TABLE IF NOT EXISTS cards
(
    id         BIGSERIAL PRIMARY KEY,
    pan        INTEGER NOT NULL UNIQUE,
    pin        INTEGER NOT NULL,
    balance    INTEGER NOT NULL,
    cvv        INTEGER NOT NULL,
    holderName TEXT    NOT NULL,
    validity   TEXT    NOT NULL,
    client_id  INTEGER NOT NULL,
	lock       bool DEFAULT FALSE
);`
