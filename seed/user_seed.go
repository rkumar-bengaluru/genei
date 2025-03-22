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

func SeedUserData(svc *service.AssigningAuthorityService) {
	// Read the JSON file
	jsonFile, err := os.Open("wellbe/one_campaign.json")
	if err != nil {
		log.Fatalf("Error opening JSON file: %v", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile) // Using io.ReadAll
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	// Unmarshal the JSON data into a struct
	var users_map = make(map[string][]models.AssigningAuthority)
	err = json.Unmarshal(byteValue, &users_map)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	users, ok := users_map["authorities"]
	if !ok {
		log.Fatalf("Error reading models.WellBeUsers")
	}
	fmt.Println("length of users", len(users))
	ctx := context.Background()
	for _, user := range users {
		fmt.Println(user.Name)
		svc.CreateAssigningAuthority(&ctx, &user)
	}

}
