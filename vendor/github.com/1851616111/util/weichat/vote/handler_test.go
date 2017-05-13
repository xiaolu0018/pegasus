package vote

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"testing"
)

func TestAddRouter(t *testing.T) {
	r := httprouter.New()
	AddRouter(r, "/home/michael/IdeaProjects/src/bjdaos/pegasus/dist/activity")

	http.ListenAndServe(":9000", r)
	select {}
}
