package models

import (
	"testing"

	th "github.com/ryanpback/raspberry_a_pi/testhelpers"

	"github.com/stretchr/testify/assert"
)

var userTC = th.BootstrapTestConfig()

var users = []struct {
	name     string
	userName string
	email    string
	password string
}{
	{
		name:     "test",
		userName: "tee",
		email:    "test1@test.com",
		password: "password",
	},
	{
		name:     "test",
		userName: "tee",
		email:    "test2@test.com",
		password: "password",
	},
	{
		name:     "test",
		userName: "tee",
		email:    "test3@test.com",
		password: "password",
	},
}

func createUsers() {
	for _, u := range users {
		hash, _ := hashAndSaltPassword([]byte(u.password))
		th.UserPersistToDB(u.name, u.userName, u.email, hash)
	}
}

func TestUserFindByIDUserNotExists(t *testing.T) {
	assert := assert.New(t)
	DBConn = userTC.DBConn
	defer th.TruncateUsers()

	createUsers()

	_, err := UserFindByID(999999)

	assert.NotNil(err)
}

func TestUserFindByIDUserExists(t *testing.T) {
	assert := assert.New(t)
	DBConn = userTC.DBConn
	defer th.TruncateUsers()

	createUsers()

	u, _ := UserFindByID(1)

	assert.Equal(int64(1), (*u).ID)
}

func TestUserCreate(t *testing.T) {
	assert := assert.New(t)
	DBConn = userTC.DBConn
	defer th.TruncateUsers()

	userData := map[string]interface{}{
		"name":     "Test",
		"userName": "tee",
		"email":    "test@test.com",
		"password": "password",
	}

	u, _ := UserCreate(userData)

	assert.Equal(userData["name"], u.Name)
}

func TestUserEdit(t *testing.T) {
	assert := assert.New(t)
	DBConn = userTC.DBConn
	defer th.TruncateUsers()

	createUsers()
	userData := map[string]interface{}{
		"name":     "Updated Name",
		"userName": "updateUserName",
		"email":    "updated@email.com",
		"password": "thispasswordshouldntupdatewhenediting",
	}

	u, _ := UserFindByID(1)
	userUpdatedAt := u.UpdatedAt
	updatedUser, _ := UserEdit(u, userData)

	assert.Equal(userData["name"], (*updatedUser).Name)
	assert.Equal(userData["userName"], (*updatedUser).UserName)
	assert.Equal(userData["email"], (*updatedUser).Email)
	assert.NotEqual(userData["password"], (*updatedUser).password)
	assert.NotEqual(userUpdatedAt, (*updatedUser).UpdatedAt)
}

func TestUserLoginFailedEmailDoesntMatch(t *testing.T) {
	assert := assert.New(t)
	DBConn = userTC.DBConn
	defer th.TruncateUsers()

	createUsers()

	userData := Payload{
		"email":    "randomfakeemail",
		"password": "password",
	}

	_, err := UserLogin(userData)

	assert.NotNil(err)
}

func TestUserLoginFailedPasswordDoesntMatch(t *testing.T) {
	assert := assert.New(t)
	DBConn = userTC.DBConn
	defer th.TruncateUsers()

	createUsers()
	u, _ := UserFindByID(1)
	email := u.Email

	userData := Payload{
		"email":    email,
		"password": "faketestpassword",
	}

	_, err := UserLogin(userData)

	assert.NotNil(err)
}

func TestUserLoginSuccessful(t *testing.T) {
	assert := assert.New(t)
	DBConn = userTC.DBConn
	defer th.TruncateUsers()

	createUsers()
	u, _ := UserFindByID(1)
	email := u.Email

	userData := Payload{
		"email":    email,
		"password": "password",
	}

	user, _ := UserLogin(userData)

	assert.Equal(u, user)
}

func TestFindByEmailUserNotExists(t *testing.T) {
	assert := assert.New(t)
	DBConn = userTC.DBConn
	defer th.TruncateUsers()

	createUsers()

	_, err := UserFindByEmail("thisisanon@existentemail.com")

	assert.NotNil(err)
}

func TestFindByEmailUserExists(t *testing.T) {
	assert := assert.New(t)
	DBConn = userTC.DBConn
	defer th.TruncateUsers()

	createUsers()

	u, _ := UserFindByEmail(users[0].email)

	assert.Equal(users[0].email, (*u).Email)
}
