package money

// Currency 통화기호를 위한 데이터형
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

func (s Currency) String() string {
	switch s {
	case AFA:
		return "AFA"
	case ALL:
		return "ALL"
	case DZD:
		return "DZD"
	case ARS:
		return "ARS"
	case AWG:
		return "AWG"
	case AUD:
		return "AUD"
	case BSD:
		return "BSD"
	case BHD:
		return "BHD"
	case BDT:
		return "BDT"
	case BBD:
		return "BBD"
	case BZD:
		return "BZD"
	case BMD:
		return "BMD"
	case BTN:
		return "BTN"
	case BOB:
		return "BOB"
	case BWP:
		return "BWP"
	case BRL:
		return "BRL"
	case GBP:
		return "GBP"
	case BND:
		return "BND"
	case BIF:
		return "BIF"
	case XOF:
		return "XOF"
	case XAF:
		return "XAF"
	case KHR:
		return "KHR"
	case CAD:
		return "CAD"
	case CVE:
		return "CVE"
	case KYD:
		return "KYD"
	case CLP:
		return "CLP"
	case CNY:
		return "CNY"
	case COP:
		return "COP"
	case KMF:
		return "KMF"
	case CRC:
		return "CRC"
	case HRK:
		return "HRK"
	case CUP:
		return "CUP"
	case CYP:
		return "CYP"
	case CZK:
		return "CZK"
	case DKK:
		return "DKK"
	case DJF:
		return "DJF"
	case DOP:
		return "DOP"
	case XCD:
		return "XCD"
	case EGP:
		return "EGP"
	case SVC:
		return "SVC"
	case EEK:
		return "EEK"
	case ETB:
		return "ETB"
	case EUR:
		return "EUR"
	case FKP:
		return "FKP"
	case GMD:
		return "GMD"
	case GHC:
		return "GHC"
	case GIP:
		return "GIP"
	case XAU:
		return "XAU"
	case GTQ:
		return "GTQ"
	case GNF:
		return "GNF"
	case GYD:
		return "GYD"
	case HTG:
		return "HTG"
	case HNL:
		return "HNL"
	case HKD:
		return "HKD"
	case HUF:
		return "HUF"
	case ISK:
		return "ISK"
	case INR:
		return "INR"
	case IDR:
		return "IDR"
	case IQD:
		return "IQD"
	case ILS:
		return "ILS"
	case JMD:
		return "JMD"
	case JPY:
		return "JPY"
	case JOD:
		return "JOD"
	case KZT:
		return "KZT"
	case KES:
		return "KES"
	case KRW:
		return "KRW"
	case KWD:
		return "KWD"
	case LAK:
		return "LAK"
	case LVL:
		return "LVL"
	case LBP:
		return "LBP"
	case LSL:
		return "LSL"
	case LRD:
		return "LRD"
	case LYD:
		return "LYD"
	case LTL:
		return "LTL"
	case MOP:
		return "MOP"
	case MKD:
		return "MKD"
	case MGF:
		return "MGF"
	case MWK:
		return "MWK"
	case MYR:
		return "MYR"
	case MVR:
		return "MVR"
	case MTL:
		return "MTL"
	case MRO:
		return "MRO"
	case MUR:
		return "MUR"
	case MXN:
		return "MXN"
	case MDL:
		return "MDL"
	case MNT:
		return "MNT"
	case MAD:
		return "MAD"
	case MZM:
		return "MZM"
	case MMK:
		return "MMK"
	case NAD:
		return "NAD"
	case NPR:
		return "NPR"
	case ANG:
		return "ANG"
	case NZD:
		return "NZD"
	case NIO:
		return "NIO"
	case NGN:
		return "NGN"
	case KPW:
		return "KPW"
	case NOK:
		return "NOK"
	case OMR:
		return "OMR"
	case XPF:
		return "XPF"
	case PKR:
		return "PKR"
	case XPD:
		return "XPD"
	case PAB:
		return "PAB"
	case PGK:
		return "PGK"
	case PYG:
		return "PYG"
	case PEN:
		return "PEN"
	case PHP:
		return "PHP"
	case XPT:
		return "XPT"
	case PLN:
		return "PLN"
	case QAR:
		return "QAR"
	case ROL:
		return "ROL"
	case RUB:
		return "RUB"
	case WST:
		return "WST"
	case STD:
		return "STD"
	case SAR:
		return "SAR"
	case SCR:
		return "SCR"
	case SLL:
		return "SLL"
	case XAG:
		return "XAG"
	case SGD:
		return "SGD"
	case SKK:
		return "SKK"
	case SIT:
		return "SIT"
	case SBD:
		return "SBD"
	case SOS:
		return "SOS"
	case ZAR:
		return "ZAR"
	case LKR:
		return "LKR"
	case SHP:
		return "SHP"
	case SDD:
		return "SDD"
	case SRG:
		return "SRG"
	case SZL:
		return "SZL"
	case SEK:
		return "SEK"
	case TRY:
		return "TRY"
	case CHF:
		return "CHF"
	case SYP:
		return "SYP"
	case TWD:
		return "TWD"
	case TZS:
		return "TZS"
	case THB:
		return "THB"
	case TOP:
		return "TOP"
	case TTD:
		return "TTD"
	case TND:
		return "TND"
	case TRL:
		return "TRL"
	case USD:
		return "USD"
	case AED:
		return "AED"
	case UGX:
		return "UGX"
	case UAH:
		return "UAH"
	case UYU:
		return "UYU"
	case VUV:
		return "VUV"
	case VEB:
		return "VEB"
	case VND:
		return "VND"
	case YER:
		return "YER"
	case YUM:
		return "YUM"
	case ZMK:
		return "ZMK"
	case ZWD:
		return "ZWD"
	default:
		return "Unknown"
	}
}
