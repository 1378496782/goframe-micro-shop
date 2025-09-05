/**
 * 基于磁盘的 FIFO 先进先出消息队列
 * source nsq
 * https://github.com/nsqio/go-diskqueue/blob/master/diskqueue.go
 */

package disk

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path"
	"sync"
	"time"
)

type LogLevel int

const (
	DEBUG = LogLevel(1)
	INFO  = LogLevel(2)
	WARN  = LogLevel(3)
	ERROR = LogLevel(4)
	FATAL = LogLevel(5)
)

type AppLogFunc func(lvl LogLevel, f string, args ...interface{})

func (l LogLevel) String() string {
	switch l {
	case 1:
		return "DEBUG"
	case 2:
		return "INFO"
	case 3:
		return "WARNING"
	case 4:
		return "ERROR"
	case 5:
		return "FATAL"
	}
	panic("invalid LogLevel")
}

type IDiskQueue interface {
	// Put 加锁检测 标志位是否退出, 如果否 则继续往文件写入数据并等待结果
	// 写入*.meta.dat 格式
	// 1.depth
	// 2.readFileNum,readPos
	// 3.writeFileNum,writePos
	Put([]byte) error
	// ReadChan 读取文件数据 返回只读的chan 可用于多消费者并发读取
	ReadChan() <-chan []byte
	PeekChan() <-chan []byte
	// Close 等待loop结束, 正常关闭 并保存元数据
	Close() error
	// Delete 等待loop结束, 直接关闭io流
	Delete() error
	// Depth 未读消息积压量
	Depth() int64
	// Empty 清空消息, 删除文件
	Empty() error
}

// diskQueue 基于文件系统的先进先出队列
type diskQueue struct {
	// 64bit atomic vars need to be first for proper alignment on 32bit platforms

	// run-time state (also persisted to disk)
	// 运行时的数据保存, 也会保存到文件
	readPos      int64 // 文件读取的位置
	writePos     int64 // 文件写入的位置
	readFileNum  int64 // 读取文件编号
	writeFileNum int64 // 写入文件编号
	depth        int64 // 未读消息积压量

	sync.RWMutex // 读写锁

	// instantiation time metadata
	name                string // 队列实例名称
	dataPath            string // 数据文件存储路径
	maxBytesPerFile     int64  // 每个文件最大字节数
	maxBytesPerFileRead int64
	minMsgSize          int32         // 最小消息长度
	maxMsgSize          int32         // 最大消息长度
	syncEvery           int64         // 刷盘频率设置
	syncTimeout         time.Duration // 刷盘时间设置（单位秒）
	exitFlag            int32         // 退出标志位
	needSync            bool          // 是否需要同步

	// 跟踪读取的位置
	nextReadPos     int64 // 下一次读取的位置
	nextReadFileNum int64 // 下一次读取对应的文件编号

	readFile  *os.File      // 读取的文件
	writeFile *os.File      // 写入的文件
	reader    *bufio.Reader // 缓冲读取
	writeBuf  bytes.Buffer  // 缓冲写入

	// 通过 ReadChan() 方法对外暴露
	readChan chan []byte // 读取的数据,可以多消费者进行通信消费

	// 通过 PeekChan() 方法对外暴露
	peekChan chan []byte // 探查数据

	// 内部管道
	depthChan         chan int64
	writeChan         chan []byte // 写入通道
	writeResponseChan chan error  // 写入结果反馈通道
	emptyChan         chan int    // 清空队列通道
	emptyResponseChan chan error  // 清空反馈通道
	exitChan          chan int    // 结束信号通道
	exitSyncChan      chan int    // 退出同步通道
	logf              AppLogFunc  // 日志记录封装
}

// NewDiskQueue 实例化一个diskQueue，从文件系统中检索元数据并启动预计程序
func NewDiskQueue(name, dataPath string, maxBytesPerFile int64, minMsgSize, maxMsgSize int32, syncEvery int64, syncTimeout time.Duration, logf AppLogFunc) IDiskQueue {
	d := diskQueue{
		name:              name,
		dataPath:          dataPath,
		maxBytesPerFile:   maxBytesPerFile,
		minMsgSize:        minMsgSize,
		maxMsgSize:        maxMsgSize,
		syncEvery:         syncEvery,
		syncTimeout:       syncTimeout,
		readChan:          make(chan []byte),
		peekChan:          make(chan []byte),
		depthChan:         make(chan int64),
		writeChan:         make(chan []byte),
		writeResponseChan: make(chan error),
		emptyChan:         make(chan int),
		emptyResponseChan: make(chan error),
		exitChan:          make(chan int),
		exitSyncChan:      make(chan int),
		logf:              logf,
	}
	// 从磁盘加载元数据
	err := d.retrieveMetaData()
	if err != nil && !os.IsNotExist(err) {
		d.logf(ERROR, "DISKQUEUE(%s) failed to retrieveMetaData - %s", d.name, err)
	}
	go d.ioLoop()
	return &d
}

// Depth 获取消息深度(条数)
func (d *diskQueue) Depth() int64 {
	depth, ok := <-d.depthChan
	if !ok {
		// ioLoop exited
		depth = d.depth
	}
	return depth
}

// ReadChan 返回一个用于读取字节型数据的通道
func (d *diskQueue) ReadChan() <-chan []byte {
	return d.readChan
}

// PeekChan 返回一个用于探测字数据的通道，不会对影响消息队列状态
func (d *diskQueue) PeekChan() <-chan []byte {
	return d.peekChan
}

// Put 写入字节型数据到队列中
func (d *diskQueue) Put(data []byte) error {
	d.RLock()
	defer d.RUnlock()

	if d.exitFlag == 1 {
		return errors.New("exiting")
	}

	d.writeChan <- data
	return <-d.writeResponseChan
}

// Close cleans up the queue and persists metadata
// 清理队列并持久化元数据
func (d *diskQueue) Close() error {
	err := d.exit(false)
	if err != nil {
		return err
	}
	return d.sync()
}

func (d *diskQueue) Delete() error {
	return d.exit(true)
}

func (d *diskQueue) exit(deleted bool) error {
	d.Lock()
	defer d.Unlock()

	d.exitFlag = 1

	if deleted {
		d.logf(INFO, "DISKQUEUE(%s): deleting", d.name)
	} else {
		d.logf(INFO, "DISKQUEUE(%s): closing", d.name)
	}
	// 关闭退出通道
	close(d.exitChan)

	// 确保ioLoop已经退出
	<-d.exitSyncChan

	// 关闭文件读写
	close(d.depthChan)
	if d.readFile != nil {
		d.readFile.Close()
		d.readFile = nil
	}
	if d.writeFile != nil {
		d.writeFile.Close()
		d.writeFile = nil
	}

	return nil
}

// Empty destructively clears out any pending data in the queue
// by fast forwarding read positions and removing intermediate files
// 破坏性的清空数据，删除文件
func (d *diskQueue) Empty() error {
	d.RLock()
	defer d.RUnlock()

	if d.exitFlag == 1 {
		return errors.New("exiting")
	}

	d.logf(INFO, "DISKQUEUE(%s): emptying", d.name)

	d.emptyChan <- 1
	return <-d.emptyResponseChan
}

// 删除所有文件
func (d *diskQueue) deleteAllFiles() error {
	err := d.skipToNextRWFile()

	innerErr := os.Remove(d.metaDataFileName())
	if innerErr != nil && !os.IsNotExist(innerErr) {
		d.logf(ERROR, "DISKQUEUE(%s) failed to remove metadata file - %s", d.name, innerErr)
		return innerErr
	}

	return err
}

// 跳到下一个读写文件
func (d *diskQueue) skipToNextRWFile() error {
	var err error

	if d.readFile != nil {
		d.readFile.Close()
		d.readFile = nil
	}

	if d.writeFile != nil {
		d.writeFile.Close()
		d.writeFile = nil
	}

	for i := d.readFileNum; i <= d.writeFileNum; i++ {
		fileName := d.fileName(i)
		innerErr := os.Remove(fileName)
		if innerErr != nil && !os.IsNotExist(innerErr) {
			d.logf(ERROR, "DISKQUEUE(%s) failed to remove data file - %s", d.name, innerErr)
			err = innerErr
		}
	}

	d.writeFileNum++
	d.writePos = 0
	d.readFileNum = d.writeFileNum
	d.readPos = 0
	d.nextReadFileNum = d.writeFileNum
	d.nextReadPos = 0
	d.depth = 0

	return err
}

func (d *diskQueue) readOne() ([]byte, error) {
	var err error
	var msgSize int32

	if d.readFile == nil {
		curFileName := d.fileName(d.readFileNum)
		d.readFile, err = os.OpenFile(curFileName, os.O_RDONLY, 0600)
		if err != nil {
			return nil, err
		}

		d.logf(INFO, "DISKQUEUE(%s): readOne() opened %s", d.name, curFileName)

		if d.readPos > 0 {
			// 设置读取偏移
			_, err = d.readFile.Seek(d.readPos, 0)
			if err != nil {
				d.readFile.Close()
				d.readFile = nil
				return nil, err
			}
		}

		// for "complete" files (i.e. not the "current" file), maxBytesPerFileRead
		// should be initialized to the file's size, or default to maxBytesPerFile
		// maxBytesPerFileRead应该初始化为文件的大小，或者默认为maxBytesPerFile
		d.maxBytesPerFileRead = d.maxBytesPerFile
		if d.readFileNum < d.writeFileNum {
			stat, err := d.readFile.Stat()
			if err == nil {
				d.maxBytesPerFileRead = stat.Size()
			}
		}

		d.reader = bufio.NewReader(d.readFile)
	}

	err = binary.Read(d.reader, binary.BigEndian, &msgSize)
	if err != nil {
		d.readFile.Close()
		d.readFile = nil
		return nil, err
	}

	if msgSize < d.minMsgSize || msgSize > d.maxMsgSize {
		// this file is corrupt and we have no reasonable guarantee on
		// where a new message should begin
		// 文件损坏，无法读取
		d.readFile.Close()
		d.readFile = nil
		return nil, fmt.Errorf("invalid message read size (%d)", msgSize)
	}

	readBuf := make([]byte, msgSize)
	_, err = io.ReadFull(d.reader, readBuf)
	if err != nil {
		d.readFile.Close()
		d.readFile = nil
		return nil, err
	}

	totalBytes := int64(4 + msgSize)

	// we only advance next* because we have not yet sent this to consumers
	// (where readFileNum, readPos will actually be advanced)
	d.nextReadPos = d.readPos + totalBytes
	d.nextReadFileNum = d.readFileNum

	// we only consider rotating if we're reading a "complete" file
	// and since we cannot know the size at which it was rotated, we
	// rely on maxBytesPerFileRead rather than maxBytesPerFile
	// if d.readFileNum < d.writeFileNum && d.nextReadPos >= d.maxBytesPerFileRead {
	if d.readFileNum <= d.writeFileNum && d.nextReadPos >= d.maxBytesPerFileRead {
		if d.readFile != nil {
			d.readFile.Close()
			d.readFile = nil
		}

		d.nextReadFileNum++
		d.nextReadPos = 0
	}

	return readBuf, nil
}

// writeOne performs a low level filesystem write for a single []byte
// while advancing write positions and rolling files, if necessary
func (d *diskQueue) writeOne(data []byte) error {
	var err error

	dataLen := int32(len(data))
	totalBytes := int64(4 + dataLen)

	if dataLen < d.minMsgSize || dataLen > d.maxMsgSize {
		return fmt.Errorf("invalid message write size (%d) minMsgSize=%d maxMsgSize=%d", dataLen, d.minMsgSize, d.maxMsgSize)
	}

	// will not wrap-around if maxBytesPerFile + maxMsgSize < Int64Max
	// 如果文件大小超过了设置的最大值，关闭当前写入文件
	if d.writePos > 0 && d.writePos+totalBytes > d.maxBytesPerFile {
		if d.readFileNum == d.writeFileNum {
			d.maxBytesPerFileRead = d.writePos
		}

		d.writeFileNum++
		d.writePos = 0

		// sync every time we start writing to a new file
		err = d.sync()
		if err != nil {
			d.logf(ERROR, "DISKQUEUE(%s) failed to sync - %s", d.name, err)
		}

		if d.writeFile != nil {
			d.writeFile.Close()
			d.writeFile = nil
		}
	}

	if d.writeFile == nil {
		curFileName := d.fileName(d.writeFileNum)
		d.writeFile, err = os.OpenFile(curFileName, os.O_RDWR|os.O_CREATE, 0600)
		if err != nil {
			return err
		}

		d.logf(INFO, "DISKQUEUE(%s): writeOne() opened %s", d.name, curFileName)

		if d.writePos > 0 {
			_, err = d.writeFile.Seek(d.writePos, 0)
			if err != nil {
				d.writeFile.Close()
				d.writeFile = nil
				return err
			}
		}
	}

	d.writeBuf.Reset()
	// 缓冲区写入二进制数据
	err = binary.Write(&d.writeBuf, binary.BigEndian, dataLen)
	if err != nil {
		return err
	}

	_, err = d.writeBuf.Write(data)
	if err != nil {
		return err
	}

	// only write to the file once
	// 写入到文件
	_, err = d.writeFile.Write(d.writeBuf.Bytes())
	if err != nil {
		d.writeFile.Close()
		d.writeFile = nil
		return err
	}

	d.writePos += totalBytes
	d.depth += 1

	return err
}

// sync fsyncs the current writeFile and persists metadata
// 刷盘，持久化数据
func (d *diskQueue) sync() error {
	if d.writeFile != nil {
		// 数据写入磁盘
		err := d.writeFile.Sync()
		if err != nil {
			d.writeFile.Close()
			d.writeFile = nil
			return err
		}
	}

	// 持久化元数据
	err := d.persistMetaData()
	if err != nil {
		return err
	}

	d.needSync = false
	return nil
}

// retrieveMetaData initializes state from the filesystem
// 从本地文件取回元数据
func (d *diskQueue) retrieveMetaData() error {
	var f *os.File
	var err error

	// 存储路径.diskqueue.meta.dat
	fileName := d.metaDataFileName()
	f, err = os.OpenFile(fileName, os.O_RDONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	var depth int64
	_, err = fmt.Fscanf(f, "%d\n%d,%d\n%d,%d\n",
		&depth,          // 待读取消息数量
		&d.readFileNum,  // 待读取文件编号
		&d.readPos,      // 待读取文件位置
		&d.writeFileNum, // 待写入文件编号
		&d.writePos,     // 待写入文件位置
	)
	if err != nil {
		return err
	}

	d.depth = depth
	// 更新读取位置和文件编号
	d.nextReadFileNum = d.readFileNum
	d.nextReadPos = d.readPos

	// 如果元数据在最后一次关闭系统时没有同步，
	// 那么实际文件的大小实际上可能大于writePos写入位置，
	// 在这种情况下，最安全的做法是跳到下一个文件进行写操作，
	// 并让读取器从磁盘队列中的消息中尽可能地抢救出元数据(可能也是过时的readPos)之外的内容
	fileName = d.fileName(d.writeFileNum)
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		return err
	}
	fileSize := fileInfo.Size()
	if d.writePos < fileSize {
		d.logf(WARN, "DISKQUEUE(%s) %s metadata writePos %d < file size of %d, skipping to new file", d.name, fileName, d.writePos, fileSize)
		d.writeFileNum += 1
		d.writePos = 0
		if d.writeFile != nil {
			d.writeFile.Close()
			d.writeFile = nil
		}
	}
	return nil
}

// persistMetaData atomically writes state to the filesystem
// 同步元数据到本地文件
func (d *diskQueue) persistMetaData() error {
	var f *os.File
	var err error
	// metaDat 临时文件
	fileName := d.metaDataFileName()
	tmpFileName := fmt.Sprintf("%s.%d.tmp", fileName, rand.Int())

	// 写入到临时文件
	f, err = os.OpenFile(tmpFileName, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(f, "%d\n%d,%d\n%d,%d\n", d.depth, d.readFileNum, d.readPos, d.writeFileNum, d.writePos)
	if err != nil {
		f.Close()
		return err
	}
	f.Sync()
	f.Close()
	// 成功往临时文件写入数据, 在进行替换源文件
	// atomically rename
	return os.Rename(tmpFileName, fileName)
}

// 获取元数据文件名称
func (d *diskQueue) metaDataFileName() string {
	return fmt.Sprintf(path.Join(d.dataPath, "%s.diskqueue.meta.dat"), d.name)
}

// 获取文件名称
func (d *diskQueue) fileName(fileNum int64) string {
	return fmt.Sprintf(path.Join(d.dataPath, "%s.diskqueue.%06d.dat"), d.name, fileNum)
}

func (d *diskQueue) checkTailCorruption(depth int64) {
	if d.readFileNum < d.writeFileNum || d.readPos < d.writePos {
		return
	}

	// we've reached the end of the diskqueue
	// if depth isn't 0 something went wrong
	if depth != 0 {
		if depth < 0 {
			d.logf(ERROR, "DISKQUEUE(%s) negative depth at tail (%d), metadata corruption, resetting 0...", d.name, depth)
		} else if depth > 0 {
			d.logf(ERROR, "DISKQUEUE(%s) positive depth at tail (%d), data loss, resetting 0...", d.name, depth)
		}
		// 强制设置为0
		d.depth = 0
		d.needSync = true
	}

	if d.readFileNum != d.writeFileNum || d.readPos != d.writePos {
		if d.readFileNum > d.writeFileNum {
			d.logf(ERROR, "DISKQUEUE(%s) readFileNum > writeFileNum (%d > %d), corruption, skipping to next writeFileNum and resetting 0...", d.name, d.readFileNum, d.writeFileNum)
		}
		if d.readPos > d.writePos {
			d.logf(ERROR, "DISKQUEUE(%s) readPos > writePos (%d > %d), corruption, skipping to next writeFileNum and resetting 0...", d.name, d.readPos, d.writePos)
		}

		// 强制跳到下一个文件
		d.skipToNextRWFile()
		d.needSync = true
	}
}

func (d *diskQueue) moveForward() {
	oldReadFileNum := d.readFileNum
	d.readFileNum = d.nextReadFileNum
	d.readPos = d.nextReadPos
	d.depth -= 1

	//see if we need to clean up the old file
	if oldReadFileNum != d.nextReadFileNum {
		// sync every time we start reading from a new file
		// 每次读取新文件时同步
		d.needSync = true

		fileName := d.fileName(oldReadFileNum)
		if d.writeFile != nil && d.writeFile.Name() == fileName {
			d.writeFile.Close()
			d.writeFile = nil
		}
		err := os.Remove(fileName)
		if err != nil {
			d.logf(ERROR, "DISKQUEUE(%s) failed to Remove(%s) - %s", d.name, fileName, err)
		}
	}
	d.checkTailCorruption(d.depth)
}

func (d *diskQueue) handleReadEOFError() {
	fileName := d.fileName(d.readFileNum)
	err := os.Remove(fileName)
	if err != nil {
		d.logf(ERROR, "DISKQUEUE(%s) failed to Remove(%s) - %s", d.name, fileName, err)
	}
	d.readFileNum++
	d.readPos = 0
	d.nextReadFileNum++
	d.nextReadPos = 0
}

// 处理读取错误
func (d *diskQueue) handleReadError() {
	// 跳转到下一个读取文件并重命名当前(损坏的)文件
	if d.readFileNum == d.writeFileNum {
		// 如果你不能正确地从当前的写文件中读取，那就应该跳过当前的文件
		if d.writeFile != nil {
			d.writeFile.Close()
			d.writeFile = nil
		}
		d.writeFileNum++
		d.writePos = 0
	}
	badFileName := d.fileName(d.readFileNum)
	badRenameFilename := badFileName + ".bad"
	d.logf(WARN, "DISKQUEUE(%s) jump to next file and saving bad file as %s", d.name, badRenameFilename)
	err := os.Rename(badFileName, badRenameFilename)
	if err != nil {
		d.logf(ERROR, "DISKQUEUE(%s) failed to rename bad diskqueue file %s to %s", d.name, badFileName, badRenameFilename)
	}
	d.readFileNum++
	d.readPos = 0
	d.nextReadFileNum = d.readFileNum
	d.nextReadPos = 0

	// 重大状态变更，在下一次迭代中安排同步
	d.needSync = true

	d.checkTailCorruption(d.depth)
}

// ioLoop 提供了暴露go通道的后端(通过ReadChan())，以支持多个并发队列消费者
// 它的工作原理是基于队列是否有数据可读和阻塞，直到数据通过适当的go通道读取或写入
// 简单来讲，这也意味着我们正在异步地从文件系统读取数据
func (d *diskQueue) ioLoop() {
	var (
		dataRead []byte
		err      error
		count    int64
		r        chan []byte
		p        chan []byte
	)
	// 设置定时器
	syncTicker := time.NewTicker(d.syncTimeout)

	for {
		// 若到达刷盘频次，标记等待刷盘
		if count == d.syncEvery {
			d.needSync = true
		}

		if d.needSync {
			err = d.sync()
			if err != nil {
				d.logf(ERROR, "DISKQUEUE(%s) failed to sync - %s", d.name, err)
			}
			count = 0
		}
		// 有可读数据，并且当前读chan的数据已经被读走，则读取下一条数据
		//if (d.readFileNum < d.writeFileNum) || (d.readPos < d.writePos) {
		if d.readFileNum < d.writeFileNum || (d.readFileNum == d.writeFileNum && d.readPos < d.writePos) {
			if d.nextReadPos == d.readPos {
				dataRead, err = d.readOne()
				if err != nil {
					if io.EOF == err {
						d.logf(WARN, "DISKQUEUE(%s) reading at %d of %s - %s", d.name, d.readPos, d.fileName(d.readFileNum), err)
						// 读取了文件未尾，同时进行读写时可能会出现这个问题
						d.handleReadEOFError()
					} else {
						d.logf(ERROR, "DISKQUEUE(%s) reading at %d of %s - %s", d.name, d.readPos, d.fileName(d.readFileNum), err)
						d.handleReadError()
					}
					continue
				}
			}
			r = d.readChan
			p = d.peekChan
		} else {
			// 如果无可读数据，那么设置 r 为nil, 防止将dataRead数据重复传入readChan中
			r = nil
			p = nil
		}

		// Go通道规范规定，跳过select中的nil通道操作(读或写)，只有在有数据要读时，我们才将r设置为d.readChan
		select {
		// peekChan 不需要改变状态
		case p <- dataRead:
		// readChan 需要改变状态
		case r <- dataRead:
			count++
			// 如果读readChan成功，则会修改读的偏移
			d.moveForward()
		case d.depthChan <- d.depth:
		case <-d.emptyChan:
			d.emptyResponseChan <- d.deleteAllFiles()
			count = 0
		case dataWrite := <-d.writeChan:
			count++
			d.writeResponseChan <- d.writeOne(dataWrite)
		case <-syncTicker.C:
			// 到刷盘时间，则修改needSync = true
			if count == 0 {
				// 没有活动时不要刷盘
				continue
			}
			d.needSync = true
		case <-d.exitChan:
			// 退出
			goto exit
		}
	}

exit:
	d.logf(INFO, "DISKQUEUE(%s): closing ... ioLoop", d.name)
	syncTicker.Stop()
	d.exitSyncChan <- 1
}
