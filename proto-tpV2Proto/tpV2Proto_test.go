package tpV2Proto_test

import (
	"testing"
	"time"

	tp "github.com/henrylee2cn/teleport"
	tpV2Proto "github.com/henrylee2cn/tp-ext/proto-tpV2Proto"
)

type Home struct {
	tp.PullCtx
}

func (h *Home) Test(arg *map[string]interface{}) (map[string]interface{}, *tp.Rerror) {
	return map[string]interface{}{
		"your_id": h.Query().Get("peer_id"),
	}, nil
}

func TestTpV2Proto(t *testing.T) {
	// Server
	srv := tp.NewPeer(tp.PeerConfig{ListenPort: 9090})
	srv.RoutePull(new(Home))
	go srv.ListenAndServe(tpV2Proto.NewProtoFunc)
	time.Sleep(1e9)

	// Client
	cli := tp.NewPeer(tp.PeerConfig{})
	sess, err := cli.Dial(":9090", tpV2Proto.NewProtoFunc)
	if err != nil {
		if err != nil {
			t.Error(err)
		}
	}
	var result interface{}
	rerr := sess.Pull("/home/test?peer_id=110",
		// map[string]interface{}{
		// 	"bytes": []byte("test bytes"),
		// },
		nil,
		&result,
		tp.WithAddMeta("add", "1"),
		tp.WithSetMeta("set", "2"),
	).Rerror()
	if rerr != nil {
		t.Error(rerr)
	}
	t.Logf("=========result:%v", result)
}
