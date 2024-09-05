package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"user-service/internal/entity/courier"
	"user-service/internal/entity/user"
	"user-service/internal/infrastructura/repository"

	"github.com/Masterminds/squirrel"
)

type UserPostgres struct {
	db *sql.DB
}

func NewUserPostgres(db *sql.DB) repository.UserRepository {
	return &UserPostgres{db: db}
}

func (u *UserPostgres) AddUser(req user.RegisterReq) (*user.UserRes, error) {
	sql, args, err := squirrel.Insert("users").
		Columns("firstname", "lastname", "phone", "email", "password").
		Values(req.FirstName, req.LastName, req.Phone, req.Email, req.Password).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("error generating SQL for adduser: %v", err)
	}

	row := u.db.QueryRow(sql, args...)
	var res user.UserRes

	if err := row.Scan(&res.ID); err != nil {
		return nil, fmt.Errorf("error querying SQL for adduser: %v", err)
	}

	return &res, nil
}

func (u *UserPostgres) GetbyEmail(req user.LoginReq) (string, error) {
	sql, args, err := squirrel.
		Select("password").
		From("users").
		Where(squirrel.Eq{"email": req.Email}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return "", fmt.Errorf("email not found")
	}
	var password string
	row := u.db.QueryRow(sql, args...)
	if err := row.Scan(&password); err != nil {
		return "", fmt.Errorf("scan error getbyEMail: %v", err)
	}

	return password, nil
}

func (u *UserPostgres) Getbyiduser(req user.UserRes) (*user.User, error) {
	var res user.User
	sql, args, err := squirrel.
		Select("*").
		From("users").
		Where(squirrel.Eq{"id": req.ID}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	row := u.db.QueryRow(sql, args...)
	if err := row.Scan(&res.ID, &res.FirstName, &res.LastName, &res.Phone, &res.Email, &res.Password); err != nil {
		return nil, fmt.Errorf("scan error: %v", err)
	}

	return &res, nil
}

func (u *UserPostgres) UpdateUser(req user.User) error {
	sql, args, err := squirrel.Update("users").
		Set("firstname", req.FirstName).
		Set("lastname", req.LastName).
		Set("phone", req.Phone).
		Set("email", req.Email).
		Set("password", req.Password).
		Where(squirrel.Eq{"id": req.ID}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return fmt.Errorf("error generating sql for UpdateUser: %v", err)
	}

	_, err = u.db.Exec(sql, args...)
	if err != nil {
		return fmt.Errorf("error executing sql for UpdateUser: %v", err)
	}
	return nil
}

func (u *UserPostgres) DeleteUser(req user.UserRes) error {
	var exists bool
	err := u.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE userid = $1)", req.ID).Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking user existence: %v", err)
	}

	if !exists {
		return fmt.Errorf("user with ID %v does not exist", req.ID)
	}

	sql, args, err := squirrel.Delete("users").
		Where(squirrel.Eq{"id": req.ID}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return fmt.Errorf("error generating SQL for deleteUser: %v", err)
	}
	_, err = u.db.Exec(sql, args...)
	if err != nil {
		return fmt.Errorf("error executing sql for deleteuser:%v", err)
	}

	return nil
}

func (u *UserPostgres) AddCourier(courier courier.Courier) error {
	query, args, err := squirrel.Insert("couriers").
		Columns("name", "email", "password_hash", "phone_number", "delivery_area").
		Values(courier.Name, courier.Email, courier.PasswordHash, courier.PhoneNumber, courier.DeliveryArea).
		ToSql()
	if err != nil {
		log.Println("Error building query:", err)
		return err
	}

	_, err = u.db.Exec(query, args...)
	if err != nil {
		log.Println("Error inserting courier:", err)
		return err
	}
	return nil
}

func (u *UserPostgres) GetbyCourierEmail(email string) (string, error) {
	sql, args, err := squirrel.
		Select("password").
		From("couriers").
		Where(squirrel.Eq{"email": email}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return "", fmt.Errorf("email not found")
	}
	var password string
	row := u.db.QueryRow(sql, args...)
	if err := row.Scan(&password); err != nil {
		return "", fmt.Errorf("scan error getbycourierEMail: %v", err)
	}

	return password, nil
}

func (u *UserPostgres) GetCourierByID(id int) (courier.Courier, error) {
	var courier courier.Courier
	query, args, err := squirrel.Select("id", "name", "email", "password_hash", "phone_number", "delivery_area", "created_at").
		From("couriers").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		log.Println("Error building query:", err)
		return courier, err
	}

	row := u.db.QueryRow(query, args...)
	err = row.Scan(&courier.ID, &courier.Name, &courier.Email, &courier.PasswordHash, &courier.PhoneNumber, &courier.DeliveryArea, &courier.CreatedAt)
	if err != nil {
		log.Println("Error finding courier:", err)
		return courier, err
	}
	return courier, nil
}

func (u *UserPostgres) UpdateCourier(courier courier.Courier) error{
	query, args, err := squirrel.Update("couriers").
        Set("phone_number", courier.PhoneNumber).
        Set("delivery_area", courier.DeliveryArea).
        Where(squirrel.Eq{"id": courier.ID}).
        ToSql()
    if err != nil {
        log.Println("Error building query:", err)
        return err
    }
    
    _, err =u.db.Exec(query, args...)
    if err != nil {
        log.Println("Error updating courier:", err)
        return err
    }
    return nil
}


func (u *UserPostgres) DeleteCourier(id int) error{
	query, args, err := squirrel.Delete("couriers").
        Where(squirrel.Eq{"id": id}).
        ToSql()
    if err != nil {
        log.Println("Error building query:", err)
        return err
    }
    
    _, err = u.db.Exec(query, args...)
    if err != nil {
        log.Println("Error deleting courier:", err)
        return err
    }
    return nil
}
