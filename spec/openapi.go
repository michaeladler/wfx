package spec

/*
 * SPDX-FileCopyrightText: 2024 Siemens AG
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * Author: Michael Adler <michael.adler@siemens.com>
 */

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/goccy/go-yaml"
)

//go:embed wfx.openapi.yml
var openapiYAML string

func init() {
	var yamlObj map[string]any
	if err := yaml.Unmarshal([]byte(openapiYAML), &yamlObj); err != nil {
		panic(err)
	}
	servers := yamlObj["servers"].([]any)
	servers2 := servers[0].(map[string]any)
	basePath := servers2["url"]
	specEndpoint := fmt.Sprintf("%s/openapi.json", basePath)

	jsonData, err := json.MarshalIndent(yamlObj, "", "  ")
	if err != nil {
		panic(err)
	}

	Handlers[fmt.Sprintf("GET %s", specEndpoint)] = http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(jsonData)
	})

	Handlers["GET /"] = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}
		w.Header().Set("Link", fmt.Sprintf(`<%s://%s%s>; rel="service-desc"`, scheme, r.Host, specEndpoint))
		w.WriteHeader(http.StatusNoContent)
	})
}
