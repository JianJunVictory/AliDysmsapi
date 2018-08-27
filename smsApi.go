// 参考https://help.aliyun.com/document_detail/56189.html?spm=a2c4g.11186623.6.583.611c7cf6UU3WdX,实现阿里云短信接口
package dysmsapi
import (
	"sort"
	"fmt"
	"time"
	"math/rand"
	"net/http"
	"net/url"
	"encoding/json"
	"crypto/hmac"
	"encoding/base64"
	"crypto/sha1"
	"strings"
)
type SendSmsReply struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
	BizId string `json:"BizId"`
	RequestId string `json:"RequestId"`
}
/**
* @param phone 发送短信的手机号
* @param accessKeyId 开通阿里短信服务时的秘钥Id
* @param accessKeySecret 开通阿里短信服务时的秘钥
* @param TemplateParam 短信模版
* @param TemplateCode 短信编码
* @param SignName 短信签名
* @return SendSmsReply 结构体实例
* @return error 错误信息
*/
func SendSms(phone, accessKeyId, accessKeySecret, TemplateParam,TemplateCode,SignName string) (*SendSmsReply,error) {
	// 第一步：请求参数
	paras := make(map[string]string)
	// 系统参数
	paras["SignatureMethod"]= "HMAC-SHA1"
	paras["SignatureNonce"]=fmt.Sprintf("%d", rand.Int63())
	paras["AccessKeyId"]= accessKeyId
	paras["SignatureVersion"]= "1.0"
	paras["Timestamp"]=time.Now().UTC().Format("2006-01-02T15:04:05Z")
	paras["Format"]= "JSON";
	// 业务API参数
	paras["Action"]= "SendSms"
	paras["Version"]="2017-05-25"
	paras["RegionId"]= "cn-hangzhou"
	paras["PhoneNumbers"]= phone
	paras["SignName"]=SignName
	paras["TemplateParam"]=TemplateParam
	paras["TemplateCode"]=TemplateCode
	paras["OutId"]="yourOutId"
	
	// 第二步：根据参数Key排序（顺序）
	var keys []string
	for k := range paras {
        keys = append(keys, k)
    }
	sort.Strings(keys)

	//第三步：构造待签名的请求串
	var sortedQueryString  string
	for _, v := range keys {
		sortedQueryString  = fmt.Sprintf("%s&%s=%s", sortedQueryString , specialUrlEncode(v), specialUrlEncode(paras[v]))
	}
	stringToSign := fmt.Sprintf("GET&%s&%s", specialUrlEncode("/"), specialUrlEncode(sortedQueryString [1:]))
	
	signature := sign(accessKeySecret,stringToSign)

	apiUrl := fmt.Sprintf("http://dysmsapi.aliyuncs.com/?Signature=%s%s", signature, sortedQueryString )

	resp, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println(err)
		return nil,err
	}
	ssr := new(SendSmsReply)
	json.NewDecoder(resp.Body).Decode(ssr)
	return ssr,nil
}

//签名函数
func sign(accessKeySecret,stringToSign string)string{
	h :=hmac.New(sha1.New,[]byte(fmt.Sprintf("%s&", accessKeySecret)))
	h.Write([]byte(stringToSign))
	str := specialUrlEncode(base64.StdEncoding.EncodeToString(h.Sum(nil)))
	return str
}
// url处理函数
func specialUrlEncode(in string) string {
	rep := strings.NewReplacer("+", "%20", "*", "%2A", "%7E", "~")
	return rep.Replace(url.QueryEscape(in))
}