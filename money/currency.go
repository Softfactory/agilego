package money

type Currency int

const (
	AFA Currency = iota + 6
	ALL
	DZD
	ARS
	AWG
	AUD
	BSD
	BHD
	BDT
	BBD
	BZD
	BMD
	BTN
	BOB
	BWP
	BRL
	GBP
	BND
	BIF
	XOF
	XAF
	KHR
	CAD
	CVE
	KYD
	CLP
	CNY
	COP
	KMF
	CRC
	HRK
	CUP
	CYP
	CZK
	DKK
	DJF
	DOP
	XCD
	EGP
	SVC
	EEK
	ETB
	EUR
	FKP
	GMD
	GHC
	GIP
	XAU
	GTQ
	GNF
	GYD
	HTG
	HNL
	HKD
	HUF
	ISK
	INR
	IDR
	IQD
	ILS
	JMD
	JPY
	JOD
	KZT
	KES
	KRW
	KWD
	LAK
	LVL
	LBP
	LSL
	LRD
	LYD
	LTL
	MOP
	MKD
	MGF
	MWK
	MYR
	MVR
	MTL
	MRO
	MUR
	MXN
	MDL
	MNT
	MAD
	MZM
	MMK
	NAD
	NPR
	ANG
	NZD
	NIO
	NGN
	KPW
	NOK
	OMR
	XPF
	PKR
	XPD
	PAB
	PGK
	PYG
	PEN
	PHP
	XPT
	PLN
	QAR
	ROL
	RUB
	WST
	STD
	SAR
	SCR
	SLL
	XAG
	SGD
	SKK
	SIT
	SBD
	SOS
	ZAR
	LKR
	SHP
	SDD
	SRG
	SZL
	SEK
	TRY
	CHF
	SYP
	TWD
	TZS
	THB
	TOP
	TTD
	TND
	TRL
	USD
	AED
	UGX
	UAH
	UYU
	VUV
	VEB
	VND
	YER
	YUM
	ZMK
	ZWD
	MAX
)

// type Currency string

// funcCurrency(currency string)Currency {
// 	returnCurrency(currency)
// }

// const (
// 	AFA = Currency("AFA")
// 	ALL = Currency("ALL")
// 	DZD = Currency("DZD")
// 	ARS = Currency("ARS")
// 	AWG = Currency("AWG")
// 	AUD = Currency("AUD")
// 	BSD = Currency("BSD")
// 	BHD = Currency("BHD")
// 	BDT = Currency("BDT")
// 	BBD = Currency("BBD")
// 	BZD = Currency("BZD")
// 	BMD = Currency("BMD")
// 	BTN = Currency("BTN")
// 	BOB = Currency("BOB")
// 	BWP = Currency("BWP")
// 	BRL = Currency("BRL")
// 	GBP = Currency("GBP")
// 	BND = Currency("BND")
// 	BIF = Currency("BIF")
// 	XOF = Currency("XOF")
// 	XAF = Currency("XAF")
// 	KHR = Currency("KHR")
// 	CAD = Currency("CAD")
// 	CVE = Currency("CVE")
// 	KYD = Currency("KYD")
// 	CLP = Currency("CLP")
// 	CNY = Currency("CNY")
// 	COP = Currency("COP")
// 	KMF = Currency("KMF")
// 	CRC = Currency("CRC")
// 	HRK = Currency("HRK")
// 	CUP = Currency("CUP")
// 	CYP = Currency("CYP")
// 	CZK = Currency("CZK")
// 	DKK = Currency("DKK")
// 	DJF = Currency("DJF")
// 	DOP = Currency("DOP")
// 	XCD = Currency("XCD")
// 	EGP = Currency("EGP")
// 	SVC = Currency("SVC")
// 	EEK = Currency("EEK")
// 	ETB = Currency("ETB")
// 	EUR = Currency("EUR")
// 	FKP = Currency("FKP")
// 	GMD = Currency("GMD")
// 	GHC = Currency("GHC")
// 	GIP = Currency("GIP")
// 	XAU = Currency("XAU")
// 	GTQ = Currency("GTQ")
// 	GNF = Currency("GNF")
// 	GYD = Currency("GYD")
// 	HTG = Currency("HTG")
// 	HNL = Currency("HNL")
// 	HKD = Currency("HKD")
// 	HUF = Currency("HUF")
// 	ISK = Currency("ISK")
// 	INR = Currency("INR")
// 	IDR = Currency("IDR")
// 	IQD = Currency("IQD")
// 	ILS = Currency("ILS")
// 	JMD = Currency("JMD")
// 	JPY = Currency("JPY")
// 	JOD = Currency("JOD")
// 	KZT = Currency("KZT")
// 	KES = Currency("KES")
// 	//KRW Korean Won)
// 	KRW = Currency("KRW")
// 	KWD = Currency("KWD")
// 	LAK = Currency("LAK")
// 	LVL = Currency("LVL")
// 	LBP = Currency("LBP")
// 	LSL = Currency("LSL")
// 	LRD = Currency("LRD")
// 	LYD = Currency("LYD")
// 	LTL = Currency("LTL")
// 	MOP = Currency("MOP")
// 	MKD = Currency("MKD")
// 	MGF = Currency("MGF")
// 	MWK = Currency("MWK")
// 	MYR = Currency("MYR")
// 	MVR = Currency("MVR")
// 	MTL = Currency("MTL")
// 	MRO = Currency("MRO")
// 	MUR = Currency("MUR")
// 	MXN = Currency("MXN")
// 	MDL = Currency("MDL")
// 	MNT = Currency("MNT")
// 	MAD = Currency("MAD")
// 	MZM = Currency("MZM")
// 	MMK = Currency("MMK")
// 	NAD = Currency("NAD")
// 	NPR = Currency("NPR")
// 	ANG = Currency("ANG")
// 	NZD = Currency("NZD")
// 	NIO = Currency("NIO")
// 	NGN = Currency("NGN")
// 	KPW = Currency("KPW")
// 	NOK = Currency("NOK")
// 	OMR = Currency("OMR")
// 	XPF = Currency("XPF")
// 	PKR = Currency("PKR")
// 	XPD = Currency("XPD")
// 	PAB = Currency("PAB")
// 	PGK = Currency("PGK")
// 	PYG = Currency("PYG")
// 	PEN = Currency("PEN")
// 	PHP = Currency("PHP")
// 	XPT = Currency("XPT")
// 	PLN = Currency("PLN")
// 	QAR = Currency("QAR")
// 	ROL = Currency("ROL")
// 	RUB = Currency("RUB")
// 	WST = Currency("WST")
// 	STD = Currency("STD")
// 	SAR = Currency("SAR")
// 	SCR = Currency("SCR")
// 	SLL = Currency("SLL")
// 	XAG = Currency("XAG")
// 	SGD = Currency("SGD")
// 	SKK = Currency("SKK")
// 	SIT = Currency("SIT")
// 	SBD = Currency("SBD")
// 	SOS = Currency("SOS")
// 	ZAR = Currency("ZAR")
// 	LKR = Currency("LKR")
// 	SHP = Currency("SHP")
// 	SDD = Currency("SDD")
// 	SRG = Currency("SRG")
// 	SZL = Currency("SZL")
// 	SEK = Currency("SEK")
// 	TRY = Currency("TRY")
// 	CHF = Currency("CHF")
// 	SYP = Currency("SYP")
// 	TWD = Currency("TWD")
// 	TZS = Currency("TZS")
// 	THB = Currency("THB")
// 	TOP = Currency("TOP")
// 	TTD = Currency("TTD")
// 	TND = Currency("TND")
// 	TRL = Currency("TRL")
// 	USD = Currency("USD")
// 	AED = Currency("AED")
// 	UGX = Currency("UGX")
// 	UAH = Currency("UAH")
// 	UYU = Currency("UYU")
// 	VUV = Currency("VUV")
// 	VEB = Currency("VEB")
// 	VND = Currency("VND")
// 	YER = Currency("YER")
// 	YUM = Currency("YUM")
// 	ZMK = Currency("ZMK")
// 	ZWD = Currency("ZWD")
// )
