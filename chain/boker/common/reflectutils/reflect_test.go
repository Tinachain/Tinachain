package reflectutils

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type FileInfo struct {
	FileId       int64     "Id"
	Uploader     int64     "uploader"
	Owner        int64     "owner"
	Name         string    "name"
	Ext          string    "ext"
	Size         int64     "size"
	Sha          string    "sha"
	Title        string    "title"
	Tag          string    "tag"
	Description  string    "description"
	IpfsHash     string    "ipfsHash"
	IpfsUrl      string    "ipfsUrl"
	AliDnaJobId  string    "aliDnaJobId"
	AliDnaFileId string    "aliDnaFileId"
	Status       int       "status"
	CreateTime   time.Time "createTime"
}

func Test(t *testing.T) {
	fmt.Println("Test")
	//	fInfo := &FileInfo{}
	//	helper, err := NewStructHelper(fInfo)
	//	fmt.Println(helper, err)
	//	values := map[string]interface{}{
	//		"ipfsHash":   "hash1111",
	//		"uploader":   12,
	//		"size":       1024,
	//		"createTime": time.Now(),
	//	}
	//	fmt.Printf("%v\n", fInfo)
	objectType := reflect.TypeOf(FileInfo{})
	object := reflect.New(objectType)
	fmt.Println(object)
	fmt.Println(object.Elem().Addr())
}
