package encrypt

import "encoding/base64"

var (
	_key = 913
)

func Encrypt(pt []byte) string {
	payload := make([]byte, len(pt))
	key := byte(_key)
	for i, v := range pt {
		payload[i] = v ^ key
		key = payload[i]
	}
	return base64.StdEncoding.EncodeToString(payload)
}

func Decrypt(ct string) ([]byte, error) {
	ds, err := base64.StdEncoding.DecodeString(ct)
	if err != nil {
		return nil, err
	}

	var nextkey byte
	key := byte(_key)
	result := make([]byte, len(ct))
	for i, v := range ds {
		nextkey = v
		result[i] = v ^ key
		key = nextkey
	}
	return result, nil
}