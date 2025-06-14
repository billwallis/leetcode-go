package samples_test

import (
	"github.com/billwallis/leetcode-go/internal/samples"
	"testing"
)

func TestNewServerDefault(t *testing.T) {
	got := samples.NewServer()
	want := samples.Server{
		Options: samples.ServerOptions{
			Id:             "default",
			MaxConnections: 10,
			Tls:            false,
		},
	}

	if got.Options != want.Options {
		t.Errorf("got %+v want %+v", got.Options, want.Options)
	}
}

func TestNewServerWithId(t *testing.T) {
	got := samples.NewServer(samples.WithId("test-id"))
	want := samples.Server{
		Options: samples.ServerOptions{
			Id:             "test-id",
			MaxConnections: 10,
			Tls:            false,
		},
	}

	if got.Options != want.Options {
		t.Errorf("got %+v want %+v", got.Options, want.Options)
	}
}

func TestNewServerWithMaxConnections(t *testing.T) {
	got := samples.NewServer(samples.WithMaxConnections(200))
	want := samples.Server{
		Options: samples.ServerOptions{
			Id:             "default",
			MaxConnections: 200,
			Tls:            false,
		},
	}

	if got.Options != want.Options {
		t.Errorf("got %+v want %+v", got.Options, want.Options)
	}
}

func TestNewServerWithTls(t *testing.T) {
	got := samples.NewServer(samples.WithTls(true))
	want := samples.Server{
		Options: samples.ServerOptions{
			Id:             "default",
			MaxConnections: 10,
			Tls:            true,
		},
	}

	if got.Options != want.Options {
		t.Errorf("got %+v want %+v", got.Options, want.Options)
	}
}

func TestNewServerWithOptions(t *testing.T) {
	got := samples.NewServer(
		samples.WithId("test-id"),
		samples.WithMaxConnections(200),
		samples.WithTls(true),
	)
	want := samples.Server{
		Options: samples.ServerOptions{
			Id:             "test-id",
			MaxConnections: 200,
			Tls:            true,
		},
	}

	if got.Options != want.Options {
		t.Errorf("got %+v want %+v", got.Options, want.Options)
	}
}
