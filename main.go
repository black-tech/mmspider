package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type Size struct {
	Width  int
	Height int
}

type Image struct {
	Id   int64
	Name string
	Type string
	Href string
	Src  string
	Size *Size
}

var (
	img_count int64 = 0
	imgs            = make(map[int64]*Image)
	o_url           = "http://www.22mm.cc"
	reg_links       = make(map[string]*regexp.Regexp)
	reg_imgs        = make(map[string]*regexp.Regexp)
)

// 初始化： 正则表达式
func init() {
	reg_imgs["imga"] = regexp.MustCompile(`pictureContent(?s:.+?)"([^"]+.(?:jpg|png))"`)
	// <a href="http://22mm.xiuna.com/mm/bagua/PmaeaHmePHaaCabPP.html" title="夏午"><img src="http://22mm-img.xiuna.com/pic/bagua/2015-8-3/1/h0.jpg">夏午</a>
	reg_imgs["imgb"] = regexp.MustCompile(`/<a href="".*title="(\w)*">`)
	reg_imgs["imgc"] = regexp.MustCompile(`/<img src="(\w+)*">(\w+)</a>`)
}

// 主程序
func main() {
	fmt.Println("Start a spider...")
	// count, err := GetImgLinks(o_url)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println(count)
	// for k, v := range imgs {
	// 	fmt.Printf("%d: %#v\n", k, v)
	// }
	// var urls []string
	// GetImgLinks(o_url)
	// for k, v := range imgs {
	// 	fmt.Println("At: ", k)
	// 	urls = append(urls, GetImg(o_url+v.Href)...)
	// }
	// for i, v := range urls {
	// 	fmt.Printf("%d: %s\n", i, v)
	// }

	fmt.Println("Over 1 ")
}

func GetImgLinks(url string) (count int64, err error) {
	count = 0
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		img_href, exists1 := s.Parent().Attr("href")
		img_title, exists2 := s.Parent().Attr("title")
		img_src, exists3 := s.Attr("src")

		if exists1 && exists2 && exists3 {
			img_count++
			count++
			imgs[img_count] = &Image{
				Id:   img_count,
				Name: img_title,
				Src:  img_src,
				Href: img_href,
			}
		}
	})
	return
}

// box-inner imgString
func GetImg(url string) (urls []string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		if s.Text() != "" {
			urls = append(urls, testString(s.Text())...)
		}
	})
	return
}

func testString(str string) (urls []string) {
	reg, err := regexp.Compile(`\"(.+?)\.(jpg|png|gif)\"`)
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}
	bss := reg.FindAll([]byte(str), -1)
	for _, bs := range bss {
		s := strings.Replace(string(bs[1:len(bs)-1]), "big", "pic", -1)
		urls = append(urls, s)
	}
	return
}

func get(url string) (content []byte, err error) {
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()
	content, err = ioutil.ReadAll(res.Body)
	return
}

//
func getimglinks(url string) {
	re := regexp.MustCompile("a(x*)b(y|z)c")
	fmt.Printf("%#v\n", re.FindStringSubmatch("-asdf aa xxxbyc-"))

	body, err := get(url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, v := range reg_imgs {
		img := v.FindSubmatch(body)
		for _, vv := range img {
			fmt.Printf("%s\n", vv)
		}
	}
}
