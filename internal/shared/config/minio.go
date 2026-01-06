package config

import (
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

type MinioConfig struct {
	MinioClient *minio.Client
	Bucket      string
}

func NewMinioClient(viper *viper.Viper) *MinioConfig {
	host := viper.GetString("MINIO_HOST")
	port := viper.GetString("MINIO_PORT")
	endpoint := fmt.Sprintf("%s:%s", host, port)
	accessKeyID := viper.GetString("MINIO_USERNAME")
	secretAccessKey := viper.GetString("MINIO_PASSWORD")
	useSSL := viper.GetBool("MINIO_SSL")
	bucketName := viper.GetString("MINIO_BUCKET")

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatal("[Minio] ❌ Gagal koneksi ke Minio")
	}

	log.Printf("[Minio] ✅ Koneksi berhasil ke Minio")
	return &MinioConfig{
		MinioClient: minioClient,
		Bucket:      bucketName,
	}
}
