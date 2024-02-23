package api

import (
	"net/url"
	"regexp"
)

func getPathParams(URL *url.URL, urlModel string) map[string]string {
	// Get param ids
	re := regexp.MustCompile(`{[a-z0-9]+}`)
	paramIds := re.FindAllString(urlModel, -1)

	// Get param values
	var paramValues []string

	for i, r := range urlModel {
		partialPath := ""

		if r != '{' {
			continue
		}

		partialPath = partialPath + urlModel[:i]

		paramValue := ""

		for _, s := range URL.Path[len(partialPath):] {
			if s == '/' {
				break
			}

			paramValue = paramValue + string(s)
		}

		paramValues = append(paramValues, paramValue)
	}

	if len(paramValues) != len(paramIds) {
		return nil
	}

	// Mix ids and values together
	paramMap := make(map[string]string)
	for i := 0; i < len(paramIds); i++ {
		id := paramIds[i]
		key := id[1 : len(id)-1]

		paramMap[key] = paramValues[i]
	}

	return paramMap
}
