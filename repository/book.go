package repository

import (
	"sesi_8/model"
)

// clean architectures -> handler->service->repo

// interface employee
type BookRepo interface {
	//GetBooks() ([]model.Books, error)
	CreateBook(in model.Books) (res model.Books, err error)
	GetBookByID(id int64) (res model.Books, err error)
	UpdateBook(in model.Books) (res model.Books, err error)
	DeleteBook(id int64) (err error)
}

func (r Repo) GetBooks() ([]model.Books, error) {
	var books []model.Books
	result := r.gorm.Find(&books)
	if result.Error != nil {
		return books, result.Error
	}
	return books, nil
}

func (r Repo) CreateBook(in model.Books) (res model.Books, err error) {
	err = r.gorm.Create(&in).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetBookByID(id int64) (res model.Books, err error) {
	err = r.gorm.First(&res, id).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) UpdateBook(in model.Books) (res model.Books, err error) {
	err = r.gorm.Model(&res).Where("id = ?", in.ID).Updates(model.Books{
		Title:  in.Title,
		Author: in.Author,
	}).Scan(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) DeleteBook(id int64) (err error) {
	book := model.Books{}
	err = r.gorm.Where("id = ?", id).Delete(&book).Error
	if err != nil {
		return err
	}

	return nil
}

// func (r Repo) CreateEmployee(in model.Employee) (res model.Employee, err error) {
// 	err = r.db.QueryRow(
// 		query.AddEmployee,
// 		in.Fullname,
// 		in.Email,
// 		in.Age,
// 		in.Division,
// 	).Scan(
// 		&res.ID,
// 		&res.Fullname,
// 		&res.Email,
// 		&res.Age,
// 		&res.Division,
// 	)
// 	if err != nil {
// 		return res, err
// 	}

// 	return res, nil
// }
