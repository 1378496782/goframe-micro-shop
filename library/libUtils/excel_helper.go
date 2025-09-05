package libUtils

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
)

const (
	ExcelHelperRowLimit = 65536
	ExcelHelperColLimit = 256
)

type ExcelHelper struct{}

type ExcelHelperFile struct {
	*excelize.File
	err error
}

func (helperFile *ExcelHelperFile) MapToExcel(sheet string, cells map[string]interface{}) *ExcelHelperFile {
	if helperFile.err != nil {
		return helperFile
	}
	for axis, item := range cells {
		err := helperFile.File.SetCellValue(sheet, axis, item)
		if err != nil {
			helperFile.err = err
			return helperFile
		}
	}
	return helperFile
}

func (helperFile *ExcelHelperFile) SetCellBorder(sheet, hcell, vcell string) *ExcelHelperFile {
	style, err := helperFile.File.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 1},
		},
	})
	if err != nil {
		helperFile.err = err
		return helperFile
	}
	_ = helperFile.File.SetCellStyle(sheet, hcell, vcell, style)
	return helperFile
}

func (helperFile *ExcelHelperFile) ArrToExcel(sheet string, axis string, rows [][]interface{}) *ExcelHelperFile {
	if helperFile.err != nil {
		return helperFile
	}

	rowNum := len(rows)
	if rowNum == 0 {
		helperFile.err = errors.New("没有需要导出的数据")
		return helperFile
	}
	if rowNum > ExcelHelperRowLimit {
		helperFile.err = errors.New(fmt.Sprintf("导出数据行数超过%d", ExcelHelperRowLimit))
		return helperFile
	}
	colNum := len(rows[0])
	if colNum > ExcelHelperColLimit {
		helperFile.err = errors.New(fmt.Sprintf("导出数据列数超过%d", ExcelHelperColLimit))
		return helperFile
	}

	axisCol, axisRow, err := excelize.SplitCellName(axis)
	if err != nil {
		helperFile.err = err
		return helperFile
	}

	for index, item := range rows {
		cell, _ := excelize.JoinCellName(axisCol, index+axisRow)
		err = helperFile.File.SetSheetRow(sheet, cell, &item)
		if err != nil {
			helperFile.err = err
			return helperFile
		}
	}
	return helperFile
}

func (helperFile *ExcelHelperFile) SaveByPath(path string) (savePath string, err error) {
	if helperFile.err != nil {
		return "", helperFile.err
	}
	path = path + ".xlsx"
	err = helperFile.File.SaveAs(path)
	if err != nil {
		return
	}
	return path, nil
}

func (excel *ExcelHelper) OpenFile(tplPath string) (file *ExcelHelperFile, err error) {
	temp, err := excelize.OpenFile(tplPath)
	if err != nil {
		return
	}
	file = &ExcelHelperFile{
		File: temp,
	}
	return
}

func (excel *ExcelHelper) CreateFile() (file *ExcelHelperFile) {
	f := excelize.NewFile()
	file = &ExcelHelperFile{
		File: f,
	}
	return
}
