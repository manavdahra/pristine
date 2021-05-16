package providers

import (
	"github.com/segmentio/ksuid"
)

type IdGenerator struct {
}

func NewIdGenerator() *IdGenerator {
	return &IdGenerator{}
}

func (idGen *IdGenerator) GenerateNewId() string {
	return ksuid.New().String()
}
