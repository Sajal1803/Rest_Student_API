package datastore

import (
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"simple-rest-api/model"
)

type student struct{}

func New() *student {
	return &student{}
}

func (s student) GetByID(ctx *gofr.Context) ([]model.Student, error) {
	rows, err := ctx.DB().QueryContext(ctx, " SELECT id,name,age,branch FROM students")

	if err != nil {
		return nil, errors.DB{Err: err}
	}

	defer rows.Close()
	students := make([]model.Student, 0)
	for rows.Next() {
		var c model.Student
		err = rows.Scan(&c.ID, &c.Name, &c.Age, &c.Branch)
		if err != nil {
			return nil, errors.DB{Err: err}
		}
		students = append(students, c)
	}
	err = rows.Err()

	if err != nil {
		return nil, errors.DB{Err: err}
	}
	return students, nil
}

func (s *student) Create(ctx *gofr.Context, student model.Student) (model.Student, error) {
	var resp model.Student
	queryInsert := "INSERT INTO students (name,age,branch) VALUES (?,?,?)"

	result, err := ctx.DB().ExecContext(ctx, queryInsert, student.Name, student.Age, student.Branch)

	if err != nil {
		return model.Student{}, errors.DB{Err: err}
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return model.Student{}, errors.DB{Err: err}
	}

	querySelect := "SELECT id,name,age,branch FROM students WHERE id = ?"
	err = ctx.DB().QueryRowContext(ctx, querySelect, lastInsertID).Scan(&resp.ID, &resp.Name, &resp.Age, &resp.Branch)

	if err != nil {
		return model.Student{}, errors.DB{Err: err}
	}

	return resp, nil
}

func (s *student) Update(ctx *gofr.Context, student model.Student) (model.Student, error) {
	_, err := ctx.DB().ExecContext(ctx, "UPDATE students SET name=?,age=?,branch=? WHERE id=?",
		student.Name, student.Age, student.Branch, student.ID)
	if err != nil {
		return model.Student{}, errors.DB{Err: err}
	}

	return student, nil
}

func (s *student) Delete(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM students where id=?", id)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}
