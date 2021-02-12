package main

import (
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"

	"fmt"
	"os"
)

func main() {
	// Initialize a session that the SDK uses to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and configuration from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create KMS service client
	svc := kms.New(sess)

	// Encrypted data
	blob, err := ioutil.ReadFile("ciphertext")
	if err != nil {
		fmt.Println("Error reading 'ciphertext'")
		os.Exit(1)
	}

	// Decrypt the data
	// How does the client know which CMK to use to decrypt the data?
	// It must be that the CMK id is encrypted with the data.
	result, err := svc.Decrypt(&kms.DecryptInput{CiphertextBlob: blob})

	if err != nil {
		fmt.Println("Got error decrypting data: ", err)
		os.Exit(1)
	}

	blob_string := string(result.Plaintext)

	fmt.Println(blob_string)
}
