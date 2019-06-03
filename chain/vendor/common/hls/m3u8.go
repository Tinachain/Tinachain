package hls

import (
	"common/httputils"
	"common/utils"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
)

const (
	TsStatusInit = 0
	TsStatusOk   = 1
	TsStatusFail = -1
)

type Ts struct {
	Raw              string
	UrlInfo          *url.URL
	Name             string
	RelativeToParent string
	ExtX             string
	ExtInf           string
	Duration         int64 //毫秒
	Size             int64

	TsUrl     string
	LocalFile string
	Status    int
}

type M3u8Entry struct {
	Raw              string
	UrlInfo          *url.URL
	Name             string
	RelativeToParent string
	Bandwidth        int64
}

type M3u8 struct {
	M3u8Url    string
	UrlInfo    *url.URL
	Antileech  string
	Timeout    int
	M3u8String string
	LocalFile  string

	//top m3u8
	M3u8Entries []*M3u8Entry

	//second m3u8
	Version   string
	Sequence  int64
	TsEntries []*Ts
}

func NewM3u8(m3u8Url string) *M3u8 {
	m3u8 := &M3u8{
		M3u8Url: m3u8Url,
		Timeout: 5,
	}
	m3u8.UrlInfo, _ = url.Parse(m3u8Url)
	return m3u8
}

func (m3u8 *M3u8) SetAntileech(antileech string) *M3u8 {
	m3u8.Antileech = antileech
	return m3u8
}

func (m3u8 *M3u8) SetTimeout(timeout int) *M3u8 {
	m3u8.Timeout = timeout
	return m3u8
}

func (m3u8 *M3u8) Download() (err error) {
	if m3u8.Antileech != "" {
		headers := http.Header{}
		headers.Set("Strm-Uri", m3u8.M3u8Url)
		antileechUrl := AntileechUrl(m3u8.Antileech)
		m3u8.M3u8String, err = httputils.HttpGet(antileechUrl, m3u8.Timeout, headers)
	} else {
		m3u8.M3u8String, err = httputils.HttpGet(m3u8.M3u8Url, m3u8.Timeout, nil)
	}

	if err != nil {
		return err
	}

	m3u8.Parse(m3u8.M3u8String)

	if !m3u8.IsSecond() && !m3u8.IsTop() {
		return fmt.Errorf("invalid m3u8 format!")
	}

	return nil
}

func (m3u8 *M3u8) Parse(m3u8String string) {
	m3u8.M3u8String = m3u8String
	lines := utils.SplitLine(m3u8String)
	for i := range lines {
		line := lines[i]

		//version
		{
			index := strings.LastIndex(line, "#EXT-X-VERSION:")
			if index >= 0 {
				m3u8.Version = line[index+len("#EXT-X-VERSION:") : len(line)]
			}
		}

		//sequence
		{
			index := strings.LastIndex(line, "#EXT-X-MEDIA-SEQUENCE:")
			if index >= 0 {
				m3u8.Sequence, _ = strconv.ParseInt(line[index+len("#EXT-X-MEDIA-SEQUENCE:"):len(line)], 10, 64)
			}
		}

		//ts
		{
			index := strings.LastIndex(line, "#EXTINF:")
			if index >= 0 {
				durationString := line[index+len("#EXTINF:"):]
				titleString := ""
				commaIndex := strings.LastIndex(durationString, ",")
				if commaIndex >= 0 {
					titleString = durationString[commaIndex+1:]
					durationString = durationString[:commaIndex]
				}
				durationFloat, _ := strconv.ParseFloat(durationString, 32)
				duration := int64(durationFloat * 1000)

				if len(titleString) > 0 {
					fields := strings.Fields(titleString)
					if len(fields) > 0 {
						durationFloat, err := strconv.ParseFloat(fields[0], 32)
						if nil == err {
							duration = int64(durationFloat * 1000)
						}
					}
				}

				ts := &Ts{}

				if i > 0 {
					extx := lines[i-1]
					if strings.LastIndex(extx, "#EXT-X-PROGRAM-DATE-TIME") >= 0 {
						ts.ExtX = extx
					}
				}

				ts.ExtInf = line
				ts.Raw = lines[i+1]
				ts.UrlInfo, _ = url.Parse(ts.Raw)
				ts.Name = path.Base(ts.UrlInfo.Path)
				ts.RelativeToParent = PathRelativeTo(ts.UrlInfo.Path, m3u8.UrlInfo.Path)
				ts.Duration = duration
				m3u8.TsEntries = append(m3u8.TsEntries, ts)
			}
		}

		//m3u8
		{
			index := strings.LastIndex(line, "#EXT-X-STREAM-INF:")
			if index >= 0 {
				entry := &M3u8Entry{}

				bandIndex := strings.LastIndex(line, "BANDWIDTH=")
				if bandIndex >= 0 {
					entry.Bandwidth, _ = strconv.ParseInt(line[bandIndex+len("BANDWIDTH="):], 10, 64)
				}
				entry.Raw = lines[i+1]
				entry.UrlInfo, _ = url.Parse(entry.Raw)
				entry.Name = path.Base(entry.UrlInfo.Path)
				entry.RelativeToParent = PathRelativeTo(entry.UrlInfo.Path, m3u8.UrlInfo.Path)
				m3u8.M3u8Entries = append(m3u8.M3u8Entries, entry)
			}
		}
	}
}

func (m3u8 *M3u8) IsTop() bool {
	return len(m3u8.M3u8Entries) > 0
}

func (m3u8 *M3u8) IsSecond() bool {
	return len(m3u8.TsEntries) > 0
}

func PathDir(p string) string {
	if p == "" {
		return ""
	}

	dir := path.Dir(p)
	if dir == "." || dir == "/" {
		dir = ""
	}
	if strings.HasPrefix(dir, "/") {
		dir = dir[1:]
	}
	return dir
}

func PathRelativeTo(path1, path2 string) string {
	if path1 == "" {
		return ""
	}
	path1UrlInfo, _ := url.Parse(path1)
	path1Dir := PathDir(path1UrlInfo.Path)

	if path2 == "" {
		return path1Dir
	}
	path2UrlInfo, _ := url.Parse(path2)
	path2Dir := PathDir(path2UrlInfo.Path)

	pathRelative := strings.TrimPrefix(path1Dir, path2Dir)
	if strings.HasPrefix(pathRelative, "/") {
		pathRelative = pathRelative[1:]
	}
	return pathRelative
}
