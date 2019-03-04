package v1

import (
	"flag"
	"fmt"
	cdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	"io/ioutil"
	"os"
	"os/exec"
)

type app struct {
	SecretID  string
	SecretKey string
}

type existBackupFiles struct {
	Directory string
	Num       int
	Items     []string
}

func (existBackupFiles *existBackupFiles) New(path string) *existBackupFiles {
	fileInfo, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	existBackupFiles.Directory = path
	for _, file := range fileInfo {
		fmt.Println(file.Name())
		existBackupFiles.Items = append(existBackupFiles.Items, file.Name())
		existBackupFiles.Num++
	}

	return existBackupFiles
}

/**
 * 确认是否需要删除一些备份文件
 */
func (existBackupFiles *existBackupFiles) checkFilesNum() (needDel bool) {
	if len(existBackupFiles.Items) >= fileNum {
		return true
	} else {
		return false
	}
}

/**
 * 删除多余的备份文件
 */
func (existBackupFiles *existBackupFiles) DeleteOldFile() {
	if existBackupFiles.checkFilesNum() {
		needDelFiles := existBackupFiles.Items[:fileNum-2]

		for _, name := range needDelFiles {
			err := os.Remove(existBackupFiles.Directory + name)

			if err != nil {
				fmt.Println(err)
			}
		}
		existBackupFiles.Items = existBackupFiles.Items[fileNum-1:]
	}
}

var (
	mode      int
	directory string
	fileNum   int
	action    string
)

func init() {
	// 初始化命令行参数
	flag.StringVar(&directory, "path", "/data/soft/backup/mysql/", "Backup File Storage Path") // 文件存放路径
	flag.IntVar(&mode, "mode", 1, "Download mode: 1 Intranet 2 Internet")                      // 下载网络模式 1内网 2外网
	flag.IntVar(&fileNum, "num", 3, "Backup File Num Limit")                                   // 日志文件数量限制 超过即删除
	flag.StringVar(&action, "action", "backups", "API Action: backups | binlogs")
}

func main() {
	// 解析参数
	flag.Parse()

	// 初始化API client
	db := app{"AKIDysbDhzBzsdFpspFUy9HEkLvKZfZRtk8K", "CD7wb7TdbN0TkqrmaJwY3dBe4NdmpncG"}
	credential := common.NewCredential(db.SecretID, db.SecretKey)
	cpf := profile.NewClientProfile()
	client, _ := cdb.NewClient(credential, regions.Shanghai, cpf)
	instanceId := "cdb-cjb8gn4z"

	var (
		downloadPath string
		date         string
	)

	switch action {
	case "binlogs":
		// 下载 binlogs
		request := cdb.NewDescribeBinlogsRequest()
		request.InstanceId = &instanceId

		response, err := client.DescribeBinlogs(request)

		if err != nil {
			fmt.Println(err)
			return
		}

		for _, item := range response.Response.Items {
			fmt.Println(*item.InternetUrl)
		}

		// 解析下载URL，拿第二个
		if mode == 1 {
			downloadPath = *response.Response.Items[1].IntranetUrl
		} else {
			downloadPath = *response.Response.Items[1].InternetUrl
		}

		date = *response.Response.Items[1].Date
	default:
		// 下载 backups
		request := cdb.NewDescribeBackupsRequest()
		request.InstanceId = &instanceId

		response, err := client.DescribeBackups(request)

		if err != nil {
			fmt.Println(err)
			return
		}

		for _, item := range response.Response.Items {
			fmt.Println(*item.InternetUrl)
		}

		// 解析下载URL
		if mode == 1 {
			downloadPath = *response.Response.Items[0].IntranetUrl
		} else {
			downloadPath = *response.Response.Items[0].InternetUrl
		}

		date = *response.Response.Items[0].Date
	}

	// 下载文件
	cmd := exec.Command("wget", "-O", directory+date[0:10], downloadPath)

	cmd.Start()
	cmd.Wait()

	// 删除多余文件
	existBackupFiles := existBackupFiles{}
	existBackupFiles.New(directory)
	existBackupFiles.DeleteOldFile()

	fmt.Println(date)
}
