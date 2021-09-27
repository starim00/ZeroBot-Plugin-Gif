package plugin_gif

import (
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func (c *context) prepareLogos(s ...string) {
	for i, v := range s {
		_, err := strconv.Atoi(v)
		if err != nil {
			download("https://gchat.qpic.cn/gchatpic_new//--"+strings.ToUpper(v)+"/0", c.usrdir+strconv.Itoa(i)+".gif")
		} else {
			download("http://q4.qlogo.cn/g?b=qq&nk="+v+"&s=640", c.usrdir+strconv.Itoa(i)+".gif")
		}
	}
}

func (c *context) exists(name string) bool {
	file := c.usrdir + name
	_, err := os.Stat(file)
	e := err == nil || os.IsExist(err)
	logrus.Debugln("[gif]", name, "exists:", e)
	return e
}
