package btax

import (
	"testing"
)

func TestParseTXTRecord(t *testing.T) {
	cases := []struct {
		in   []byte
		want *TXTRecord
	}{
		{
			[]byte("221234567890000001104095483479512345678RY12345678000000000080100000000001        \n"),
			&TXTRecord{
				FormatCode:          22,
				TaxSerialNumber:     123456789,
				SerialNumber:        "0000001",
				Year:                104,
				Month:               9,
				BuyerVATNumber:      "54834795",
				SalerVATNumber:      "12345678",
				InvoiceWord:         "RY",
				InvoiceSerialNumber: "12345678",
				SalesAmounts:        80,
				TaxableBase:         80,
				TaxAmounts:          0,
				TaxType:             "1",
				CreditCode:          "1",
			},
		},
		{
			[]byte("259876543210000005104095483479587654321YN12340000000000000941100000000471        \n"),
			&TXTRecord{
				FormatCode:          25,
				TaxSerialNumber:     987654321,
				SerialNumber:        "0000005",
				Year:                104,
				Month:               9,
				BuyerVATNumber:      "54834795",
				SalerVATNumber:      "87654321",
				InvoiceWord:         "YN",
				InvoiceSerialNumber: "12340000",
				SalesAmounts:        941,
				TaxableBase:         941,
				TaxAmounts:          47,
				TaxType:             "1",
				CreditCode:          "1",
			},
		},
		{
			[]byte("31987654321000001410410        54834795RN12340000000000000000F0000000000         \n"),
			&TXTRecord{
				FormatCode:          31,
				TaxSerialNumber:     987654321,
				SerialNumber:        "0000014",
				Year:                104,
				Month:               10,
				BuyerVATNumber:      "        ",
				SalerVATNumber:      "54834795",
				InvoiceWord:         "RN",
				InvoiceSerialNumber: "12340000",
				SalesAmounts:        0,
				TaxableBase:         0,
				TaxAmounts:          0,
				TaxType:             "F",
				CreditCode:          " ",
			},
		},
	}

	for _, c := range cases {
		actual := parseTXTRecord(c.in)
		if actual.FormatCode != c.want.FormatCode {
			t.Errorf("format code incorrect, parse %v, want %v, got %v",
				string(c.in), c.want.FormatCode, actual.FormatCode)
		}

		if actual.TaxSerialNumber != c.want.TaxSerialNumber {
			t.Errorf("tax serial number incorrect, parse %v, want %v, got %v",
				string(c.in), c.want.TaxSerialNumber, actual.TaxSerialNumber)
		}

		if actual.SerialNumber != c.want.SerialNumber {
			t.Errorf("serial number incorrect, parse %v, want %v, got %v",
				string(c.in), c.want.SerialNumber, actual.SerialNumber)
		}

		if actual.Year != c.want.Year {
			t.Errorf("year incorrect, parse %v, want %v, got %v",
				string(c.in), c.want.Year, actual.Year)
		}

		if actual.Month != c.want.Month {
			t.Errorf("month incorrect, parse %v, want %v, got %v",
				string(c.in), c.want.Month, actual.Month)
		}

		if actual.BuyerVATNumber != c.want.BuyerVATNumber {
			t.Errorf("buyer vat number incorrect, parse %v, want %v, got %v",
				string(c.in), c.want.BuyerVATNumber, actual.BuyerVATNumber)
		}

		if actual.FormatCode == 22 || actual.FormatCode == 25 || actual.FormatCode == 31 {

			if actual.SalerVATNumber != c.want.SalerVATNumber {
				t.Errorf("saler vat number incorrect, parse %v, want %v, got %v",
					string(c.in), c.want.SalerVATNumber, actual.SalerVATNumber)
			}
			if actual.InvoiceWord != c.want.InvoiceWord {
				t.Errorf("invoice word incorrect, parse %v, want %v, got %v",
					string(c.in), c.want.InvoiceWord, actual.InvoiceWord)
			}
			if actual.InvoiceSerialNumber != c.want.InvoiceSerialNumber {
				t.Errorf("invoice serial number incorrect, parse %v, want %v, got %v",
					string(c.in), c.want.InvoiceSerialNumber, actual.InvoiceSerialNumber)
			}
			if actual.SalesAmounts != c.want.SalesAmounts {
				t.Errorf("sales amounts incorrect, parse %v, want %v, got %v",
					string(c.in), c.want.SalesAmounts, actual.SalesAmounts)
			}
			if actual.TaxableBase != c.want.TaxableBase {
				t.Errorf("taxable base incorrect, parse %v, want %v, got %v",
					string(c.in), c.want.TaxableBase, actual.TaxableBase)
			}

		}
		if actual.TaxType != c.want.TaxType {
			t.Errorf("tax type incorrect, parse %v, want %v, got %v",
				string(c.in), c.want.TaxType, actual.TaxType)
		}
		if actual.TaxAmounts != c.want.TaxAmounts {
			t.Errorf("tax amounts incorrect, parse %v, want %v, got %v",
				string(c.in), c.want.TaxAmounts, actual.TaxAmounts)
		}
		if actual.CreditCode != c.want.CreditCode {
			t.Errorf("credit code incorrect, parse %v, want %v, got %v",
				string(c.in), c.want.CreditCode, actual.CreditCode)
		}

	}

}

func TestMarshalTXTRecord(t *testing.T) {
	cases := []struct {
		in   *TXTRecord
		want []byte
	}{
		{
			want: append([]byte("221234567890000001104095483479512345678RY12345678000000000080100000000001        "), 0x0d, 0x0a),
			in: &TXTRecord{
				FormatCode:          22,
				TaxSerialNumber:     123456789,
				SerialNumber:        "0000001",
				Year:                104,
				Month:               9,
				BuyerVATNumber:      "54834795",
				SalerVATNumber:      "12345678",
				InvoiceWord:         "RY",
				InvoiceSerialNumber: "12345678",
				SalesAmounts:        80,
				TaxableBase:         80,
				TaxAmounts:          0,
				TaxType:             "1",
				CreditCode:          "1",
			},
		},
		{
			want: append([]byte("259876543210000005104095483479587654321YN12340000000000000941100000000471        "), 0x0d, 0x0a),
			in: &TXTRecord{
				FormatCode:          25,
				TaxSerialNumber:     987654321,
				SerialNumber:        "0000005",
				Year:                104,
				Month:               9,
				BuyerVATNumber:      "54834795",
				SalerVATNumber:      "87654321",
				InvoiceWord:         "YN",
				InvoiceSerialNumber: "12340000",
				SalesAmounts:        941,
				TaxableBase:         941,
				TaxAmounts:          47,
				TaxType:             "1",
				CreditCode:          "1",
			},
		},
		{
			want: append([]byte("31987654321000001410410        54834795RN12340000000000000000F0000000000         "), 0x0d, 0x0a),
			in: &TXTRecord{
				FormatCode:          31,
				TaxSerialNumber:     987654321,
				SerialNumber:        "0000014",
				Year:                104,
				Month:               10,
				BuyerVATNumber:      "        ",
				SalerVATNumber:      "54834795",
				InvoiceWord:         "RN",
				InvoiceSerialNumber: "12340000",
				SalesAmounts:        0,
				TaxableBase:         0,
				TaxAmounts:          0,
				TaxType:             "F",
				CreditCode:          " ",
			},
		},
	}

	for i, c := range cases {
		actual := marshalTXTRecord(c.in)

		if string(actual) != string(c.want) {
			t.Errorf("marshal test case %d error, want %v(len: %d), but got %v(len: %d)",
				i, string(c.want), len(c.want), string(actual), len(actual))
		}

	}

}
