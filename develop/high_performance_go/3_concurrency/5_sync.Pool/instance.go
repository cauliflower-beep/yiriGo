package main

/*
	Go 语言标准库大量使用了 sync.Pool，例如 fmt 和 encoding/json。
	以下是 fmt.Printf 的源代码(go/src/fmt/print.go)
*/

// go 1.13.6

// pp is used to store a printer's state and is reused with sync.Pool to avoid allocations.

/*
	fmt.Printf 的调用是非常频繁的，利用 sync.Pool 复用 pp 对象能够极大的提升性能，减少内存分配次数及占用，降低 GC 压力
	type pp struct {
	buf bytes
	...
	}

	var ppFree = sync.Pool{
		New: func() interface{} { return new(pp) },
	}

	// newPrinter allocates a new pp struct or grabs a cached one.
	func newPrinter() *pp {
		p := ppFree.Get().(*pp)
		p.panicking = false
		p.erroring = false
		p.wrapErrs = false
		p.fmt.init(&p.buf)
		return p
	}

	// free saves used pp structs in ppFree; avoids an allocation per invocation.
	func (p *pp) free() {
		if cap(p.buf) > 64<<10 {
			return
		}

		p.buf = p.buf[:0]
		p.arg = nil
		p.value = _reflect.Value{}
		p.wrappedErr = nil
		ppFree.Put(p)
	}

	func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
		p := newPrinter()
		p.doPrintf(format, a)
		n, err = w.Write(p.buf)
		p.free()
		return
	}

	// Printf formats according to a format specifier and writes to standard output.
	// It returns the number of bytes written and any write error encountered.
	func Printf(format string, a ...interface{}) (n int, err error) {
		return Fprintf(os.Stdout, format, a...)
	}
*/
