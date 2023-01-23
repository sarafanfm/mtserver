package static

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"strings"
)

type NotFoundRedirectRespWr struct {
	http.ResponseWriter
	status      int
	docsRoot    http.Dir
	hasNotFound bool
	requestPath string
	indexFile   string
}

func (w *NotFoundRedirectRespWr) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w *NotFoundRedirectRespWr) WriteHeader(status int) {
	w.status = status
	if status != http.StatusNotFound {
		w.ResponseWriter.WriteHeader(status)
	} else {
		w.hasNotFound = true
		w.ResponseWriter.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.ResponseWriter.WriteHeader(http.StatusOK)
	}
}

func (w *NotFoundRedirectRespWr) Write(p []byte) (int, error) {
	if w.hasNotFound {
		p, err := w.docsRoot.Open(w.indexFile)
		if err != nil {
			return 0, err
		}
		defer p.Close()

		buf := bytes.NewBuffer(nil)
		io.Copy(buf, p)

		// fill og: tags if needed
		if post, err := w.isPostPage(); err == nil {
			buf = w.ogTagsPost(buf, post)
		}

		return w.ResponseWriter.Write(buf.Bytes())
	} else {
		return w.ResponseWriter.Write(p)
	}
}

func (w *NotFoundRedirectRespWr) isPostPage() (interface{}, error) {
	pathArr := strings.Split(w.requestPath, "/")
	if len(pathArr) > 1 && pathArr[0] == "" {
		pathArr = pathArr[1:]
	}

	if pathArr[0] != "post" || pathArr[1] != "view" || len(pathArr) < 4 {
		return nil, errors.New("not post")
	}

	return nil, errors.New("not implemented")
}

func (w *NotFoundRedirectRespWr) ogTagsPost(buf *bytes.Buffer, post interface{}) *bytes.Buffer {
	ogTags := []byte(`<meta property="og:title" content="Post Author">`)
	ogTags = append(ogTags, []byte(`<meta property="og:description" content="Post description">`)...)

	if false { //if post.Images != nil && len(post.Images) > 0 {
		ogTags = append(ogTags, []byte(`<meta property="og:image" content="post.Images[0]">`)...)
	} else if false == true { //} else if post.Author.Avatar != "" {
		ogTags = append(ogTags, []byte(`<meta property="og:image" content="post.Author.Avatar">`)...)
	}
	ogTags = append(ogTags, []byte(`</head>`)...)

	replaced := bytes.ReplaceAll(buf.Bytes(), []byte(`</head>`), ogTags)
	return bytes.NewBuffer(replaced)
}
