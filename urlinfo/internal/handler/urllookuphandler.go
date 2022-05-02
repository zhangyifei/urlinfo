package handler

import (
	"net/http"

	"urlinfo/urlinfo/internal/logic"
	"urlinfo/urlinfo/internal/svc"
	"urlinfo/urlinfo/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UrlLookupHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		// Get all querys from the url
		if len(r.URL.RawQuery) > 0 {
			req.Queryparamter += "?" + r.URL.RawQuery
		}

		l := logic.NewurlinfoLogic(r.Context(), svcCtx)
		resp, err := l.UrlLookup(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
