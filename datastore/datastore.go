package datastore

import (
	"database/sql"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"simple-rest-api/model"
)

type student struct{}

func New() *student {
	return &student{}
}

func (s *student) GetByID(ctx *gofr.Context, id string) (*model.Student, error) {
	var resp model.Student

	err := ctx.DB().QueryRowContext(ctx, " SELECT id,name,age,branch FROM students where id=$1", id).
		Scan(&resp.ID, &resp.Name, &resp.Age, &resp.Branch)
	switch err {
	case sql.ErrNoRows:
		return &model.Student{}, errors.EntityNotFound{Entity: "student", ID: id}
	case nil:
		return &resp, nil
	default:
		return &model.Student{}, err
	}
}

func (s *student) Create(ctx *gofr.Context, student *model.Student) (*model.Student, error) {
	var resp model.Student

	err := ctx.DB().QueryRowContext(ctx, "INSERT INTO students (name, age, branch) VALUES($1,$2,$3)"+
		" RETURNING  id,name,age,branch", student.ID, student.Name, student.Age, student.Branch).Scan(
		&resp.ID, &resp.Name, &resp.Age, &resp.Branch)

	if err != nil {
		return &model.Student{}, errors.DB{Err: err}
	}

	return &resp, nil
}

func (s *student) Update(ctx *gofr.Context, student *model.Student) (*model.Student, error) {
	_, err := ctx.DB().ExecContext(ctx, "UPDATE students SET name=$1,age=$2,branch=$3 WHERE id=$4",
		student.Name, student.Age, student.Branch, student.ID)
	if err != nil {
		return &model.Student{}, errors.DB{Err: err}
	}

	return student, nil
}

func (s *student) Delete(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM students where id=$1", id)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}
