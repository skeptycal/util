package fileutils

// File Utilities
// * use these with some adjustments:
// * func NewFile(fd uintptr, name string) *File
// *    if fd non-blocking, NewFile returns pollable, so ...

// * func (f *File) Fd() uintptr (always returns in blocking mode)
// *    change to non-blocking to pass to NewFile as needed per
// *    go/1.15.3/libexec/src/os/file_unix.go newFile():

// *    "If the caller passed a non-blocking filedes
// *     (kindNonBlock), we assume they know what they
// *     are doing so we allow it to be used with kqueue."

//  use these os utilities directly:
//      const DevNull = "/dev/null"
//      func Symlink(oldname, newname string) error
//      func Link(oldname, newname string) error
//      func tempDir() string
//      func Remove(name string) error
//      func Truncate(name string, size int64) error

// todo - something like this
// func NewFile(fd uintptr, name string) *File {
// 	kind := kindNewFile
// 	if nb, err := unix.IsNonblock(int(fd)); err == nil && nb {
// 		kind = kindNonBlock
// 	}
// 	return newFile(fd, name, kind)
// }

// todo - something like this ...
// Fd returns the integer Unix file descriptor referencing the open file.
// The file descriptor is valid only until f.Close is called or f is garbage collected.
// On Unix systems this will cause the SetDeadline methods to stop working.
// func (f *os.File) FdNb() uintptr {
// 	if f == nil {
// 		return ^(uintptr(0))
// 	}

// 	// Historically we have always returned a descriptor
// 	// opened in blocking mode, but this function is specifically
// 	// designed to return a non-blocking fd.
// 	f.pfd.SetBlocking()
// 	return uintptr(f.pfd.Sysfd)
// }
