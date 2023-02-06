package static

import (
	"log"
	"net/http"
	"os"
)

type StaticServer struct {
	mux     *http.ServeMux
	rootURI string
}

func NewServer(mux *http.ServeMux, rootURI string) *StaticServer {
	s := &StaticServer{mux: mux, rootURI: rootURI}

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	s.handle(path + "/web")

	return s
}

func (s *StaticServer) handle(path string) {
	root := http.Dir(path)
	fs := s.wrapHandler(http.FileServer(root), root)
	s.mux.Handle(s.rootURI, fs)
}

func (s *StaticServer) wrapHandler(h http.Handler, root http.Dir) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// auto redirect to index.html
		nfrw := &NotFoundRedirectRespWr{ResponseWriter: w, docsRoot: root, requestPath: r.URL.Path, indexFile: "index.html"}
		h.ServeHTTP(nfrw, r)
	}
}
