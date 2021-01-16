package gofile

import "fmt"

// A FileMode represents a file's mode and permission bits.
/*The bits have the same definition on all systems, so that
information about files can be moved from one system
to another portably. Not all bits apply to all systems.

The only required bit is ModeDir for directories.

The defined file mode bits are the most significant bits of the FileMode.
The nine least-significant bits are the standard Unix rwxrwxrwx permissions.
The values of these bits should be considered part of the public API and
may be used in wire protocols or disk representations: they must not be
changed, although new bits might be added.

 type FileMode uint32

	ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
	ModeAppend                                     // a: append-only
	ModeExclusive                                  // l: exclusive use
	ModeTemporary                                  // T: temporary file; Plan 9 only
	ModeSymlink                                    // L: symbolic link
	ModeDevice                                     // D: device file
	ModeNamedPipe                                  // p: named pipe (FIFO)
	ModeSocket                                     // S: Unix domain socket
	ModeSetuid                                     // u: setuid
	ModeSetgid                                     // g: setgid
	ModeCharDevice                                 // c: Unix character device, when ModeDevice is set
	ModeSticky                                     // t: sticky
	ModeIrregular                                  // ?: non-regular file; nothing else is known about this file

Mask for the type bits. For regular files, none will be set.
	ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice | ModeCharDevice | ModeIrregular

	ModePerm FileMode = 0777 // Unix permission bits
)
*/
const cFileMode = ""

// type bufio.Writer struct
/* fields
    err error
    buf []byte
    n   int
    wr  io.Writer

   func NewWriter(w io.Writer) *Writer
   func NewWriterSize(w io.Writer, size int) *Writer
   func (b *Writer) Available() int
   func (b *Writer) Buffered() int
   func (b *Writer) Flush() error
   func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)
   func (b *Writer) Reset(w io.Writer)
   func (b *Writer) Size() int
   func (b *Writer) Write(p []byte) (n int, err error)
   func (b *Writer) WriteByte(c byte) error
   func (b *Writer) WriteRune(r rune) (size int, err error)
   func (b *Writer) WriteString(s string) (int, error)*/
const cbufioWriter = ""

// type File (for build aix darwin dragonfly freebsd js,wasm linux netbsd openbsd solaris)
/* pfd         poll.FD
   name        string
   dirinfo     *dirInfo // nil unless directory being read
   nonblock    bool     // whether we set nonblocking mode
   stdoutOrErr bool     // whether this is stdout or stderr
   appendMode  bool     // whether file is opened for appending*/
const cFileunixdarwinwasm = ""

// type FileInfo interface {
/* Name() string       // base name of the file
   Size() int64        // length in bytes for regular files; system-dependent for others
   Mode() FileMode     // file mode bits
   ModTime() time.Time // modification time
   IsDir() bool        // abbreviation for Mode().IsDir()
   Sys() interface{}   // underlying data source (can return nil)*/
const cFileInfo = ""

// type File
/* func Create(name string) (*File, error)
   func NewFile(fd uintptr, name string) *File
   func Open(name string) (*File, error)
   func OpenFile(name string, flag int, perm FileMode) (*File, error)
   func (f *File) Chdir() error
   func (f *File) Chmod(mode FileMode) error
   func (f *File) Chown(uid, gid int) error
   func (f *File) Close() error
   func (f *File) Fd() uintptr
   func (f *File) Name() string
   func (f *File) Read(b []byte) (n int, err error)
   func (f *File) ReadAt(b []byte, off int64) (n int, err error)
   func (f *File) ReadFrom(r io.Reader) (n int64, err error)
   func (f *File) Readdir(n int) ([]FileInfo, error)
   func (f *File) Readdirnames(n int) (names []string, err error)
   func (f *File) Seek(offset int64, whence int) (ret int64, err error)
   func (f *File) SetDeadline(t time.Time) error
   func (f *File) SetReadDeadline(t time.Time) error
   func (f *File) SetWriteDeadline(t time.Time) error
   func (f *File) Stat() (FileInfo, error)
   func (f *File) Sync() error
   func (f *File) SyscallConn() (syscall.RawConn, error)
   func (f *File) Truncate(size int64) error
   func (f *File) Write(b []byte) (n int, err error)
   func (f *File) WriteAt(b []byte, off int64) (n int, err error)
   func (f *File) WriteString(s string) (n int, err error)*/
const cFileMethods = ""

func ShowDocs() {
	fmt.Println(cbufioWriter)
	fmt.Println(cFileInfo)
	fmt.Println(cFileMethods)
	fmt.Println(cFileunixdarwinwasm)
	fmt.Println(cFileMode)
}
