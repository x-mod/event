package event

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEvent(t *testing.T) {
	ev := New()
	assert.NotNil(t, ev)
	r1 := ev.Fire()
	<-ev.Done()
	assert.Equal(t, int32(1), r1)
	assert.Equal(t, true, ev.HasFired())
	r2 := ev.Fire()
	assert.Equal(t, int32(2), r2)
	assert.Equal(t, true, ev.HasFired())

	e2 := New()
	time.AfterFunc(time.Second, func() {
		e2.Fire()
	})
	log.Println("waiting 1 sec")
	<-Done()
}
