package hubspot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
)

type CrmImportConfig struct {
	Name                    string                `json:"name"`
	MarketableContactImport bool                  `json:"marketableContactImport"`
	ImportOperations        map[string]string     `json:"importOperations"`
	Files                   []CrmImportFileConfig `json:"files"`
}

type CrmImportFilePageConfig struct {
	HasHeader      bool                     `json:"hasHeader"`
	ColumnMappings []CrmImportColumnMapping `json:"columnMappings"`
}

type CrmImportFileConfig struct {
	FileName       string                  `json:"fileName"`
	FileFormat     string                  `json:"fileFormat"`
	DateFormat     string                  `json:"dateFormat"`
	FileImportPage CrmImportFilePageConfig `json:"fileImportPage"`
	// Data is the CSV or Spreadsheet data for this file.
	Data io.Reader `json:"-"`
}

type CrmImportColumnMapping struct {
	ColumnObjectTypeId string `json:"columnObjectTypeId"`
	ColumnName         string `json:"columnName"`
	PropertyName       string `json:"propertyName"`
	IdColumnType       string `json:"idColumnType,omitempty"`
}

func addJSONtoImportMultipart(writer *multipart.Writer, importRequest *CrmImportConfig) error {
	data, err := json.Marshal(importRequest)
	if err != nil {
		return err
	}
	header := textproto.MIMEHeader{}
	header.Set("Content-Disposition", "form-data; name=\"importRequest\"")
	part, err := writer.CreatePart(header)
	if err != nil {
		return err
	}
	if _, err := part.Write(data); err != nil {
		return err
	}
	return nil
}

func addFilesToImportMultipart(writer *multipart.Writer, importRequest *CrmImportConfig) error {
	for _, fileDef := range importRequest.Files {
		csvHeader := textproto.MIMEHeader{}
		csvHeader.Set("Content-Disposition", fmt.Sprintf("form-data; name=\"files\"; filename=\"%s\"", fileDef.FileName))
		csvPart, err := writer.CreatePart(csvHeader)
		if err != nil {
			return err
		}
		fileData, err := io.ReadAll(fileDef.Data)
		if err != nil {
			return err
		}
		if _, err := csvPart.Write(fileData); err != nil {
			return err
		}
	}
	return nil
}

func (s *CrmImportsServiceOp) Start(importRequest *CrmImportConfig) (interface{}, error) {
	resource := make(map[string]interface{})

	// Body is our final result that we pass to postMultipart
	body := &bytes.Buffer{}

	// Write the importRequest to the multipart message
	writer := multipart.NewWriter(body)

	// Write a part for the JSON metadata.
	if err := addJSONtoImportMultipart(writer, importRequest); err != nil {
		return nil, err
	}

	// Write file data to multipart.
	if err := addFilesToImportMultipart(writer, importRequest); err != nil {
		return nil, err
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	if err := s.client.PostMultipart(s.crmImportsPath, writer.Boundary(), body.Bytes(), &resource); err != nil {
		return nil, err
	}
	return resource, nil
}
