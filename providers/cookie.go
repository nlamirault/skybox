// Copyright (C) 2016, 2017 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package providers

import (
	"bytes"
	"fmt"
	//"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
)

// TimeFormat is the time format to use with
// time.Parse and time.Time.Format when parsing
// or generating times in HTTP headers.
// It is like time.RFC1123 but hard codes GMT as the time zone.
const TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"

var cookieNameSanitizer = strings.NewReplacer("\n", "-", "\r", "-")

func sanitizeCookieName(n string) string {
	return cookieNameSanitizer.Replace(n)
}

// path-av           = "Path=" path-value
// path-value        = <any CHAR except CTLs or ";">
func sanitizeCookiePath(v string) string {
	return sanitizeOrWarn("Cookie.Path", validCookiePathByte, v)
}

// http://tools.ietf.org/html/rfc6265#section-4.1.1
// cookie-value      = *cookie-octet / ( DQUOTE *cookie-octet DQUOTE )
// cookie-octet      = %x21 / %x23-2B / %x2D-3A / %x3C-5B / %x5D-7E
//           ; US-ASCII characters excluding CTLs,
//           ; whitespace DQUOTE, comma, semicolon,
//           ; and backslash
// We loosen this as spaces and commas are common in cookie values
// but we produce a quoted cookie-value in when value starts or ends
// with a comma or space.
// See https://golang.org/issue/7243 for the discussion.
func sanitizeCookieValue(v string) string {
	v = sanitizeOrWarn("Cookie.Value", validCookieValueByte, v)
	if len(v) == 0 {
		return v
	}
	if v[0] == ' ' || v[0] == ',' || v[len(v)-1] == ' ' || v[len(v)-1] == ',' {
		return `"` + v + `"`
	}
	return v
}

func sanitizeOrWarn(fieldName string, valid func(byte) bool, v string) string {
	ok := true
	for i := 0; i < len(v); i++ {
		if valid(v[i]) {
			continue
		}
		logrus.Debugf("net/http: invalid byte %q in %s; dropping invalid bytes", v[i], fieldName)
		ok = false
		break
	}
	if ok {
		return v
	}
	buf := make([]byte, 0, len(v))
	for i := 0; i < len(v); i++ {
		if b := v[i]; valid(b) {
			buf = append(buf, b)
		}
	}
	return string(buf)
}

func validCookieValueByte(b byte) bool {
	return 0x20 <= b && b < 0x7f && b != '"' && b != ';' && b != '\\'
}

func validCookiePathByte(b byte) bool {
	return 0x20 <= b && b < 0x7f && b != ';'
}

func parseCookieValue(raw string, allowDoubleQuote bool) (string, bool) {
	// Strip the quotes, if present.
	if allowDoubleQuote && len(raw) > 1 && raw[0] == '"' && raw[len(raw)-1] == '"' {
		raw = raw[1 : len(raw)-1]
	}
	for i := 0; i < len(raw); i++ {
		if !validCookieValueByte(raw[i]) {
			return "", false
		}
	}
	return raw, true
}

func generateCookie(c *http.Cookie) string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "%s=%s", sanitizeCookieName(c.Name), sanitizeCookieValue(c.Value))
	if len(c.Path) > 0 {
		fmt.Fprintf(&b, "; Path=%s", sanitizeCookiePath(c.Path))
	}
	if len(c.Domain) > 0 {
		// if validCookieDomain(c.Domain) {
		// A c.Domain containing illegal characters is not
		// sanitized but simply dropped which turns the cookie
		// into a host-only cookie. A leading dot is okay
		// but won't be sent.
		d := c.Domain
		if d[0] == '.' {
			d = d[1:]
		}
		fmt.Fprintf(&b, "; Domain=%s", d)
		// } else {
		// 	log.Printf("net/http: invalid Cookie.Domain %q; dropping domain attribute",
		// 		c.Domain)
		// }
	}
	if c.Expires.Unix() > 0 {
		fmt.Fprintf(&b, "; Expires=%s", c.Expires.UTC().Format(TimeFormat))
	}
	if c.MaxAge > 0 {
		fmt.Fprintf(&b, "; Max-Age=%d", c.MaxAge)
	} else if c.MaxAge < 0 {
		fmt.Fprintf(&b, "; Max-Age=0")
	}
	if c.HttpOnly {
		fmt.Fprintf(&b, "; HttpOnly")
	}
	if c.Secure {
		fmt.Fprintf(&b, "; Secure")
	}
	return b.String()
}

func readCookie(line string) *http.Cookie {
	logrus.Debugf("Line: %s", line)
	parts := strings.Split(strings.TrimSpace(line), ";")
	if len(parts) == 1 && parts[0] == "" {
		return nil
	}
	parts[0] = strings.TrimSpace(parts[0])
	j := strings.Index(parts[0], "=")
	if j < 0 {
		return nil
	}
	name, value := parts[0][:j], parts[0][j+1:]
	// if !isCookieNameValid(name) {
	// 	continue
	// }
	value, success := parseCookieValue(value, true)
	if !success {
		return nil
	}

	c := &http.Cookie{
		Name:  name,
		Value: value,
		Raw:   line,
	}
	logrus.Debugf("Find cookie: %#v %s %s %s", c, name, value, line)
	for i := 1; i < len(parts); i++ {
		parts[i] = strings.TrimSpace(parts[i])
		if len(parts[i]) == 0 {
			continue
		}

		attr, val := parts[i], ""
		if j := strings.Index(attr, "="); j >= 0 {
			attr, val = attr[:j], attr[j+1:]
		}
		lowerAttr := strings.ToLower(attr)
		val, success = parseCookieValue(val, false)
		if !success {
			c.Unparsed = append(c.Unparsed, parts[i])
			continue
		}
		switch lowerAttr {
		case "secure":
			c.Secure = true
			continue
		case "httponly":
			c.HttpOnly = true
			continue
		case "domain":
			c.Domain = val
			continue
		case "max-age":
			secs, err := strconv.Atoi(val)
			if err != nil || secs != 0 && val[0] == '0' {
				break
			}
			if secs <= 0 {
				c.MaxAge = -1
			} else {
				c.MaxAge = secs
			}
			continue
		case "expires":
			c.RawExpires = val
			exptime, err := time.Parse(time.RFC1123, val)
			if err != nil {
				exptime, err = time.Parse("Mon, 02-Jan-2006 15:04:05 MST", val)
				if err != nil {
					c.Expires = time.Time{}
					break
				}
			}
			c.Expires = exptime.UTC()
			continue
		case "path":
			c.Path = val
			continue
		}
		c.Unparsed = append(c.Unparsed, parts[i])
	}
	logrus.Debugf("Cookie ok: %#v", c)
	return c
}
