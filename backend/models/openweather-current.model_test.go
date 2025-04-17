package models

import (
	"io"
	"os"
	"testing"

	"github.com/Jason2924/scanner/backend/ultilities"
	"github.com/stretchr/testify/require"
)

func Test_ParseOpenWeatherCurrent(t *testing.T) {
	t.Log("Test_ParseOpenWeatherCurrent start")
	jsonFile, erro := os.Open("../data/openweather/openweather-current.data.json")
	require.NoError(t, erro, "Open json file error")
	defer func() {
		jsonFile.Close()
	}()
	jsonData, erro := io.ReadAll(jsonFile)
	require.NoError(t, erro, "Read json file error")
	model := &OpenWeatherCurrentResp{}
	erro = ultilities.ParseObjectFromJson(jsonData, model)
	require.NoError(t, erro, "Parse json file to object error")
}
