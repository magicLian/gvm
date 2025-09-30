package utils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"sort"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/PuerkitoBio/goquery"
)

type GoVersion struct {
	Version string
	File    string
	OS      string
	Arch    string
	Size    string
	URL     string
}

func FetchGoVersions(osFilter, archFilter string) ([]GoVersion, error) {
	resp, err := http.Get("https://go.dev/dl/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	results := make([]GoVersion, 0)
	cache := make(map[string]bool)
	limit := 15
	stop := false

	doc.Find("table.downloadtable").Each(func(i int, s *goquery.Selection) {
		if stop {
			return
		}

		s.Find("tr").Each(func(_ int, tr *goquery.Selection) {
			if stop {
				return
			}

			tds := tr.Find("td")
			if tds.Length() != 6 {
				return
			}

			a := tds.Eq(0).Find("a.download").First()
			if a == nil {
				return
			}
			if !strings.HasPrefix(a.AttrOr("href", ""), "/dl/go") {
				return
			}

			href := a.AttrOr("href", "")
			filename := href[strings.LastIndex(href, "/")+1:]

			// 排除源码包、安装包、测试包等
			if strings.Contains(filename, "src.") || strings.Contains(filename, "pkg.") {
				return
			}
			re := regexp.MustCompile(`go([0-9]+\.[0-9]+(?:\.[0-9]+)?)\.([a-z0-9]+)-([a-z0-9]+)`)
			m := re.FindStringSubmatch(filename)
			if len(m) != 4 {
				return
			}
			version := m[1]
			osName := m[2]
			arch := m[3]
			size := tds.Eq(4).Text()

			if osFilter != "" && osName != osFilter {
				return
			}
			if archFilter != "" && arch != archFilter {
				return
			}

			// 排除重复版本
			if _, ok := cache[version]; ok {
				return
			}
			cache[version] = true

			results = append(results, GoVersion{
				Version: version,
				OS:      osName,
				Arch:    arch,
				File:    filename,
				Size:    size,
				URL:     fmt.Sprintf("https://go.dev%s", href),
			})

			if len(results) >= limit {
				stop = true
			}
		})
	})

	// ✅ 排序（按版本号倒序）
	sort.Slice(results, func(i, j int) bool {
		v1, err1 := semver.NewVersion(results[i].Version)
		v2, err2 := semver.NewVersion(results[j].Version)
		if err1 != nil || err2 != nil {
			// 如果解析失败，回退到字符串比较
			return results[i].Version > results[j].Version
		}
		// 倒序排序
		return v1.GreaterThan(v2)
	})

	return results, nil
}
