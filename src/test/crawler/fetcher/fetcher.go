package fetcher

//获取器
import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	_ "io"
	"io/ioutil"
	"log"
	"net/http"
)

//自动检查字符编码
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	//将前1024字节给到自动检查编码函数，为了不使字节移动，使用bufio 的peek方法
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func Fetch(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close() //关闭body体

	//返回错误的状态码
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code : %d", response.StatusCode)
	}
	//如果是gbk编码的 utf8Reader := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())
	bodyReader := bufio.NewReader(response.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader) //读取全部body

}
