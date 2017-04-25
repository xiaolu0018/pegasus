package util

import "github.com/julienschmidt/httprouter"

func AddParam(ps httprouter.Params, key, value string) httprouter.Params {
	newPs := ([]httprouter.Param)(ps)
	newPs = append(newPs, httprouter.Param{key, value})
	return httprouter.Params(newPs)
}
