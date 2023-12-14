package alert

type Server struct {
	POP3     string
	POP3Port int
	SMTP     string
	SMTPProt int
}

var QQ = &Server{
	POP3:     "pop.qq.com",
	POP3Port: 995,
	SMTP:     "smtp.qq.cocm",
	SMTPProt: 465, //587
}

var NetEase = &Server{
	POP3:     "pop.163.com",
	POP3Port: 110,
	SMTP:     "smtp.163.com",
	SMTPProt: 25,
}

var Gmail = &Server{
	POP3:     "pop.gmail.com",
	POP3Port: 995,
	SMTP:     "smtp.gmail.com",
	SMTPProt: 587,
}

const (
	LF         = "\u000A"               // \n
	CR         = "\u000D"               // \r
	CRLf       = "\U000D000A"           // \r\n
	DOUBLECRLF = "\U000D000A\U000D000A" // \r\n\r\n
	NULL       = "\u0000"               // null
	NBSP       = "\u00A0"               // No-Break Space
	BR         = "<br>"
)
const (
	Default  = Serena
	Allison  = "Allison"  // 深沉美式女声
	Ava      = "Ava"      // 深沉美式女声
	Daniel   = "Daniel"   // 正式英式男声
	Lanlan   = "Lanlan"   // 童声中文女声
	Meijia   = "Meijia"   // 正式中文女声
	Lilian   = "Lilian"   // 柔和中文女声
	Samantha = "Samantha" // 正经美式女声
	Serena   = "Serena"   // 沉稳英式女声
	Shanshan = "Shanshan" // 浑厚中文女声
	Shasha   = "Shasha"   // 成熟中文女声
	Sinji    = "Sinji"    // 粤语中文女声
	Tingting = "Tingting" // 机械中文女声
	Victoria = "Victoria" // 压缩美式女声
)
