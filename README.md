This repository contains programs to demonstrate AWS KMS manual key rotation
functionality.

The original ask was how to handle rotating CMKs. From the AWS documentation,
automatic CMK rotation occurs every 365 days. Manually CMK rotation is required
for faster rotation frequencies. 

Manual CMK rotation means creating an alias for the CMK. When it is time to 
rotate, a new CMK is created, and the alias is updated to point to that CMK.
Data encrypted with the original CMK will still be available for decryption...
<b> as long as the original CMK is kept. </b>

To use this repo, modify demo_cmds.sh to point to your CMK alias, then source:
`
$ . ./demo_cmds.sh
`
This will create bash functions that call the go programs.

1. Encrypt some data:
`
$ aws-encrypt "Hello there :)"
`
This will create a file called ciphertext

2. Rotate the key:
`
$ aws-rotate
`
This will create a new key and update the alias to point to that key.

3. Decrypt the data:
`
$ aws-decrypt
`
Even though the alias has been updated, the data was still able to be decrypted.
This is likely because the CMK id gets enbedded with the data in ciphertext.
