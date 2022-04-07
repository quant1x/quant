//Package Cache comment
// This file war generated by tars2go 1.1
// Generated from cache.tars
package Cache

import (
	"encoding/csv"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
	"github.com/mymmsc/gox/api"
)

//DayKLine struct implement
type DayKLine struct {
	Date   string  `json:"date" csv:"date" array:"0"`
	Open   float64 `json:"open" csv:"open" array:"1"`
	High   float64 `json:"high" csv:"high" array:"2"`
	Low    float64 `json:"low" csv:"low" array:"3"`
	Close  float64 `json:"close" csv:"close" array:"4"`
	Volume int64   `json:"volume" csv:"volume" array:"5"`
}

func (st *DayKLine) resetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *DayKLine) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.resetDefault()

	err = _is.Read_string(&st.Date, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_float64(&st.Open, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_float64(&st.High, 2, true)
	if err != nil {
		return err
	}

	err = _is.Read_float64(&st.Low, 3, true)
	if err != nil {
		return err
	}

	err = _is.Read_float64(&st.Close, 4, true)
	if err != nil {
		return err
	}

	err = _is.Read_int64(&st.Volume, 5, true)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *DayKLine) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require DayKLine, but not exist. tag %d", tag)
		}
		return nil

	}

	st.ReadFrom(_is)

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

func (st *DayKLine) Update(_writer interface{}) error {

	if _w, ok := _writer.(*codec.Buffer); ok {
		return st.WriteTo(_w)
	}
	if _w, ok := _writer.(*csv.Writer); ok {
		return st.WriteCSV(_w)
	}
	return nil
}

func (st *DayKLine) WriteCSV(_writer *csv.Writer) error {
	line := []string{}
	line = append(line, st.Date)
	line = append(line, api.ToString(st.Open))
	line = append(line, api.ToString(st.High))
	line = append(line, api.ToString(st.Low))
	line = append(line, api.ToString(st.Close))
	line = append(line, api.ToString(st.Volume))
	return _writer.Write(line)
}

//WriteTo encode struct to buffer
func (st *DayKLine) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.Date, 0)
	if err != nil {
		return err
	}

	err = _os.Write_float64(st.Open, 1)
	if err != nil {
		return err
	}

	err = _os.Write_float64(st.High, 2)
	if err != nil {
		return err
	}

	err = _os.Write_float64(st.Low, 3)
	if err != nil {
		return err
	}

	err = _os.Write_float64(st.Close, 4)
	if err != nil {
		return err
	}

	err = _os.Write_int64(st.Volume, 5)
	if err != nil {
		return err
	}

	return nil
}

//WriteBlock encode struct
func (st *DayKLine) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	st.WriteTo(_os)

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}
