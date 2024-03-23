package vo

import "github.com/google/uuid"

const TTL_SECOND int = 604800

type Token struct {
	value uuid.UUID
}

func NewToken() (Token, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return Token{}, err
	}

	return Token{uuid}, nil
}

func (token Token) String() string {
	return token.value.String()
}
