# Nafue Security Services
# Menklab LLC

# Nafue Security Services
# Menklab LLC

## Requirements
- Go Version >= 1.6.0

## Setep Env
1. Clone this repository.
2. add .env to root of project directory (contents at bottom)
3. Install Reflex: go get -u github.com/cespare/reflex
4. Install GoVendor: go get -u github.com/kardianos/govendor
5. Install Deps: ./utility.sh deps


## Run
- reflex -c reflex.config


## .env ##
# app config
PORT=9090
CORS_HOST=http://localhost:8080
GIN_MODE=debug

# DB Local
DB_NAME=dbName
DB_USER=dbUser
DB_PASSWORD=dbPassword
DB_SERVER=tcp(localhost:3306)

# s3
S3_KEY=files
S3_PUT_TTL=15
S3_BUCKET=s3Bucket
S3_LOCATION=us-east-1

# aws creds
AWS_ACCESS_KEY_ID=keyId
AWS_SECRET_ACCESS_KEY=key

# braintree sandbox
BT_ENV=sandbox
BT_MERCH_ID=merchId
BT_MERCH_ACT_ID=actId
BT_PUB_KEY=pubKey
BT_PRIV_KEY=privKey
