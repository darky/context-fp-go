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

func TestGenerics1(t *testing.T) {
	type Context struct {
		a int
	}
	type ContextA struct {
		a int
	}
	fn1 := func(ctx *Context) *ContextA {
		return &ContextA{a: ctx.a}
	}
	fn := Cfp1(
		&fn1,
		func(dep1 *ContextA) int {
			return dep1.a
		},
	)
	assert.Equal(t,
		fn(&Context{
			a: 1,
		}),
		1,
	)
}

func TestGenerics2(t *testing.T) {
	type Context struct {
		a int
		b int
	}
	type ContextA struct {
		a int
	}
	type ContextB struct {
		b int
	}
	fn1 := func(ctx *Context) *ContextA {
		return &ContextA{a: ctx.a}
	}
	fn2 := func(ctx *Context) *ContextB {
		return &ContextB{b: ctx.b}
	}
	fn := Cfp2(
		&fn1,
		&fn2,
		func(dep1 *ContextA, dep2 *ContextB) int {
			return dep1.a + dep2.b
		},
	)
	assert.Equal(t,
		fn(&Context{
			a: 1,
			b: 1,
		}),
		2,
	)
}

func TestGenerics3(t *testing.T) {
	type Context struct {
		a int
		b int
		c int
	}
	type ContextA struct {
		a int
	}
	type ContextB struct {
		b int
	}
	type ContextC struct {
		c int
	}
	fn1 := func(ctx *Context) *ContextA {
		return &ContextA{a: ctx.a}
	}
	fn2 := func(ctx *Context) *ContextB {
		return &ContextB{b: ctx.b}
	}
	fn3 := func(ctx *Context) *ContextC {
		return &ContextC{c: ctx.c}
	}
	fn := Cfp3(
		&fn1,
		&fn2,
		&fn3,
		func(dep1 *ContextA, dep2 *ContextB, dep3 *ContextC) int {
			return dep1.a + dep2.b + dep3.c
		},
	)
	assert.Equal(t,
		fn(&Context{
			a: 1,
			b: 1,
			c: 1,
		}),
		3,
	)
}

func TestGenerics4(t *testing.T) {
	type Context struct {
		a int
		b int
		c int
		d int
	}
	type ContextA struct {
		a int
	}
	type ContextB struct {
		b int
	}
	type ContextC struct {
		c int
	}
	type ContextD struct {
		d int
	}
	fn1 := func(ctx *Context) *ContextA {
		return &ContextA{a: ctx.a}
	}
	fn2 := func(ctx *Context) *ContextB {
		return &ContextB{b: ctx.b}
	}
	fn3 := func(ctx *Context) *ContextC {
		return &ContextC{c: ctx.c}
	}
	fn4 := func(ctx *Context) *ContextD {
		return &ContextD{d: ctx.d}
	}
	fn := Cfp4(
		&fn1,
		&fn2,
		&fn3,
		&fn4,
		func(dep1 *ContextA, dep2 *ContextB, dep3 *ContextC, dep4 *ContextD) int {
			return dep1.a + dep2.b + dep3.c + dep4.d
		},
	)
	assert.Equal(t,
		fn(&Context{
			a: 1,
			b: 1,
			c: 1,
			d: 1,
		}),
		4,
	)
}

func TestGenerics5(t *testing.T) {
	type Context struct {
		a int
		b int
		c int
		d int
		e int
	}
	type ContextA struct {
		a int
	}
	type ContextB struct {
		b int
	}
	type ContextC struct {
		c int
	}
	type ContextD struct {
		d int
	}
	type ContextE struct {
		e int
	}
	fn1 := func(ctx *Context) *ContextA {
		return &ContextA{a: ctx.a}
	}
	fn2 := func(ctx *Context) *ContextB {
		return &ContextB{b: ctx.b}
	}
	fn3 := func(ctx *Context) *ContextC {
		return &ContextC{c: ctx.c}
	}
	fn4 := func(ctx *Context) *ContextD {
		return &ContextD{d: ctx.d}
	}
	fn5 := func(ctx *Context) *ContextE {
		return &ContextE{e: ctx.e}
	}
	fn := Cfp5(
		&fn1,
		&fn2,
		&fn3,
		&fn4,
		&fn5,
		func(dep1 *ContextA, dep2 *ContextB, dep3 *ContextC, dep4 *ContextD, dep5 *ContextE) int {
			return dep1.a + dep2.b + dep3.c + dep4.d + dep5.e
		},
	)
	assert.Equal(t,
		fn(&Context{
			a: 1,
			b: 1,
			c: 1,
			d: 1,
			e: 1,
		}),
		5,
	)
}

func TestGenerics6(t *testing.T) {
	type Context struct {
		a int
		b int
		c int
		d int
		e int
		f int
	}
	type ContextA struct {
		a int
	}
	type ContextB struct {
		b int
	}
	type ContextC struct {
		c int
	}
	type ContextD struct {
		d int
	}
	type ContextE struct {
		e int
	}
	type ContextF struct {
		f int
	}
	fn1 := func(ctx *Context) *ContextA {
		return &ContextA{a: ctx.a}
	}
	fn2 := func(ctx *Context) *ContextB {
		return &ContextB{b: ctx.b}
	}
	fn3 := func(ctx *Context) *ContextC {
		return &ContextC{c: ctx.c}
	}
	fn4 := func(ctx *Context) *ContextD {
		return &ContextD{d: ctx.d}
	}
	fn5 := func(ctx *Context) *ContextE {
		return &ContextE{e: ctx.e}
	}
	fn6 := func(ctx *Context) *ContextF {
		return &ContextF{f: ctx.f}
	}
	fn := Cfp6(
		&fn1,
		&fn2,
		&fn3,
		&fn4,
		&fn5,
		&fn6,
		func(
			dep1 *ContextA,
			dep2 *ContextB,
			dep3 *ContextC,
			dep4 *ContextD,
			dep5 *ContextE,
			dep6 *ContextF,
		) int {
			return dep1.a + dep2.b + dep3.c + dep4.d + dep5.e + dep6.f
		},
	)
	assert.Equal(t,
		fn(&Context{
			a: 1,
			b: 1,
			c: 1,
			d: 1,
			e: 1,
			f: 1,
		}),
		6,
	)
}

func TestGenerics7(t *testing.T) {
	type Context struct {
		a int
		b int
		c int
		d int
		e int
		f int
		g int
	}
	type ContextA struct {
		a int
	}
	type ContextB struct {
		b int
	}
	type ContextC struct {
		c int
	}
	type ContextD struct {
		d int
	}
	type ContextE struct {
		e int
	}
	type ContextF struct {
		f int
	}
	type ContextG struct {
		g int
	}
	fn1 := func(ctx *Context) *ContextA {
		return &ContextA{a: ctx.a}
	}
	fn2 := func(ctx *Context) *ContextB {
		return &ContextB{b: ctx.b}
	}
	fn3 := func(ctx *Context) *ContextC {
		return &ContextC{c: ctx.c}
	}
	fn4 := func(ctx *Context) *ContextD {
		return &ContextD{d: ctx.d}
	}
	fn5 := func(ctx *Context) *ContextE {
		return &ContextE{e: ctx.e}
	}
	fn6 := func(ctx *Context) *ContextF {
		return &ContextF{f: ctx.f}
	}
	fn7 := func(ctx *Context) *ContextG {
		return &ContextG{g: ctx.g}
	}
	fn := Cfp7(
		&fn1,
		&fn2,
		&fn3,
		&fn4,
		&fn5,
		&fn6,
		&fn7,
		func(
			dep1 *ContextA,
			dep2 *ContextB,
			dep3 *ContextC,
			dep4 *ContextD,
			dep5 *ContextE,
			dep6 *ContextF,
			dep7 *ContextG,
		) int {
			return dep1.a + dep2.b + dep3.c + dep4.d + dep5.e + dep6.f + dep7.g
		},
	)
	assert.Equal(t,
		fn(&Context{
			a: 1,
			b: 1,
			c: 1,
			d: 1,
			e: 1,
			f: 1,
			g: 1,
		}),
		7,
	)
}

func TestGenerics8(t *testing.T) {
	type Context struct {
		a int
		b int
		c int
		d int
		e int
		f int
		g int
		h int
	}
	type ContextA struct {
		a int
	}
	type ContextB struct {
		b int
	}
	type ContextC struct {
		c int
	}
	type ContextD struct {
		d int
	}
	type ContextE struct {
		e int
	}
	type ContextF struct {
		f int
	}
	type ContextG struct {
		g int
	}
	type ContextH struct {
		h int
	}
	fn1 := func(ctx *Context) *ContextA {
		return &ContextA{a: ctx.a}
	}
	fn2 := func(ctx *Context) *ContextB {
		return &ContextB{b: ctx.b}
	}
	fn3 := func(ctx *Context) *ContextC {
		return &ContextC{c: ctx.c}
	}
	fn4 := func(ctx *Context) *ContextD {
		return &ContextD{d: ctx.d}
	}
	fn5 := func(ctx *Context) *ContextE {
		return &ContextE{e: ctx.e}
	}
	fn6 := func(ctx *Context) *ContextF {
		return &ContextF{f: ctx.f}
	}
	fn7 := func(ctx *Context) *ContextG {
		return &ContextG{g: ctx.g}
	}
	fn8 := func(ctx *Context) *ContextH {
		return &ContextH{h: ctx.h}
	}
	fn := Cfp8(
		&fn1,
		&fn2,
		&fn3,
		&fn4,
		&fn5,
		&fn6,
		&fn7,
		&fn8,
		func(
			dep1 *ContextA,
			dep2 *ContextB,
			dep3 *ContextC,
			dep4 *ContextD,
			dep5 *ContextE,
			dep6 *ContextF,
			dep7 *ContextG,
			dep8 *ContextH,
		) int {
			return dep1.a + dep2.b + dep3.c + dep4.d + dep5.e + dep6.f + dep7.g + dep8.h
		},
	)
	assert.Equal(t,
		fn(&Context{
			a: 1,
			b: 1,
			c: 1,
			d: 1,
			e: 1,
			f: 1,
			g: 1,
			h: 1,
		}),
		8,
	)
}

func TestGenerics9(t *testing.T) {
	type Context struct {
		a int
		b int
		c int
		d int
		e int
		f int
		g int
		h int
		i int
	}
	type ContextA struct {
		a int
	}
	type ContextB struct {
		b int
	}
	type ContextC struct {
		c int
	}
	type ContextD struct {
		d int
	}
	type ContextE struct {
		e int
	}
	type ContextF struct {
		f int
	}
	type ContextG struct {
		g int
	}
	type ContextH struct {
		h int
	}
	type ContextI struct {
		i int
	}
	fn1 := func(ctx *Context) *ContextA {
		return &ContextA{a: ctx.a}
	}
	fn2 := func(ctx *Context) *ContextB {
		return &ContextB{b: ctx.b}
	}
	fn3 := func(ctx *Context) *ContextC {
		return &ContextC{c: ctx.c}
	}
	fn4 := func(ctx *Context) *ContextD {
		return &ContextD{d: ctx.d}
	}
	fn5 := func(ctx *Context) *ContextE {
		return &ContextE{e: ctx.e}
	}
	fn6 := func(ctx *Context) *ContextF {
		return &ContextF{f: ctx.f}
	}
	fn7 := func(ctx *Context) *ContextG {
		return &ContextG{g: ctx.g}
	}
	fn8 := func(ctx *Context) *ContextH {
		return &ContextH{h: ctx.h}
	}
	fn9 := func(ctx *Context) *ContextI {
		return &ContextI{i: ctx.i}
	}
	fn := Cfp9(
		&fn1,
		&fn2,
		&fn3,
		&fn4,
		&fn5,
		&fn6,
		&fn7,
		&fn8,
		&fn9,
		func(
			dep1 *ContextA,
			dep2 *ContextB,
			dep3 *ContextC,
			dep4 *ContextD,
			dep5 *ContextE,
			dep6 *ContextF,
			dep7 *ContextG,
			dep8 *ContextH,
			dep9 *ContextI,
		) int {
			return dep1.a + dep2.b + dep3.c + dep4.d + dep5.e + dep6.f + dep7.g + dep8.h + dep9.i
		},
	)
	assert.Equal(t,
		fn(&Context{
			a: 1,
			b: 1,
			c: 1,
			d: 1,
			e: 1,
			f: 1,
			g: 1,
			h: 1,
			i: 1,
		}),
		9,
	)
}

func TestSfpBasic(t *testing.T) {
	type N struct {
		n int
	}
	numbers := Sfp(
		func(ns []int, payload *N) []int {
			return append(ns, payload.n)
		},
		[]int{},
	)
	numbers(&N{1})
	numbers(&N{2})
	numbers(&N{3})
	assert.Equal(t,
		len(numbers(nil)),
		3,
	)
	assert.Equal(t,
		numbers(nil),
		[]int{1, 2, 3},
	)
}

func TestSfpWithCfp(t *testing.T) {
	type N struct {
		n int
	}
	type Context struct {
		incNumber int
	}
	numbers := func(ctx *Context) func(payload *N) []int {
		return Sfp(
			func(ns []int, payload *N) []int {
				return append(ns, payload.n+ctx.incNumber)
			},
			[]int{},
		)
	}
	addNumber1 := Cfp1(&numbers, func(ns func(payload *N) []int) []int {
		return ns(&N{1})
	})
	addNumber2 := Cfp1(&numbers, func(ns func(payload *N) []int) []int {
		return ns(&N{2})
	})
	addNumber3 := Cfp1(&numbers, func(ns func(payload *N) []int) []int {
		return ns(&N{3})
	})
	numbersToString := Cfp4(
		&numbers,
		&addNumber1,
		&addNumber2,
		&addNumber3,
		func(ns func(payload *N) []int, n1 []int, n2 []int, n3 []int) string {
			return strings.Join(
				lo.Map(ns(nil), func(n int, i int) string { return fmt.Sprintf("%v", n) }),
				",",
			)
		},
	)
	assert.Equal(t,
		numbersToString(&Context{incNumber: 1}),
		"2,3,4",
	)
}
