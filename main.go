package main

import (
	"fmt"

	"github.com/Lesspion/user-manager-api/Core"
)

// import (
// 	"encoding/json"
//     "fmt"
//     "log"
//     "net/http"

// 	"gopkg.in/mgo.v2"
//     "gopkg.in/mgo.v2/bson"
//     "github.com/spf13/viper"
// )

func main() {
	conf := Core.GetConfig("./config.toml")
	fmt.Printf("%+v\n", conf)
	// session, err := mgo.Dial("localhost")
}
