package middleware

import (
	"bytes"
	"compress/gzip"
	"knife"
)

// GzipDefault 当响应体的大小超过100kb后，进行gzip压缩
func GzipDefault() knife.MiddlewareFunc {
	return Gzip(100 * 1024)
}

// Gzip 对响应的数据进行压缩
// response Header 设置 Accept-Encoding=gzip ，说明响应的数据是gzip压缩过的，需要客户端进行gzip解压
// gzipMinLength 是设置响应体大小的阈值，超过这个阈值会进行响应数据的压缩
func Gzip(gzipMinLength int) knife.MiddlewareFunc {
	return func(context *knife.Context) {
		writer := context.Writer
		context.Writer = &gzipResponseWriter{
			HttpResponseWriter: writer,
			gzipMinLength:      gzipMinLength,
		}
		context.Next()
	}
}

// 自定义的ResponseWriter，将数据写入gzip.Writer
type gzipResponseWriter struct {
	knife.HttpResponseWriter
	gzipMinLength int
}

// 重写Write方法，将数据写入gzip.Writer
func (w *gzipResponseWriter) Write(data []byte) (int, error) {

	beforeSize := len(data)
	if beforeSize > w.gzipMinLength {
		// 添加 gzip 头部
		w.Header().Set("Content-Encoding", "gzip")

		// 调用gzip.Writer的Write方法，将数据压缩后写入到底层的http.ResponseWriter
		var zBuf bytes.Buffer
		zw := gzip.NewWriter(&zBuf)
		if n, err := zw.Write(data); err != nil {
			err := zw.Close()
			if err != nil {
				return n, err
			}
			knife.Logger.Printf("gzip is failed,err: %s", err)
		}
		err := zw.Close()
		if err != nil {
			return 0, err
		}
		afterData := zBuf.Bytes()
		knife.Logger.Printf("gzip success,befor size: %d, after size: %d", beforeSize, len(afterData))

		return w.HttpResponseWriter.Write(afterData)
	}
	return w.HttpResponseWriter.Write(data)
}
