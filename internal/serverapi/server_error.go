package serverapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func bodyToError(body io.ReadCloser) error {
	data, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	response := RespError{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return err
	}

	return fmt.Errorf("server error: %v", response.Error)
}
