package repository

import (
	"sesi_8/model"
)

// clean architectures -> handler->service->repo

// interface employee
type EmployeeRepo interface {
	CreateEmployee(in model.Employee) (res model.Employee, err error)
	GetEmployeeByID(id int64) (res model.Employee, err error)
	UpdateEmployee(in model.Employee) (res model.Employee, err error)
	DeleteEmployee(id int64) (err error)
}

func (r Repo) CreateEmployee(in model.Employee) (res model.Employee, err error) {
	err = r.gorm.Create(&in).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetEmployeeByID(id int64) (res model.Employee, err error) {
	err = r.gorm.First(&res, id).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) UpdateEmployee(in model.Employee) (res model.Employee, err error) {
	err = r.gorm.Model(&res).Where("id = ?", in.ID).Updates(model.Employee{
		Fullname: in.Fullname,
		Email:    in.Email,
		Age:      in.Age,
		Division: in.Division,
	}).Scan(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) DeleteEmployee(id int64) (err error) {
	employee := model.Employee{}
	err = r.gorm.Where("id = ?", id).Delete(&employee).Error
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
