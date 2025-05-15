package api

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	netUrl "net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/wgngoo/ops-cmdb/pkg/cmdb_sdk_go/utils"
)

// NewCmdbClient 创建新的 CMDB 客户端
func NewCmdbClient(config *Config) *CmdbClient {
	return &CmdbClient{Config: *config}
}

// NewResourceService 创建新的资源服务
func NewResourceService(client *CmdbClient) *ResourceService {
	return &ResourceService{cmdbClient: client}
}

// String 字符串指针辅助函数
func String(s string) *string {
	return &s
}

// StringValue 获取字符串指针的值
func StringValue(a *string) string {
	if a == nil {
		return ""
	}
	return *a
}

// sort 对参数进行升序排序
func (client *CmdbClient) sort(params map[string]interface{}) string {
	if params == nil {
		return ""
	}

	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	var result strings.Builder
	for _, key := range keys {
		result.WriteString(fmt.Sprintf("%s=%v&", key, params[key]))
	}

	paramsStr := result.String()
	return paramsStr[:len(paramsStr)-1]
}

// signature 生成签名
func (client *CmdbClient) signature(method, url, contentType string, expires int64, data map[string]interface{}) (string, error) {
	var params, contentMd5 string

	if contentType == "" {
		contentType = "application/json"
	}

	if method == http.MethodGet {
		params = client.sort(data)
		params = strings.ReplaceAll(params, "=", "")
		params = strings.ReplaceAll(params, "&", "")
	}

	if method == http.MethodPost || method == http.MethodPut {
		ds, err := json.Marshal(data)
		if err != nil {
			return "", err
		}

		hash := md5.Sum([]byte(ds))
		contentMd5 = hex.EncodeToString(hash[:])
	}

	signStr := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%d\n%s", method, url, params, contentType, contentMd5, expires, StringValue(client.AccessKey))

	h := hmac.New(sha1.New, []byte(StringValue(client.SecretKey)))
	h.Write([]byte(signStr))
	hashed := h.Sum(nil)
	signature := hex.EncodeToString(hashed)

	return signature, nil
}

// Do 执行 HTTP 请求
func (client *CmdbClient) Do(method, url string, data map[string]interface{}) (string, error) {
	expires := time.Now().Unix()
	signature, err := client.signature(method, url, "", expires, data)
	if err != nil {
		return "", err
	}

	if StringValue(client.Host) == "" {
		return "", errors.New("param err: host is nil")
	}

	Url, err := netUrl.Parse(fmt.Sprintf("https://%s%s", StringValue(client.Host), url))
	if err != nil {
		return "", err
	}

	params := netUrl.Values{}
	params.Add("accesskey", StringValue(client.AccessKey))
	params.Add("signature", signature)
	params.Add("expires", strconv.FormatInt(expires, 10))
	Url.RawQuery = params.Encode()

	if method == http.MethodGet {
		s := client.sort(data)
		if s != "" {
			Url.RawQuery = s + "&" + params.Encode()
		}
	}
	url = Url.String()

	bytesData, _ := json.Marshal(data)
	req, err := http.NewRequest(method, url, bytes.NewReader(bytesData))
	if err != nil {
		return "", err
	}

	req.Host = "openapi.easyops-only.com"
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	return string(res), err
}

// SearchInstanceV3Page 搜索实例
func (client *CmdbClient) SearchInstanceV3Page(objectId string, req *InstanceApi_PostSearchV3RequestBody) (string, error) {
	url := fmt.Sprintf("/cmdbservice/v3/object/%s/instance/_search", objectId)
	byteReq, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	data := make(map[string]interface{})
	if err = json.Unmarshal(byteReq, &data); err != nil {
		return "", err
	}

	resp, err := client.Do(http.MethodPost, url, data)
	if err != nil {
		return "", err
	}
	return resp, nil
}

// GetHostForCmdb 获取主机信息
func (s *ResourceService) GetCmdbResource(objectId string, fields []string, query map[string]interface{}) (*CmdbListResponse, error) {
	response, err := s.fetchCmdbResources(objectId, fields, query)
	if err != nil {
		return nil, err
	}
	fmt.Println("response: ", string(response.Data.List))
	return response, nil
}

func (s *ResourceService) CheckDomainAvailability(domainName string) (bool, error) {
	// 设置查询参数
	objectId := "DNSRECORD"
	fields := []string{"name"}

	// 构建查询条件
	queryStr := fmt.Sprintf(`{
		"$and": [{
			"$or": [{
				"name": {
					"$eq": "%s"
				}
			}]
		}]
	}`, domainName)

	query, err := utils.StringFormat(queryStr)
	if err != nil {
		log.Fatalf("解析查询条件失败: %v", err)
		return false, err
	}
	result, err := s.GetCmdbResource(objectId, fields, query)
	if err != nil {
		log.Fatalf("获取主机信息失败: %v", err)
		return false, err
	}
	if len(result.Data.List) > 0 {
		return true, nil
	}
	return false, nil
}

// fetchCmdbResources 获取 CMDB 资源
func (s *ResourceService) fetchCmdbResources(objectId string, fields []string, query map[string]interface{}) (*CmdbListResponse, error) {
	pageSize := 1000
	currentPage := 1
	var allData CmdbListResponse
	allData.Data.List = json.RawMessage("[]")

	for {
		cmdbResponse, err := s.cmdbClient.SearchInstanceV3Page(objectId, &InstanceApi_PostSearchV3RequestBody{
			Fields:   fields,
			Page:     currentPage,
			PageSize: pageSize,
			Query:    query,
		})
		if err != nil {
			return nil, err
		}

		var response CmdbListResponse
		if err := json.Unmarshal([]byte(cmdbResponse), &response); err != nil {
			return nil, err
		}

		if currentPage == 1 {
			allData = response
		} else {
			var existingList, newList []json.RawMessage
			json.Unmarshal(allData.Data.List, &existingList)
			json.Unmarshal(response.Data.List, &newList)
			existingList = append(existingList, newList...)
			allData.Data.List, _ = json.Marshal(existingList)
		}

		if currentPage*pageSize >= response.Data.Total {
			break
		}
		currentPage++
	}

	return &allData, nil
}
