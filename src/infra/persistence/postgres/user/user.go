package user

import (
	"fmt"

	dto "github.com/medivh13/dating_app/src/app/dto/user"
	"golang.org/x/crypto/bcrypt"

	"log"

	"github.com/jmoiron/sqlx"

	"errors"

	model "github.com/medivh13/dating_app/src/infra/models"
)

type UserRepository interface {
	Register(data *dto.RegisterReqDTO) (int64, error)
	Login(data *dto.LoginReqDTO) (*model.LoginModel, error)
}

const (
	Register = `INSERT INTO public.user (
		email, password, created_at) 
		values ($1, $2, now()) returning id`

	Login = `select id, email, password, is_premium, verified from public.user where email = $1`
)

var statement PreparedStatement

type PreparedStatement struct {
	register *sqlx.Stmt
	login    *sqlx.Stmt
}

type userRepo struct {
	Connection *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	repo := &userRepo{
		Connection: db,
	}
	InitPreparedStatement(repo)
	return repo
}

func (p *userRepo) Preparex(query string) *sqlx.Stmt {
	statement, err := p.Connection.Preparex(query)
	if err != nil {
		log.Fatalf("Failed to preparex query: %s. Error: %s", query, err.Error())
	}

	return statement
}

func InitPreparedStatement(m *userRepo) {
	statement = PreparedStatement{
		register: m.Preparex(Register),
		login:    m.Preparex(Login),
	}
}

func (p *userRepo) Register(data *dto.RegisterReqDTO) (int64, error) {
	var resultId int64

	pwd, err := hashPassword(data.Password)
	fmt.Println("1 sini")
	if err != nil {
		log.Println(err)
		return 0, err
	}

	result, err := statement.register.Query(data.Email, pwd)

	if err != nil {
		log.Println(err)
		return 0, err
	}

	defer result.Close() // pastikan untuk menutup hasil query setelah digunakan

	if result.Next() { // periksa apakah ada baris hasil

		err = result.Scan(&resultId)
		if err != nil {
			log.Println("Failed Query Register : ", err.Error())
			return 0, err
		}
	} else {
		log.Println("No rows returned from the query")
		return 0, err
	}

	return resultId, err
}

func (p *userRepo) Login(data *dto.LoginReqDTO) (*model.LoginModel, error) {
	var resultData []*model.LoginModel

	err := statement.login.Select(&resultData, data.Email)

	if err != nil {
		return nil, err
	}

	if len(resultData) < 1 {
		return nil, errors.New("no rows returned from the query")
	}

	hashPwd := resultData[0].Password

	err = verifyPassword(hashPwd, data.Password)

	if err != nil {
		return nil, err
	}

	return resultData[0], nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func verifyPassword(hashedPassword, inputPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
}
