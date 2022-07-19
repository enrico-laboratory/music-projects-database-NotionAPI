package notion

import (
	"encoding/json"
	"io"
	"net/http"
)

func readJSON(r *http.Response, model any) error {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyBytes, &model)
	if err != nil {
		return err
	}
	return nil
}
