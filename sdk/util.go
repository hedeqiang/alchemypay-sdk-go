package sdk

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
)

// OrderParam order params
func OrderParam(p map[string]interface{}, bizKey string) (returnStr string) {
	keys := make([]string, 0, len(p))
	for k := range p {
		if k == "sign" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var buf bytes.Buffer
	for _, k := range keys {
		if p[k] == "" {
			continue
		}
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(k)
		buf.WriteByte('=')

		buf.WriteString(fmt.Sprint(p[k]))
	}
	buf.WriteString("&key=" + bizKey)
	returnStr = buf.String()
	return
}

func SignWithMd5(signStr string) string {
	d := []byte(signStr)
	md5str := fmt.Sprintf("%x", md5.Sum(d))

	return strings.ToUpper(md5str)
}
