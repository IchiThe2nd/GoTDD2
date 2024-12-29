package main

import (
	"testing"
	"time"

	acceptancetests "github.com/IchiThe2nd/GoTDD2/acceptancetesting/acceptancetests/acceptancetests/gracefulSD"
	assert "github.com/IchiThe2nd/GoTDD2/acceptancetesting/assert"
)

const (
	port = "8080"
	url  = "http://localhost: " + port
)

func TestGracefulShutdown(t *testing.T) {
	cleanup, sendInterrupt, err := acceptancetests.LaunchTestProgram(port)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(cleanup)

	//check the server works before shutting down
	assert.CanGet(t, url)

	//fire off a request and before it has a chance to respond send SIGTERM.
	time.AfterFunc(50*time.Millisecond, func() {
		assert.Noerror(t, sendInterrupt())
	})
	// withour gracefulshtdown this would fail
	assert.CanGet(t, url)
}

//assertions andc helpers
