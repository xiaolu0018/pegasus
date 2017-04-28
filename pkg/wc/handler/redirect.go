package handler

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/1851616111/util/weichat/handler"
	"github.com/julienschmidt/httprouter"
)

type redirectManager struct {
	baseUrl                string
	pathToResourceMappings map[string]string //重定向路径与重定向资源映射
}

func NewRedirectManager(base string) (*redirectManager, error) {
	if base == "" {
		return nil, errors.New("redirect basepath not found")
	}

	if strings.LastIndex(base, string(os.PathSeparator)) == len(base)-1 {
		base = base[:len(base)-1]
	}

	return &redirectManager{
		baseUrl:                base,
		pathToResourceMappings: map[string]string{},
	}, nil
}

func (m *redirectManager) Redirect(path, Resource string) {
	m.pathToResourceMappings[path] = Resource
}

func (m *redirectManager) AddRouter(router *httprouter.Router) error {
	for path := range m.pathToResourceMappings {
		redirectFn, err := m.getRedirectHandler(path)
		if err != nil {
			return err
		}

		router.GET(path, handler.AuthValidator(completeUserInfo, redirectFn))
	}

	return nil
}

func (m *redirectManager) getRedirectHandler(path string) (func(http.ResponseWriter, *http.Request, httprouter.Params), error) {
	var err error
	resource, ok := m.pathToResourceMappings[path]
	if !ok {
		return nil, redirectNotFound(path)
	}

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		redirectUrl := fmt.Sprintf("%s/%s?bear_token=%s", m.baseUrl, resource, ps.ByName("bear_token"))
		http.Redirect(w, r, redirectUrl, 302)
	}, err
}

func redirectNotFound(path string) error {
	return fmt.Errorf(" path %s to redirect not found", path)
}
