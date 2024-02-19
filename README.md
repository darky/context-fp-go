# context-fp-go

![logo](logo.png)

Functional programming context for Golang <br/>

## Motivation

if you want Dependency injection for Golang using plain functions, you are in the right place.

## Features

- 🤏 Tiny size and most of it is Golang generics
- 💉 Dependency injection without magic, only functions
- 🤌 Functions cached during workflow, no excess cost of CPU
- 💡 Smart type inference, only think about Context type and rest will be checked
- ♻️ Unit tests friendly, feel free to pass mocked function in the Context
- 📦 Tiny Redux like state manager also attached

## How to

#### Basic example

```golang
import (
  "github.com/samber/lo"
  "github.com/darky/context-fp-go"
)

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
```

#### Calculations cached example

```golang
import (
  "github.com/samber/lo"
  "github.com/darky/context-fp-go"
)

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
```

#### Dependency injection example

```golang
import (
  "github.com/darky/context-fp-go"
)

type User struct {
  name string
}
type Context struct {
  fetchUser func() User
}
fetchUserFromDB := func() User {
  // some production implementation
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
```

#### State manager example

```golang
import (
  "github.com/darky/context-fp-go"
)

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
```

## See also

- [context-fp](https://github.com/darky/context-fp) - Functional programming context for TypeScript
