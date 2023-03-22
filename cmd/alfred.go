package cmd

import "fmt"

type AlfredResult struct {
	Items []struct {
		UID          string `json:"uid"`
		Type         string `json:"type"`
		Title        string `json:"title"`
		Subtitle     string `json:"subtitle"`
		Arg          string `json:"arg"`
		Autocomplete string `json:"autocomplete"`
		Icon         struct {
			Type string `json:"type"`
			Path string `json:"path"`
		} `json:"icon"`
	} `json:"items"`
}

func ToSimpleAlfredResult(arg string) string {
	return fmt.Sprintf(`{"items":[{"title":"%s","subtitle":"%s","arg":"%s"}]}`, arg, arg, arg)
}
