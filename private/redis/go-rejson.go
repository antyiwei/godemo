package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	// "github.com/garyburd/redigo/redis"
	"github.com/gomodule/redigo/redis"
	rejson "github.com/nitishm/go-rejson"
)

var addr = flag.String("Server", "localhost:6379", "Redis server address")

type Name struct {
	First  string `json:"first,omitempty"`
	Middle string `json:"middle,omitempty"`
	Last   string `json:"last,omitempty"`
}

type Student struct {
	Name Name `json:"name,omitempty"`
	Rank int  `json:"rank,omitempty"`
}

func main() {
	// flag.Parse()

	// conn, err := redis.Dial("tcp", *addr)
	// if err != nil {
	// 	log.Fatalf("Failed to connect to redis-server @ %s", *addr)
	// }
	// defer func() {
	// 	conn.Do("FLUSHALL")
	// 	conn.Close()
	// }()

	// student := Student{
	// 	Name: Name{
	// 		"Mark",
	// 		"S",
	// 		"Pronto",
	// 	},
	// 	Rank: 1,
	// }
	// res, err := rejson.JSONSet(conn, "student", ".", student, false, false)
	// if err != nil {
	// 	log.Fatalf("Failed to JSONSet")
	// 	return
	// }

	// log.Printf("Success if - %s\n", res)

	// studentJSON, err := redis.Bytes(rejson.JSONGet(conn, "student", ""))
	// if err != nil {
	// 	log.Fatalf("Failed to JSONGet")
	// 	return
	// }

	// readStudent := Student{}
	// err = json.Unmarshal(studentJSON, &readStudent)
	// if err != nil {
	// 	log.Fatalf("Failed to JSON Unmarshal")
	// 	return
	// }

	// log.Printf("Student read from redis : %#v\n", readStudent)

	// log.Println(" ====== test json.set =========")
	// nameByte, err := redis.Bytes(rejson.JSONGet(conn, "student", ".name.last"))
	// if err != nil {
	// 	log.Fatalf("Failed to JSONGet ")
	// 	return
	// }
	// log.Println(" ====== test result =========")
	// fmt.Println(fmt.Sprintf("%s", string(nameByte)))
	// log.Println(" ============================")
	testMain()
}

type StudentInfo struct {
	Info *StudentDetails `json:”info,omitempty”`
	Rank int             `json:”rank,omitempty”`
}
type StudentDetails struct {
	FirstName string
	LastName  string
	Major     string
}

func testMain() {
	flag.Parse()

	conn, err := redis.Dial("tcp", *addr)
	if err != nil {
		log.Fatalf("Failed to connect to redis-server @ %s", *addr)
	}
	defer func() {
		// conn.Do("FLUSHALL")
		conn.Close()
	}()

	studentJD := StudentInfo{
		Info: &StudentDetails{
			FirstName: "John",
			LastName:  "Doe",
			Major:     "CSE",
		},
		Rank: 1,
	}
	b, err := json.Marshal(&studentJD)
	if err != nil {
		return
	}
	// _, err = conn.Do(“SET”, “JohnDoe”, string(b))
	// if err != nil {
	//     return
	// }

	fmt.Println(string(b))

	_, err = rejson.JSONSet(conn, "JohnDoeJSON", ".", studentJD, false, false)
	if err != nil {
		return
	}

	outJSON, err := rejson.JSONGet(conn, "JohnDoeJSON", "")
	if err != nil {
		return
	}
	outStudent := &StudentInfo{}
	err = json.Unmarshal(outJSON.([]byte), outStudent)
	if err != nil {
		return
	}
	v, _ := json.Marshal(outStudent)
	fmt.Println(string(v))
	fmt.Println(fmt.Sprintf("Rank:%d", outStudent.Rank))
	fmt.Println(fmt.Sprintf("Info:%s", outStudent.Info))
	fmt.Println(fmt.Sprintf("Info.FirstName:%s", outStudent.Info.FirstName))
	fmt.Println(fmt.Sprintf("Info.LastName:%s", outStudent.Info.LastName))
	fmt.Println(fmt.Sprintf("Info.Major:%s", outStudent.Info.Major))

	//======= test json.set =====
	log.Println(" ====== test json.set =========")
	_, err = rejson.JSONSet(conn, "JohnDoeJSON", ".Info.Major", "EE", false, false)
	if err != nil {
		return
	}
	outJSON, err = rejson.JSONGet(conn, "JohnDoeJSON", "")
	if err != nil {
		return
	}
	outStudent = &StudentInfo{}
	err = json.Unmarshal(outJSON.([]byte), outStudent)
	if err != nil {
		return
	}
	v, _ = json.Marshal(outStudent)
	fmt.Println(string(v))
	fmt.Println(fmt.Sprintf("Rank:%d", outStudent.Rank))
	fmt.Println(fmt.Sprintf("Info:%s", outStudent.Info))
	fmt.Println(fmt.Sprintf("Info.FirstName:%s", outStudent.Info.FirstName))
	fmt.Println(fmt.Sprintf("Info.LastName:%s", outStudent.Info.LastName))
	fmt.Println(fmt.Sprintf("Info.Major:%s", outStudent.Info.Major))

}
