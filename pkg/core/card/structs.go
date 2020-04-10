package card

type StructCard struct {
	Pan        int64  `json:"pan"`
	Pin        int64  `json:"pin"`
	Balance    int64  `json:"balance"`
	Cvv        int64  `json:"cvv"`
	HolderName string `json:"holder_name"`
	Validity   string `json:"validity"`
	ClientId   int64  `json:"client_id"`
}

type CreateCard struct {
	ClientId   int64  `json:"client_id"`
	Pin        int64  `json:"pin"`
	Balance    int64  `json:"balance"`
	Cvv        int64  `json:"cvv"`
	HolderName string `json:"holder_name"`
	Validity   string `json:"validity"`
}

type LockUnLock struct {
	PAN int64 `json:"pan"`
}
