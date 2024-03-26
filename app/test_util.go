package app

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"text/tabwriter"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"grest.dev/grest"
)

const (
	testMainDB = "main_test.db"

	TestInvalidToken        = "invalidToken"
	TestForbiddenToken      = "forbiddenToken"
	TestReadOnlyToken       = "detail,list"
	TestCreateReadOnlyToken = "detail,list,create"
	TestEditReadOnlyToken   = "detail,list,edit"
	TestDeleteReadOnlyToken = "detail,list,delete"
	TestFullAccessToken     = "fullAccessToken"
)

func Test() *testUtil {
	if tu == nil {
		tu = &testUtil{}
		tu.configure()
	}
	return tu
}

var tu *testUtil

type testUtil struct {
	Tx *gorm.DB
}

func (t *testUtil) configure() {
	var err error
	conf := grest.DBConfig{}
	conf.Driver = DB_DRIVER
	conf.Host = DB_HOST
	conf.Port = DB_PORT
	conf.User = DB_USERNAME
	conf.Password = DB_PASSWORD

	t.Tx, err = DB().Conn(testMainDB)
	if err != nil {
		conf.DbName = testMainDB
		err = DB().Connect(testMainDB, conf)
		if err != nil {
			panic(err)
		}
		t.Tx, err = DB().Conn(testMainDB)
		if err != nil {
			panic(err)
		}
	}
}

func (t *testUtil) NewCtx(aclKeys []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := Ctx{
			mainTx: t.Tx,
			Lang:   "en",
			Action: Action{
				Method:   c.Method(),
				EndPoint: c.Path(),
			},
		}

		c.Locals(CtxKey, &ctx)
		return c.Next()
	}
}

// AssertMatchJSONElement checks if values are MatchElementJSON.
//
//	TODO :
//	1. cocokan masing-masing elemen json yang ada di expected
//	2. Untuk element yang ada di actual tapi tidak ada di expected maka diabaikan.
//	3. UUID bisa dicocokan dengan {uuid}
//	4. NullDate bisa dicocokan dengan {date} atau {current_date}
//	5. NullTime bisa dicocokan dengan {time} atau {current_time}
//	6. NullDateTime bisa dicocokan dengan {datetime} atau {current_datetime}
func (*testUtil) AssertMatchJSONElement(tb testing.TB, expected, actual []byte, description ...string) {
	if reflect.DeepEqual(expected, actual) {
		return
	}

	if tb != nil {
		tb.Helper()
	}

	testName := "AssertMatchJSONElement"
	if tb != nil {
		testName = tb.Name()
	}

	_, file, line, _ := runtime.Caller(1)

	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 5, ' ', 0)
	fmt.Fprintf(w, "\nTest:\t%s", testName)
	fmt.Fprintf(w, "\nTrace:\t%s:%d", filepath.Base(file), line)
	if len(description) > 0 {
		fmt.Fprintf(w, "\nDescription:\t%s", description[0])
	}
	fmt.Fprintf(w, "\nExpect:\t%v", expected)
	fmt.Fprintf(w, "\nResult:\t%v", actual)

	result := ""
	if err := w.Flush(); err != nil {
		result = err.Error()
	} else {
		result = buf.String()
	}

	if tb != nil {
		tb.Fatal(result)
	} else {
		log.Fatal(result)
	}
}
