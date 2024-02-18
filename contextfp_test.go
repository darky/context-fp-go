package contextfp

import (
	"fmt"
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	type Context struct {
		numbers []int
	}
	positiveNumbers := func(ctx *Context) []int {
		return lo.Filter(ctx.numbers, func(n int, i int) bool { return n > 0 })
	}
	numbersPrefix := func(ctx *Context) string {
		return "Here is numbers:"
	}
	positiveNumbersAsString := Cfp2(
		&positiveNumbers,
		&numbersPrefix,
		func(ns []int, prefix string) string {
			ns2str := lo.Map(ns, func(n int, i int) string { return fmt.Sprintf("%v", n) })
			return prefix + " " + strings.Join(ns2str, ",")
		},
	)
	assert.Equal(t,
		positiveNumbersAsString(&Context{
			numbers: []int{-1, -5, 7, 0, 4},
		}),
		"Here is numbers: 7,4",
	)
}

func TestDI(t *testing.T) {
	type User struct {
		name string
	}
	type Context struct {
		fetchUser func() User
	}
	fetchUserFromDB := func() User {
		panic("should not be called")
	}
	fetchUser := func(ctx *Context) User {
		if ctx.fetchUser == nil {
			return fetchUserFromDB()
		}
		return ctx.fetchUser()
	}
	helloWorldUser := Cfp1(&fetchUser, func(user User) string {
		return "Hello world, " + user.name
	})
	assert.Equal(
		t,
		helloWorldUser(&Context{fetchUser: func() User { return User{name: "Vasya"} }}),
		"Hello world, Vasya",
	)
}

func TestCache(t *testing.T) {
	type Context struct {
		numbers []int
	}
	called := 0
	positiveNumbers := func(ctx *Context) []int {
		called++
		return lo.Filter(ctx.numbers, func(n int, i int) bool { return n > 0 })
	}
	positiveNumbersLength := Cfp1(
		&positiveNumbers,
		func(ns []int) int {
			return len(ns)
		},
	)
	positiveNumbersAsString := Cfp2(
		&positiveNumbers,
		&positiveNumbersLength,
		func(ns []int, length int) string {
			ns2str := lo.Map(ns, func(n int, i int) string { return fmt.Sprintf("%v", n) })
			return strings.Join(ns2str, ",") + "; length - " + fmt.Sprintf("%v", length)
		},
	)
	assert.Equal(t,
		positiveNumbersAsString(&Context{
			numbers: []int{-1, -5, 7, 0, 4},
		}),
		"7,4; length - 2",
	)
	assert.Equal(t, called, 1)
}
