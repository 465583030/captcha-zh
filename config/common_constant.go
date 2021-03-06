package config

const (
	//这是个奇葩,必须是这个时间点, 据说是go诞生之日, 记忆方法:6-1-2-3-4-5
	DATE_FORMAT_DEFAULT string = "2006-01-02 15:04:05"
	DATE_FORMAT_DEFAULT_LEN int = len(DATE_FORMAT_DEFAULT)

	// 分隔符
	SEPARATOR_UNDER_LINE = "_"
	SEPARATOR_VERTICAL_LINE = "|"
	SEPARATOR_SPACE = " "

	// 绝对路径，根
	PATH_ROOT = "/Volumes/IntelSSD/Dev/GoRepository/src/captcha-zh/"
	// 配置文件路径
	// TODO 测试用例场景读不到toml路径，要写绝对路径
	PATH_CONFIG_TOML = PATH_ROOT + "resource/conf/conf.toml"
	PATH_CONFIG_JSON = "resource/conf/config.json"
	// 字库文件路径
	PATH_CONFIG_FONT = "resource/assets/fonts/"
	// 背景图片路径
	PATH_CONFIG_IMAGE_BG = "resource/assets/images/bg.gif"
	// 测试用临时图片路径
	PATH_CONFIG_IMAGE_TEMP = "bin/tmp/"
)

// 月份
var Months = [...]string {
	1	:	"January",	// 1
	2	:	"February",
	3	:	"March",
	4	:	"April",
	5	:	"May",
	6	:	"June",
	7	:	"July",
	8	:	"August",
	9	:	"September",
	10	:	"October",
	11	:	"November",
	12	:	"December",
}

// 星期
var Weeks = [...]string {
	"Sunday",	// 0
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

// 性别的枚举
var SexConstant = map[string]string {
	"0":"非男非女",
	"1":"女",
	"2":"男",
}

var (
	Fonts []string
)
