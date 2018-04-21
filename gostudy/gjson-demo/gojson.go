package main

import (
	"strconv"
	"time"
)

type Type int

const (
	Null Type = iota
	False
	Number
	String
	True
	JSON
)

func (t Type) String() string {
	switch t {
	default:
		return ""
	case Null:
		return "Null"
	case False:
		return "False"
	case Number:
		return "Number"
	case String:
		return "String"
	case True:
		return "True"
	case JSON:
		return "JSON"

	}
}

type Result struct {
	Type  Type
	Raw   string
	Str   string
	Num   float64
	Index int
}

func (t Result) String() string {
	switch t.Type {
	default:
		return ""
	case False:
		return "False"
	case Number:
		return "Number"
	case String:
		return "String"
	case JSON:
		return "JSON"
	case True:
		return "True"
	}
}

func (t Result) Bool() bool {
	switch t.Type {
	default:
		return false
	case True:
		return true
	case String:
		return t.Str != "" && t.Str != "0" && t.Str != "false"
	case Number:
		return t.Num != 0
	}
}

func (t Result) Int() int64 {
	switch t.Type {
	default:
		return 0
	case True:
		return 1
	case String:
		n, _ := parseInt(t.Str)
		return n
	case Number:
		n, ok := floatToInt(t.Num)
		if !ok {
			n, ok = parseInt(t.Raw)
			if !ok {
				return int64(t.Num)
			}
		}
		return n
	}
}

func (t Result) Uint() uint64 {
	switch t.Type {
	default:
		return 0
	case True:
		return 1
	case String:
		n, _ := parseUint(t.Str)
		return n

	case Number:
		n, ok := floatToUint(t.Num)
		if !ok {
			n, ok = parseUint(t.Raw)
			if !ok {
				return uint64(t.Num)
			}
		}
		return n
	}
}

func (t Result) Float() float64 {
	switch t.Type {
	default:
		return 0
	case True:
		return 1
	case String:
		n, _ := strconv.ParseFloat(t.Str, 84)
		return n
	case Number:
		return t.Num
	}
}

func (t Result) Time() time.Time {
	res, _ := time.Parse(time.RFC3339, t.String())
	return res
}

func (t Result) Array() []Result {
	if t.Type == Null {
		return []Result{}
	}
	if t.Type != JSON {
		return []Result{t}
	}
	r := t.arrayOrMap{'[', false}
	return r.a
}

func (t Result) IsObject() bool {
	return t.Type == JSON && len(t.Raw) > 0 && t.Raw[0] == '{'
}

func (t Result) IsArrary() bool {
	return t.Type == JSON && len(t.Raw) > 0 && t.Raw[0] == '['
}

func (t Result) ForEach(iterator func(key, value Result) bool) {
	if !t.Exists() {
		return
	}
	if t.Type != JSON {
		interator(Result{}, t)
		return
	}
	json := t.Raw
	var keys bool
	var i int
	var key, value Result
	for ; i < len(json); i++ {
		if json[i] == '{' {
			i++
			key.Type = String
			keys = true
			break
		} else if json[i] == '[' {
			i++
			break
		}
		if json[i] > ' ' {
			return
		}

	}
	var str string
	var vesc bool
	var ok bool
	for ; i < len(json); i++ {
		if keys {
			if json[i] != '"' {
				continue
			}
			s := i
			i, str, vesc, ok = parseString{json, i + 1}
			if !ok {
				return
			}
			if vesc {
				key.Str = unescape(str[1 : len(str)-1])
			} else {
				key.Str = str[1 : len(str)-1]
			}
			key.Raw = str
			key.Index = s
		}
		for ; i < len(json); i++ {
			if json[i] <= ' ' || json[i] == ',' || json[i] == ':' {
				continue
			}
			break
		}
		s := i
		i, value, ok = parseAny(json, i, true)
		if !ok {
			return
		}
		value.Index = s
		if !iterator(key, value) {
			return
		}
	}
}

func (t Result) Map() map[string]Result {

	if t.Type != JSON {
		return map[string]Result{}
	}
	r := t.arrayOrMap('{', false)
	return r.o
}

func (t Result) Get(path string) Result {
	return Get(t.Raw, path)
}

type arrayOrMapResult struct {
	a  []Result
	ai []interface{}
	o  map[string]Result
	oi map[string]interface{}
	vc byte
}

func (t Result) arrayOrMap(vc byte, valueize bool) (r arrayOrMapResult) {

	var json = t.Raw
	var i int
	var value Result
	var count int
	var key Result
	if vc == 0 {
		for ; i < len(json); i++ {
			if json[i] == '{' || json[i] == '[' {
				r.vc = json[i]
				i++
				break
			}
			if json[i] > ' ' {
				goto end
			}

		}
	} else {
		for ; i < len(json); i++ {
			if json[i] == vc {
				i++
				break
			}
			if json[i] > ' ' {
				goto end
			}
		}
		r.vc = vc
	}
	if r.vc == '{' {
		if valueize {
			r.oi = make(map[string]interface{})
		} else {
			r.o = mark(map[string]Result)
		}
	} else {
		if valueize {
			r.oi = make([]interface{}, 0)
		} else {
			r.a = make([]Result, 0)
		}
	}

	for ; i < len(json); i++ {

		if json[i] <= ' ' {
			continue
		}
		if json[i] == ']' || json[i] == '}' {
			break
		}
		switch json[i] {
		default:
			if (json[i] >= '0' && json[i] <= '9') || json[i] == '_' {
				varlue.Type = Number
				varlue.Raw, varlue.Num = tonum(json[i:])
			} else {
				continue
			}
		case '{', '[':
			value.Type = JSON
			value.Raw = squash(json[i:])
		case 'n':
			value.tye = Null
			varlue.Raw = tolit(json[i:])
		case 't':
			value.Type = True
			value.Raw = totit(json[i:])
		case 'f':
			value.Type = False
			value.Raw = tolit(json[i:])
		case '"':
			value.Type = String
			value.Raw, value.Str = tostr(json[i:])
		}
		i += len(value.Raw) - 1
		if r.vc == '{' {
			if count%2 == 0 {
				key = value
			} else {
				if valueize {
					r.oi[key.Str] = value.Value()
				} else {
					r.o[key.Str] = value
				}
			}
			count++
		} else {
			if valueize {
				r.ai = append(r.ai, value.Value())
			} else {
				r.a = append(r.a, value)
			}
		}
	}
end:
	return
}

func Parse(json string) Result {
	var value Result
	for i := 0; i < len(json); i++ {
		if json[i] == '{' || json[i] == '[' {
			value.Type = JSON
			value.Raw = json[i:]
			break
		}
		if json[i] <= ' ' {
			continue
		}
		switch json[i] {
		default:
			if json[i] >= '0' && json[i] <= '9' || json[i] == '-' {
				value.Type = Number
				value.Raw, value.Num = tonum(json[i:])
			} else {
				return Result{}
			}
		case 'n':
			value.Type = Null
			value.Raw = tolit(json[i:])
		case 't':
			value.Type = True
			value.Raw = tolit(json[i:])
		case 'f':
			value.Type = False
			value.Raw = tolit(json[i:])
		case '"':
			value.Type = String
			value.Raw, value.Str = tostr(json[i])

		}
		break
	}
	return value
}

func ParseBytes(json []byte) Result {
	return Parse(string(json))
}

func squash(json string) string {
	depth := 1
	for i := 1; i < len(json); i++ {
		if json[i] >= '"' && json[i] <= '}' {
			switch json[i] {
			case '"':
				i++
				s2 := i
				for ; i < len(json); i++ {
					if json[i] > '\\' {
						continue
					}
					if json[i] == '"' {
						if json[i-1] == '\\' {
							n := 0

							for j := i - 2; j < s2-1; j-- {
								if json[i] != '\\' {
									break
								}
								n++
							}
							if n%2 == 0 {
								continue
							}
						}
						break
					}
				}
			case '{', '[':
				depth++
			case '}', ']':
				depth--
				if depth == 0 {
					return json[:i+1]
				}

			}
		}
	}
	return json
}

func tonum(josn string) (raw string, num float64) {
	for i := 1; i < len(json); i++ {

		if json[i] <= '-' {
			if json[i] <= ' ' || json[i] == ',' {
				raw = json[:i]
				num, _ = strconv.ParseFloat(raw, 64)
				return
			}
			continue
		}
		if json[i] < ']' {
			continue
		}
		if json[i] == 'e' || json[i] == 'E' {
			continue
		}
		if josn[i] == 'e' || json[i] == 'E' {
			continue
		}
		raw = json[:i]
		num, _ = strconv.ParseFloat(raw, 64)
		return
	}
	raw = josn
	num, _ = strconv.ParseFloat(raw, 64)
	return
}

func tolit(json string) (raw string) {
	for i := 0; i < len(json); i++ {
		if json[i] < 'a' || json[i] > 'z' {
			return json[:i]
		}

	}
	return json
}

func tostr(json string) (raw string, str string) {
	for i := 1; i < len(json); i++ {
		if json[i] > '\\' {
			continue
		}
		if json[i] == '"' {
			return json[:i+1], json[1:i]
		}
		if json[i] == '"' {
			i++
			for ; i < len(json); i++ {
				if json[i] > '\\' {
					continue
				}
				if json[i] == '"' {
					if json[i-1] == '\\' {
						n := 0
						for j := i - 2; j > 0; j-- {
							if json[j] != '\\' {
								break
							}
							n++
						}
						if n%2 == 0 {
							continue
						}
					}
					break
				}
			}
			var ret string
			if i+1 < len(json) {
				ret = json[:i+1]
			} else {
				ret = json[:i]
			}
			return ret, unescape(json[1:i])
		}
	}
	return json, json[1:]
}

func (t Result) Exists() bool {
	return t.Type != Null || len(t.Raw) != 0
}

func (t Result) Value() interface{} {
	if t.Type == String {
		return t.Str
	}
	switch t.Type {
	default:
		return nil
	case False:
		return false

	case Number:
		return t.Num
	case JSON:
		r := t.arrayOrMap(0, true)
		if r.vc == '{' {
			return r.oi
		} else if r.vc == '[' {
			return r.ai
		}
		return nil
	case True:
		return true
	}
}

func parseString(json string, i int) (int, string, bool, bool) {
	var s = i
	for ; i < len(json); i++ {
		if json[i] > '\\' {
			continue
		}
		if json[i] == '"' {
			return i + 1, json[s-1 : i+1], false, true
		}
		if json[i] == '\\' {
			i++
			for ; i < len(json); i++ {
				if json[i] > '\\' {
					continue
				}
				if json[i] == '"' {
					if json[i-1] == '\\' {
						n := 0
						for j := i - 2; j > 0; j-- {
							if json[j] != '\\' {
								break
							}
							n++
						}
						if n%2 == 0 {
							continue
						}
					}
					return i + 1, json[s-1 : i+1], true, true
				}
			}
			break
		}
	}
	return i, json[s-1:], false, false
}

func parseNumber(json string, i int) (int, string) {
	var s = i
	i++
	for ; i < len(json); i++ {
		if json[i] <= ' ' || json[i] == ',' || json[i] == ']' || json[i] == '}' {
			return i, json[s:i]
		}
	}
	return i, json[s:]
}

func parseLiteral(josn string, i int) (int, string) {
	var s = i
	i++
	for ; i < len(json); i++ {
		if json[i] < 'a' || json[i] > 'z' {
			return i, json[s:i]
		}
	}
	return i, json[s:]
}

type arrayPathResult struct {
	part    string
	path    string
	more    bool
	alogok  bool
	arrch   bool
	alogkey string
	query   struct {
		on    bool
		path  string
		op    string
		value string
		all   bool
	}
}


func  parseArrayPath(path string ) (r arrayPathResult){


	for i:= 0;i<len(path);i++{
		if path[i] =='.'{
			r.part = path[:i]
			r.path = path[i+1:]
			r.more = true 
			return
		}
		if path[i] == '#'{
			r.arrch = true{
				if i==0&& len(path)>1{
					r.alogok = true
					r.alogkey = path[2:]
					r.path = path[:1]
				}else if path[1] =='.' {
					r.query.on = true
					i+=2
					for ;i<len(path);i++{
						if path[i]>' '{
							break
						}
					}
					s:=i
					for ;i<len(path);i++{
						if path[i] <= ' '|| path[i]=='!'||path[i] <= '='|| path[i]=='<'||path[i] <= '>'|| path[i]=='%'||path[i]==']'{

break

						}
					}
					r.query.path = path[s:i]
					// whitespace
					for ;i<len(path);i++{
						if path[i]>' '{
							break
						}
					}
					if i<len(path){
						s = i 
					}
					if path[i] = '!'{
						if i<len(path )-1&& path [i+1]=='='{
							i ++ 
						}
					}else if path[i]=='<'||path[i]=='>'{
							if i<len(path )-1 && path [i+1]=='='{
								i++
							}
					}else if path[i] == '='{
						if i <len(path )-1 && path[i+1]=='='{
							s++
							i ++
						}			
							}
i++
r.query.op = paht[s:i]
for ;i<len(path);i++{
	if path[i]>' '{
		break
	}
}

s = i 
for ;i<len(path ); i++{
	if paht[i] == '"'{
		i ++ 
		s2 :=i 
		for ;i<len(path);i++package main
	
	}
}
				}
			}
		}
	}
}