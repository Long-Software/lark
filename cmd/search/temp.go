package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"net/http"
// 	"net/url"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/Long-Software/lark/pkg/log"
// 	"github.com/PuerkitoBio/goquery"
// )

// var domains = map[string]string{
// 	"com": "https://www.google.com/search?q=",
// 	"ac":  "https://www.google.ac/search?q=",
// 	"ad":  "https://www.google.ad/search?q=",
// 	"ae":  "https://www.google.ae/search?q=",
// 	"af":  "https://www.google.com.af/search?q=",
// 	"ag":  "https://www.google.com.ag/search?q=",
// 	"ai":  "https://www.google.com.ai/search?q=",
// 	"al":  "https://www.google.al/search?q=",
// 	"am":  "https://www.google.am/search?q=",
// 	"ao":  "https://www.google.co.ao/search?q=",
// 	"ar":  "https://www.google.com.ar/search?q=",
// 	"as":  "https://www.google.as/search?q=",
// 	"at":  "https://www.google.at/search?q=",
// 	"au":  "https://www.google.com.au/search?q=",
// 	"az":  "https://www.google.az/search?q=",
// 	"ba":  "https://www.google.ba/search?q=",
// 	"bd":  "https://www.google.com.bd/search?q=",
// 	"be":  "https://www.google.be/search?q=",
// 	"bf":  "https://www.google.bf/search?q=",
// 	"bg":  "https://www.google.bg/search?q=",
// 	"bh":  "https://www.google.com.bh/search?q=",
// 	"bi":  "https://www.google.bi/search?q=",
// 	"bj":  "https://www.google.bj/search?q=",
// 	"bn":  "https://www.google.com.bn/search?q=",
// 	"bo":  "https://www.google.com.bo/search?q=",
// 	"br":  "https://www.google.com.br/search?q=",
// 	"bs":  "https://www.google.bs/search?q=",
// 	"bt":  "https://www.google.bt/search?q=",
// 	"bw":  "https://www.google.co.bw/search?q=",
// 	"by":  "https://www.google.by/search?q=",
// 	"bz":  "https://www.google.com.bz/search?q=",
// 	"ca":  "https://www.google.ca/search?q=",
// 	"kh":  "https://www.google.com.kh/search?q=",
// 	"cc":  "https://www.google.cc/search?q=",
// 	"cd":  "https://www.google.cd/search?q=",
// 	"cf":  "https://www.google.cf/search?q=",
// 	"cat": "https://www.google.cat/search?q=",
// 	"cg":  "https://www.google.cg/search?q=",
// 	"ch":  "https://www.google.ch/search?q=",
// 	"ci":  "https://www.google.ci/search?q=",
// 	"ck":  "https://www.google.co.ck/search?q=",
// 	"cl":  "https://www.google.cl/search?q=",
// 	"cm":  "https://www.google.cm/search?q=",
// 	"co":  "https://www.google.com.co/search?q=",
// 	"cr":  "https://www.google.co.cr/search?q=",
// 	"cu":  "https://www.google.com.cu/search?q=",
// 	"cv":  "https://www.google.cv/search?q=",
// 	"cy":  "https://www.google.com.cy/search?q=",
// 	"cz":  "https://www.google.cz/search?q=",
// 	"de":  "https://www.google.de/search?q=",
// 	"dj":  "https://www.google.dj/search?q=",
// 	"dk":  "https://www.google.dk/search?q=",
// 	"dm":  "https://www.google.dm/search?q=",
// 	"do":  "https://www.google.com.do/search?q=",
// 	"dz":  "https://www.google.dz/search?q=",
// 	"ec":  "https://www.google.com.ec/search?q=",
// 	"ee":  "https://www.google.ee/search?q=",
// 	"eg":  "https://www.google.com.eg/search?q=",
// 	"es":  "https://www.google.es/search?q=",
// 	"et":  "https://www.google.com.et/search?q=",
// 	"fi":  "https://www.google.fi/search?q=",
// 	"fj":  "https://www.google.com.fj/search?q=",
// 	"fm":  "https://www.google.fm/search?q=",
// 	"fr":  "https://www.google.fr/search?q=",
// 	"ga":  "https://www.google.ga/search?q=",
// 	"ge":  "https://www.google.ge/search?q=",
// 	"gf":  "https://www.google.gf/search?q=",
// 	"gg":  "https://www.google.gg/search?q=",
// 	"gh":  "https://www.google.com.gh/search?q=",
// 	"gi":  "https://www.google.com.gi/search?q=",
// 	"gl":  "https://www.google.gl/search?q=",
// 	"gm":  "https://www.google.gm/search?q=",
// 	"gp":  "https://www.google.gp/search?q=",
// 	"gr":  "https://www.google.gr/search?q=",
// 	"gt":  "https://www.google.com.gt/search?q=",
// 	"gy":  "https://www.google.gy/search?q=",
// 	"hk":  "https://www.google.com.hk/search?q=",
// 	"hn":  "https://www.google.hn/search?q=",
// 	"hr":  "https://www.google.hr/search?q=",
// 	"ht":  "https://www.google.ht/search?q=",
// 	"hu":  "https://www.google.hu/search?q=",
// 	"id":  "https://www.google.co.id/search?q=",
// 	"iq":  "https://www.google.iq/search?q=",
// 	"ie":  "https://www.google.ie/search?q=",
// 	"il":  "https://www.google.co.il/search?q=",
// 	"im":  "https://www.google.im/search?q=",
// 	"in":  "https://www.google.co.in/search?q=",
// 	"io":  "https://www.google.io/search?q=",
// 	"is":  "https://www.google.is/search?q=",
// 	"it":  "https://www.google.it/search?q=",
// 	"je":  "https://www.google.je/search?q=",
// 	"jm":  "https://www.google.com.jm/search?q=",
// 	"jo":  "https://www.google.jo/search?q=",
// 	"jp":  "https://www.google.co.jp/search?q=",
// 	"ke":  "https://www.google.co.ke/search?q=",
// 	"ki":  "https://www.google.ki/search?q=",
// 	"kg":  "https://www.google.kg/search?q=",
// 	"kr":  "https://www.google.co.kr/search?q=",
// 	"kw":  "https://www.google.com.kw/search?q=",
// 	"kz":  "https://www.google.kz/search?q=",
// 	"la":  "https://www.google.la/search?q=",
// 	"lb":  "https://www.google.com.lb/search?q=",
// 	"lc":  "https://www.google.com.lc/search?q=",
// 	"li":  "https://www.google.li/search?q=",
// 	"lk":  "https://www.google.lk/search?q=",
// 	"ls":  "https://www.google.co.ls/search?q=",
// 	"lt":  "https://www.google.lt/search?q=",
// 	"lu":  "https://www.google.lu/search?q=",
// 	"lv":  "https://www.google.lv/search?q=",
// 	"ly":  "https://www.google.com.ly/search?q=",
// 	"ma":  "https://www.google.co.ma/search?q=",
// 	"md":  "https://www.google.md/search?q=",
// 	"me":  "https://www.google.me/search?q=",
// 	"mg":  "https://www.google.mg/search?q=",
// 	"mk":  "https://www.google.mk/search?q=",
// 	"ml":  "https://www.google.ml/search?q=",
// 	"mm":  "https://www.google.com.mm/search?q=",
// 	"mn":  "https://www.google.mn/search?q=",
// 	"ms":  "https://www.google.ms/search?q=",
// 	"mt":  "https://www.google.com.mt/search?q=",
// 	"mu":  "https://www.google.mu/search?q=",
// 	"mv":  "https://www.google.mv/search?q=",
// 	"mw":  "https://www.google.mw/search?q=",
// 	"mx":  "https://www.google.com.mx/search?q=",
// 	"my":  "https://www.google.com.my/search?q=",
// 	"mz":  "https://www.google.co.mz/search?q=",
// 	"na":  "https://www.google.com.na/search?q=",
// 	"ne":  "https://www.google.ne/search?q=",
// 	"nf":  "https://www.google.com.nf/search?q=",
// 	"ng":  "https://www.google.com.ng/search?q=",
// 	"ni":  "https://www.google.com.ni/search?q=",
// 	"nl":  "https://www.google.nl/search?q=",
// 	"no":  "https://www.google.no/search?q=",
// 	"np":  "https://www.google.com.np/search?q=",
// 	"nr":  "https://www.google.nr/search?q=",
// 	"nu":  "https://www.google.nu/search?q=",
// 	"nz":  "https://www.google.co.nz/search?q=",
// 	"om":  "https://www.google.com.om/search?q=",
// 	"pk":  "https://www.google.com.pk/search?q=",
// 	"pa":  "https://www.google.com.pa/search?q=",
// 	"pe":  "https://www.google.com.pe/search?q=",
// 	"ph":  "https://www.google.com.ph/search?q=",
// 	"pl":  "https://www.google.pl/search?q=",
// 	"pg":  "https://www.google.com.pg/search?q=",
// 	"pn":  "https://www.google.pn/search?q=",
// 	"pr":  "https://www.google.com.pr/search?q=",
// 	"ps":  "https://www.google.ps/search?q=",
// 	"pt":  "https://www.google.pt/search?q=",
// 	"py":  "https://www.google.com.py/search?q=",
// 	"qa":  "https://www.google.com.qa/search?q=",
// 	"ro":  "https://www.google.ro/search?q=",
// 	"rs":  "https://www.google.rs/search?q=",
// 	"ru":  "https://www.google.ru/search?q=",
// 	"rw":  "https://www.google.rw/search?q=",
// 	"sa":  "https://www.google.com.sa/search?q=",
// 	"sb":  "https://www.google.com.sb/search?q=",
// 	"sc":  "https://www.google.sc/search?q=",
// 	"se":  "https://www.google.se/search?q=",
// 	"sg":  "https://www.google.com.sg/search?q=",
// 	"sh":  "https://www.google.sh/search?q=",
// 	"si":  "https://www.google.si/search?q=",
// 	"sk":  "https://www.google.sk/search?q=",
// 	"sl":  "https://www.google.com.sl/search?q=",
// 	"sn":  "https://www.google.sn/search?q=",
// 	"sm":  "https://www.google.sm/search?q=",
// 	"so":  "https://www.google.so/search?q=",
// 	"st":  "https://www.google.st/search?q=",
// 	"sr":  "https://www.google.sr/search?q=",
// 	"sv":  "https://www.google.com.sv/search?q=",
// 	"td":  "https://www.google.td/search?q=",
// 	"tg":  "https://www.google.tg/search?q=",
// 	"th":  "https://www.google.co.th/search?q=",
// 	"tj":  "https://www.google.com.tj/search?q=",
// 	"tk":  "https://www.google.tk/search?q=",
// 	"tl":  "https://www.google.tl/search?q=",
// 	"tm":  "https://www.google.tm/search?q=",
// 	"to":  "https://www.google.to/search?q=",
// 	"tn":  "https://www.google.tn/search?q=",
// 	"tr":  "https://www.google.com.tr/search?q=",
// 	"tt":  "https://www.google.tt/search?q=",
// 	"tw":  "https://www.google.com.tw/search?q=",
// 	"tz":  "https://www.google.co.tz/search?q=",
// 	"ua":  "https://www.google.com.ua/search?q=",
// 	"ug":  "https://www.google.co.ug/search?q=",
// 	"uk":  "https://www.google.co.uk/search?q=",
// 	"us":  "https://www.google.com/search?q=",
// 	"uy":  "https://www.google.com.uy/search?q=",
// 	"uz":  "https://www.google.co.uz/search?q=",
// 	"vc":  "https://www.google.com.vc/search?q=",
// 	"ve":  "https://www.google.co.ve/search?q=",
// 	"vg":  "https://www.google.vg/search?q=",
// 	"vi":  "https://www.google.co.vi/search?q=",
// 	"vn":  "https://www.google.com.vn/search?q=",
// 	"vu":  "https://www.google.vu/search?q=",
// 	"ws":  "https://www.google.ws/search?q=",
// 	"za":  "https://www.google.co.za/search?q=",
// 	"zm":  "https://www.google.co.zm/search?q=",
// 	"zw":  "https://www.google.co.zw/search?q=",
// }

// type Result struct {
// 	Rank  int
// 	URL   string
// 	Title string
// 	Desc  string
// }

// type Agent struct {
// 	names []string
// }

// func NewAgent() *Agent {
// 	return &Agent{
// 		names: []string{
// 			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
// 			"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
// 			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
// 			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
// 			"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:56.0) Gecko/20100101 Firefox/56.0",
// 			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
// 		},
// 	}
// }

// func (a *Agent) Random() string {
// 	rand.Seed(time.Now().Unix())
// 	idx := rand.Int() % len(a.names)
// 	return a.names[idx]
// }

// type Scraper struct {
// 	terms        string
// 	TLD          string
// 	pages        int
// 	count        int
// 	proxy        string
// 	timeout      int
// 	languageCode string
// }

// func (s *Scraper) Run() ([]Result, error) {
// 	results := []Result{}
// 	counter := 0
// 	urls, err := s.build()
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, url := range urls {
// 		res, err := s.request(url)
// 		if err != nil {
// 			lg.NewLog(log.ERROR, "Request failed: "+err.Error())
// 			continue
// 		}
// 		lg.NewLog(log.DEBUG, res.Status)
// 		data, err := s.parse(res, counter)
// 		if err != nil {
// 			lg.NewLog(log.ERROR, "Parse failed: "+err.Error())
// 			continue
// 		}
// 		counter += len(data)
// 		for _, item := range data {
// 			results = append(results, item)
// 		}
// 		time.Sleep(time.Duration(s.timeout) * time.Second)
// 	}
// 	lg.NewLog(log.INFO, strconv.Itoa(len(results)))
// 	return results, nil
// }

// func (s *Scraper) request(url string) (*http.Response, error) {
// 	a := NewAgent()
	
// 	// Handle proxy if provided
// 	var client *http.Client
// 	if s.proxy != "" {
// 		proxyUrl, err := url.Parse(s.proxy)
// 		if err != nil {
// 			return nil, fmt.Errorf("invalid proxy URL: %w", err)
// 		}
// 		client = &http.Client{
// 			Transport: &http.Transport{
// 				Proxy: http.ProxyURL(proxyUrl),
// 			},
// 		}
// 	} else {
// 		client = &http.Client{}
// 	}
	
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	req.Header.Set("User-Agent", a.Random())

// 	res, err := client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
	
// 	if res.StatusCode != 200 {
// 		res.Body.Close()
// 		return nil, fmt.Errorf("scraper received a non-200 status code (%d) suggesting a ban", res.StatusCode)
// 	}

// 	return res, nil
// }

// func (s *Scraper) parse(res *http.Response, rank int) ([]Result, error) {
// 	defer res.Body.Close()
	
// 	doc, err := goquery.NewDocumentFromReader(res.Body)
// 	if err != nil {
// 		return nil, err
// 	}
	
// 	results := []Result{}
// 	sel := doc.Find("div.g")
// 	rank++
	
// 	for i := range sel.Nodes {
// 		item := sel.Eq(i)
// 		linkTag := item.Find("a")
// 		link, _ := linkTag.Attr("href")
// 		titleTag := item.Find("h3.r")
// 		descTag := item.Find("span.st")
// 		desc := descTag.Text()
// 		title := titleTag.Text()
// 		link = strings.Trim(link, " ")
		
// 		if link != "" && link != "#" && !strings.HasPrefix(link, "/") {
// 			result := Result{
// 				Rank:  rank,
// 				URL:   link,
// 				Title: title,
// 				Desc:  desc,
// 			}
// 			results = append(results, result)
// 			rank++
// 		}
// 	}
// 	return results, nil
// }

// func (s *Scraper) build() ([]string, error) {
// 	urls := []string{}
// 	s.terms = strings.Trim(s.terms, " ")
// 	s.terms = strings.Replace(s.terms, " ", "+", -1)
	
// 	if base, found := domains[s.TLD]; found {
// 		for i := 0; i < s.pages; i++ {
// 			start := i * s.count
// 			url := fmt.Sprintf("%s%s&num=%d&hl=%s&start=%d&filter=0", base, s.terms, s.count, s.languageCode, start)
// 			urls = append(urls, url)
// 		}
// 	} else {
// 		return nil, fmt.Errorf("country (%s) is not supported", s.TLD)
// 	}
// 	return urls, nil
// }

// func main() {
// 	s := Scraper{
// 		terms:        "akhil sharma",
// 		TLD:          "com",
// 		languageCode: "en",
// 		pages:        1,
// 		count:        30,
// 		timeout:      10,
// 	}

// 	res, err := s.Run()
// 	if err != nil {
// 		lg.NewLog(log.ERROR, err.Error())
// 	}
// 	for _, r := range res {
// 		lg.NewLog(log.INFO, r.URL)
// 	}
// }

// var lg = log.Logger{
// 	IsProduction: true,
// 	HasTimestamp: true,
// 	HasFilepath:  true,
// 	HasMethod:    true,
// }