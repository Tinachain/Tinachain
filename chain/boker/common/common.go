package common

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Tinachain/Tina/chain/boker/common/chainware"

	log "github.com/Tinachain/Tina/chain/boker/common/log4go"
)

// execute cmd line
func ShellExecute(s string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", s)

	var cout bytes.Buffer
	cmd.Stdout = &cout

	var cerr bytes.Buffer
	cmd.Stderr = &cerr

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return cout.String(), nil
}

func MD5(b []byte) string {
	md5Ctx := md5.New()
	md5Ctx.Write(b)
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

func RandomStringByPattern(pattern []byte, n int) string {
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, pattern[r.Intn(len(pattern))])
	}
	return string(result)
}

func RandomString(n int) string {
	pattern := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLOMNOPQRSTUVWXYZ1234567890"
	return RandomStringByPattern([]byte(pattern), n)
}

func RandomNumber(n int) string {
	pattern := "1234567890"
	return RandomStringByPattern([]byte(pattern), n)
}

func HttpError(w http.ResponseWriter, result int, msg string) {
	w.Write([]byte(fmt.Sprintf("{\"result\":%d,\"msg\":\"%s\"}", result, msg)))
}

func HttpErrorEx(w http.ResponseWriter, err chainware.Error) {
	w.Write([]byte(fmt.Sprintf("{\"result\":%d,\"msg\":\"%s\"}", err.Code, err.Message)))
}

func HttpErrorParamNotFound(w http.ResponseWriter, msg string) {
	w.Write([]byte(fmt.Sprintf("{\"result\":%d,\"msg\":\"%s\"}", chainware.ErrorParamNotFound, msg)))
}

func HttpFormGetString(r *http.Request, param string) (value string, err chainware.Error) {
	if len(r.Form[param]) <= 0 {
		return "", chainware.NewError(chainware.ErrorParamNotFound, "param %s not found", param)
	}
	return strings.TrimSpace(r.Form[param][0]), nil
}

func HttpFormGetInt(r *http.Request, param string) (value int, err chainware.Error) {
	if len(r.Form[param]) <= 0 {
		return 0, chainware.NewError(chainware.ErrorParamNotFound, "param %s not found", param)
	}
	valueStr := strings.TrimSpace(r.Form[param][0])
	value, e := strconv.Atoi(valueStr)
	if e != nil {
		return 0, chainware.NewError(chainware.ErrorParamInvalid, "param %s invalid err=%s", param, e.Error())
	}
	return value, nil
}

func HttpFormGetInt64(r *http.Request, param string) (value int64, err chainware.Error) {
	if len(r.Form[param]) <= 0 {
		return 0, chainware.NewError(chainware.ErrorParamNotFound, "param %s not found", param)
	}
	valueStr := strings.TrimSpace(r.Form[param][0])
	value, e := strconv.ParseInt(valueStr, 10, 64)
	if e != nil {
		return 0, chainware.NewError(chainware.ErrorParamInvalid, "param %s invalid err=%s", param, e.Error())
	}
	return value, nil
}

func PathBase(filePath string) string {
	filePath = strings.TrimRight(filePath, "/\\")
	if filePath == "" {
		return "."
	}

	idx1 := strings.LastIndex(filePath, "/")
	idx2 := strings.LastIndex(filePath, "\\")
	idx := idx1
	if idx2 > idx {
		idx = idx2
	}

	if idx >= 0 {
		filePath = filePath[idx+1:]
	}
	if filePath == "" {
		return "/"
	}
	return filePath
}

func GetExeName() string {
	ret := ""
	ex, err := os.Executable()
	if err == nil {
		ret = filepath.Base(ex)
	}
	return ret
}

func InitLog() {

	logJson := "log.json"
	set := false
	if bExist := PathExist(logJson); bExist {
		if err := log.SetupLogWithConf(logJson); err == nil {
			set = true
		}
	}

	if !set {
		fileWriter := log.NewFileWriter()
		exeName := GetExeName()
		fileWriter.SetPathPattern("log/" + exeName + "-%Y%M%D.log")
		log.Register(fileWriter)
		log.SetLevel(log.DEBUG)
	}
}

func Sha1File(file *os.File) (string, error) {
	//record current pos and reset to begining
	curPos, _ := file.Seek(0, os.SEEK_CUR)
	file.Seek(0, os.SEEK_SET)
	sha1Ctx := sha1.New()
	_, err := io.Copy(sha1Ctx, file)

	//recover pos
	file.Seek(curPos, os.SEEK_SET)

	if err != nil {
		return "", err
	}
	return hex.EncodeToString(sha1Ctx.Sum(nil)), nil
}

func Sha256(b []byte) string {
	sha256Ctx := sha256.New()
	sha256Ctx.Write(b)
	return hex.EncodeToString(sha256Ctx.Sum(nil))
}

func FileSize(file *os.File) int64 {
	curPos, _ := file.Seek(0, os.SEEK_CUR)
	fileSize, _ := file.Seek(0, os.SEEK_END)
	file.Seek(curPos, os.SEEK_SET)
	return fileSize
}

func TimeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func TimestampToString(ts int64) string {
	return time.Unix(ts, 0).Format("2006-01-02 15:04:05")
}
func TimeFromString(str string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
}

func NowString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func PageRange(total, page, pageSize int) (start, end int) {
	if total <= 0 {
		return 0, 0
	}

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		start = 0
		end = total - 1
		return start, end
	}

	start = (page - 1) * pageSize
	if start >= total {
		page := total/pageSize + 1
		if 0 == total%pageSize {
			page = page - 1
		}
		start = (page - 1) * pageSize
	}

	end = start + pageSize - 1
	if end >= total {
		end = total - 1
	}

	return start, end
}

func PathExist(fullPath string) bool {
	_, err := os.Stat(fullPath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
