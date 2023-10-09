package model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCSVJson(t *testing.T) {
	path := "../data-test/csv1.csv"

	err := CSV(path)
	require.NoError(t, err)

}

func TestJsonCSV(t *testing.T) {
	path := "../data-test/data2.json"

	err := Json(path)
	require.NoError(t, err)
}
