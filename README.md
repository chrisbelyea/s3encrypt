[![Go Report Card](https://goreportcard.com/badge/github.com/donmills/s3encrypt)](https://goreportcard.com/report/github.com/donmills/s3encrypt)
# s3encrypt
A tool designed to work with the [ruby-kms-s3-gem](https://github.com/DonMills/ruby-kms-s3-gem).  
Fully compatable with the gem - can encrypt and upload files or download and decrypt files.  Also can do all forms of S3 SSE.

#### New: KMSencrypt
Like the concept of S3encrypt but wish it didn't use S3?  Want to use KMS generated keys to encrypt files for local storage, email, or other places like github?  Well, here you go:

https://github.com/DonMills/kmsencrypt
___
## Decryption
1. takes a file that you have uploaded via the ruby gem or this go program
2. fetches the envelope key and decrypts the envelope key with the appropriate EncryptionContext via KMS
3. then takes that key and unencrypts the data key stored with the file in metadata
4. and then uses _that_ key to decrypt the file and save it in the location specified.

## Encryption
1. This takes the file,
2. generates a KMS envelope key tied to a supplied EncryptionContext value
3. generates a local data encryption key
4. encrypts the file with the data key
5. encrypts the data key with the envelope key
6. stores the encrypted envelope key in s3 with a (filename).key value
7. stores the encrypted data key in the s3 metadata and uploads the encrypted file

___

## How to build:

#### git clone into the $GOPATH/src/github.com/DonMills directory
```
mkdir $GOPATH/src/github.com/DonMills
cd $GOPATH/src/github.com/DonMills
git clone https://github.com/DonMills/s3encrypt.git
```
_or_
```
go get github.com/DonMills/s3encrypt
```

#### This tool requires the "aws-sdk-go" and the ["urfave/cli"](https://github.com/urfave/cli) packages be installed.
```
go get github.com/aws/aws-sdk-go/
go get github.com/urfave/cli
```
Alternatively, if you have [glide](https://github.com/Masterminds/glide) installed, you can just get the deps like this:
```
glide up
```

#### Then just build or install it...
```
go install
```
_or_
```bash
go build -o s3encrypt
./s3encrypt 
```
## Usage Notes
The tool has a full help system, but in general usage is 
```
 s3encrypt [command] {command specific options}
```
where commands are 
```
s3encrypt encrypt [localfilename] [remotefilename] [bucket] [context]
OPTIONS:
   -c value  The customer master key id - can set with S3ENCRYPT_CMKID environment variable [$S3ENCRYPT_CMKID]
   -s value  The ServerSideEncryption method to use - default is none, valid options are "AES256" or "aws:kms" (default: "nil")
```
or
```
s3encrypt decrypt [localfilename] [remotefilename] [bucket] [context]
```
#### Dealing with an "AWS Error: NoCredentialProviders" error or needing ~/.aws/config
In some situations (like needing a STS token to work on an environment) or if you have entries in your ~/.aws/config file that are needed, you may need to set the following environment variable:
```
AWS_SDK_LOAD_CONFIG=1
```
This is a function of the aws sdk for go discussed here: http://docs.aws.amazon.com/sdk-for-go/api/aws/session/
## Mac installation via homebrew
New!  Now you can install on a mac by using homebrew.
```
brew install DonMills/tools/s3encrypt
```

