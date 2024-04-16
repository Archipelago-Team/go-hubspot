package hubspot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
)

type CrmExportConfig struct {
	Name                    string                `json:"name"`
	MarketableContactExport bool                  `json:"marketableContactExport"`
	ExportOperations        map[string]string     `json:"exportOperations"`
	Files                   []CrmExportFileConfig `json:"files"`
}

type CrmExportFilePageConfig struct {
	HasHeader      bool                     `json:"hasHeader"`
	ColumnMappings []CrmExportColumnMapping `json:"columnMappings"`
}

type CrmExportFileConfig struct {
	FileName       string                  `json:"fileName"`
	FileFormat     string                  `json:"fileFormat"`
	DateFormat     string                  `json:"dateFormat"`
	FileExportPage CrmExportFilePageConfig `json:"fileExportPage"`
	// Data is the CSV or Spreadsheet data for this file.
	Data io.Reader `json:"-"`
}

type CrmExportColumnMapping struct {
	ColumnObjectTypeId string `json:"columnObjectTypeId"`
	ColumnName         string `json:"columnName"`
	PropertyName       string `json:"propertyName"`
	IdColumnType       string `json:"idColumnType,omitempty"`
}

func (s *CrmExportsServiceOp) Start(exportRequest *CrmExportConfig) (interface{}, error) {
	resource := make(map[string]interface{})

	// Body is our final result that we pass to postMultipart
	body := &bytes.Buffer{}

	// Write the exportRequest to the multipart message
	writer := multipart.NewWriter(body)

	// Write a part for the JSON metadata.
	if err := addJSONtoExportMultipart(writer, exportRequest); err != nil {
		return nil, err
	}

	// Write file data to multipart.
	if err := addFilesToExportMultipart(writer, exportRequest); err != nil {
		return nil, err
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	if err := s.client.PostMultipart(s.crmExportsPath, writer.Boundary(), body.Bytes(), &resource); err != nil {
		return nil, err
	}
	return resource, nil
}

func addJSONtoExportMultipart(writer *multipart.Writer, exportRequest *CrmExportConfig) error {
	data, err := json.Marshal(exportRequest)
	if err != nil {
		return err
	}
	header := textproto.MIMEHeader{}
	header.Set("Content-Disposition", "form-data; name=\"exportRequest\"")
	part, err := writer.CreatePart(header)
	if err != nil {
		return err
	}
	if _, err := part.Write(data); err != nil {
		return err
	}
	return nil
}

func addFilesToExportMultipart(writer *multipart.Writer, exportRequest *CrmExportConfig) error {
	for _, fileDef := range exportRequest.Files {
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
