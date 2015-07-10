package main

import "io"

type argWriter struct {
	writer io.WriteCloser
	err    error
}

func newArgWriterPtr(writer io.WriteCloser, err error) *argWriter {
	return &argWriter{writer, err}
}

func newArgWriterVal(writer io.WriteCloser, err error) argWriter {
	return argWriter{writer, err}
}

func (w argWriter) writeVal(f func() error) error {
	if w.err != nil {
		return w.err
	}

	if err := f(); err != nil {
		return err
	}

	return w.writer.Close()
}

func (w argWriter) writePtr(f func() error) error {
	if w.err != nil {
		return w.err
	}

	if err := f(); err != nil {
		return err
	}

	return w.writer.Close()
}

func (w argWriter) WriteIndirectVal(bs []byte) error {
	return w.writeVal(func() error {
		_, err := w.writer.Write(bs)
		return err
	})
}

func (w *argWriter) WriteIndirectPtr(bs []byte) error {
	return w.writePtr(func() error {
		_, err := w.writer.Write(bs)
		return err
	})
}

func (w argWriter) WriteDirectVal(bs []byte) error {
	if w.err != nil {
		return w.err
	}

	if _, err := w.writer.Write(bs); err != nil {
		return err
	}

	return w.writer.Close()
}

func (w *argWriter) WriteDirectPtr(bs []byte) error {
	if w.err != nil {
		return w.err
	}

	if _, err := w.writer.Write(bs); err != nil {
		return err
	}

	return w.writer.Close()
}
