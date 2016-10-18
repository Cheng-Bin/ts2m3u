package utils

import (
	"fmt"
)

//PlayListType : EXT-X-PLAYLIST-TYPE
type PlayListType string

// VOD and EVENT
const (
	VOD   PlayListType = "VOD"
	EVENT PlayListType = "EVENT"
)

// M3U8Writer object
type M3U8Writer struct {
	m3u8 string
}

// NewM3U8 ： 构造方法
// autoHeader :  是否自动生成header
func NewM3U8(autoHeader bool) *M3U8Writer {

	m3u8 := new(M3U8Writer)

	if autoHeader {
		m3u8.WriteHeader()
	}

	return m3u8
}

// GetM3U8 : m3u8
func (that *M3U8Writer) GetM3U8() string {
	return that.m3u8
}

// Writer a line data
func (that *M3U8Writer) Writer(line string) {
	if line != "" {
		that.m3u8 += line + "\n"
	}
}

// WriteHeader : 4
func (that *M3U8Writer) WriteHeader() {
	that.WriteM3U()
	that.WriteVersion(3)
	that.WriteSequence(0)
	that.SetEnableCache(true)
	that.WritePlayListType(VOD)
}

// WriteM3U : #EXTM3U
func (that *M3U8Writer) WriteM3U() {
	that.Writer("#EXTM3U")
}

// WriteVersion : #EXT-X-VERSION
func (that *M3U8Writer) WriteVersion(version int) {
	that.Writer(fmt.Sprintf("#EXT-X-VERSION:%d", version))
}

// WriteTargetDuration : EXT-X-TARGETDURATION
func (that *M3U8Writer) WriteTargetDuration(duration int) {
	that.Writer(fmt.Sprintf("#EXT-X-TARGETDURATION:%d", duration))
}

// WriteSequence : #EXT-X-MEDIA-SEQUENCE
func (that *M3U8Writer) WriteSequence(sequence int) {
	that.Writer(fmt.Sprintf("#EXT-X-MEDIA-SEQUENCE:%d", sequence))
}

// SetEnableCache : #EXT-X-ALLOW-CACHE
func (that *M3U8Writer) SetEnableCache(cache bool) {
	if cache {
		that.Writer("#EXT-X-ALLOW-CACHE:YES")
	} else {
		that.Writer("#EXT-X-ALLOW-CACHE:NO")
	}

}

// WritePlayListType : #EXT-X-PLAYLIST-TYPE
func (that *M3U8Writer) WritePlayListType(t PlayListType) {
	that.Writer("#EXT-X-PLAYLIST-TYPE:" + string(t))
}

// WriteTs  写入ts文件到列表
//
func (that *M3U8Writer) WriteTs(duration string, ts string) {
	that.Writer(fmt.Sprintf("EXTINF:%s,", duration))
	that.Writer(ts)
}

// WriterEnd : #EXT-X-ENDLIST
func (that *M3U8Writer) WriterEnd() {
	that.Writer("#EXT-X-ENDLIST")
}
