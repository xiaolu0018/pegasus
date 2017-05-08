package handler

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/1851616111/util/weichat/handler"
	"github.com/julienschmidt/httprouter"

	"bjdaos/pegasus/pkg/wc/user"

	tk "github.com/1851616111/util/weichat/util/user-token"
)

type redirectManager struct {
	paramName                       string
	baseUrl                         string
	pathToResourceMappings          map[string]string //重定向路径与重定向资源映射
	pathToResourceCompleterMappings map[string]func(*tk.Token, *httprouter.Params) error
}

func NewMenuRedirectManager(base, paramName string) (*redirectManager, error) {
	if base == "" {
		return nil, errors.New("redirect basepath not found")
	}

	if strings.LastIndex(base, string(os.PathSeparator)) == len(base)-1 {
		base = base[:len(base)-1]
	}

	return &redirectManager{
		paramName:                       paramName,
		baseUrl:                         base,
		pathToResourceMappings:          map[string]string{},
		pathToResourceCompleterMappings: map[string]func(*tk.Token, *httprouter.Params) error{},
	}, nil
}

func (m *redirectManager) Redirect(path, Resource string, completer func(*tk.Token, *httprouter.Params) error) {
	m.pathToResourceMappings[path] = Resource
	m.pathToResourceCompleterMappings[path] = completer
}

func (m *redirectManager) AddRedirectToRouter(router *httprouter.Router) error {
	for path := range m.pathToResourceMappings {
		pathTmp := path
		redirectFn, err := m.getRedirectHandler(path)
		if err != nil {
			return err
		}

		router.GET(path, handler.AuthValidator(m.pathToResourceCompleterMappings[pathTmp], redirectFn))
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
		if path == "/api/appoint" {
			if ok, id := user.IDCache.Auth(ps.ByName(m.paramName)); ok {
				if u, err := user.GetUserByid(id); err == nil {
					if u.CardNo != "" {
						resource = "package.html"
					}
				}
			}
		}

		redirectUrl := fmt.Sprintf("%s/%s?%s=%s", m.baseUrl, resource, m.paramName, ps.ByName(m.paramName))

		http.Redirect(w, r, redirectUrl, 302)
	}, err
}

func redirectNotFound(path string) error {
	return fmt.Errorf(" path %s to redirect not found", path)
}
