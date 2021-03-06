// xslx is a package designed to help with reading data from
// spreadsheets stored in the XLSX format used in recent versions of
// Microsoft's Excel spreadsheet.
//
// For a concise example of how to use this library why not check out
// the source for xlsx2csv here: https://github.com/tealeg/xlsx2csv

package xlsx

import (
	"encoding/xml"
	"strconv"
	"strings"
)

// xlsxStyle directly maps the styleSheet element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxStyleSheet struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/spreadsheetml/2006/main styleSheet"`

	Fonts        xlsxFonts        `xml:"fonts,omitempty"`
	Fills        xlsxFills        `xml:"fills,omitempty"`
	Borders      xlsxBorders      `xml:"borders,omitempty"`
	CellStyleXfs xlsxCellStyleXfs `xml:"cellStyleXfs,omitempty"`
	CellXfs      xlsxCellXfs      `xml:"cellXfs,omitempty"`
	NumFmts      xlsxNumFmts      `xml:"numFmts,omitempty"`
}

// xlsxNumFmts directly maps the numFmts element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxNumFmts struct {
	Count  int          `xml:"count,attr"`
	NumFmt []xlsxNumFmt `xml:"numFmt,omitempty"`
}

// xlsxNumFmt directly maps the numFmt element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxNumFmt struct {
	NumFmtId   int    `xml:"numFmtId,omitempty"`
	FormatCode string `xml:"formatCode,omitempty"`
}

// xlsxFonts directly maps the fonts element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxFonts struct {
	Count int        `xml:"count,attr"`
	Font  []xlsxFont `xml:"font,omitempty"`
}

// xlsxFont directly maps the font element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxFont struct {
	Sz      xlsxVal   `xml:"sz,omitempty"`
	Name    xlsxVal   `xml:"name,omitempty"`
	Family  xlsxVal   `xml:"family,omitempty"`
	Charset xlsxVal   `xml:"charset,omitempty"`
	Color   xlsxColor `xml:"color,omitempty"`
}

// xlsxVal directly maps the val element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxVal struct {
	Val string `xml:"val,attr,omitempty"`
}

// xlsxFills directly maps the fills element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxFills struct {
	Count int        `xml:"count,attr"`
	Fill  []xlsxFill `xml:"fill,omitempty"`
}

// xlsxFill directly maps the fill element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxFill struct {
	PatternFill xlsxPatternFill `xml:"patternFill,omitempty"`
}

// xlsxPatternFill directly maps the patternFill element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxPatternFill struct {
	PatternType string    `xml:"patternType,attr,omitempty"`
	FgColor     xlsxColor `xml:"fgColor,omitempty"`
	BgColor     xlsxColor `xml:"bgColor,omitempty"`
}

// xlsxColor is a common mapping used for both the fgColor and bgColor
// elements in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxColor struct {
	RGB string `xml:"rgb,attr,omitempty"`
}

// xlsxBorders directly maps the borders element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxBorders struct {
	Count  int          `xml:"count,attr"`
	Border []xlsxBorder `xml:"border,omitempty"`
}

// xlsxBorder directly maps the border element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxBorder struct {
	Left   xlsxLine `xml:"left,omitempty"`
	Right  xlsxLine `xml:"right,omitempty"`
	Top    xlsxLine `xml:"top,omitempty"`
	Bottom xlsxLine `xml:"bottom,omitempty"`
}

// xlsxLine directly maps the line style element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxLine struct {
	Style string `xml:"style,attr,omitempty"`
}

// xlsxCellStyleXfs directly maps the cellStyleXfs element in the
// namespace http://schemas.openxmlformats.org/spreadsheetml/2006/main
// - currently I have not checked it for completeness - it does as
// much as I need.
type xlsxCellStyleXfs struct {
	Count int      `xml:"count,attr"`
	Xf    []xlsxXf `xml:"xf,omitempty"`
}

// xlsxCellXfs directly maps the cellXfs element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxCellXfs struct {
	Count int      `xml:"count,attr"`
	Xf    []xlsxXf `xml:"xf,omitempty"`
}

// xlsxXf directly maps the xf element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type xlsxXf struct {
	ApplyAlignment  bool          `xml:"applyAlignment,attr"`
	ApplyBorder     bool          `xml:"applyBorder,attr"`
	ApplyFont       bool          `xml:"applyFont,attr"`
	ApplyFill       bool          `xml:"applyFill,attr"`
	ApplyProtection bool          `xml:"applyProtection,attr"`
	BorderId        int           `xml:"borderId,attr"`
	FillId          int           `xml:"fillId,attr"`
	FontId          int           `xml:"fontId,attr"`
	NumFmtId        int           `xml:"numFmtId,attr"`
	alignment       xlsxAlignment `xml:"alignment"`
}

type xlsxAlignment struct {
	Horizontal   string `xml:"horizontal,attr"`
	Indent       int    `xml:"indent,attr"`
	ShrinkToFit  bool   `xml:"shrinkToFit,attr"`
	TextRotation int    `xml:"textRotation,attr"`
	Vertical     string `xml:"vertical,attr"`
	WrapText     bool   `xml:"wrapText,attr"`
}

func (styles *xlsxStyleSheet) getStyle(styleIndex int) (style Style) {
	var styleXf xlsxXf
	style = Style{}
	style.Border = Border{}
	style.Fill = Fill{}
	style.Font = Font{}

	xfCount := styles.CellXfs.Count
	if styleIndex > -1 && xfCount > 0 && styleIndex <= xfCount {
		xf := styles.CellXfs.Xf[styleIndex]

		// Google docs can produce output that has fewer
		// CellStyleXfs than CellXfs - this copes with that.
		if styleIndex < styles.CellStyleXfs.Count {
			styleXf = styles.CellStyleXfs.Xf[styleIndex]
		} else {
			styleXf = xlsxXf{}
		}

		style.ApplyBorder = xf.ApplyBorder || styleXf.ApplyBorder
		style.ApplyFill = xf.ApplyFill || styleXf.ApplyFill
		style.ApplyFont = xf.ApplyFont || styleXf.ApplyFont

		if xf.BorderId > -1 && xf.BorderId < styles.Borders.Count {
			style.Border.Left = styles.Borders.Border[xf.BorderId].Left.Style
			style.Border.Right = styles.Borders.Border[xf.BorderId].Right.Style
			style.Border.Top = styles.Borders.Border[xf.BorderId].Top.Style
			style.Border.Bottom = styles.Borders.Border[xf.BorderId].Bottom.Style
		}

		if xf.FillId > -1 && xf.FillId < styles.Fills.Count {
			xFill := styles.Fills.Fill[xf.FillId]
			style.Fill.PatternType = xFill.PatternFill.PatternType
			style.Fill.FgColor = xFill.PatternFill.FgColor.RGB
			style.Fill.BgColor = xFill.PatternFill.BgColor.RGB
		}

		if xf.FontId > -1 && xf.FontId < styles.Fonts.Count {
			xfont := styles.Fonts.Font[xf.FontId]
			style.Font.Size, _ = strconv.Atoi(xfont.Sz.Val)
			style.Font.Name = xfont.Name.Val
			style.Font.Family, _ = strconv.Atoi(xfont.Family.Val)
			style.Font.Charset, _ = strconv.Atoi(xfont.Charset.Val)
		}
	}
	return style

}

func (styles *xlsxStyleSheet) getNumberFormat(styleIndex int, numFmtRefTable map[int]xlsxNumFmt) string {
	if styles.CellXfs.Xf == nil {
		return ""
	}
	var numberFormat string = ""
	if styleIndex > -1 && styleIndex <= styles.CellXfs.Count {
		xf := styles.CellXfs.Xf[styleIndex]
		numFmt := numFmtRefTable[xf.NumFmtId]
		numberFormat = numFmt.FormatCode
	}
	return strings.ToLower(numberFormat)
}

func (styles *xlsxStyleSheet) addFont(xFont xlsxFont) (index int) {
	styles.Fonts.Font = append(styles.Fonts.Font, xFont)
	index = styles.Fonts.Count
	styles.Fonts.Count += 1
	return
}

func (styles *xlsxStyleSheet) addFill(xFill xlsxFill) (index int) {
	styles.Fills.Fill = append(styles.Fills.Fill, xFill)
	index = styles.Fills.Count
	styles.Fills.Count += 1
	return
}

func (styles *xlsxStyleSheet) addBorder(xBorder xlsxBorder) (index int) {
	styles.Borders.Border = append(styles.Borders.Border, xBorder)
	index = styles.Borders.Count
	styles.Borders.Count += 1
	return
}

func (styles *xlsxStyleSheet) addCellStyleXf(xCellStyleXf xlsxXf) (index int) {
	styles.CellStyleXfs.Xf = append(styles.CellStyleXfs.Xf, xCellStyleXf)
	index = styles.CellStyleXfs.Count
	styles.CellStyleXfs.Count += 1
	return
}

func (styles *xlsxStyleSheet) addCellXf(xCellXf xlsxXf) (index int) {
	styles.CellXfs.Xf = append(styles.CellXfs.Xf, xCellXf)
	index = styles.CellXfs.Count
	styles.CellXfs.Count += 1
	return
}
