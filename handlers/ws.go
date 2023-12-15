package handlers

import (
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/wsserver"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func WebsocketUpgrade(r *http.Request, w http.ResponseWriter) {
	mid := r.Context().Value(models.CTX_user_id).(int64)
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		// handle error
	}
	go func() {
		defer conn.Close()

		for {
			msg, op, err := wsutil.ReadClientData(conn)
			if err != nil {
				// handle error
			}
			err = wsutil.WriteServerMessage(conn, op, msg)
			if err != nil {
				// handle error
			}
		}
	}()

	wsserver.Global.Clients[mid] = &wsserver.Client{
		UserId:   mid,
		Platform: "web",
		// conn:     conn,
	}
}
