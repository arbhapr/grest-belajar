package i18n

func IdID() map[string]string {
	return map[string]string{
		"400_bad_request":              "Permintaan tidak dapat dilakukan karena ada parameter yang salah atau tidak lengkap.",
		"401_unauthorized":             "Token otentikasi tidak valid. Silakan logout dan login ulang",
		"403_forbidden":                "Pengguna tidak memiliki izin untuk :action.",
		"404_not_found":                "The resource you have specified cannot be found.",
		"500_internal_error":           "Gagal terhubung ke server, silakan coba lagi nanti.",
		"invalid_username_or_password": "Username atau kata sandi tidak valid",
	}
}
