package hypertables

import (
	"log/slog"
	"testing"

	"github.com/leometzger/timescale-cli/internal/domain"
	"github.com/leometzger/timescale-cli/testlib"
	"github.com/stretchr/testify/require"
)

func TestGetHypertableInformationFromTimescale(t *testing.T) {
	conn := testlib.GetConnection()
	repository := NewHypertablesRepository(conn, slog.Default().WithGroup("hypertables"))
	filter := HypertablesFilter{}

	hypertables, err := repository.GetHypertables(&filter)

	require.Nil(t, err)
	require.Equal(t, 1, len(hypertables))
	require.Equal(t, "metrics", hypertables[0].HypertableName)
}

func TestGetHypertableInformationFromTimescaleWithNameFilter(t *testing.T) {
	conn := testlib.GetConnection()
	repository := NewHypertablesRepository(conn, slog.Default().WithGroup("hypertables"))
	filter := HypertablesFilter{Name: "metricsaa"}

	hypertables, err := repository.GetHypertables(&filter)

	require.Nil(t, err)
	require.Equal(t, 0, len(hypertables))
}

func TestGetHypertableInformationFromTimescaleWithCompressFilter(t *testing.T) {
	conn := testlib.GetConnection()
	repository := NewHypertablesRepository(conn, slog.Default().WithGroup("hypertables"))
	filter := HypertablesFilter{Compressed: domain.OptionFlagTrue}

	hypertables, err := repository.GetHypertables(&filter)

	require.Nil(t, err)
	require.Equal(t, 0, len(hypertables))
}

func TestGetHypertableInformationFromTimescaleWithCompoundFilter(t *testing.T) {
	conn := testlib.GetConnection()
	repository := NewHypertablesRepository(conn, slog.Default().WithGroup("hypertables"))
	filter := HypertablesFilter{Name: "metrics", Compressed: domain.OptionFlagFalse}

	hypertables, err := repository.GetHypertables(&filter)

	require.Nil(t, err)
	require.Equal(t, 1, len(hypertables))
}
