package heartbeat_test

import (
	"testing"
	"time"

	tp "github.com/henrylee2cn/teleport"
	heartbeat "github.com/henrylee2cn/tp-ext/plugin-heartbeat"
)

func TestHeartbeat11(t *testing.T) {
	srv := tp.NewPeer(
		tp.PeerConfig{ListenAddress: ":9090"},
		heartbeat.NewPong(),
	)
	go srv.Listen()
	time.Sleep(time.Second)

	cli := tp.NewPeer(
		tp.PeerConfig{},
		heartbeat.NewPing(3),
	)
	cli.Dial(":9090")
	time.Sleep(time.Second * 10)
}

func TestHeartbeat12(t *testing.T) {
	srv := tp.NewPeer(
		tp.PeerConfig{ListenAddress: ":9090"},
		heartbeat.NewPong(),
	)
	go srv.Listen()
	time.Sleep(time.Second)

	cli := tp.NewPeer(
		tp.PeerConfig{},
		heartbeat.NewPing(3),
	)
	sess, _ := cli.Dial(":9090")
	for i := 0; i < 8; i++ {
		sess.Pull("/", nil, nil)
		time.Sleep(time.Second)
	}
	time.Sleep(time.Second * 5)
}

func TestHeartbeat21(t *testing.T) {
	srv := tp.NewPeer(
		tp.PeerConfig{ListenAddress: ":9090"},
		heartbeat.NewPing(3),
	)
	go srv.Listen()
	time.Sleep(time.Second)

	cli := tp.NewPeer(
		tp.PeerConfig{},
		heartbeat.NewPong(),
	)
	cli.Dial(":9090")
	time.Sleep(time.Second * 10)
}

func TestHeartbeat22(t *testing.T) {
	srv := tp.NewPeer(
		tp.PeerConfig{ListenAddress: ":9090"},
		heartbeat.NewPing(3),
	)
	go srv.Listen()
	time.Sleep(time.Second)

	cli := tp.NewPeer(
		tp.PeerConfig{},
		heartbeat.NewPong(),
	)
	sess, _ := cli.Dial(":9090")
	for i := 0; i < 8; i++ {
		sess.Push("/", nil)
		time.Sleep(time.Second)
	}
	time.Sleep(time.Second * 5)
}
