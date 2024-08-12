/*
Copyright © 2024 Micael Mota <micaelomota@gmail.com>
*/
package cmd

import (
	"context"
	"encoding/json"
	"log"

	firebase "firebase.google.com/go"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a document to Firestore",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		collection, _ := cmd.Flags().GetString("collection")
		jsonString, _ := cmd.Flags().GetString("data")

		var jsonData map[string]interface{}
		if err := json.Unmarshal([]byte(jsonString), &jsonData); err != nil {
			log.Fatalf("Failed to parse JSON: %v", err)
		}

		ctx := context.Background()
		app, err := firebase.NewApp(ctx, nil)
		if err != nil {
			log.Fatalf("error initializing app: %v\n", err)
		}

		client, err := app.Firestore(ctx)
		if err != nil {
			log.Fatalf("error initializing Firestore client: %v\n", err)
		}
		defer client.Close()

		_, _, err = client.Collection(collection).Add(ctx, jsonData)
		if err != nil {
			log.Fatalf("Failed to add document to Firestore: %v", err)
		}

		docRef := client.Collection(collection).NewDoc()
		log.Printf("Document added to Firestore: %s", docRef.ID)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("collection", "c", "", "The name of the Firestore collection")
	addCmd.Flags().StringP("data", "d", "", "The name of the field to update")

	addCmd.MarkFlagRequired("collection")
	addCmd.MarkFlagRequired("data")
}
