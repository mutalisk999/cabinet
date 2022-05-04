package tutorial

import "fyne.io/fyne/v2"

type Tutorial struct {
	Title string
	Intro string
	View  func(w fyne.Window) fyne.CanvasObject
}

var (
	Tutorials = map[string]Tutorial{
		"convert": {"convert",
			"convert",
			convertScreen},
		"base_convert": {"base_convert",
			"base_convert",
			baseConvertScreen},
		"time_convert": {"time_convert",
			"time_convert",
			timeConvertScreen},
		"case_convert": {"case_convert",
			"case_convert",
			caseConvertScreen},
		"encode": {"encode",
			"encode",
			encodeScreen},
		"base64_encode": {"base64_encode",
			"base64_encode",
			base64EncodeScreen},
		"url_encode": {"url_encode",
			"url_encode",
			convertScreen},
		"html_encode": {"html_encode",
			"html_encode",
			convertScreen},
		"image": {"image",
			"image",
			convertScreen},
		"image_convert": {"image_convert",
			"image_convert",
			convertScreen},
		"image_compress": {"image_compress",
			"image_compress",
			convertScreen},
		"image_to_json": {"image_to_json",
			"image_to_json",
			convertScreen},
		"qrcode": {"qrcode",
			"qrcode",
			convertScreen},
		"json": {"json",
			"json",
			convertScreen},
		"json_compress": {"json_compress",
			"json_compress",
			convertScreen},
		"json_to_yaml": {"json_to_yaml",
			"json_to_yaml",
			convertScreen},
		"digest": {"digest",
			"digest",
			convertScreen},
		"calc_hash": {"calc_hash",
			"calc_hash",
			convertScreen},
		"file_checksum": {"file_checksum",
			"file_checksum",
			convertScreen},
		"crypto": {"crypto",
			"crypto",
			convertScreen},
		"encrypt_aes": {"encrypt_aes",
			"encrypt_aes",
			convertScreen},
		"encrypt_des": {"encrypt_des",
			"encrypt_des",
			convertScreen},
		"signature": {"signature",
			"signature",
			convertScreen},
		"sign_rsa": {"sign_rsa",
			"sign_rsa",
			convertScreen},
		"sign_ecc": {"sign_ecc",
			"sign_ecc",
			convertScreen},
		"network": {"network",
			"network",
			networkScreen},
		"get_my_ip": {"get_my_ip",
			"get_my_ip",
			networkGetIPScreen},
		"ip_mask_calc": {"ip_mask_calc",
			"ip_mask_calc",
			networkIPMaskScreen},
		"web_server": {"web_server",
			"web_server",
			networkWebServerScreen},
		"other": {"other",
			"other",
			otherScreen},
		"uuid": {"uuid",
			"uuid",
			uuidScreen},
		"random_pass": {"random_pass",
			"random_pass",
			randomPasswordScreen},
		"rsa_key_pair": {"rsa_key_pair",
			"rsa_key_pair",
			convertScreen},
		"ecc_key_pair": {"ecc_key_pair",
			"ecc_key_pair",
			convertScreen},
		"expression_calc": {"expression_calc",
			"expression_calc",
			convertScreen},
		"markdown": {"markdown",
			"markdown",
			convertScreen},
	}

	TutorialIndex = map[string][]string{
		"":          {"convert", "encode", "image", "json", "digest", "crypto", "signature", "network", "other"},
		"convert":   {"base_convert", "time_convert", "case_convert"},
		"encode":    {"base64_encode", "url_encode", "html_encode"},
		"image":     {"image_convert", "image_compress", "image_to_json", "qrcode"},
		"json":      {"json_compress", "json_to_yaml"},
		"digest":    {"calc_hash", "file_checksum"},
		"crypto":    {"encrypt_aes", "encrypt_des"},
		"signature": {"sign_rsa", "sign_ecc"},
		"network":   {"get_my_ip", "ip_mask_calc", "web_server"},
		"other":     {"uuid", "random_pass", "rsa_key_pair", "ecc_key_pair", "expression_calc", "markdown"},
	}
)
