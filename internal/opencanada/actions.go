package opencanada

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func ShowPackageMetadata(package_id string) interface{} {

	// TODO: take in dataset title and perform lookup against datasets.json
	// TODO: datasets.json should be exposed to a frontend as an options list; might require moving

	return action("package_show", "id", package_id)
}

func SearchDataStoreResource(resource_id string) interface{} {

	return action("datastore_search", "resource_id", resource_id)
}

func action(args ...string) interface{} {

	// TODO: need size handling for args
	// TODO: additional kv pair filters can be accepted, join each additional pair as "&key=value"

	// TODO: need to handle pagination

	action := args[0]
	field := args[1]
	value := args[2]

	const api_action_url = "https://open.canada.ca/data/en/api/3/action"

	var target_url = fmt.Sprintf("%s/%s?%s=%s", api_action_url, action, field, value)

	resp, err := http.Get(target_url)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var parsed interface{}
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		log.Fatal(err)
	}
	return parsed
}
