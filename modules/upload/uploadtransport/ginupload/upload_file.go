package ginupload

import (
	"github.com/gin-gonic/gin"
	"go-food-delivery/common"
	"go-food-delivery/component"
	"go-food-delivery/modules/upload/uploadbiz"
	"net/http"
)

func Upload(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		//upload static
		//fileHeader, err := c.FormFile("file")
		//
		//if err!= nil {
		//	panic(common.ErrInvalidRequest(err))
		//}

		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close() // defer to close file here

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		//if upload file successful -> save file into db
		//imgStore := uploadNewSqLStore(db)
		biz := uploadbiz.NewUploadBiz(appCtx.UploadProvider(), nil)
		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))

	}
}
