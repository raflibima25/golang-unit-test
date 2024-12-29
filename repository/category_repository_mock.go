package repository

import (
	"belajar-golang-unit-test/entity"
	"github.com/stretchr/testify/mock"
)

type CategoryRepositoyMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoyMock) FindById(id string) *entity.Category {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	}

	category := arguments.Get(0).(entity.Category)
	return &category
}
