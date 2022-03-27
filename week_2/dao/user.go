package dao

import (
	"database/sql"
	"errors"
)

func DaoDemo(query string) (interface{}, error) {
	// do something
	err := mockError()

	if err == sql.ErrNoRows {
		// handle error
		return nil, err
	} else if err != nil {
		return nil, err
	}
	return 1, nil
}

func mockError() error {
	return errors.New("this is a error")
}