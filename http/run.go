package http

import (
	"io/ioutil"
	"net/http"

	"github.com/baishancloud/octopux-swcollector/g"
	"github.com/toolkits/sys"
)

func configRunRoutes() {
	http.HandleFunc("/run", func(w http.ResponseWriter, r *http.Request) {
		if g.IsTrustable(r.RemoteAddr) {
			if r.ContentLength == 0 {
				http.Error(w, "body is blank", http.StatusBadRequest)
				return
			}

			bs, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			body := string(bs)
			out, err := sys.CmdOutBytes("sh", "-c", body)
			if err != nil {
				w.Write([]byte("exec fail: " + err.Error()))
				return
			}

			w.Write(out)
		} else {
			w.Write([]byte("no privilege"))
		}
	})
}
