package controller

import (
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"go-service/domain/document/constants"
	"go-service/domain/document/dtos"
	documentInterfaces "go-service/domain/document/interfaces"
	"go-service/infrastructure/exceptions"
	"go-service/infrastructure/utils"
	"go-service/infrastructure/validators"

	"github.com/gin-gonic/gin"
)

type DocumentController struct {
	documentService documentInterfaces.DocumentServiceInterface
}

func NewDocumentController(router *gin.Engine, documentService documentInterfaces.DocumentServiceInterface) {
	controller := &DocumentController{documentService: documentService}

	router.POST("/api/v1/projects/:id/documents", controller.Upload())
}

func (c *DocumentController) Upload() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		projectId := validators.ValidateUUID(httpContext.Param("id"))

		fileHeader, err := httpContext.FormFile("file")
		if err != nil {
			panic(*exceptions.BadRequestException(constants.DOCUMENT_FILE_REQUIRED))
		}

		filename := filepath.Base(fileHeader.Filename)
		if filename == "." || filename == ".." || strings.ContainsAny(filename, "/\\") || filepath.Ext(filename) != ".pdf" {
			panic(*exceptions.BadRequestException(constants.DOCUMENT_FILE_INVALID_TYPE))
		}

		file, err := fileHeader.Open()
		if err != nil {
			panic(*exceptions.BadRequestException(constants.DOCUMENT_FILE_REQUIRED))
		}
		defer file.Close()

		fileBuffer, err := io.ReadAll(file)
		if err != nil {
			panic(*exceptions.BadRequestException(constants.DOCUMENT_FILE_REQUIRED))
		}

		document := c.documentService.Upload(ctx, projectId, filename, fileBuffer)
		result := dtos.DocumentResponseDtoFromEntity(document)
		response := utils.SuccessResponse(http.StatusCreated, constants.DOCUMENT_UPLOAD_SUCCESS, result)
		httpContext.JSON(http.StatusCreated, response)
	}
}
