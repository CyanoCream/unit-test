package service

// usecase

import (
	"errors"
	"sesi_8/model"
)

type BookService interface {
	//CreateEmployee(in model.Employee) (res model.Employee, err error)
	//GetEmployeeByID(id int64) (res model.Employee, err error)
	//UpdateEmployee(in model.Employee) (res model.Employee, err error)
	//DeleteEmployee(id int64) (err error)
	//GetBooks() ([]model.Books, error)
	CreateBook(in model.Books) (res model.Books, err error)
	GetBookByID(id int64) (res model.Books, err error)
	UpdateBook(in model.Books) (res model.Books, err error)
	DeleteBook(id int64) (err error)
}

//	func (s *Service) GetBooks() ([]model.Books, error) {
//		book, err := s.repo.GetBooks()
//		if err != nil {
//			return book, err
//		}
//
//		return book, nil
//	}
func (s *Service) CreateBook(in model.Books) (res model.Books, err error) {
	if len(in.Author) < 5 {
		return res, errors.New("invalid fullname length")
	}

	//if in.Title != "Tech" {
	//	return res, errors.New("invalid division")
	//}

	res, err = s.repo.CreateBook(in)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) GetBookByID(id int64) (res model.Books, err error) {
	res, err = s.repo.GetBookByID(id)
	if err != nil {
		return res, err
	}
	// res.Fullname = "Rizqi salah"
	return res, nil
	// return s.repo.GetEmployeeByID(id)
}

func (s *Service) UpdateBook(in model.Books) (res model.Books, err error) {
	return s.repo.UpdateBook(in)
}

func (s *Service) DeleteBook(id int64) (err error) {
	return s.repo.DeleteBook(id)
}
