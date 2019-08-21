package packit

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"strconv"
)

const (
	intT     byte = 0
	float64T byte = 1
	bytesT   byte = 2
	stringT  byte = 3
)

func Pack(values ...interface{}) ([]byte, error) {
	w := new(bytes.Buffer)
	for _, value := range values {
		switch v := value.(type) {
		case int:
			if err := packInt(w, v); err != nil {
				return nil, fmt.Errorf("failed to pack int: %s", err)
			}
		case float64:
			if err := packFloat64(w, v); err != nil {
				return nil, fmt.Errorf("failed to pack float64: %s", err)
			}
		case []byte:
			if err := packBytes(w, v); err != nil {
				return nil, fmt.Errorf("failed to pack bytes: %s", err)
			}
		case string:
			if err := packString(w, v); err != nil {
				return nil, fmt.Errorf("failed to pack string: %s", err)
			}
		default:
			return nil, fmt.Errorf("don't know how to pack %T", v)
		}
	}
	return w.Bytes(), nil
}

func Unpack(p []byte) ([]interface{}, error) {
	values := []interface{}{}

	// create a reader
	r := bytes.NewBuffer(p)
	for {
		b, err := r.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		switch b {
		case intT:
			i, err := unpackInt(r)
			if err != nil {
				return nil, err
			}
			values = append(values, i)
		case float64T:
			f, err := unpackFloat64(r)
			if err != nil {
				return nil, err
			}
			values = append(values, f)
		case bytesT:
			b, err := unpackBytes(r)
			if err != nil {
				return nil, err
			}
			values = append(values, b)
		case stringT:
			s, err := unpackString(r)
			if err != nil {
				return nil, err
			}
			values = append(values, s)
		default:
			return nil, fmt.Errorf("unknown type %d", b)
		}
	}
	return values, nil
}

func packInt(w io.Writer, i int) error {
	// write the type
	if _, err := w.Write([]byte{intT}); err != nil {
		return fmt.Errorf("failed to write type: %s", err)
	}
	// encode int to binary
	buf := []byte(strconv.Itoa(i))

	// write the length
	if err := binary.Write(w, binary.BigEndian, int64(len(buf))); err != nil {
		return fmt.Errorf("failed to write size: %s", err)
	}
	if _, err := w.Write(buf); err != nil {
		return fmt.Errorf("failed to write value: %s", err)
	}
	return nil
}

func unpackInt(r io.Reader) (int, error) {
	// Read the size of the message.
	var sz int64
	if err := binary.Read(r, binary.BigEndian, &sz); err != nil {
		return 0, fmt.Errorf("read message size: %s", err)
	}

	// Read the value.
	buf := make([]byte, sz)
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, fmt.Errorf("read message value: %s", err)
	}
	i, err := strconv.Atoi(string(buf))
	if err != nil {
		return 0, fmt.Errorf("unable to convert %q to int.", string(buf))
	}
	return i, nil
}

func packString(w io.Writer, s string) error {
	// write the type
	if _, err := w.Write([]byte{stringT}); err != nil {
		return fmt.Errorf("failed to write type: %s", err)
	}
	// encode int to binary
	buf := []byte(s)

	// write the length
	if err := binary.Write(w, binary.BigEndian, int64(len(buf))); err != nil {
		return fmt.Errorf("failed to write size: %s", err)
	}
	if _, err := w.Write(buf); err != nil {
		return fmt.Errorf("failed to write value: %s", err)
	}
	return nil
}

func unpackString(r io.Reader) (string, error) {
	// Read the size of the message.
	var sz int64
	if err := binary.Read(r, binary.BigEndian, &sz); err != nil {
		return "", fmt.Errorf("read message size: %s", err)
	}

	// Read the value.
	buf := make([]byte, sz)
	if _, err := io.ReadFull(r, buf); err != nil {
		return "", fmt.Errorf("read message value: %s", err)
	}
	return string(buf), nil
}

func packFloat64(w io.Writer, f float64) error {
	// write the type
	if _, err := w.Write([]byte{float64T}); err != nil {
		return fmt.Errorf("failed to write type: %s", err)
	}
	// encode int to binary
	buf := []byte(strconv.FormatFloat(f, 'E', -1, 64))

	// write the length
	if err := binary.Write(w, binary.BigEndian, int64(len(buf))); err != nil {
		return fmt.Errorf("failed to write size: %s", err)
	}
	if _, err := w.Write(buf); err != nil {
		return fmt.Errorf("failed to write value: %s", err)
	}
	return nil
}

func unpackFloat64(r io.Reader) (float64, error) {
	// Read the size of the message.
	var sz int64
	if err := binary.Read(r, binary.BigEndian, &sz); err != nil {
		return 0, fmt.Errorf("read message size: %s", err)
	}

	// Read the value.
	buf := make([]byte, sz)
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, fmt.Errorf("read message value: %s", err)
	}
	f, err := strconv.ParseFloat(string(buf), 64)
	if err != nil {
		return 0, fmt.Errorf("unable to convert %q to int.", string(buf))
	}
	return f, nil
}

func packBytes(w io.Writer, buf []byte) error {
	// write the type
	if _, err := w.Write([]byte{bytesT}); err != nil {
		return fmt.Errorf("failed to write type: %s", err)
	}
	// write the length
	if err := binary.Write(w, binary.BigEndian, int64(len(buf))); err != nil {
		return fmt.Errorf("failed to write size: %s", err)
	}
	if _, err := w.Write(buf); err != nil {
		return fmt.Errorf("failed to write value: %s", err)
	}
	return nil
}

func unpackBytes(r io.Reader) ([]byte, error) {
	// Read the size of the message.
	var sz int64
	if err := binary.Read(r, binary.BigEndian, &sz); err != nil {
		return nil, fmt.Errorf("read message size: %s", err)
	}

	// Read the value.
	buf := make([]byte, sz)
	if _, err := io.ReadFull(r, buf); err != nil {
		return nil, fmt.Errorf("read message value: %s", err)
	}
	return buf, nil
}
