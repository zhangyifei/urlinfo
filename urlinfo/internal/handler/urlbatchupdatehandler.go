package handler

import (
	"net/http"

	"urlinfo/urlinfo/internal/logic"
	"urlinfo/urlinfo/internal/svc"
	"urlinfo/urlinfo/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UrlBatchUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BatchUpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUrlBatchUpdateLogic(r.Context(), svcCtx)
		resp, err := l.UrlBatchUpdate(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
