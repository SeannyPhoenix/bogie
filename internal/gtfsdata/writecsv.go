package gtfsdata

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path"

	"github.com/seannyphoenix/bogie/internal/gtfsdata/cardinality"
	"github.com/seannyphoenix/bogie/pkg/csvmum"
)

type writeCsvOptions[T keyedData] struct {
	records []T
	path    string
	gzip    bool
}

type meta struct {
	Headers string `json:"headers"`
	Bits    int    `json:"bits"`
}

func writeCsv[T keyedData](ctx context.Context, opts writeCsvOptions[T]) (meta, error) {
	slog.InfoContext(ctx, fmt.Sprintf("Writing %d records to %s", len(opts.records), opts.path))

	var m meta

	err := preparePath(opts)
	if err != nil {
		return m, fmt.Errorf("error preparing path %s: %w", opts.path, err)
	}

	var buf bytes.Buffer
	cm, err := csvmum.NewMarshaler[T](&buf)
	if err != nil {
		return m, fmt.Errorf("error creating CSV marshaler: %w", err)
	}
	err = cm.Flush()
	if err != nil {
		return m, fmt.Errorf("error flushing CSV marshaler: %w", err)
	}

	headers, err := buf.ReadBytes('\n')
	if err != nil && err != io.EOF {
		return m, fmt.Errorf("error reading headers: %w", err)
	}
	buf.Reset()
	m.Headers = string(headers)

	// err = writeHeaders(&buf, opts)
	// if err != nil {
	// 	return fmt.Errorf("error writing headers: %w", err)
	// }

	var ids [][]byte
	uniqueIds := make(map[string]struct{})

	for _, r := range opts.records {
		err := cm.Marshal(r)
		if err != nil {
			return m, fmt.Errorf("error marshaling record: %w", err)
		}

		key := r.partitionKey()
		ids = append(ids, key)
		uniqueIds[string(key)] = struct{}{}
	}
	err = cm.Flush()
	if err != nil {
		return m, fmt.Errorf("error flushing CSV marshaler: %w", err)
	}

	c := cardinality.GetCardinality(cardinality.MB, buf.Len(), len(uniqueIds))
	m.Bits = c.Bits

	recordPosition := 0
	for i := range c.Buckets {
		var data bytes.Buffer

		for recordPosition < len(ids) {
			pre := cardinality.GetPrefix(ids[recordPosition], c.Bits)
			if pre != uint64(i) {
				break
			}

			line, err := buf.ReadBytes('\n')
			recordPosition++
			if err == io.EOF && len(line) == 0 {
				break
			}
			if err != nil {
				return m, fmt.Errorf("error reading line: %w", err)
			}

			_, err = data.Write(line)
			if err != nil {
				return m, fmt.Errorf("error writing line to partition buffer: %w", err)
			}
		}

		if data.Len() > 0 {
			err := writePartition(partitionOptions{
				path:        opts.path,
				data:        data,
				partition:   i,
				cardinality: c,
				gzip:        opts.gzip,
			})
			if err != nil {
				return m, fmt.Errorf("error writing partition: %w", err)
			}
		}
	}

	return m, nil
}

func preparePath[T keyedData](opts writeCsvOptions[T]) error {
	err := os.RemoveAll(opts.path)
	if err != nil {
		return fmt.Errorf("error removing directory %s: %w", opts.path, err)
	}

	err = os.MkdirAll(opts.path, 0755)
	if err != nil {
		return fmt.Errorf("error creating directory %s: %w", opts.path, err)
	}

	return nil
}

// func writeHeaders[T keyedData](buf *bytes.Buffer, opts writeCsvOptions[T]) error {
// 	file, err := os.Create(path.Join(opts.path, "header.csv"))
// 	if err != nil {
// 		return fmt.Errorf("error creating header file: %w", err)
// 	}
// 	defer file.Close()
// 	_, err = buf.WriteTo(file)
// 	if err != nil {
// 		return fmt.Errorf("error creating header file: %w", err)
// 	}

// 	return nil
// }

type partitionOptions struct {
	path        string
	data        bytes.Buffer
	partition   int
	cardinality cardinality.Cardinality
	gzip        bool
}

func writePartition(opts partitionOptions) error {
	file, err := openPartitionFile(opts)
	if err != nil {
		return fmt.Errorf("error opening partition file: %w", err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			slog.Error("Error closing partition file", "error", err)
		}
	}()

	if opts.gzip {
		gz := gzip.NewWriter(file)
		defer func() {
			err := gz.Close()
			if err != nil {
				slog.Error("Error closing gzip writer", "error", err)
			}
		}()

		_, err = gz.Write(opts.data.Bytes())
		if err != nil {
			return fmt.Errorf("error writing line to partition file: %w", err)
		}
	} else {
		_, err = file.Write(opts.data.Bytes())
		if err != nil {
			return fmt.Errorf("error writing line to partition file: %w", err)
		}
	}

	return nil
}

func openPartitionFile(opts partitionOptions) (*os.File, error) {
	var gz string
	if opts.gzip {
		gz = ".gz"
	}

	filePath := path.Join(opts.path, fmt.Sprintf("%0*d.csv%s", opts.cardinality.Bits/2, opts.partition, gz))
	file, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening partition file %s: %w", filePath, err)
	}
	return file, nil
}

func writeMeta(pth string, meta map[string]meta) error {
	file, err := os.Create(path.Join(pth, "meta.json"))
	if err != nil {
		return fmt.Errorf("error creating meta file: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			slog.Error("Error closing meta file", "error", err)
		}
	}()

	m, err := json.Marshal(meta)
	if err != nil {
		return fmt.Errorf("error marshaling meta: %w", err)
	}

	_, err = file.Write(m)
	if err != nil {
		return fmt.Errorf("error writing meta file: %w", err)
	}

	return nil
}
