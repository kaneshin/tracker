package tracker

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestParsePaths runs
func TestParsePaths(t *testing.T) {
	assert := assert.New(t)

	// Default
	assert.NotNil(Logger)
	assert.NotNil(Format)
	assert.Nil(Hook)

	// Format tests
	d := 12 * time.Second
	assert.Equal("12s", Format(d))

	// Time tests
	var buf bytes.Buffer
	Logger = log.New(&buf, "", 0)
	Format = func(d time.Duration) string {
		return fmt.Sprintf("Result: %ds", int(d.Seconds()))
	}
	t1 := time.Now().Add(-d)
	Time(t1)
	assert.Equal("Result: 12s\n", buf.String())

	buf = bytes.Buffer{}
	Logger = log.New(&buf, "", 0)
	Format = nil
	Hook = nil
	Time(t1)
	assert.True(strings.HasPrefix(buf.String(), "12."))
}
