package adapters

import (
	"github.com/Najah7/task2schedule/internal/domain/common"
	"github.com/oklog/ulid/v2"
)

var _ common.IDGenerator = ULIDGenerator{}

type ULIDGenerator struct{}

func NewULIDGenerator() ULIDGenerator {
	return ULIDGenerator{}
}

func (g ULIDGenerator) Generate() string {
	return ulid.Make().String()
}
