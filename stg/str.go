package stg

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/Nhanderu/brdoc"
	_uuid "github.com/google/uuid"
)

// ToLowerTrim - remove caracteres e executa trim to lower
func ToLowerTrim(b []byte) string {
	return strings.ToLower(strings.TrimSpace(strings.Trim(string(b), `"`)))
}

// ToUpperTrimAll - remove escapes to upper string
func ToUpperTrimAll(value string) string {
	return strings.ToUpper(strings.Replace(strings.Replace(value, "'", "", -1), "\"", "", -1))
}

// ToTrimAllUnscape - remove escapes string
func ToTrimAllUnscape(value string) string {
	return strings.Replace(strings.Replace(value, "'", "", -1), "\"", "", -1)
}

// IsEmpty - verifica se a string eh vazia
func IsEmpty(s *string) bool {
	if s == nil {
		return true
	}
	if len(strings.TrimSpace(*s)) == 0 {
		return true
	}
	return false
}

// ToTitleCase titleccase
func ToTitleCase(value string) string {
	return ToCamel(strings.Replace(value, "_", "", -1))
}

// IsEmail - verifica se eh um email valido
func IsEmail(value string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(value)
}

// ToCamel camelcase
func ToCamel(value string) string {
	return strings.Title(strings.ToLower(value))
}

// FirstChar exported
func FirstChar(capacity int, value string) string {
	if capacity <= 0 {
		return value
	}

	if capacity < len(value) {
		return value[:capacity]
	}
	return value
}

// RemoveScapes exported
func RemoveScapes(s string, r string, n string) string {
	return strings.Replace(s, r, n, -1)
}

// ToJSON -convert para json string
func ToJSON(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}

// ToJSONIndent -convert para json string
func ToJSONIndent(v interface{}) string {
	data, _ := json.MarshalIndent(v, "", "\t")
	return string(data)
}

// StringToMap convert to map
func StringToMap(in string) map[string]interface{} {
	out := make(map[string]interface{})
	if err := json.Unmarshal([]byte(in), &out); err != nil {
		return nil
	}
	return out
}

// ToMap convert to map
func ToMap(in interface{}) map[string]interface{} {
	out := make(map[string]interface{})
	data, err := json.Marshal(in)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(data, &out)
	if err != nil {
		return nil
	}
	return out
}

// MergeMaps join to maps
func MergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

// ConvertMapArray exported
func ConvertMapArray(val interface{}) []map[string]interface{} {
	if val == nil {
		return nil
	}
	if reflect.TypeOf(val).Kind() == reflect.Ptr {
		if reflect.Slice == reflect.TypeOf(val).Elem().Kind() {
			items := make([]map[string]interface{}, 0)
			element := reflect.ValueOf(val).Elem()
			for i := 0; i < element.Len(); i++ {
				value := ToMap(element.Index(i).Interface())
				for k, v := range value {
					items = append(items, map[string]interface{}{
						k: v,
					})
				}
			}
			return items
		}
		if reflect.Struct == reflect.TypeOf(val).Elem().Kind() {
			items := make([]map[string]interface{}, 0)
			element := reflect.ValueOf(val).Elem()
			value := ToMap(element.Interface())
			for k, v := range value {
				items = append(items, map[string]interface{}{
					k: v,
				})
			}
			return items
		}

	} else if reflect.Struct == reflect.TypeOf(val).Kind() {
		items := make([]map[string]interface{}, 0)
		value := ToMap(val)
		for k, v := range value {
			items = append(items, map[string]interface{}{
				k: v,
			})
		}
		return items
	} else if reflect.Slice == reflect.TypeOf(val).Kind() {
		element := reflect.ValueOf(val)
		items := make([]map[string]interface{}, 0)
		for i := 0; i < element.Len(); i++ {
			value := ToMap(element.Index(i).Interface())
			for k, v := range value {
				items = append(items, map[string]interface{}{
					k: v,
				})
			}
		}
		return items
	} else if reflect.Map == reflect.TypeOf(val).Kind() {
		items := make([]map[string]interface{}, 0)
		if value, ok := val.(map[string]interface{}); ok {
			for k, v := range value {
				items = append(items, map[string]interface{}{
					k: v,
				})
			}
		}
		return items
	}
	return nil
}

// ToMapArray convert []map
func ToMapArray(in interface{}) []map[string]interface{} {
	out := make([]map[string]interface{}, 0)
	data, err := json.Marshal(in)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(data, &out)
	if err != nil {
		return nil
	}
	return out
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

// RandStringBytes - gera um numero randomico alfa numerico
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		rand.Seed(time.Now().UnixNano())
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// Pad - Pad string
func Pad(input string, padLength int, padString string, padType string) string {
	var output string

	inputLength := len(input)
	padStringLength := len(padString)

	if inputLength >= padLength {
		return input
	}

	repeat := math.Ceil(float64(1) + (float64(padLength-padStringLength))/float64(padStringLength))

	switch padType {
	case "RIGHT":
		output = input + strings.Repeat(padString, int(repeat))
		output = output[:padLength]
	case "LEFT":
		output = strings.Repeat(padString, int(repeat)) + input
		output = output[len(output)-padLength:]
	case "BOTH":
		length := (float64(padLength - inputLength)) / float64(2)
		repeat = math.Ceil(length / float64(padStringLength))
		output = strings.Repeat(padString, int(repeat))[:int(math.Floor(float64(length)))] + input + strings.Repeat(padString, int(repeat))[:int(math.Ceil(float64(length)))]
	}

	return output
}

// String returns a pointer to the string value passed in.
func String(v string) *string {
	return &v
}

// StringValue returns the value of the string pointer passed in or
// "" if the pointer is nil.
func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}

// // CurrencyBRL - formata dinheiro em real
// func CurrencyBRL(cents int64) string {
// 	ac := accounting.Accounting{Symbol: "R$ ", Precision: 2, Thousand: ".", Decimal: ","}
// 	return fmt.Sprintf(ac.FormatMoney(float64(cents) / 100))
// }

// PadLimit - Pad string and limit to N
func PadLimit(input string, padLength int, padString string, padType string) string {
	var output string

	inputLength := len(input)
	padStringLength := len(padString)

	length := 0.0
	repeat := 0.0
	if inputLength < padLength {
		repeat = math.Ceil(float64(1) + (float64(padLength-padStringLength))/float64(padStringLength))
	} else {
		output = input
	}

	switch padType {
	case "RIGHT":
		if inputLength < padLength {
			output = input + strings.Repeat(padString, int(repeat))
		}
		output = output[:padLength]
	case "LEFT":
		if inputLength < padLength {
			output = strings.Repeat(padString, int(repeat)) + input
		}
		output = output[len(output)-padLength:]
	case "BOTH":
		if inputLength < padLength {
			length = (float64(padLength - inputLength)) / float64(2)
			repeat = math.Ceil(length / float64(padStringLength))
		}
		output = strings.Repeat(padString, int(repeat))[:int(math.Floor(float64(length)))] + input + strings.Repeat(padString, int(repeat))[:int(math.Ceil(float64(length)))]
	}

	return output
}

// MaskDocument - mascara um cpf/cnpj
func MaskDocument(cpfcnpj string) string {
	// Remove qualquer caractere não numérico
	cpfcnpj = regexp.MustCompile(`\D`).ReplaceAllString(cpfcnpj, "")

	// Verifica se é um CPF (11 dígitos) ou CNPJ (14 dígitos)
	if brdoc.IsCPF(cpfcnpj) {
		// Máscara para CPF: XXX***XXX** (coloca asteriscos nos caracteres do meio)
		return fmt.Sprintf("%s***%s**", cpfcnpj[:3], cpfcnpj[6:9])
	}

	// Máscara para CNPJ: XX***XXX****XX (coloca asteriscos nos caracteres do meio)
	return fmt.Sprintf("%s***%s****%s", cpfcnpj[:2], cpfcnpj[5:8], cpfcnpj[12:])
}

// DocumentFormat - formatar um cpf/cnpj
func DocumentFormat(value string) string {
	if brdoc.IsCPF(value) {
		return CPFFormat(value)
	}
	return CNPJFormat(value)
}

// CPFFormat - formatar uma string cpf
func CPFFormat(value string) string {
	value = regexp.MustCompile(`\D`).ReplaceAllString(value, "")
	if brdoc.IsCPF(value) {
		return regexp.MustCompile(`^(\d{1,3})\.?(\d{1,3})\.?(\d{1,3})-?(\d{1,2})$`).ReplaceAllString(value, "$1.$2.$3-$4")
	}
	return value
}

// CNPJFormat - formatar uma string cnpj
func CNPJFormat(value string) string {
	value = regexp.MustCompile(`\D`).ReplaceAllString(value, "")
	if brdoc.IsCNPJ(value) {
		return regexp.MustCompile(`^(\d{1,2})\.?(\d{1,3})\.?(\d{1,3})/?(\d{1,4})-?(\d{1,2})$`).ReplaceAllString(value, "$1.$2.$3/$4-$5")
	}
	return value
}

// Truncate - Truncate string
func Truncate(name string, limit int) string {
	result := name
	chars := 0
	for i := range name {
		if chars >= limit {
			result = name[:i]
			break
		}
		chars++
	}
	return result
}

// GenerateIdempotencyKey - gerando nova chave aleatoria
// init - inicia a chave com um valor
// limit - tamanho maximo da chave
func GenerateIdempotencyKey(init string, limit int) string {
	return Truncate(init+_uuid.New().String(), limit)
}

// RemoveSpecialCharacter - remove caracter esepcial
func RemoveSpecialCharacter(value string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9 ]+")
	return reg.ReplaceAllString(value, "")
}

// RemoveSpecialCharacterWithRegex - remove caracter esepcial
func RemoveSpecialCharacterWithRegex(value, regex string) string {
	reg, _ := regexp.Compile(regex)
	return reg.ReplaceAllString(value, "")
}

// TruncateDirection - truncate string por direccao
func TruncateDirection(name string, limit int, direction string) string {
	result := name
	chars := 0
	if direction == "RIGHT" {
		for i := range name {
			if chars >= limit {
				result = name[:i]
				break
			}
			chars++
		}
	}
	if direction == "LEFT" {
		runes := []rune(name)
		result = ""
		for i := len(name) - 1; i >= 0; i = i - 1 {
			if chars >= limit {
				break
			}
			result = string(runes[i]) + result
			chars++
		}
	}
	return result
}

// ToLike - link query
func ToLike(item string) string {
	return fmt.Sprintf("%%%s%%", item)
}

// BasicAuth - codifica base64
func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
