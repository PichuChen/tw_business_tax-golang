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

func ParseTXTFile() {
	file, err := os.Open("54834795.TXT")
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}
	// len: 81
	rowLen := 81
	buf := make([]byte, rowLen+2, rowLen+2)
	for {
		n, _ := file.Read(buf)
		if n == 0 {
			break
		}
		fmt.Println(n)
		fmt.Println(string(buf))
		r := parseTXTRecord(buf)
		fmt.Println(r)
		fmt.Println(string(marshalTXTRecord(r)))
	}

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
