package Ecpayby77

import (
	"crypto/sha256"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func generateMerchantTradeNo(MemberId string) (No string) {
	time := time.Now().UnixNano() / 1e6
	timeString := strconv.FormatInt(time, 16)

	sMemberId := MemberId

	No = timeString + sMemberId
	Count := strings.Count(No, "")

	fmt.Print("string Count=", Count, " MerchantTradeNo=", No, "\n")
	return No
}

func FormUrlEncode(s string) string {
	s = url.QueryEscape(s)
	//s = strings.ReplaceAll(s, "%2d", "-")
	//s = strings.ReplaceAll(s, "%5f", "_")
	//s = strings.ReplaceAll(s, "%2e", ".")
	s = strings.ReplaceAll(s, "%21", "!")
	s = strings.ReplaceAll(s, "%2A", "*")
	s = strings.ReplaceAll(s, "%28", "(")
	s = strings.ReplaceAll(s, "%29", ")")
	return s
}

type EcPayParm struct {
	Parameter string
	Value     string
}

func EcpayCheckMacValue(RecvCheckMacValue string, slice []EcPayParm, HashKey string, HashIV string) (IsCheckOK bool) {

	CheckMacValue := ""
	for i := 0; i < len(slice); i++ {
		if slice[i].Parameter == "CheckMacValue" {
			continue
		}

		if CheckMacValue != "" {
			CheckMacValue = CheckMacValue + "&"
		}

		CheckMacValue = CheckMacValue + slice[i].Parameter + "=" + slice[i].Value
	}

	CheckMacValue = "HashKey=" + HashKey + "&" + CheckMacValue + "&HashIV=" + HashIV
	fmt.Print("\nCheckMacValue=", CheckMacValue)

	CheckMacValue = FormUrlEncode(CheckMacValue)
	fmt.Print("\nCheckMacValue=", CheckMacValue)

	CheckMacValue = strings.ToLower(CheckMacValue)
	fmt.Print("\nCheckMacValue=", CheckMacValue)

	sum := sha256.Sum256([]byte(CheckMacValue))
	fmt.Printf("\n%x", sum)

	CheckMacValue = fmt.Sprintf("%x", sum)
	fmt.Print("\nCheckMacValue=", CheckMacValue)

	CheckMacValue = strings.ToUpper(CheckMacValue)
	fmt.Print("\nCheckMacValue=", CheckMacValue)

	if RecvCheckMacValue == CheckMacValue {
		return true
	}
	return false
}

/*
綠界的參數請參考 https://www.ecpay.com.tw/Content/files/ecpay_011.pdf
*/

func SendPostToEcPayPeriod(CustomField1 string, CustomField2 string, CustomField3 string, MerchantID string, ITotalAmount int, TradeDesc string, ItemName string, ReturnURL string, ClientBackURL string, PeriodReturnURL string, CustomerIdentifier string, CustomerEmail string, CarruerType string, CarruerNum string, Donation string, LoveCode string, Print string, InvoiceItemName string, InvoiceItemCount string, InvoiceItemWord string, InvoiceItemPrice string, CustomerName string, CustomerAddr string, HashKey string, HashIV string) (CheckMacValue string, slice []EcPayParm) {
	MerchantTradeNo := generateMerchantTradeNo(CustomField1)
	MerchantTradeDate := time.Now().Format("2006/01/02 15:04:05")
	PaymentType := "aio"
	RelateNumber := MerchantTradeNo
	ChoosePayment := "Credit"
	ItemURL := ClientBackURL
	InvoiceMark := "Y"
	EncryptType := "1"
	PeriodAmount := strconv.Itoa(ITotalAmount)
	TotalAmount := strconv.Itoa(ITotalAmount)
	PeriodType := "D"
	Frequency := strconv.Itoa(30)
	ExecTimes := strconv.Itoa(999)
	TaxType := "1"
	DelayDay := "0"
	InvType := "07"

	/*
		TradeDesc = FormUrlEncode(TradeDesc)
		CustomerEmail = FormUrlEncode(CustomerEmail)
		InvoiceItemName = FormUrlEncode(InvoiceItemName)
		InvoiceItemWord = FormUrlEncode(InvoiceItemWord)
	*/

	slice = []EcPayParm{}
	//按照字母排列 //

	slice = append(slice, EcPayParm{"CarruerNum", CarruerNum})
	slice = append(slice, EcPayParm{"CarruerType", CarruerType})
	slice = append(slice, EcPayParm{"ChoosePayment", ChoosePayment})
	slice = append(slice, EcPayParm{"ClientBackURL", ClientBackURL})
	slice = append(slice, EcPayParm{"CustomerAddr", CustomerAddr})
	slice = append(slice, EcPayParm{"CustomerEmail", CustomerEmail})
	slice = append(slice, EcPayParm{"CustomerIdentifier", CustomerIdentifier})
	slice = append(slice, EcPayParm{"CustomerName", CustomerName})
	slice = append(slice, EcPayParm{"CustomField1", CustomField1})
	slice = append(slice, EcPayParm{"CustomField2", CustomField2})
	slice = append(slice, EcPayParm{"CustomField3", CustomField3})

	slice = append(slice, EcPayParm{"DelayDay", DelayDay})
	slice = append(slice, EcPayParm{"Donation", Donation})

	slice = append(slice, EcPayParm{"EncryptType", EncryptType})
	slice = append(slice, EcPayParm{"ExecTimes", ExecTimes})

	slice = append(slice, EcPayParm{"Frequency", Frequency})

	slice = append(slice, EcPayParm{"InvoiceItemCount", InvoiceItemCount})
	slice = append(slice, EcPayParm{"InvoiceItemName", InvoiceItemName})
	slice = append(slice, EcPayParm{"InvoiceItemPrice", InvoiceItemPrice})
	slice = append(slice, EcPayParm{"InvoiceItemWord", InvoiceItemWord})
	slice = append(slice, EcPayParm{"InvoiceMark", InvoiceMark})
	slice = append(slice, EcPayParm{"InvType", InvType})
	slice = append(slice, EcPayParm{"ItemName", ItemName})
	slice = append(slice, EcPayParm{"ItemURL", ItemURL})

	slice = append(slice, EcPayParm{"LoveCode", LoveCode})

	slice = append(slice, EcPayParm{"MerchantID", MerchantID})
	slice = append(slice, EcPayParm{"MerchantTradeDate", MerchantTradeDate})
	slice = append(slice, EcPayParm{"MerchantTradeNo", MerchantTradeNo})

	slice = append(slice, EcPayParm{"PaymentType", PaymentType})
	slice = append(slice, EcPayParm{"PeriodAmount", PeriodAmount})
	slice = append(slice, EcPayParm{"PeriodReturnURL", PeriodReturnURL})
	slice = append(slice, EcPayParm{"PeriodType", PeriodType})
	slice = append(slice, EcPayParm{"Print", Print})

	slice = append(slice, EcPayParm{"RelateNumber", RelateNumber})
	slice = append(slice, EcPayParm{"ReturnURL", ReturnURL})

	slice = append(slice, EcPayParm{"TaxType", TaxType})
	slice = append(slice, EcPayParm{"TotalAmount", TotalAmount})
	slice = append(slice, EcPayParm{"TradeDesc", TradeDesc})

	SliceFinal := []EcPayParm{}

	for i := 0; i < len(slice); i++ {
		if slice[i].Value == "" {
			continue
		}
		SliceFinal = append(SliceFinal, slice[i])
	}

	CheckMacValue = ""

	for i := 0; i < len(SliceFinal); i++ {
		if SliceFinal[i].Value == "" {
			continue
		}

		if CheckMacValue != "" {
			CheckMacValue = CheckMacValue + "&"
		}

		CheckMacValue = CheckMacValue + SliceFinal[i].Parameter + "=" + SliceFinal[i].Value
	}

	CheckMacValue = "HashKey=" + HashKey + "&" + CheckMacValue + "&HashIV=" + HashIV
	fmt.Print("\nCheckMacValue=", CheckMacValue)

	CheckMacValue = FormUrlEncode(CheckMacValue)
	fmt.Print("\nCheckMacValue=", CheckMacValue)

	CheckMacValue = strings.ToLower(CheckMacValue)
	fmt.Print("\nCheckMacValue=", CheckMacValue)

	sum := sha256.Sum256([]byte(CheckMacValue))
	fmt.Printf("\n%x", sum)

	CheckMacValue = fmt.Sprintf("%x", sum)
	fmt.Print("\nCheckMacValue=", CheckMacValue)

	CheckMacValue = strings.ToUpper(CheckMacValue)
	fmt.Print("\nCheckMacValue=", CheckMacValue)

	fmt.Print("\nMerchantTradeNo=", MerchantTradeNo)
	fmt.Print("\nMerchantTradeDate=", MerchantTradeDate)
	fmt.Print("\nPaymentType=", PaymentType)
	fmt.Print("\nTradeDesc=", TradeDesc)
	fmt.Print("\nItemName=", ItemName)
	fmt.Print("\nChoosePayment=", ChoosePayment)
	fmt.Print("\nReturnURL=", ReturnURL)
	fmt.Print("\nClientBackURL=", ClientBackURL)
	fmt.Print("\nItemURL=", ItemURL)
	fmt.Print("\nInvoiceMark=", InvoiceMark)
	fmt.Print("\nCustomField1=", CustomField1)
	fmt.Print("\nEncryptType=", EncryptType)
	fmt.Print("\nPeriodAmount=", PeriodAmount)
	fmt.Print("\nPeriodType=", PeriodType)
	fmt.Print("\nFrequency=", Frequency)
	fmt.Print("\nExecTimes=", ExecTimes)
	fmt.Print("\nPeriodReturnURL=", PeriodReturnURL)
	fmt.Print("\nTaxType=", TaxType)
	fmt.Print("\nDelayDay=", DelayDay)
	fmt.Print("\nInvType=", InvType)
	fmt.Print("\nCarruerNum=", CarruerNum)
	fmt.Print("\n")
	return CheckMacValue, SliceFinal
}

func SendPostToEcPayOnce(CustomField1 string, CustomField2 string, CustomField3 string, MerchantID string, ITotalAmount int, TradeDesc string, ItemName string, ReturnURL string, ClientBackURL string, CustomerIdentifier string, CustomerEmail string, CarruerType string, CarruerNum string, Donation string, LoveCode string, Print string, InvoiceItemName string, InvoiceItemCount string, InvoiceItemWord string, InvoiceItemPrice string, CustomerName string, CustomerAddr string, HashKey string, HashIV string) (CheckMacValue string, slice []EcPayParm) {
	MerchantTradeNo := generateMerchantTradeNo(CustomField1)
	MerchantTradeDate := time.Now().Format("2006/01/02 15:04:05")
	PaymentType := "aio"
	RelateNumber := MerchantTradeNo
	ChoosePayment := "ALL"
	ItemURL := ClientBackURL
	InvoiceMark := "Y"
	EncryptType := "1"
	TotalAmount := strconv.Itoa(ITotalAmount)

	TaxType := "1"
	DelayDay := "0"
	InvType := "07"

	slice = []EcPayParm{}
	//按照字母排列 //

	slice = append(slice, EcPayParm{"CarruerNum", CarruerNum})
	slice = append(slice, EcPayParm{"CarruerType", CarruerType})
	slice = append(slice, EcPayParm{"ChoosePayment", ChoosePayment})
	slice = append(slice, EcPayParm{"ClientBackURL", ClientBackURL})
	slice = append(slice, EcPayParm{"CustomerAddr", CustomerAddr})
	slice = append(slice, EcPayParm{"CustomerEmail", CustomerEmail})
	slice = append(slice, EcPayParm{"CustomerIdentifier", CustomerIdentifier})
	slice = append(slice, EcPayParm{"CustomerName", CustomerName})
	slice = append(slice, EcPayParm{"CustomField1", CustomField1})
	slice = append(slice, EcPayParm{"CustomField2", CustomField2})
	slice = append(slice, EcPayParm{"CustomField3", CustomField3})

	slice = append(slice, EcPayParm{"DelayDay", DelayDay})
	slice = append(slice, EcPayParm{"Donation", Donation})

	slice = append(slice, EcPayParm{"EncryptType", EncryptType})

	slice = append(slice, EcPayParm{"InvoiceItemCount", InvoiceItemCount})
	slice = append(slice, EcPayParm{"InvoiceItemName", InvoiceItemName})
	slice = append(slice, EcPayParm{"InvoiceItemPrice", InvoiceItemPrice})
	slice = append(slice, EcPayParm{"InvoiceItemWord", InvoiceItemWord})
	slice = append(slice, EcPayParm{"InvoiceMark", InvoiceMark})
	slice = append(slice, EcPayParm{"InvType", InvType})
	slice = append(slice, EcPayParm{"ItemName", ItemName})
	slice = append(slice, EcPayParm{"ItemURL", ItemURL})

	slice = append(slice, EcPayParm{"LoveCode", LoveCode})

	slice = append(slice, EcPayParm{"MerchantID", MerchantID})
	slice = append(slice, EcPayParm{"MerchantTradeDate", MerchantTradeDate})
	slice = append(slice, EcPayParm{"MerchantTradeNo", MerchantTradeNo})

	slice = append(slice, EcPayParm{"PaymentType", PaymentType})

	slice = append(slice, EcPayParm{"Print", Print})

	slice = append(slice, EcPayParm{"RelateNumber", RelateNumber})
	slice = append(slice, EcPayParm{"ReturnURL", ReturnURL})

	slice = append(slice, EcPayParm{"TaxType", TaxType})
	slice = append(slice, EcPayParm{"TotalAmount", TotalAmount})
	slice = append(slice, EcPayParm{"TradeDesc", TradeDesc})

	SliceFinal := []EcPayParm{}

	for i := 0; i < len(slice); i++ {
		if slice[i].Value == "" {
			continue
		}
		SliceFinal = append(SliceFinal, slice[i])
	}

	CheckMacValue = ""

	for i := 0; i < len(SliceFinal); i++ {
		if SliceFinal[i].Value == "" {
			continue
		}

		if CheckMacValue != "" {
			CheckMacValue = CheckMacValue + "&"
		}

		CheckMacValue = CheckMacValue + SliceFinal[i].Parameter + "=" + SliceFinal[i].Value
	}

	CheckMacValue = "HashKey=" + HashKey + "&" + CheckMacValue + "&HashIV=" + HashIV
	fmt.Print("\nCheckMacValue=", CheckMacValue)

	CheckMacValue = FormUrlEncode(CheckMacValue)
	fmt.Print("\nCheckMacValue=", CheckMacValue)

	CheckMacValue = strings.ToLower(CheckMacValue)
	fmt.Print("\nCheckMacValue=", CheckMacValue)

	sum := sha256.Sum256([]byte(CheckMacValue))
	fmt.Printf("\n%x", sum)

	CheckMacValue = fmt.Sprintf("%x", sum)
	fmt.Print("\nCheckMacValue=", CheckMacValue)

	CheckMacValue = strings.ToUpper(CheckMacValue)
	fmt.Print("\nCheckMacValue=", CheckMacValue)

	fmt.Print("\nMerchantTradeNo=", MerchantTradeNo)
	fmt.Print("\nMerchantTradeDate=", MerchantTradeDate)
	fmt.Print("\nPaymentType=", PaymentType)
	fmt.Print("\nTradeDesc=", TradeDesc)
	fmt.Print("\nItemName=", ItemName)
	fmt.Print("\nChoosePayment=", ChoosePayment)
	fmt.Print("\nReturnURL=", ReturnURL)
	fmt.Print("\nClientBackURL=", ClientBackURL)
	fmt.Print("\nItemURL=", ItemURL)
	fmt.Print("\nInvoiceMark=", InvoiceMark)
	fmt.Print("\nCustomField1=", CustomField1)
	fmt.Print("\nEncryptType=", EncryptType)
	fmt.Print("\nTaxType=", TaxType)
	fmt.Print("\nDelayDay=", DelayDay)
	fmt.Print("\nInvType=", InvType)
	fmt.Print("\nCarruerNum=", CarruerNum)
	fmt.Print("\n")
	return CheckMacValue, SliceFinal

}
