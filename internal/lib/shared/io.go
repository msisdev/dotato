package shared

import (
	"github.com/msisdev/dotato/internal/lib/io"
	"github.com/msisdev/dotato/pkg/dotato"
)

func (s Shared) ImportFile(pre dotato.Preview) error {
	err := io.ImportFile(s.fs, s.d, pre)
	if err != nil {
		return err
	}
	return nil
}

func (s Shared) ImportLink(pre dotato.Preview) error {
	err := io.ImportLink(s.fs, s.d, pre)
	if err != nil {
		return err
	}
	return nil
}

func (s Shared) ExportFile(pre dotato.Preview) error {
	err := io.ExportFile(s.fs, s.d, pre)
	if err != nil {
		return err
	}
	return nil
}

func (s Shared) ExportLink(pre dotato.Preview) error {
	err := io.ExportLink(s.fs, s.d, pre)
	if err != nil {
		return err
	}
	return nil
}

func (s Shared) Unlink(pre dotato.Preview) error {
	err := io.Unlink(s.fs, s.d, pre)
	if err != nil {
		return err
	}
	return nil
}