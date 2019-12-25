package component

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestMetadata_UnmarshalJSON(t *testing.T) {
	data, err := ioutil.ReadFile(filepath.Join("testdata", "metadata.json"))
	require.NoError(t, err)

	got := Metadata{}
	require.NoError(t, got.UnmarshalJSON(data))

	expected := Metadata{
		Type: "type",
		Title: []TitleComponent{
			NewText("title"),
		},
		Accessor: "accessor",
	}
	require.Equal(t, expected, got)
}
