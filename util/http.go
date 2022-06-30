package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type id struct {
	Mdns string
}

type CfgResponse struct {
	Id id
}

func GetHostname(host string) (string, error) {
	var r CfgResponse

	res, err := http.Get(fmt.Sprintf("http://%s/json/cfg", host))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Failed to get hostname for host '%s' (HTTP %d)", host, res.StatusCode)
	}

	err = json.NewDecoder(res.Body).Decode(&r)
	return r.Id.Mdns, err
}

func DownloadFile(filepath string, url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to download file '%s' (HTTP %d)", url, res.StatusCode)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, res.Body)
	return err
}
