package contextfp

var cache = map[interface{}]map[interface{}]interface{}{}

func Cfp1[C, T1, R any](
	dep1 func(ctx C) T1,
	fn func(x1 T1) R) func(ctx C) R {
	return func(ctx C) R {
		cacheInitialized := false
		if cache[ctx] == nil {
			cacheInitialized = true
			cache[ctx] = map[interface{}]interface{}{}
		}
		var resp R
		if cache[ctx][fn] == nil {
			resp = fn(dep1(ctx))
		} else {
			resp = cache[ctx][fn].(R)
		}
		if cacheInitialized {
			delete(cache, ctx)
		}
		return resp
	}
}

func Cfp2[C, T1, T2, R any](
	dep1 func(ctx C) T1,
	dep2 func(ctx C) T2,
	fn func(x1 T1, x2 T2) R) func(ctx C) R {
	return func(ctx C) R {
		cacheInitialized := false
		if cache[ctx] == nil {
			cacheInitialized = true
			cache[ctx] = map[interface{}]interface{}{}
		}
		var resp R
		if cache[ctx][fn] == nil {
			resp = fn(dep1(ctx), dep2(ctx))
		} else {
			resp = cache[ctx][fn].(R)
		}
		if cacheInitialized {
			delete(cache, ctx)
		}
		return resp
	}
}

func Cfp3[C, T1, T2, T3, R any](
	dep1 func(ctx C) T1,
	dep2 func(ctx C) T2,
	dep3 func(ctx C) T3,
	fn func(x1 T1, x2 T2, x3 T3) R) func(ctx C) R {
	return func(ctx C) R {
		return fn(dep1(ctx), dep2(ctx), dep3(ctx))
	}
}
