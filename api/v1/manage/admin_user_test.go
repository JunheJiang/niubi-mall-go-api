package manage

import (
	"github.com/gavv/httpexpect"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestAdminUserApi_CreateAdminUser(t *testing.T) {

}

func Test(t *testing.T) {
	t.Run("TestBasic1", TestBasic1)
	t.Run("TestBasic2", TestBasic2)
	t.Run("TestBasic3", TestBasic3)
}

func TestMain(m *testing.M) {
	tearDownAll := setUpAll()
	code := m.Run()
	tearDownAll() // you cannot use defer tearDownAll()
	os.Exit(code)
}

func setUpAll() func() {
	log.Printf("LLLLLLLL: setUpAll")
	return func() {
		log.Printf("LLLLLLLL: tearDownAll")
	}
}

func setUp(t *testing.T) func(t *testing.T) {
	log.Printf("LLLLLLLL: setUp")
	return func(t *testing.T) {
		log.Printf("LLLLLLLL: tearDown")
	}
}

func TestBasic1(t *testing.T) {
	tearDown := setUp(t)
	defer tearDown(t)
	log.Printf("LLLLLLLL: TestBasic1")
}

func TestBasic2(t *testing.T) {
	tearDown := setUp(t)
	defer tearDown(t)
	log.Printf("LLLLLLLL: TestBasic2")
}

func TestBasic3(t *testing.T) {
	tearDown := setUp(t)
	defer tearDown(t)
	log.Printf("LLLLLLLL: TestBasic3")
}

var testUrl string = "http://127.0.0.1:12345"

func TestHttpGetPass(t *testing.T) {
	e := httpexpect.New(t, testUrl) //创建一个httpexpect实例
	e.GET("/checkon").              //ge请求
					Expect().
					Status(http.StatusOK). //判断请求是否200
					JSON().
					Object().                   //json body实例化
					ContainsKey("msg").         //检验是否包括key
					ValueEqual("msg", "online") //对比key的value
}

func TestHttpGetFail(t *testing.T) {
	e := httpexpect.New(t, testUrl) //创建一个httpexpect实例
	e.GET("/checkon").              //ge请求
					Expect().
					Status(http.StatusOK). //判断请求是否200
					JSON().
					Object().                   //json body实例化
					ContainsKey("msg1").        //检验是否包括key,使用不存在的key
					ValueEqual("msg", "online") //对比key的value
}

func TestHttpPostPass(t *testing.T) {
	e := httpexpect.New(t, testUrl)     //创建一个httpexpect实例
	postdata := map[string]interface{}{ //创建一个json变量
		"flag": 1,
		"msg":  "terrychow",
	}
	contentType := "application/json;charset=utf-8"

	e.POST("/postdata"). //post 请求
				WithHeader("ContentType", contentType). //定义头信息
				WithJSON(postdata).                     //传入json body
				Expect().
				Status(http.StatusOK). //判断请求是否200
				JSON().
				Object().                      //json body实例化
				ContainsKey("msg").            //检验是否包括key
				ValueEqual("msg", "terrychow") //对比key的value

}

func TestHttpPostFail(t *testing.T) {
	e := httpexpect.New(t, testUrl)     //创建一个httpexpect实例
	postData := map[string]interface{}{ //创建一个json变量
		"flag": 1,
		"msg":  "terrychow",
	}
	contentType := "application/json;charset=utf-8"

	e.POST("/postdata"). //post 请求
				WithHeader("ContentType", contentType). //定义头信息
				WithJSON(postData).                     //传入json body
				Expect().
				Status(http.StatusOK). //判断请求是否200
				JSON().
				Object().                      //json body实例化
				ContainsKey("msg").            //检验是否包括key
				ValueEqual("msg", "terryzhou") //对比key的value，value不匹配

}
