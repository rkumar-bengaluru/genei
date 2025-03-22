package seed

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"example.com/rest-api/models"
	"example.com/rest-api/service"
)

func SeedCampaignData(svc *service.CampaignService) {
	// Read the JSON file
	jsonFile, err := os.Open("wellbe/camp_detail_list.json")
	if err != nil {
		log.Fatalf("Error opening JSON file: %v", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile) // Using io.ReadAll
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	// Unmarshal the JSON data into a struct
	var campaigns_map = make(map[string][]models.ArogyaCampaign)
	err = json.Unmarshal(byteValue, &campaigns_map)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	docs, ok := campaigns_map["docs"]
	if !ok {
		log.Fatalf("Error reading models.WellBeUsers")
	}

	fmt.Println("length of users", len(docs))
	ctx := context.Background()
	for _, doc := range docs {
		fmt.Println(doc)
		svc.CreateCampaign(&ctx, &doc)
	}

}
