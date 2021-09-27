package plugin_gif

import (
	"os"
	"strconv"
	"strings"
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
	return err == nil || os.IsExist(err)
}
