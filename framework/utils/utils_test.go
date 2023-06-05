package utils_test

import (
	"testing"

	"github.com/raphaelmb/go-ms-encoder/framework/utils"
	"github.com/stretchr/testify/require"
)

func TestIsJson(t *testing.T) {
	json := `
	{
		"id": "0c02c92a-581d-4721-9550-60682b2b9a40",
		"file_path": "file.mp4",
		"status": "pending"
	}
	`
	err := utils.IsJSON(json)
	require.Nil(t, err)

	json = `wrong`
	err = utils.IsJSON(json)
	require.Error(t, err)
}
