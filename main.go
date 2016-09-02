package btax

import (
	"fmt"
	"os"
	"strconv"
)

type Reader struct{}

const (
	// FILE_TYPE_EXTENSIONS
	// 營業人進、銷資料檔
	EXT_BUSINESS_INPUT_OUTPUT_FILE = "TXT"

	// 營業人銷售額與稅額申報書檔
	EXT_營業人銷售額與稅額申報書檔 = "TET"

	// 兼營營業人採用直接扣抵法附表檔
	EXT_兼營營業人採用直接扣抵法附表檔 = "T01"

	// 營業人零稅率銷售額資料檔
	EXT_營業人零稅率銷售額資料檔 = "T02"

	// 營業人購買舊乘人小汽車及機車進項憑證明細資料檔
	EXT_營業人購買舊乘人小汽車及機車進項憑證明細資料檔 = "T03"

	// 特定營業人辦理外籍旅客現場小額退稅代墊稅款及代為繳納稅款彙總表資料檔
	EXT_特定營業人辦理外籍旅客現場小額退稅代墊稅款及代為繳納稅款彙總表資料檔 = "T04"

	// 特定營業人辦理外籍旅客現場小額退稅之代墊稅款申報扣減清冊資料檔
	EXT_特定營業人辦理外籍旅客現場小額退稅之代墊稅款申報扣減清冊資料檔 = "T05"

	// 特定營業人為外籍旅客代為繳納稅款清冊資料檔
	EXT_特定營業人為外籍旅客代為繳納稅款清冊資料檔 = "T06"

	// 兼營營業人（含採用直接扣抵法）營業稅額調整計算表資料檔
	EXT_兼營營業人含採用直接扣抵法營業稅額調整計算表資料檔 = "T07"

	// 營業人申報固定資產退稅清單資料檔
	EXT_營業人申報固定資產退稅清單資料檔 = "T08"

	// 營業人向法院或行政執行機關拍定或承受貨物申報進項憑證明細表資料檔
	EXT_營業人向法院或行政執行機關拍定或承受貨物申報進項憑證明細表資料檔 = "T09"
)

type TXTRecord struct {
	FormatCode             int    // len: 2
	TaxSerialNumber        int    // len: 9
	SerialNumber           string // len: 7
	Year                   int    // len: 3
	Month                  int    // len: 2
	BuyerVATNumber         string // len: 8
	SalerVATNumber         string // len: 8
	CollectiveNumber       int    // len: 4
	InvoiceWord            string // len: 2
	InvoiceSerialNumber    string // len: 8
	OtherCertificateNumber string // len: 10
	CustomsVATCode         string // len: 14
	TaxableBase            int    // len: 12
	SalesAmounts           int    // len: 12
	TaxType                string // len: 1
	TaxAmounts             int    // len: 10
	CreditCode             string // len: 1
	SpecialTaxRate         string // len: 1
	CollectiveNote         string // len: 1
	CustomTypeNote         string // len: 1

}

type TETRecord struct {
	FormatCode       string // len: 1
	reserve1         string // len: 2
	FileSerialNumber string // len: 8
	VATNumber        string // len: 8
	YearMonth        string // len: 5
	ReturnCode       string // len: 1
	TaxSerialNumber  string // len: 9
	reserve2         string // len: 5
	CollectiveCode   string // len: 1
	reserve3         string // len: 6
	InvoiceAmounts   int    // len: 10
	// 銷項 應稅 銷售額
	TriplicateInvoiceTaxableSaleAmounts   int // len: 12, code: 1
	CashRegisterInvoiceTaxableSaleAmounts int // len: 12, code: 5
	DuplicateInvoiceTaxableSaleAmounts    int // len: 12, code: 9
	NonInvoiceTaxableSaleAmounts          int // len: 12, code: 13
	ReturnsTaxableSaleAmounts             int // len: 12, code: 17
	TotalTaxableSaleAmounts               int // len: 12, code: 21
	// 銷項 應稅 稅額
	TriplicateInvoiceTaxableTaxAmounts   int // len: 10, code: 2
	CashRegisterInvoiceTaxableTaxAmounts int // len: 10, code: 6
	DuplicateInvoiceTaxableTaxAmounts    int // len: 10, code: 10
	NonInvoiceTaxableTaxAmounts          int // len: 10, code: 14
	ReturnsTaxableTaxAmounts             int // len: 10, code: 18
	TotalTaxableTaxAmounts               int // len: 10, code: 22

	Code82Amounts int // len: 12, code: 82
	// 銷項 零稅率銷售額
	NoCostumsOutputAmounts int // len: 12, code: 7
	reserve4               int // len: 12
	CostumsOutputAmounts   int // len: 12, code: 15
	OutputReturnAmounts    int // len: 12, code: 19
	OutputTotalAmounts     int // len: 12, code: 23

	// 銷稅 免稅 銷售額
	TriplicateInvoiceTaxFreeSaleAmounts   int // len: 12, code: 4
	CashRegisterInvoiceTaxFreeSaleAmounts int // len: 12, code: 8
	DuplicateInvoiceTaxFreeSaleAmounts    int // len: 12, code: 12
	NonInvoiceTaxFreeSaleAmounts          int // len: 12, code: 16
	ReturnsTaxFreeSaleAmounts             int // len: 12, code: 20
	TotalTaxFreeSaleAmounts               int // len: 12, code: 24

	// 特種飲食
	SpecialFood25SaleAmounts int // len: 12, code: 52
	SpecialFood25TaxAmounts  int // len: 10, code: 53
	SpecialFood15SaleAmounts int // len: 12, code: 54
	SpecialFood15TaxAmounts  int // len: 10, code: 55

	// 銀行，保險及信託投資業
	SpecialBank2SaleAmounts int // len: 12, code: 56
	SpecialBank2TaxAmounts  int // len: 10, code: 57
	SpecialBank5SaleAmounts int // len: 12, code: 58
	SpecialBank5TaxAmounts  int // len: 10, code: 59
	SpecialBank1SaleAmounts int // len: 12, code: 60
	SpecialBank1TaxAmounts  int // len: 10, code: 61

	Special0SaleAmounts        int // len: 12, code: 62
	SpecialReturnsSaleAmounts  int // len: 12, code: 63
	SpecialReturnsTaxAmounts   int // len: 10, code: 64
	SubtotalSpecialSaleAmounts int // len: 12, code: 65
	SubtotalSpecialTaxAmounts  int // len: 10, code; 66

	TotalSaleAmounts            int // len: 12, code: 25
	TotalLandSaleAmounts        int // len: 12, code: 26
	TotalFixedAssestSaleAmounts int // len: 12, code: 27

	// 應比例計算得扣抵進項金額
	InvoiceCostCreditAmounts                     int // len: 12, code: 28
	InvoiceFixedAssestSaleCreditAmount           int // len: 12, code: 30
	TriplicateInvoiceCostCreditAmounts           int // len: 12, code: 32
	TriplicateInvoiceFixedAssestSaleCreditAmount int // len: 12, code: 34
	OtherInvoiceCostCreditAmounts                int // len: 12, code: 36
	OtherInvoiceFixedAssestCreditAmounts         int // len: 12, code: 38
	ReturnsCostCreditAmounts                     int // len: 12, code: 40
	ReturnsFixedAssestCreditAmounts              int // len: 12, code: 42
	TotalCostCreditAmounts                       int // len: 12, code: 44
	TotalFixedAssestCreditAmounts                int // len: 12, code: 46

	InvoiceCostCreditTaxAmounts                     int // len: 10, code: 29
	InvoiceFixedAssestSaleCreditTaxAmount           int // len: 10, code: 31
	TriplicateInvoiceCostCreditTaxAmounts           int // len: 10, code: 33
	TriplicateInvoiceFixedAssestSaleCreditTaxAmount int // len: 10, code: 35
	OtherInvoiceCostCreditTaxAmounts                int // len: 10, code: 37
	OtherInvoiceFixedAssestCreditTaxAmounts         int // len: 10, code: 39
	ReturnsCostCreditTaxAmounts                     int // len: 10, code: 41
	ReturnsFixedAssestCreditTaxAmounts              int // len: 10, code: 43
	TotalCostCreditTaxAmounts                       int // len: 10, code: 45
	TotalFixedAssestCreditTaxAmounts                int // len: 10, code: 47

	TotalCostIncludeNonDeductibleCreditAmounts int // len: 12, code: 48
	TotalFixedAssestNonDeductibleCreditAmounts int // len: 12, code: 49

	NonDeductibleRatio             int // len: 3, code: 50
	DeductibleCreditTaxAmounts     int // len: 10, code: 51
	CostumsCostAmounts             int // len: 12, code: 78
	CostumsFixedAssestAmounts      int // len: 12, code: 80
	ImportTaxExemptGoodsAmounts    int // len: 12, code: 73
	ImportTaxExemptServicesAmounts int // len: 12, code: 74
	CostumsCostTaxAmounts          int // len: 10, code: 79
	CostumsFixedAssestTaxAmounts   int // len: 10, code: 81
	ImportRatioServicesTaxAmounts  int // len: 10, code: 75
	reserve5                       int // len: 10
	reserve6                       int // len: 10
	ImportServicesTaxAmounts       int // len: 10, code: 76

	TotalSaleTaxAmounts                 int // len: 10, code: 101
	reserve7                            int // len: 10
	TotalImportServicesTaxAmounts       int // len: 10, code: 103
	TotalSpecialTaxAmounts              int // len: 10, code: 104
	AdjustmentTaxAmounts                int // len: 10, code: 105
	SubTotal106                         int // len: 10, code: 106
	TotalBusinessTaxPaidAmounts         int // len: 10, code: 107
	LastPeriodExcessTaxPaidAmounts      int // len: 10, code: 108
	Field109                            int // len: 10, code: 109
	SubTotal110                         int // len: 10, code: 110
	ThisPeriodTaxPayableAmounts         int // len: 10, code: 111
	ThisPeriodExcessTaxPaidAmounts      int // len: 10, code: 112
	ReturnableTaxAmounts                int // len: 10, code: 113
	ThisPeriodReturnTaxAmounts          int // len: 10, code: 114
	ThisPeriodTotalExcessTaxPaidAmounts int // len: 10, code: 115

	reserve8               string // len: 23
	ReportType             string // len: 1
	reserve9               string // len: 1
	County                 string // len: 1
	reserve10              string // len: 6
	BySelfOrTheOther       string // len: 1
	ReporterID             string // len: 10
	ReporterPhoneLocalCode string // len: 4
	ReporterPhoneNumber    string // len: 11
	ReporterPhoneExt       string // len: 5
	ReporterSerialNumber   string // len: 50
	reserve11              string // len: 92

}

// http://gazette.nat.gov.tw/EG_FileManager/eguploadpub/eg017152/ch04/type2/gov30/num3/images/Eg01.htm
const (
	TaipeiCity       = "A"
	TaichungCity     = "B"
	KeelungCity      = "C"
	TainanCity       = "D"
	KaohsiungCity    = "E"
	NewTaipeiCity    = "F"
	YilanCounty      = "G"
	TaoyuanCounty    = "H"
	ChiayiCity       = "I"
	HsinchuCounty    = "J"
	MiaoliCounty     = "K"
	NantouCounty     = "M"
	ChunghuaCounty   = "N"
	HsinchuCity      = "O"
	YunlinCounty     = "P"
	ChiayiCounty     = "Q"
	PingtungCounty   = "T"
	HualienCounty    = "U"
	TaitungCounty    = "V"
	KinmenCounty     = "W"
	PenghuCounty     = "X"
	LienchiangCounty = "Z"
)

func ParseTXTFile(file *os.File) []*TXTRecord {
	// len: 81
	rowLen := 81
	buf := make([]byte, rowLen+2, rowLen+2)
	ret := make([]*TXTRecord, 0)
	for {
		n, _ := file.Read(buf)
		if n == 0 {
			break
		}
		r := parseTXTRecord(buf)
		ret = append(ret, r)
	}
	return ret

}

func parseTXTRecord(b []byte) *TXTRecord {
	record := &TXTRecord{}
	record.FormatCode, _ = strconv.Atoi(string(b[0:2]))
	record.TaxSerialNumber, _ = strconv.Atoi(string(b[2:11]))
	record.SerialNumber = string(b[11:18])
	record.Year, _ = strconv.Atoi(string(b[18:21]))
	record.Month, _ = strconv.Atoi(string(b[21:23]))

	record.BuyerVATNumber = string(b[23:31])
	// 31 ~ 49
	record.SalerVATNumber = string(b[31:39])

	record.InvoiceWord = string(b[39:41])
	record.InvoiceSerialNumber = string(b[41:49])
	// 49 ~ 61
	record.SalesAmounts, _ = strconv.Atoi(string(b[49:61]))
	record.TaxableBase, _ = strconv.Atoi(string(b[49:61]))

	record.TaxType = string(b[61:62])
	record.TaxAmounts, _ = strconv.Atoi(string(b[62:72]))
	record.CreditCode = string(b[72:73])

	// record.CollectiveNumber, _ = strconv.Atoi(string(b[31:35]))
	return record
}

func marshalTXTRecord(r *TXTRecord) (b []byte) {
	b = make([]byte, 81+2, 81+2)
	copy(b[0:31],
		fmt.Sprintf("%02d%09d%07s%03d%02d% 8s",
			r.FormatCode,
			r.TaxSerialNumber,
			r.SerialNumber,
			r.Year,
			r.Month,
			r.BuyerVATNumber,
		))

	if r.FormatCode == 22 || r.FormatCode == 25 || r.FormatCode == 31 {
		copy(b[31:73],
			fmt.Sprintf("% 8s%2s%08s%012d%1s%010d%1s",
				r.SalerVATNumber,
				r.InvoiceWord,
				r.InvoiceSerialNumber,
				r.SalesAmounts,
				r.TaxType,
				r.TaxAmounts,
				r.CreditCode,
			))
		copy(b[73:81], fmt.Sprintf("% 8s", ""))
		b[81] = 0x0d
		b[82] = 0x0a
	}

	return b
}
