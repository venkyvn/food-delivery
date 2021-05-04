package uploadbiz

import (
	"bytes"
	"context"
	"fmt"
	"go-food-delivery/common"
	"go-food-delivery/component/uploadprovider"
	"go-food-delivery/modules/upload/uploadmodel"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"
)

/*
note: using mimetype check to validate input data
*/
type CreateImageStorage interface {
	CreateImage(ctx context.Context, data *common.Image) error
}

type uploadBiz struct {
	provider uploadprovider.UploadProvider
	imgStore CreateImageStorage
}

func NewUploadBiz(provider uploadprovider.UploadProvider, imgStore CreateImageStorage) *uploadBiz {
	return &uploadBiz{
		provider: provider,
		imgStore: imgStore,
	}
}

func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)

	w, h, err := getImageDimension(fileBytes)

	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage(err)
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)                                    // "abc.jpg" => ".jpg"
	newFileName := fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt) // 9129324893248.jpg
	dest := fmt.Sprintf("%s/%s", folder, newFileName)

	img, err := biz.provider.SaveFileUploaded(ctx, data, dest)
	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	img.Width = w
	img.Height = h
	//img.CouldName = "s3" // should be set in provider
	img.Extension = fileExt

	//if err := biz.imgStore.CreateImage(ctx, img); err != nil {
	//	// delete img on S3
	//	return nil, uploadmodel.ErrCannotSaveFile(err)
	//}

	return img, nil
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)

	if err != nil {
		log.Println("err: ", err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
