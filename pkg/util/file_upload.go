package util

import (
	"bytes"
	"context"
	"fmt"
	"image/jpeg"
	"io"
	"mime/multipart"
	"path/filepath"

	"github.com/disintegration/imaging"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"golang.org/x/sync/errgroup"
)

type UploadResult struct {
	Type       string
	ObjectName string
	Info       minio.UploadInfo
}

type MinioObject struct {
	Context    context.Context
	Client     *minio.Client
	Bucket     string
	Path       string
	FileHeader *multipart.FileHeader
}

func (m *MinioObject) UploadToS3() (*minio.UploadInfo, error) {
	file, err := m.FileHeader.Open()
	if err != nil {
		return nil, err
	}

	defer file.Close()

	id := uuid.New()
	ext := filepath.Ext(m.FileHeader.Filename)
	objectName := fmt.Sprintf("%s/%s%s", m.Path, id, ext)

	info, err := m.Client.PutObject(
		m.Context,
		m.Bucket,
		objectName,
		file,
		m.FileHeader.Size,
		minio.PutObjectOptions{
			ContentType: m.FileHeader.Header.Get("Content-Type"),
		},
	)

	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (m *MinioObject) UploadImageWithThumbnail() (string, []UploadResult, error) {
	groupID := uuid.New().String()
	ext := filepath.Ext(m.FileHeader.Filename)

	g, ctx := errgroup.WithContext(m.Context)
	results := make([]UploadResult, 2)

	file, err := m.FileHeader.Open()
	if err != nil {
		return "", nil, err
	}

	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return "", nil, err
	}

	g.Go(func() error {
		objName := fmt.Sprintf("%s/%s-original%s", m.Path, groupID, ext)
		info, err := m.Client.PutObject(
			ctx,
			m.Bucket,
			objName,
			bytes.NewReader(fileBytes),
			int64(len(fileBytes)),
			minio.PutObjectOptions{
				ContentType: m.FileHeader.Header.Get("Content-Type"),
			},
		)
		if err == nil {
			results[0] = UploadResult{Type: "original", ObjectName: objName, Info: info}
		}
		return err
	})

	g.Go(func() error {
		src, err := imaging.Decode(bytes.NewReader(fileBytes))
		if err != nil {
			return err
		}

		thumb := imaging.Fill(src, 200, 200, imaging.Center, imaging.Lanczos)

		buf := new(bytes.Buffer)
		if err := jpeg.Encode(buf, thumb, nil); err != nil {
			return err
		}

		objName := fmt.Sprintf("%s/%s-thumb.jpg", m.Path, groupID)
		info, err := m.Client.PutObject(
			ctx,
			m.Bucket,
			objName,
			buf,
			int64(buf.Len()),
			minio.PutObjectOptions{
				ContentType: "image/jpeg",
			})
		if err == nil {
			results[1] = UploadResult{Type: "thumbnail", ObjectName: objName, Info: info}
		}
		return err
	})

	if err := g.Wait(); err != nil {
		return "", nil, err
	}

	return groupID, results, nil
}
