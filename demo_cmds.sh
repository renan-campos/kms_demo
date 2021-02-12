# If you are not me, you should change this:
KEY="alias/rcampos-test-key-1"

aws-encrypt() {
  if [ -z "$1" ]
  then
    MSG="Hello World"
  else
    MSG="$1"
  fi

  go run encrypt.go \
    -keyId "$KEY" \
    -msg "$MSG"
}

aws-decrypt() {
  go run decrypt.go
}

aws-rotate() {
  go run rotate.go \
    -keyAlias "$KEY"

  echo "Showing results of 'aws kms list-aliases':"
  aws kms list-aliases | grep -B 1 -A 3 "\"$KEY\""
}
