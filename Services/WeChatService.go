package Services

import (
	db "brocaedu/Database"
	"brocaedu/Models/Article"
	"brocaedu/Models/Site"
	"brocaedu/Pkg/e"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

var wt sync.WaitGroup

// @Summer 获取token
func GetToken() (string, error) {
	u, err := url.Parse("https://api.weixin.qq.com/cgi-bin/token")

	if err != nil {
		log.Fatal(err)
	}

	weChatConfig := Site.GetWeChatConfig()

	parse := url.Values{}
	parse.Set("grant_type", weChatConfig.GRANTTYPE)
	parse.Set("appid", "wx86181b3a0022cc1f")
	parse.Set("secret", "9fdeab8a71dfdd8ae9db13ee3db7fbcf")
	u.RawQuery = parse.Encode()

	resp, err := http.Get(u.String())

	jMap := make(map[string]interface{})

	if err != nil {
		return "", errors.New("request token err :" + err.Error())
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	err = json.NewDecoder(resp.Body).Decode(&jMap)
	if err != nil {
		return "", errors.New("request token response json parse err :" + err.Error())
	}

	if jMap["errcode"] == nil || jMap["errcode"] == 0 {
		accessToken, _ := jMap["access_token"].(string)
		e.SetAccessToken(accessToken) //设置缓存
		return accessToken, nil
	} else {
		errcode := jMap["errcode"].(string)
		errmsg := jMap["errmsg"].(string)
		err = errors.New(errcode + ":" + errmsg)
		return "", err
	}
}

type BatChGetMaterial struct {
	Item []struct {
		MediaId string `json:"media_id"`
		Content struct {
			NewsItem []struct {
				Title              string `json:"title"`
				Author             string `json:"author"`
				Digest             string `json:"digest"`
				Content            string `json:"content"`
				ContentSourceUrl   string `json:"content_source_url"`
				ThumbMediaId       string `json:"thumb_media_id"`
				ShowCoverPic       int    `json:"show_cover_pic"`
				Url                string `json:"url"`
				ThumbUrl           string `json:"thumb_url"`
				NeedOpenComment    int    `json:"need_open_comment"`
				OnlyFansCanComment int    `json:"only_fans_can_comment"`
			} `json:"news_item"`
			CreateTime int64 `json:"create_time"`
			UpdateTime int64 `json:"update_time"`
		}
		UpdateTime int `json:"update_time"`
	}
	TotalCount int `json:"total_count"`
	ItemCount  int `json:"item_count"`
}

// @Summer 微信获取文章
func GetArticle(begin, count int) {
	//total := Article.GetArticleTotal()
	result, err := ResolveUrl(begin, count)
	var article = make(map[string]interface{})

	if err != nil {
		fmt.Printf("read resp.body failed,err:%v\n", err)
	} else {
		stu := &BatChGetMaterial{}
		res := json.Unmarshal(result, &stu)
		if res == nil {
			for _, item := range stu.Item {
				res := item.Content.NewsItem[0]
				if res.Title != "" {
					tit := strings.TrimSpace(res.Title)
					currentTime := time.Unix(item.Content.CreateTime, 0).Format("2006-01-02 15:04:05")
					if strings.Contains("练脑时刻", tit) {
						currentTime = Article.SubTime(item.Content.UpdateTime)
					}
					// url尾部字符串
					imgType := Article.ThumbImgType(res.ThumbUrl)
					thumbImg := Article.TrimUrl(imgType, res.ThumbUrl)
					article["title"] = res.Title
					article["summary"] = res.Digest
					article["thumb_img"] = thumbImg
					article["admin"] = res.Author
					article["com"] = "weChat"
					article["is_show"] = 1
					article["content"] = Article.ReplaceContent(res.Content)
					article["hot"] = 0
					article["sort"] = 0
					article["nav_id"] = 8
					article["created_at"] = currentTime
					if thumbImg != "" {
						wt.Add(1)
						go WeAddArticle(article)
					}
				}
			}
			wg.Wait()
		}
	}
}

func GetArt() {
	total := Article.GetArticleTotal()
	GetArticle(total+1, 1)
}

func ResolveUrl(offset, count int) ([]byte, error) {
	isOk, accessToken := e.GetVal("access_token")
	if !isOk {
		token, err := GetToken()
		if err != nil {
			panic(err)
		}
		accessToken = token
	}
	data := make(map[string]interface{})
	data["type"] = "news"
	data["offset"] = offset
	data["count"] = count
	url := "https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token=" + accessToken

	bytesData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
	}
	reader := bytes.NewReader(bytesData)

	rep, err := http.NewRequest("POST", url, reader)
	resp, err := http.DefaultClient.Do(rep)

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// @Summer 添加文章
func WeAddArticle(data map[string]interface{}) bool {
	defer wt.Done()
	UpdatedAt := time.Now().Format("2006-01-02 15:04:05")
	article := db.Db.Create(&Article.Article{
		Title:     data["title"].(string),
		Summary:   data["summary"].(string),
		ThumbImg:  data["thumb_img"].(string),
		Admin:     data["admin"].(string),
		Com:       data["com"].(string),
		IsShow:    data["is_show"].(int),
		Content:   data["content"].(string),
		Hot:       data["hot"].(int),
		Sort:      data["sort"].(int),
		NavId:     data["nav_id"].(int),
		CreatedAt: data["created_at"].(string),
		UpdatedAt: UpdatedAt,
	})

	if article.Error != nil {
		fmt.Print("添加文章失败", article)
		return false
	}
	return true
}
