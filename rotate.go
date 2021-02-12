// Demonstates how to manually rotate a CMK:
//	1. Creates a new key.
//  2. Updates the alias to point to the new key.
package main

import (
	"flag"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"

	"fmt"
)

func main() {

	var pKeyAlias = flag.String("keyAlias", "", "The CMK Alias to be updated")

	flag.Parse()

	fmt.Printf("Using key alias: '%s'\n", *pKeyAlias)

	// Initialize a session in us-east-1 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	if err != nil {
		fmt.Println("Error creating session.")
		os.Exit(1)
	}

	// Create KMS service client
	svc := kms.New(sess)

	// Creating a CMK key.
	tk := "Owner"
	tv := "rcampos"
	tag := kms.Tag{
		TagKey:   &tk,
		TagValue: &tv,
	}

	createKeyInput := kms.CreateKeyInput{
		Tags: []*kms.Tag{
			&tag,
		},
	}

	createKeyOutput, err := svc.CreateKey(&createKeyInput)
	if err != nil {
		fmt.Println("Error creating key")
		os.Exit(1)
	}

	pNewKeyId := createKeyOutput.KeyMetadata.Arn

	fmt.Printf("New Key: %s\n", *pNewKeyId)

	// Updating the Alias to point to the key
	updateAliasInput := kms.UpdateAliasInput{
		AliasName:   pKeyAlias,
		TargetKeyId: pNewKeyId,
	}

	_, err = svc.UpdateAlias(&updateAliasInput)
	if err != nil {
		fmt.Println("Error updating alias.")
		os.Exit(1)
	}

	fmt.Println("Successfully updated alias!")
}
