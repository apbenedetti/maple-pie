package main

import (
	"fmt"

	"github.com/apbenedetti/maple-pie/internal/opencanada"
)

func main() {

	fmt.Print(opencanada.ShowPackageMetadata("065439a9-c194-4259-9c95-245a852be4a1"))
	// fmt.Print(opencanada.SearchDataStoreResource("d9ff2144-4591-4ef5-b313-be5012eca3be"))
}
