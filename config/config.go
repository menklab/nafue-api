package config

import (
	"os"
	"strconv"
	"fmt"
)

var (
	// Debug
	Debug bool
	SecurityOutput bool

	// S3
	S3Key string
	S3PutTTL int64
	S3Bucket string
	S3Location string

	// DB
	DbName string
	DbUser string
	DbPassword string
	DbServer string

	// App Config
	Port string
	CorsHost string
	ChunkSize int64
	PresignLimit int64

	// Braintree
	BtEnv string
	BtMerchId string
	BtMerchActId string
	BtPubKey string
	BtPrivKey string
)

func init() {

	// Debug
	Debug = getBoolOrFail("DEBUG")
	SecurityOutput= getBoolOrFail("SECURITY_OUTPUT")

	// S3
	S3Key= getStringOrFail("S3_KEY")
	S3Bucket= getStringOrFail("S3_BUCKET")
	S3Location= getStringOrFail("S3_LOCATION")
	S3PutTTL= getIntOrFail("S3_PUT_TTL")


	// DB
	DbName= getStringOrFail("DB_NAME")
	DbUser= getStringOrFail("DB_USER")
	DbPassword= getStringOrFail("DB_PASSWORD")
	DbServer= getStringOrFail("DB_SERVER")

	// App Config
	Port= getStringOrFail("PORT")
	CorsHost= getStringOrFail("CORS_HOST")
	ChunkSize= getIntOrFail("CHUNK_SIZE")
	PresignLimit= getIntOrFail("PRESIGN_LIMIT")


	// Braintree
	BtEnv= getStringOrFail("BT_ENV")
	BtMerchId= getStringOrFail("BT_MERCH_ID")
	BtMerchActId= getStringOrFail("BT_MERCH_ACT_ID")
	BtPubKey= getStringOrFail("BT_PUB_KEY")
	BtPrivKey= getStringOrFail("BT_PRIV_KEY")

}

func getIntOrFail(envVar string) int64 {
	is := os.Getenv(envVar)
	if is == "" {
		fmt.Println("Error retrieving envVar: " + envVar)
		os.Exit(1)
	}
	i, err := strconv.ParseInt(is, 10, 10)
	if err != nil {
		fmt.Println("Error parsing envVar: " + envVar + " into int: " + err.Error())
		os.Exit(0)
	}
	return i
}

func getStringOrFail(envVar string) string {
	s := os.Getenv(envVar)
	if s == "" {
		fmt.Println("Error retrieving envVar: " + envVar)
		os.Exit(1)
	}
	return s
}
func getBoolOrFail(envVar string) bool {
	bs := os.Getenv(envVar)
	if bs == "" {
		fmt.Println("Error retrieving envVar: " + envVar)
		os.Exit(1)
	}
	b, err := strconv.ParseBool(os.Getenv(envVar))
	if err != nil {
		fmt.Println("Error parsing envVar: " + envVar + " into bool: " + err.Error())
		os.Exit(0)
	}
	return b
}