package api

import (
	"encoding/json"
	"net/http"

	"github.com/apbenedetti/maple-pie/internal/opencanada"
)

func OpenCanadaHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(opencanada.ShowPackageMetadata("065439a9-c194-4259-9c95-245a852be4a1"))

	// opencanada.SearchDataStoreResource("d9ff2144-4591-4ef5-b313-be5012eca3be")

}
