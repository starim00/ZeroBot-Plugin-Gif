package plugin_gif

import (
	"image"

	img "github.com/FloatTech/imgfactory"
)

// 摸
func (cc *context) mo() string {
	tou := img.LoadFirstFrame(cc.imgs[0], 0, 0).Circle(0).Im
	mo := []*image.NRGBA{
		img.LoadFirstFrame(dlpic(`mo/0.png`), 0, 0).InsertBottom(tou, 80, 80, 32, 32).Im,
		img.LoadFirstFrame(dlpic(`mo/1.png`), 0, 0).InsertBottom(tou, 70, 90, 42, 22).Im,
		img.LoadFirstFrame(dlpic(`mo/2.png`), 0, 0).InsertBottom(tou, 75, 85, 37, 27).Im,
		img.LoadFirstFrame(dlpic(`mo/3.png`), 0, 0).InsertBottom(tou, 85, 75, 27, 37).Im,
		img.LoadFirstFrame(dlpic(`mo/4.png`), 0, 0).InsertBottom(tou, 90, 70, 22, 42).Im,
	}
	img.SaveGif(img.MergeGif(1, mo), cc.user+`摸.gif`)
	return "file:///" + cc.user + `摸.gif`
}

// 搓
func (cc *context) cuo() string {
	tou := img.LoadFirstFrame(cc.imgs[0], 110, 110).Circle(0).Im
	m1 := img.Rotate(tou, 72, 0, 0)
	m2 := img.Rotate(tou, 144, 0, 0)
	m3 := img.Rotate(tou, 216, 0, 0)
	m4 := img.Rotate(tou, 288, 0, 0)
	cuo := []*image.NRGBA{
		img.LoadFirstFrame(dlpic(`cuo/0.png`), 0, 0).InsertBottomC(tou, 0, 0, 75, 130).Im,
		img.LoadFirstFrame(dlpic(`cuo/1.png`), 0, 0).InsertBottomC(m1.Im, 0, 0, 75, 130).Im,
		img.LoadFirstFrame(dlpic(`cuo/2.png`), 0, 0).InsertBottomC(m2.Im, 0, 0, 75, 130).Im,
		img.LoadFirstFrame(dlpic(`cuo/3.png`), 0, 0).InsertBottomC(m3.Im, 0, 0, 75, 130).Im,
		img.LoadFirstFrame(dlpic(`cuo/4.png`), 0, 0).InsertBottomC(m4.Im, 0, 0, 75, 130).Im,
	}
	img.SaveGif(img.MergeGif(5, cuo), cc.user+`搓.gif`)
	return "file:///" + cc.user + `搓.gif`
}

// 敲
func (cc *context) qiao() string {
	tou := img.LoadFirstFrame(cc.imgs[0], 40, 40).Circle(0).Im
	qiao := []*image.NRGBA{
		img.LoadFirstFrame(dlpic(`qiao/0.png`), 0, 0).InsertUp(tou, 40, 33, 57, 52).Im,
		img.LoadFirstFrame(dlpic(`qiao/1.png`), 0, 0).InsertUp(tou, 38, 36, 58, 50).Im,
	}
	img.SaveGif(img.MergeGif(1, qiao), cc.user+`敲.gif`)
	return "file:///" + cc.user + `敲.gif`

}

// 吃
func (cc *context) chi() string {
	tou := img.LoadFirstFrame(cc.imgs[0], 32, 32).Im
	chi := []*image.NRGBA{
		img.LoadFirstFrame(dlpic(`chi/0.png`), 0, 0).InsertBottom(tou, 0, 0, 1, 38).Im,
		img.LoadFirstFrame(dlpic(`chi/1.png`), 0, 0).InsertBottom(tou, 0, 0, 1, 38).Im,
		img.LoadFirstFrame(dlpic(`chi/2.png`), 0, 0).InsertBottom(tou, 0, 0, 1, 38).Im,
	}
	img.SaveGif(img.MergeGif(1, chi), cc.user+`吃.gif`)
	return "file:///" + cc.user + `吃.gif`

}

// 蹭
func (cc *context) ceng() string {
	tou := img.LoadFirstFrame(cc.imgs[0], 100, 100).Circle(0).Im
	tou2 := img.LoadFirstFrame(cc.imgs[1], 100, 100).Circle(0).Im
	ceng := []*image.NRGBA{
		img.LoadFirstFrame(dlpic(`ceng/0.png`), 0, 0).InsertUp(tou, 75, 77, 40, 88).InsertUp(tou2, 77, 103, 102, 81).Im,
		img.LoadFirstFrame(dlpic(`ceng/1.png`), 0, 0).InsertUp(tou, 75, 77, 46, 100).InsertUp(img.Rotate(tou2, 10, 62, 127).Im, 0, 0, 92, 40).Im,
		img.LoadFirstFrame(dlpic(`ceng/2.png`), 0, 0).InsertUp(tou, 75, 77, 67, 99).InsertUp(tou2, 76, 117, 90, 8).Im,
		img.LoadFirstFrame(dlpic(`ceng/3.png`), 0, 0).InsertUp(tou, 75, 77, 52, 83).InsertUp(img.Rotate(tou2, -40, 94, 94).Im, 0, 0, 53, -20).Im,
		img.LoadFirstFrame(dlpic(`ceng/4.png`), 0, 0).InsertUp(tou, 75, 77, 56, 110).InsertUp(img.Rotate(tou2, -66, 132, 80).Im, 0, 0, 78, 40).Im,
		img.LoadFirstFrame(dlpic(`ceng/5.png`), 0, 0).InsertUp(tou, 75, 77, 62, 102).InsertUp(tou2, 71, 100, 110, 94).Im,
	}
	img.SaveGif(img.MergeGif(8, ceng), cc.user+`蹭.gif`)
	// return "file:////" + cc.User + `蹭.gif`
	return "file:///" + cc.user + `蹭.gif`
}

// 啃
func (cc *context) ken() string {
	tou := img.LoadFirstFrame(cc.imgs[0], 100, 100).Im
	ken := []*image.NRGBA{
		img.LoadFirstFrame(dlpic(`ken/0.png`), 0, 0).InsertBottom(tou, 90, 90, 105, 150).Im,
		img.LoadFirstFrame(dlpic(`ken/1.png`), 0, 0).InsertBottom(tou, 90, 83, 96, 172).Im,
		img.LoadFirstFrame(dlpic(`ken/2.png`), 0, 0).InsertBottom(tou, 90, 90, 106, 148).Im,
		img.LoadFirstFrame(dlpic(`ken/3.png`), 0, 0).InsertBottom(tou, 88, 88, 97, 167).Im,
		img.LoadFirstFrame(dlpic(`ken/4.png`), 0, 0).InsertBottom(tou, 90, 85, 89, 179).Im,
		img.LoadFirstFrame(dlpic(`ken/5.png`), 0, 0).InsertBottom(tou, 90, 90, 106, 151).Im,
		img.LoadFirstFrame(dlpic(`ken/6.png`), 0, 0).Im,
		img.LoadFirstFrame(dlpic(`ken/7.png`), 0, 0).Im,
		img.LoadFirstFrame(dlpic(`ken/8.png`), 0, 0).Im,
		img.LoadFirstFrame(dlpic(`ken/9.png`), 0, 0).Im,
		img.LoadFirstFrame(dlpic(`ken/10.png`), 0, 0).Im,
		img.LoadFirstFrame(dlpic(`ken/11.png`), 0, 0).Im,
		img.LoadFirstFrame(dlpic(`ken/12.png`), 0, 0).Im,
		img.LoadFirstFrame(dlpic(`ken/13.png`), 0, 0).Im,
		img.LoadFirstFrame(dlpic(`ken/14.png`), 0, 0).Im,
		img.LoadFirstFrame(dlpic(`ken/15.png`), 0, 0).Im,
	}
	img.SaveGif(img.MergeGif(7, ken), cc.user+`啃.gif`)
	return "file:///" + cc.user + `啃.gif`
}

// 拍
func (cc *context) pai() string {
	tou := img.LoadFirstFrame(cc.imgs[0], 30, 30).Circle(0).Im
	pai := []*image.NRGBA{
		img.LoadFirstFrame(dlpic(`pai/0.png`), 0, 0).InsertUp(tou, 0, 0, 1, 47).Im,
		img.LoadFirstFrame(dlpic(`pai/1.png`), 0, 0).InsertUp(tou, 0, 0, 1, 67).Im,
	}
	img.SaveGif(img.MergeGif(1, pai), cc.user+`拍.gif`)
	return "file:///" + cc.user + `拍.gif`

}

// 冲
func (cc *context) chong() string {
	tou := img.LoadFirstFrame(cc.imgs[0], 0, 0).Circle(0).Im
	chong := []*image.NRGBA{
		img.LoadFirstFrame(dlpic(`xqe/0.png`), 0, 0).InsertUp(tou, 30, 30, 15, 53).Im,
		img.LoadFirstFrame(dlpic(`xqe/1.png`), 0, 0).InsertUp(tou, 30, 30, 40, 53).Im,
	}
	img.SaveGif(img.MergeGif(1, chong), cc.user+`冲.gif`)
	return "file:///" + cc.user + `冲.gif`

}

// 丢
func (cc *context) diu() string {
	tou := img.LoadFirstFrame(cc.imgs[0], 0, 0).Circle(0).Im
	diu := []*image.NRGBA{
		img.LoadFirstFrame(dlpic(`diu/0.png`), 0, 0).InsertUp(tou, 32, 32, 108, 36).Im,
		img.LoadFirstFrame(dlpic(`diu/1.png`), 0, 0).InsertUp(tou, 32, 32, 122, 36).Im,
		img.LoadFirstFrame(dlpic(`diu/2.png`), 0, 0).Im,
		img.LoadFirstFrame(dlpic(`diu/3.png`), 0, 0).InsertUp(tou, 123, 123, 19, 129).Im,
		img.LoadFirstFrame(dlpic(`diu/4.png`), 0, 0).InsertUp(tou, 185, 185, -50, 200).InsertUp(tou, 33, 33, 289, 70).Im,
		img.LoadFirstFrame(dlpic(`diu/5.png`), 0, 0).InsertUp(tou, 32, 32, 280, 73).Im,
		img.LoadFirstFrame(dlpic(`diu/6.png`), 0, 0).InsertUp(tou, 35, 35, 259, 31).Im,
		img.LoadFirstFrame(dlpic(`diu/7.png`), 0, 0).InsertUp(tou, 175, 175, -50, 220).Im,
	}
	img.SaveGif(img.MergeGif(7, diu), cc.user+`丢.gif`)
	return "file:///" + cc.user + `丢.gif`
}
