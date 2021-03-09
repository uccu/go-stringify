package stringify_test

import (
	"log"
	"testing"

	. "github.com/uccu/go-stringify"
)

func TestBool(t *testing.T) {
	s := ToString(true)
	if s != "true" {
		t.Error("err")
	}
	log.Println("success", s)

	s = ToString(false)
	if s != "false" {
		t.Error("err")
	}
	log.Println("success", s)
}

func TestString(t *testing.T) {
	s := ToString("1d23")
	if s != "1d23" {
		t.Error("err")
	}
	log.Println("success", s)
}

func TestInt(t *testing.T) {
	s := ToString(int8(-128))
	if s != "-128" {
		t.Error("err")
	}
	log.Println("success", s)

	s = ToString(int8(127))
	if s != "127" {
		t.Error("err")
	}
	log.Println("success", s)

	s = ToString(int16(-32768))
	if s != "-32768" {
		log.Println("success", s)
		t.Error("err")
	}

	s = ToString(int16(32767))
	if s != "32767" {
		t.Error("err")
	}
	log.Println("success", s)

	s = ToString(int32(-2147483648))
	if s != "-2147483648" {
		t.Error("err")
	}
	log.Println("success", s)

	s = ToString(int32(2147483647))
	if s != "2147483647" {
		t.Error("err")
	}
	log.Println("success", s)

	s = ToString(int64(-9223372036854775808))
	if s != "-9223372036854775808" {
		t.Error("err")
	}
	log.Println("success", s)

	s = ToString(int64(9223372036854775807))
	if s != "9223372036854775807" {
		t.Error("err")
	}
	log.Println("success", s)

	s = ToString(int(9223372036854775807))
	if s != "9223372036854775807" {
		t.Error("err")
	}
	log.Println("success", s)

	s = ToString(int(-9223372036854775808))
	if s != "-9223372036854775808" {
		t.Error("err")
	}
	log.Println("success", s)
}

func TestUint(t *testing.T) {
	s := ToString(uint8(255))
	if s != "255" {
		t.Error("err")
	}
	log.Println("success", s)

	s = ToString(uint16(65535))
	if s != "65535" {
		t.Error("err")
	}
	log.Println("success", s)

	s = ToString(uint32(4294967295))
	if s != "4294967295" {
		t.Error("err")
	}
	log.Println("success", s)

	s = ToString(uint64(18446744073709551615))
	if s != "18446744073709551615" {
		t.Error("err")
	}
	log.Println("success", s)

	s = ToString(uintptr(18446744073709551615))
	if s != "18446744073709551615" {
		t.Error("err")
	}
	log.Println("success", s)

	s = ToString(uint(18446744073709551615))
	if s != "18446744073709551615" {
		t.Error("err")
	}
	log.Println("success", s)
}

type W struct{}

func (v W) String() string {
	return "1"
}

type E string
type F string

func (v E) String() string {
	return "2"
}
func TestStruct(t *testing.T) {
	s := ToString(struct{}{})
	if s != "" {
		t.Error("err")
	}
	log.Println("success", s)

	s = ToString(&W{})
	if s != "1" {
		t.Error("err")
	}
	log.Println("success", s)

	var e E = "111daw"

	s = ToString(e)
	if s != "2" {
		t.Error("err")
	}
	log.Println("success", s)

	var f F = "111daw"

	s = ToString(f)
	if s != "111daw" {
		t.Error("err")
	}
	log.Println("success", s)
}

func TestPointer(t *testing.T) {
	var i = 1
	s := ToString(&i)
	if s != "1" {
		t.Error("err")
	}
	log.Println("success", s)
}

type I int

func TestSame(t *testing.T) {
	var i I = 12
	s := ToString(&i)
	if s != "12" {
		t.Error("err")
	}
	log.Println("success", s)
}

func TestFloat(t *testing.T) {
	var i float64 = 123123.1241231204
	s := ToString(&i)
	if s != "123123.1241231204" {
		t.Error("err")
	}
	log.Println("success", s)

	i = 123123
	s = ToString(&i)
	if s != "123123" {
		t.Error("err")
	}
	log.Println("success", s)
}

func TestSlice(t *testing.T) {

	s := ToString([]string{"1", "2", "3"})
	if s != "1,2,3" {
		t.Error("err")
	}
	log.Println("success", s)

	s = ToString([]string{"1", "2", "3"}, "-")
	if s != "1-2-3" {
		t.Error("err")
	}
	log.Println("success", s)

}
