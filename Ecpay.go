package Ecpayby77

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func generateMerchantTradeNo(MemberId int) (No string) {
	time := time.Now().UnixNano() / 1e6
	timeString := strconv.FormatInt(time, 16)

	sMemberId := strconv.Itoa(MemberId)

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

/*
綠界的參數請參考 https://www.ecpay.com.tw/Content/files/ecpay_011.pdf

MemberId 自己商場會員的ID 用於自己識別
MerchantID  綠界參數
TotalAmount 綠界參數
*/

func SendPostToEcPayPeriod(MemberId int, MerchantID string, TotalAmount int, TradeDesc string, ItemName string) (err error) {
	MerchantTradeNo := generateMerchantTradeNo(MemberId)
	MerchantTradeDate := time.Now().Format("2006/01/02 15:04:05")
	PaymentType := "aio"
	TradeDesc = FormUrlEncode(TradeDesc)

	fmt.Print("\nMerchantTradeNo=", MerchantTradeNo)
	fmt.Print("\nMerchantTradeDate=", MerchantTradeDate)
	fmt.Print("\nPaymentType=", PaymentType)
	fmt.Print("\nTradeDesc=", TradeDesc)
	fmt.Print("\nItemName=", ItemName)
	fmt.Print("\n")
	return
}

func SendPostToEcPayOnce(MemberId int, MerchantID string, TotalAmount int, TradeDesc string, ItemName string) (err error) {
	return
}
