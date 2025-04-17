package ultilities

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func HttpGet(ctxt context.Context, url string, headers map[string]string, result any) error {
	reqt, erro := http.NewRequestWithContext(ctxt, http.MethodGet, url, nil)
	if erro != nil {
		return fmt.Errorf("error occured while creating request: %v", erro)
	}
	for key, value := range headers {
		reqt.Header.Set(key, value)
	}
	resp, erro := http.DefaultClient.Do(reqt)
	if erro != nil {
		return fmt.Errorf("error occured while calling request: %v", erro)
	}
	defer resp.Body.Close()
	body, erro := io.ReadAll(resp.Body)
	if erro != nil {
		return fmt.Errorf("error occured while reading request body: %v", erro)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error non-successful request: %s", string(body))
	}
	erro = ParseObjectFromJson(body, result)
	if erro != nil {
		return fmt.Errorf("error occured while parsing request body: %v", erro)
	}
	return nil
}
