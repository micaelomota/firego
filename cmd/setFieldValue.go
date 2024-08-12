/*
Copyright © 2024 Micael Mota <micaelomota@gmail.com>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/spf13/cobra"
	"google.golang.org/api/iterator"
)

// setFieldValueCmd represents the setFieldValue command
var setFieldValueCmd = &cobra.Command{
	Use:   "setFieldValue",
	Short: "Set a value for a field in all documents in a Firestore collection",
	Long:  `Set a value for a field in all documents in a Firestore collection`,
	Run: func(cmd *cobra.Command, args []string) {
		collection, _ := cmd.Flags().GetString("collection")
		field, _ := cmd.Flags().GetString("field")
		value, _ := cmd.Flags().GetString("value")

		// Print the values for debugging
		fmt.Println("Collection:", collection)
		fmt.Println("Field:", field)
		fmt.Println("Value:", value)

		// Ensure required flags are provided
		if collection == "" || field == "" || value == "" {
			fmt.Println("Usage: go run main.go -c <collection> -f <field> -v <value>")
			os.Exit(1)
		}

		// Initialize Firebase Admin SDK
		ctx := context.Background()
		// sa := option.WithCredentialsFile(*credentials)
		app, err := firebase.NewApp(ctx, nil)
		if err != nil {
			log.Fatalf("error initializing app: %v\n", err)
		}

		client, err := app.Firestore(ctx)
		if err != nil {
			log.Fatalf("error initializing Firestore client: %v\n", err)
		}
		defer client.Close()

		// Determine the value to set
		var updateValue interface{}

		if value == "null" {
			updateValue = nil
		} else {
			updateValue = value
		}

		// Get documents in the specified collection
		iter := client.Collection(collection).Documents(ctx)

		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatalf("error retrieving document: %v\n", err)
			}

			// Update the document
			_, err = doc.Ref.Update(ctx, []firestore.Update{
				{
					Path:  field,
					Value: updateValue,
				},
			})
			if err != nil {
				log.Printf("error updating document %s: %v\n", doc.Ref.ID, err)
			} else {
				log.Printf("Document %s updated successfully\n", doc.Ref.ID)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(setFieldValueCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setFieldValueCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setFieldValueCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Define flags
	setFieldValueCmd.Flags().StringP("collection", "c", "", "The name of the Firestore collection")
	setFieldValueCmd.Flags().StringP("field", "f", "", "The name of the field to update")
	setFieldValueCmd.Flags().StringP("value", "v", "", "The value to set the field to")

	// Mark flags as required
	setFieldValueCmd.MarkFlagRequired("collection")
	setFieldValueCmd.MarkFlagRequired("field")
	setFieldValueCmd.MarkFlagRequired("value")
}
