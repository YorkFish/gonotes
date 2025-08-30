package middlewares

import (
	"demo/pojo"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func UserPasd(field validator.FieldLevel) bool {
	match, _ := regexp.MatchString(`^[a-zA-z\d]{4,20}$`, field.Field().String())
	return match
}

func UserList(field validator.StructLevel) {
	users := field.Current().Interface().(pojo.Users)
	if users.UserListSize == len(users.UserList) {

	} else {
		field.ReportError(users.UserListSize, "Sizeof user list", "UserListSize", "UserListSizeMustEqualsUserList", "")
	}
}
