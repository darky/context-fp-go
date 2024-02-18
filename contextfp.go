package contextfp

var cache = map[interface{}]map[interface{}]interface{}{}

func Cfp1[C, T1, R any](
	dep1 *func(ctx *C) T1,
	fn func(x1 T1) R,
) func(ctx *C) R {
	return func(ctx *C) R {
		return cfpBase(ctx, func() R {
			return fn(
				getCacheOrCall(ctx, dep1),
			)
		})
	}
}

func Cfp2[C, T1, T2, R any](
	dep1 *func(ctx *C) T1,
	dep2 *func(ctx *C) T2,
	fn func(x1 T1, x2 T2) R,
) func(ctx *C) R {
	return func(ctx *C) R {
		return cfpBase(ctx, func() R {
			return fn(
				getCacheOrCall(ctx, dep1),
				getCacheOrCall(ctx, dep2),
			)
		})
	}
}

func Cfp3[C, T1, T2, T3, R any](
	dep1 *func(ctx *C) T1,
	dep2 *func(ctx *C) T2,
	dep3 *func(ctx *C) T3,
	fn func(x1 T1, x2 T2, x3 T3) R,
) func(ctx *C) R {
	return func(ctx *C) R {
		return cfpBase(ctx, func() R {
			return fn(
				getCacheOrCall(ctx, dep1),
				getCacheOrCall(ctx, dep2),
				getCacheOrCall(ctx, dep3),
			)
		})
	}
}

func Cfp4[C, T1, T2, T3, T4, R any](
	dep1 *func(ctx *C) T1,
	dep2 *func(ctx *C) T2,
	dep3 *func(ctx *C) T3,
	dep4 *func(ctx *C) T4,
	fn func(x1 T1, x2 T2, x3 T3, x4 T4) R,
) func(ctx *C) R {
	return func(ctx *C) R {
		return cfpBase(ctx, func() R {
			return fn(
				getCacheOrCall(ctx, dep1),
				getCacheOrCall(ctx, dep2),
				getCacheOrCall(ctx, dep3),
				getCacheOrCall(ctx, dep4),
			)
		})
	}
}

func Cfp5[C, T1, T2, T3, T4, T5, R any](
	dep1 *func(ctx *C) T1,
	dep2 *func(ctx *C) T2,
	dep3 *func(ctx *C) T3,
	dep4 *func(ctx *C) T4,
	dep5 *func(ctx *C) T5,
	fn func(x1 T1, x2 T2, x3 T3, x4 T4, x5 T5) R,
) func(ctx *C) R {
	return func(ctx *C) R {
		return cfpBase(ctx, func() R {
			return fn(
				getCacheOrCall(ctx, dep1),
				getCacheOrCall(ctx, dep2),
				getCacheOrCall(ctx, dep3),
				getCacheOrCall(ctx, dep4),
				getCacheOrCall(ctx, dep5),
			)
		})
	}
}

func Cfp6[C, T1, T2, T3, T4, T5, T6, R any](
	dep1 *func(ctx *C) T1,
	dep2 *func(ctx *C) T2,
	dep3 *func(ctx *C) T3,
	dep4 *func(ctx *C) T4,
	dep5 *func(ctx *C) T5,
	dep6 *func(ctx *C) T6,
	fn func(x1 T1, x2 T2, x3 T3, x4 T4, x5 T5, x6 T6) R,
) func(ctx *C) R {
	return func(ctx *C) R {
		return cfpBase(ctx, func() R {
			return fn(
				getCacheOrCall(ctx, dep1),
				getCacheOrCall(ctx, dep2),
				getCacheOrCall(ctx, dep3),
				getCacheOrCall(ctx, dep4),
				getCacheOrCall(ctx, dep5),
				getCacheOrCall(ctx, dep6),
			)
		})
	}
}

func Cfp7[C, T1, T2, T3, T4, T5, T6, T7, R any](
	dep1 *func(ctx *C) T1,
	dep2 *func(ctx *C) T2,
	dep3 *func(ctx *C) T3,
	dep4 *func(ctx *C) T4,
	dep5 *func(ctx *C) T5,
	dep6 *func(ctx *C) T6,
	dep7 *func(ctx *C) T7,
	fn func(x1 T1, x2 T2, x3 T3, x4 T4, x5 T5, x6 T6, x7 T7) R,
) func(ctx *C) R {
	return func(ctx *C) R {
		return cfpBase(ctx, func() R {
			return fn(
				getCacheOrCall(ctx, dep1),
				getCacheOrCall(ctx, dep2),
				getCacheOrCall(ctx, dep3),
				getCacheOrCall(ctx, dep4),
				getCacheOrCall(ctx, dep5),
				getCacheOrCall(ctx, dep6),
				getCacheOrCall(ctx, dep7),
			)
		})
	}
}

func Cfp8[C, T1, T2, T3, T4, T5, T6, T7, T8, R any](
	dep1 *func(ctx *C) T1,
	dep2 *func(ctx *C) T2,
	dep3 *func(ctx *C) T3,
	dep4 *func(ctx *C) T4,
	dep5 *func(ctx *C) T5,
	dep6 *func(ctx *C) T6,
	dep7 *func(ctx *C) T7,
	dep8 *func(ctx *C) T8,
	fn func(x1 T1, x2 T2, x3 T3, x4 T4, x5 T5, x6 T6, x7 T7, x8 T8) R,
) func(ctx *C) R {
	return func(ctx *C) R {
		return cfpBase(ctx, func() R {
			return fn(
				getCacheOrCall(ctx, dep1),
				getCacheOrCall(ctx, dep2),
				getCacheOrCall(ctx, dep3),
				getCacheOrCall(ctx, dep4),
				getCacheOrCall(ctx, dep5),
				getCacheOrCall(ctx, dep6),
				getCacheOrCall(ctx, dep7),
				getCacheOrCall(ctx, dep8),
			)
		})
	}
}

func Cfp9[C, T1, T2, T3, T4, T5, T6, T7, T8, T9, R any](
	dep1 *func(ctx *C) T1,
	dep2 *func(ctx *C) T2,
	dep3 *func(ctx *C) T3,
	dep4 *func(ctx *C) T4,
	dep5 *func(ctx *C) T5,
	dep6 *func(ctx *C) T6,
	dep7 *func(ctx *C) T7,
	dep8 *func(ctx *C) T8,
	dep9 *func(ctx *C) T9,
	fn func(x1 T1, x2 T2, x3 T3, x4 T4, x5 T5, x6 T6, x7 T7, x8 T8, x9 T9) R,
) func(ctx *C) R {
	return func(ctx *C) R {
		return cfpBase(ctx, func() R {
			return fn(
				getCacheOrCall(ctx, dep1),
				getCacheOrCall(ctx, dep2),
				getCacheOrCall(ctx, dep3),
				getCacheOrCall(ctx, dep4),
				getCacheOrCall(ctx, dep5),
				getCacheOrCall(ctx, dep6),
				getCacheOrCall(ctx, dep7),
				getCacheOrCall(ctx, dep8),
				getCacheOrCall(ctx, dep9),
			)
		})
	}
}

func cfpBase[C, R any](ctx *C, cb func() R) R {
	cacheInitialized := false
	if cache[ctx] == nil {
		cacheInitialized = true
		cache[ctx] = map[interface{}]interface{}{}
	}
	resp := cb()
	if cacheInitialized {
		delete(cache, ctx)
	}
	return resp
}

func getCacheOrCall[C, T any](ctx *C, depFn *func(ctx *C) T) T {
	var depResp T
	if cache[ctx][depFn] == nil {
		depResp = (*depFn)(ctx)
		cache[ctx][depFn] = depResp
	} else {
		depResp = cache[ctx][depFn].(T)
	}
	return depResp
}

func Sfp[T comparable, S any](fn func(state S, payload T) S, init S) func(payload T) S {
	var state S = init
	var null T // TODO https://github.com/golang/go/issues/53656
	return func(payload T) S {
		if payload == null {
			return state
		}
		state = fn(state, payload)
		return state
	}
}
