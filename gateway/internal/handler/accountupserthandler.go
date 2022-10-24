package handler

import (
	"net/http"

	"github.com/chernyakoff/telemon/gateway/internal/logic"
	"github.com/chernyakoff/telemon/gateway/internal/svc"
	"github.com/chernyakoff/telemon/gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AccountUpsertHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AccountUpsertRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAccountUpsertLogic(r.Context(), svcCtx)
		resp, err := l.AccountUpsert(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
