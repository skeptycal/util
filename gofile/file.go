package gofile

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
const c_bufioWriter = 0

// type File (for build aix darwin dragonfly freebsd js,wasm linux netbsd openbsd solaris)
/* pfd         poll.FD
   name        string
   dirinfo     *dirInfo // nil unless directory being read
   nonblock    bool     // whether we set nonblocking mode
   stdoutOrErr bool     // whether this is stdout or stderr
   appendMode  bool     // whether file is opened for appending*/
const c_File_unix_darwin_wasm = 0

// type FileInfo interface {
/* Name() string       // base name of the file
   Size() int64        // length in bytes for regular files; system-dependent for others
   Mode() FileMode     // file mode bits
   ModTime() time.Time // modification time
   IsDir() bool        // abbreviation for Mode().IsDir()
   Sys() interface{}   // underlying data source (can return nil)*/
const c_FileInfo = 0

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
const c_File_Methods = 0
