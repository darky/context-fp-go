package contextfp

var cache = map[interface{}]map[interface{}]interface{}{}

func Cfp1[C, T1, R any](
	dep1 *func(ctx *C) T1,
	fn func(x1 T1) R) func(ctx *C) R {
	return func(ctx *C) R {
		cacheInitialized := false
		if cache[ctx] == nil {
			cacheInitialized = true
			cache[ctx] = map[interface{}]interface{}{}
		}
		var resp R
		var dep1Resp T1
		if cache[ctx][dep1] == nil {
			dep1Resp = (*dep1)(ctx)
			cache[ctx][dep1] = dep1Resp
		} else {
			dep1Resp = cache[ctx][dep1].(T1)
		}
		resp = fn(dep1Resp)
		if cacheInitialized {
			delete(cache, ctx)
		}
		return resp
	}
}

func Cfp2[C, T1, T2, R any](
	dep1 *func(ctx *C) T1,
	dep2 *func(ctx *C) T2,
	fn func(x1 T1, x2 T2) R) func(ctx *C) R {
	return func(ctx *C) R {
		cacheInitialized := false
		if cache[ctx] == nil {
			cacheInitialized = true
			cache[ctx] = map[interface{}]interface{}{}
		}
		var resp R
		var dep1Resp T1
		var dep2Resp T2
		if cache[ctx][dep1] == nil {
			dep1Resp = (*dep1)(ctx)
			cache[ctx][dep1] = dep1Resp
		} else {
			dep1Resp = cache[ctx][&dep1].(T1)
		}
		if cache[ctx][dep2] == nil {
			dep2Resp = (*dep2)(ctx)
			cache[ctx][dep2] = dep2Resp
		} else {
			dep2Resp = cache[ctx][&dep2].(T2)
		}
		resp = fn(dep1Resp, dep2Resp)
		if cacheInitialized {
			delete(cache, ctx)
		}
		return resp
	}
}
