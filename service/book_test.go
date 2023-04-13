package service

import (
	"errors"
	"sesi_8/model"
	"sesi_8/repository/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_BookService_GetBookByID(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		expectedResult model.Books
		expectedError  error
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBookByID(gomock.Any()).Return(model.Books{
				ID:     1,
				Title:  "Rizqi",
				Author: "rizqi@gmail.com",
			}, nil).Times(1)
		},
		expectedResult: model.Books{
			ID:     1,
			Title:  "Rizqi",
			Author: "rizqi@gmail.com",
		},
	})

	testTable = append(testTable, testCase{
		name:          "record not found",
		wantError:     true,
		expectedError: errors.New("record not found"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBookByID(gomock.Any()).Return(model.Books{}, errors.New("record not found")).Times(1)
		},
	})

	testTable = append(testTable, testCase{
		name:          "unexpected error",
		wantError:     true,
		expectedError: errors.New("unexpected error"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBookByID(gomock.Any()).Return(model.Books{}, errors.New("unexpected error")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			bookRepo := mocks.NewBookRepo(mockCtrl)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{
				repo: bookRepo,
			}

			res, err := service.GetBookByID(1)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}
}

func Test_BookService_CreateBook(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		input          model.Books
		expectedResult model.Books
		expectedError  error
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		input: model.Books{
			Title:  "Belajar Si Golang",
			Author: "Hacktiv9 Indonesia",
		},
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().CreateBook(gomock.Any()).Return(model.Books{
				ID:     1,
				Title:  "Belajar Si Golang",
				Author: "Hacktiv9 Indonesia",
			}, nil).Times(1)
		},
		expectedResult: model.Books{
			ID:     1,
			Title:  "Belajar Si Golang",
			Author: "Hacktiv9 Indonesia",
		},
	})

	testTable = append(testTable, testCase{
		name:      "unexpected error",
		wantError: true,
		input: model.Books{
			Title:  "Belajar Si Golang",
			Author: "Hacktiv9 Indonesia",
		},
		expectedError: errors.New("unexpected error"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().CreateBook(gomock.Any()).Return(model.Books{}, errors.New("unexpected error")).Times(1)
		},
	})

	testTable = append(testTable, testCase{
		name:      "invalid Title length",
		wantError: true,
		input: model.Books{
			Title:  "Golang si Golang", // case negative
			Author: "Hacktiv10 Indonesia",
		},
		expectedError: errors.New("invalid Titel length"),
	})

	testTable = append(testTable, testCase{
		name:      "invalid Author",
		wantError: true,
		input: model.Books{
			Title:  "Golang",
			Author: "anang",
		},
		expectedError: errors.New("invalid Author"),
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			bookRepo := mocks.NewBookRepo(mockCtrl)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{
				repo: bookRepo,
			}

			res, err := service.CreateBook(testCase.input)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}
}
func Test_BookService_DeleteBook(t *testing.T) {
	type testCase struct {
		name          string
		wantError     bool
		expectedError error
		onBookRepo    func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().DeleteBook(gomock.Any()).Return(nil).Times(1)
		},
	})

	testTable = append(testTable, testCase{
		name:          "record not found",
		wantError:     true,
		expectedError: errors.New("record not found"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().DeleteBook(gomock.Any()).Return(errors.New("record not found")).Times(1)
		},
	})

	testTable = append(testTable, testCase{
		name:          "unexpected error",
		wantError:     true,
		expectedError: errors.New("unexpected error"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().DeleteBook(gomock.Any()).Return(errors.New("unexpected error")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			bookRepo := mocks.NewBookRepo(mockCtrl)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{
				repo: bookRepo,
			}

			err := service.DeleteBook(1)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
func Test_BookService_UpdateBook(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		inputID        int64
		input          model.Books
		expectedResult model.Books
		expectedError  error
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		inputID:   1,
		input: model.Books{
			Title:  "Belajar Si Golang",
			Author: "Hacktiv9 Indonesia",
		},
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().UpdateBook(gomock.Any()).Return(nil)
		},
		expectedResult: model.Books{
			ID:     1,
			Title:  "Belajar Si Golang",
			Author: "Hacktiv9 Indonesia",
		},
	})

	testTable = append(testTable, testCase{
		name:      "unexpected error",
		wantError: true,
		inputID:   1,
		input: model.Books{
			Title:  "Belajar Si Golang",
			Author: "Hacktiv9 Indonesia",
		},
		expectedError: errors.New("unexpected error"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().UpdateBook(gomock.Any()).Return(nil)
		},
	})

	testTable = append(testTable, testCase{
		name:      "invalid Title length",
		wantError: true,
		inputID:   1,
		input: model.Books{
			Title:  "Golang si Golang", // case negative
			Author: "Hacktiv10 Indonesia",
		},
		expectedError: errors.New("invalid Title length"),
	})

	testTable = append(testTable, testCase{
		name:      "invalid Author",
		wantError: true,
		inputID:   1,
		input: model.Books{
			Title:  "Golang",
			Author: "anang",
		},
		expectedError: errors.New("invalid Author"),
	})

	testTable = append(testTable, testCase{
		name:      "record not found",
		wantError: true,
		inputID:   1,
		input: model.Books{
			Title:  "Belajar Si Golang",
			Author: "Hacktiv9 Indonesia",
		},
		expectedError: errors.New("record not found"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().UpdateBook(gomock.Any()).Return(nil)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			bookRepo := mocks.NewBookRepo(mockCtrl)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{
				repo: bookRepo,
			}

			err := service.DeleteBook(1)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
