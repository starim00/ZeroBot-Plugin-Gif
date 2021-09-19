package plugin_gif

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type context struct {
	user string
	imgs []string
}

func dlpic(name string) string {
	target := datapath + `materials/` + name
	_, err := os.Stat(target)
	if err != nil {
		download(`https://codechina.csdn.net/u011570312/imagematerials/-/raw/main/`+name, target)
	}
	return target
}

// 新的上下文
func newContext(user int64) *context {
	c := new(context)
	c.user = datapath + strconv.FormatInt(user, 10) + `/`
	os.MkdirAll(c.user, 0755)
	c.imgs = make([]string, 2)
	c.imgs[0] = c.user + "yuan0.gif"
	c.imgs[1] = c.user + "yuan1.gif"
	return c
}

// 下载图片
func download(url, dlpath string) error {
	// 创建目录
	var List = strings.Split(dlpath, `/`)
	err := os.MkdirAll(strings.TrimSuffix(dlpath, List[len(List)-1]), 0755)
	if err != nil {
		logrus.Errorln("[gif] mkdir err:", err)
		return err
	}
	res, err := http.Get(url)
	if err != nil {
		logrus.Errorln("[gif] http get err:", err)
		return err
	}
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)
	// 创建文件
	file, err := os.Create(dlpath)
	if err != nil {
		logrus.Errorln("[gif] create file err:", err)
		return err
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)
	written, err := io.Copy(writer, reader)
	if err != nil {
		logrus.Errorln("[gif] copy err:", err)
		return err
	}
	res.Body.Close()
	logrus.Info("[gif] dl len:", written)
	return nil
}
