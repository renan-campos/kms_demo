// Encrypts a user provided message with a user provided CMK.
package main

import (
	"flag"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"

	"fmt"
	"os"
)

func main() {

	var pKeyId = flag.String("keyId", "", "The CMK id")
	var pMsg = flag.String("msg", "Hello World", "Message to be encrypted")

	flag.Parse()

	if *pKeyId == "" {
		fmt.Println("Must specify a CMK key id")
		os.Exit(1)
	}
	fmt.Printf("Using key: '%s'\n", *pKeyId)
	fmt.Printf("Encrypting msg: '%s'\n", *pMsg)

	// Initialize a session in us-east-1 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	// Create KMS service client
	svc := kms.New(sess)

	// Encrypt the data
	result, err := svc.Encrypt(&kms.EncryptInput{
		KeyId:     aws.String(*pKeyId),
		Plaintext: []byte(*pMsg),
	})

	if err != nil {
		fmt.Println("Got error encrypting data: ", err)
		os.Exit(1)
	}

	//fmt.Println("Blob (base-64 byte array):")
	//fmt.Println(result.CiphertextBlob)
	fmt.Println("Writing encrypted data to 'ciphertext'")

	file, err := os.Create("ciphertext")
	if err != nil {
		fmt.Println("Error creating 'ciphertext'")
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.Write(result.CiphertextBlob)
	if err != nil {
		fmt.Println("Error writing to 'ciphertext'")
		os.Exit(1)
	}

}
