/*
Copyright © 2024 Micael Mota <micaelomota@gmail.com>
*/
package cmd

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a document from Firestore",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		collection, _ := cmd.Flags().GetString("collection")
		document, _ := cmd.Flags().GetString("document")

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

		doc := client.Collection(collection).Doc(document)
		id := doc.ID

		_, err = doc.Delete(ctx)
		if err != nil {
			log.Fatalf("Failed to delete document from Firestore: %v\n", err)
		}

		log.Printf("Document deleted from Firestore: %s\n", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringP("collection", "c", "", "The name of the Firestore collection")
	deleteCmd.Flags().StringP("document", "d", "", "The name of the document to delete")

	deleteCmd.MarkFlagRequired("collection")
	deleteCmd.MarkFlagRequired("document")
}
