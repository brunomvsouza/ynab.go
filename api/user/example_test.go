package user_test

import (
	"fmt"
	"reflect"

	"go.bmvs.io/ynab"
)

func ExampleService_GetUser() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	user, _ := c.User().GetUser()
	fmt.Println(reflect.TypeOf(user))

	// Output: *user.User
}
