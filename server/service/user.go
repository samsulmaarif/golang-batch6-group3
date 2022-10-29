package service

import (
	"database/sql"
	"golang-batch6-group3/helper"
	"golang-batch6-group3/server/params"
	"golang-batch6-group3/server/repository"
	"golang-batch6-group3/server/view"
	"log"
	"time"

	"github.com/google/uuid"
)

type UserServices struct {
	repo repository.UserRepo
}

func NewUserServices(repo repository.UserRepo) *UserServices {
	return &UserServices{
		repo: repo,
	}
}

func (u *UserServices) GetUsers() *view.Response {
	users, err := u.repo.GetUsers()
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer("GET_ALL_USERS_FAIL", err.Error())
	}

	return view.SuccessFindAll(view.NewUserFindAllResponse(users))
}

func (u *UserServices) CreateUser(req *params.UserCreate) *view.Response {
	user := req.ParseToModel()

	user.Id = uuid.NewString()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Role = "member"

	hash, err := helper.GeneratePassword(user.Password)
	if err != nil {
		log.Printf("get error when try to generate password %v\n", "")
		return view.ErrInternalServer("CREATE_USERS_FAIL", err.Error())
	}

	user.Password = hash

	err = u.repo.Register(user)
	if err != nil {
		log.Printf("get error register user with error %v\n", "")
		return view.ErrInternalServer("CREATE_USERS_FAIL", err.Error())
	}

	return view.SuccessCreated(user)
}

func (u *UserServices) UpdateProfile(email string, req *params.UserCreate) *view.Response {
	user := req.ParseToModel()

	err := u.repo.UpdateUserByEmail(email, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer("UPDATE_USER_FAIL", err.Error())
	}

	return view.SuccessCreated("UPDATE_USER_SUCCESS")
}

func (u *UserServices) Login(req *params.UserLogin) *view.Response {
	user, err := u.repo.FindUserByEmail(req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer("LOGIN_FAIL", err.Error())
	}

	err = helper.ValidatePassword(user.Password, req.Password)
	if err != nil {
		return view.ErrUnauthorized()
	}

	token := helper.Token{
		UserId: user.Id,
		Email:  user.Email,
	}

	tokString, err := helper.CreateToken(&token)
	if err != nil {
		return view.ErrInternalServer("LOGIN_FAIL", err.Error())
	}

	return view.SuccessLogin("LOGIN_SUCCESS", tokString)
}

func (u *UserServices) DeleteUserById(id string) *view.Response {
	user, err := u.repo.FindUserById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer("FIND_USER_FAIL", err.Error())
	}

	if user != nil {
		delete, err := u.repo.DeleteUserById(id)
		if err != nil {
			if err == sql.ErrNoRows {
				return view.ErrNotFound()
			}
			return view.ErrInternalServer("DELETE_USER_FAIL", err.Error())
		}

		return view.SuccessFindAll(delete)
	}
	return nil
}

func (u *UserServices) FindUserByEmail(email string) *view.Response {
	user, err := u.repo.FindUserByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer("FIND_USER_FAIL", err.Error())
	}
	return view.SuccessFindAll(user)
}

func (u *UserServices) FindUserById(id string) *view.Response {
	user, err := u.repo.FindUserById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer("FIND_USER_FAIL", err.Error())
	}
	return view.SuccessFindAll(user)
}
