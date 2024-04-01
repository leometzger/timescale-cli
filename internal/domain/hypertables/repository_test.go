package hypertables

import (
	"log/slog"
	"testing"

	"github.com/leometzger/timescale-cli/testlib"
	"github.com/stretchr/testify/assert"
)

func TestGetHypertableInformationFromTimescale(t *testing.T) {
	conn := testlib.GetConnection()
	repository := NewHypertablesRepository(conn, slog.Default().WithGroup("hypertables"))

	hypertables, err := repository.GetHypertables()

	assert.Nil(t, err)
	assert.Equal(t, 1, len(hypertables))
	assert.Equal(t, "metrics", hypertables[0].HypertableName)
}
