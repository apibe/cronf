package internal

import (
	"errors"
	"fmt"
	"github.com/jlaffaye/ftp"
	"github.com/robfig/cron"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"
)

func download() {
	c := cron.New()
	for _, cr := range C.CRONDOWNLOAD {
		t := task{
			addr:      C.FTP.ADDR,
			username:  C.FTP.USERNAME,
			password:  C.FTP.PASSWORD,
			ftpPath:   cr.FTPPATH,
			localPath: cr.LOCALPATH,
			times:     cr.RETRY.TIMES,
			interval:  cr.RETRY.INTERVAL,
		}
		if cr.EXEC {
			t.downloadWithRetry()
		}
		if cr.JOINCRON {
			err := c.AddFunc(cr.CRON, t.downloadWithRetry)
			if err != nil {
				log.Printf("%s 加入定时失败，cron表达式 %s \n", cr.NAME, cr.CRON)
				panic(err)
			}
			log.Printf("%s 加入定时任务成功，cron表达式 %s \n", cr.NAME, cr.CRON)
		}
	}
	c.Start()
}

func ftpUpload() {}

type task struct {
	addr      string
	username  string
	password  string
	ftpPath   string
	localPath string
	times     int
	interval  int
}

func (t task) downloadWithRetry() {
	err := t.download()
	if err != nil && t.times > 0 && t.interval > 0 {
		counter := 1
		ticker := time.NewTicker(time.Duration(t.interval) * time.Minute)
		for {
			select {
			case <-ticker.C:
				if err := t.download(); err == nil {
					return
				}
				counter++
				if counter > t.times {
					log.Printf("ftpPath:%s 重试失败", t.ftpPath)
					return
				}
			}
		}
	}
}

func (t task) download() error {
	t.ftpPath = path(t.ftpPath)
	t.localPath = path(t.localPath)
	defer Recover()
	client, err := ftp.Dial(t.addr, ftp.DialWithDialer(net.Dialer{Timeout: 30 * time.Second}))
	defer client.Quit()
	if err != nil {
		log.Printf("ftp服务器连接失败，FTP服务不可达 %s HOST：%s \n", err.Error(), t.addr)
		return err
	}
	err = client.Login(t.username, t.password)
	if err != nil {
		log.Printf("ftp服务器连接失败，FTP账号密码错误 %s 账号：%s 密码：%s \n", err.Error(), t.username, t.password)
		return err
	}
	list, err := client.List(t.ftpPath)
	if err != nil {
		log.Printf("ftp服务器连接失败，获取数据资源列表失败: %s \n", err.Error())
		return err
	} else if len(list) < 1 {
		log.Printf("ftp服务器连接失败，资源列表不存在: %s \n", t.ftpPath)
		return errors.New("ftp服务器连接失败，资源列表不存在")
	}
	for _, entry := range list {
		ftpName := fmt.Sprintf("%s/%s", t.ftpPath, entry.Name)
		localName := fmt.Sprintf("%s/%s", t.localPath, entry.Name)
		if _, err := os.Stat(t.localPath); os.IsNotExist(err) {
			err := os.MkdirAll(t.localPath, 0755)
			if err != nil {
				log.Printf("local文件夹创建失败 localName %v err: %s \n", localName, err.Error())
				continue
			}
		}
		response, err := client.Retr(ftpName)
		fileBytes, err := ioutil.ReadAll(response)
		response.Close()
		if err != nil || len(fileBytes) == 0 {
			log.Printf("ftp文件下载失败 fileBytes %v err: %s \n", len(fileBytes), err.Error())
			continue
		}
		if _, err := os.Stat(localName); os.IsNotExist(err) {
			_, errCreate := os.Create(localName)
			if errCreate != nil {
				log.Printf("local文件创建失败 localName %v err: %s \n", localName, err.Error())
				continue
			}
		}
		if err = ioutil.WriteFile(localName, fileBytes, 0777); err != nil {
			log.Printf("ftp文件写入本地文件失败 fileBytes %v err: %s \n", len(fileBytes), err.Error())
			continue
		}
	}
	log.Printf("ftpPath:%s 文件下载成功,文件下载路径 %s \n", t.ftpPath, t.localPath)
	return err
}
