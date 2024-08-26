package websocket

import (
	"bufio"
	"net"
)

func NewConn(conn net.Conn, isServer bool, readBufferSize, writeBufferSize int, writeBufferPool BufferPool, br *bufio.Reader, writeBuf []byte) *Conn {
	return newConn(conn, isServer, readBufferSize, writeBufferSize, writeBufferPool, br, writeBuf)
}

func (c *Conn) SetupCompression(level ...int) *Conn {
	if len(level) > 0 {
		c.SetCompressionLevel(level[0])
	}
	c.enableWriteCompression = true
	c.newCompressionWriter = compressNoContextTakeover
	c.newDecompressionReader = decompressNoContextTakeover
	return c
}
