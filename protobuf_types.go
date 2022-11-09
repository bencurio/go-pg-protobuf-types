package go_pg_protobuf_types

import (
	"fmt"
	"github.com/go-pg/pg/v10/types"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"reflect"
)

func DoubleValueAppender(in []byte, v reflect.Value, flags int) []byte {
	wpb := v.Interface().(*wrapperspb.DoubleValue)
	return types.Append(in, wpb.Value, flags)
}

func DoubleValueScanner(v reflect.Value, rd types.Reader, n int) error {
	if !v.CanSet() {
		return fmt.Errorf("pg: Scan(nonsettable %s)", v.Type())
	}
	scan, err := types.ScanFloat64(rd, n)
	if err != nil {
		return err
	}
	wpb := wrapperspb.Double(scan)
	v.Set(reflect.ValueOf(wpb))
	return nil
}

func FloatValueAppender(in []byte, v reflect.Value, flags int) []byte {
	wpb := v.Interface().(*wrapperspb.FloatValue)
	return types.Append(in, wpb.Value, flags)
}

func FloatValueScanner(v reflect.Value, rd types.Reader, n int) error {
	if !v.CanSet() {
		return fmt.Errorf("pg: Scan(nonsettable %s)", v.Type())
	}
	scan, err := types.ScanFloat32(rd, n)
	if err != nil {
		return err
	}
	wpb := wrapperspb.Float(scan)
	v.Set(reflect.ValueOf(wpb))
	return nil
}

func Int64ValueAppender(in []byte, v reflect.Value, flags int) []byte {
	wpb := v.Interface().(*wrapperspb.Int64Value)
	return types.Append(in, wpb.Value, flags)
}

func Int64ValueScanner(v reflect.Value, rd types.Reader, n int) error {
	if !v.CanSet() {
		return fmt.Errorf("pg: Scan(nonsettable %s)", v.Type())
	}
	scan, err := types.ScanInt64(rd, n)
	if err != nil {
		return err
	}
	wpb := wrapperspb.Int64(scan)
	v.Set(reflect.ValueOf(wpb))
	return nil
}

func UInt64ValueAppender(in []byte, v reflect.Value, flags int) []byte {
	wpb := v.Interface().(*wrapperspb.UInt64Value)
	return types.Append(in, wpb.Value, flags)
}

func UInt64ValueScanner(v reflect.Value, rd types.Reader, n int) error {
	if !v.CanSet() {
		return fmt.Errorf("pg: Scan(nonsettable %s)", v.Type())
	}
	scan, err := types.ScanUint64(rd, n)
	if err != nil {
		return err
	}
	wpb := wrapperspb.UInt64(scan)
	v.Set(reflect.ValueOf(wpb))
	return nil
}

func Int32ValueAppender(in []byte, v reflect.Value, flags int) []byte {
	wpb := v.Interface().(*wrapperspb.Int32Value)
	return types.Append(in, wpb.Value, flags)
}

func Int32ValueScanner(v reflect.Value, rd types.Reader, n int) error {
	if !v.CanSet() {
		return fmt.Errorf("pg: Scan(nonsettable %s)", v.Type())
	}
	scan, err := types.ScanInt64(rd, n)
	if err != nil {
		return err
	}
	wpb := wrapperspb.Int32(int32(scan))
	v.Set(reflect.ValueOf(wpb))
	return nil
}

func UInt32ValueAppender(in []byte, v reflect.Value, flags int) []byte {
	wpb := v.Interface().(*wrapperspb.UInt32Value)
	return types.Append(in, wpb.Value, flags)
}

func UInt32ValueScanner(v reflect.Value, rd types.Reader, n int) error {
	if !v.CanSet() {
		return fmt.Errorf("pg: Scan(nonsettable %s)", v.Type())
	}
	scan, err := types.ScanUint64(rd, n)
	if err != nil {
		return err
	}
	wpb := wrapperspb.UInt32(uint32(scan))
	v.Set(reflect.ValueOf(wpb))
	return nil
}

func BoolValueAppender(in []byte, v reflect.Value, flags int) []byte {
	wpb := v.Interface().(*wrapperspb.BoolValue)
	return types.Append(in, wpb.Value, flags)
}

func BoolValueScanner(v reflect.Value, rd types.Reader, n int) error {
	if !v.CanSet() {
		return fmt.Errorf("pg: Scan(nonsettable %s)", v.Type())
	}
	scan, err := types.ScanBool(rd, n)
	if err != nil {
		return err
	}
	wpb := wrapperspb.Bool(scan)
	v.Set(reflect.ValueOf(wpb))
	return nil
}

func StringValueAppender(in []byte, v reflect.Value, flags int) []byte {
	wpb := v.Interface().(*wrapperspb.StringValue)
	return types.AppendString(in, wpb.Value, flags)
}

func StringValueScanner(v reflect.Value, rd types.Reader, n int) error {
	if !v.CanSet() {
		return fmt.Errorf("pg: Scan(nonsettable %s)", v.Type())
	}
	scan, err := types.ScanString(rd, n)
	if err != nil {
		return err
	}
	wpb := wrapperspb.String(scan)
	v.Set(reflect.ValueOf(wpb))
	return nil
}

func BytesValueAppender(in []byte, v reflect.Value, flags int) []byte {
	wpb := v.Interface().(*wrapperspb.BytesValue)
	return types.AppendBytes(in, wpb.Value, flags)
}

func BytesValueScanner(v reflect.Value, rd types.Reader, n int) error {
	if !v.CanSet() {
		return fmt.Errorf("pg: Scan(nonsettable %s)", v.Type())
	}
	scan, err := types.ScanBytes(rd, n)
	if err != nil {
		return err
	}
	wpb := wrapperspb.Bytes(scan)
	v.Set(reflect.ValueOf(wpb))
	return nil
}

func TimestampAppender(in []byte, v reflect.Value, flags int) []byte {
	pbts := v.Interface().(*timestamppb.Timestamp)
	if err := pbts.CheckValid(); err != nil {
		panic(err)
	}
	return types.AppendTime(in, pbts.AsTime(), flags)
}

func TimestampScanner(v reflect.Value, rd types.Reader, n int) error {
	if !v.CanSet() {
		return fmt.Errorf("pg: Scan(nonsettable %s)", v.Type())
	}
	scan, err := types.ScanTime(rd, n)
	if err != nil {
		return err
	}
	ts := timestamppb.New(scan)
	if err := ts.CheckValid(); err != nil {
		return err
	}
	v.Set(reflect.ValueOf(ts))
	return nil
}

func RegisterAppenderAndScanner() {
	var ex *timestamppb.Timestamp
	types.RegisterAppender(ex, TimestampAppender)
	types.RegisterScanner(ex, TimestampScanner)

	var dv *wrapperspb.DoubleValue
	types.RegisterAppender(dv, DoubleValueAppender)
	types.RegisterScanner(dv, DoubleValueScanner)

	var fv *wrapperspb.FloatValue
	types.RegisterAppender(fv, FloatValueAppender)
	types.RegisterScanner(fv, FloatValueScanner)

	var i64v *wrapperspb.Int64Value
	types.RegisterAppender(i64v, Int64ValueAppender)
	types.RegisterScanner(i64v, Int64ValueScanner)

	var ui64v *wrapperspb.UInt64Value
	types.RegisterAppender(ui64v, UInt64ValueAppender)
	types.RegisterScanner(ui64v, UInt64ValueScanner)

	var i32v *wrapperspb.Int32Value
	types.RegisterAppender(i32v, Int32ValueAppender)
	types.RegisterScanner(i32v, Int32ValueScanner)

	var ui32v *wrapperspb.UInt32Value
	types.RegisterAppender(ui32v, UInt32ValueAppender)
	types.RegisterScanner(ui32v, UInt32ValueScanner)

	var bov *wrapperspb.BoolValue
	types.RegisterAppender(bov, BoolValueAppender)
	types.RegisterScanner(bov, BoolValueScanner)

	var sv *wrapperspb.StringValue
	types.RegisterAppender(sv, StringValueAppender)
	types.RegisterScanner(sv, StringValueScanner)

	var bv *wrapperspb.BytesValue
	types.RegisterAppender(bv, BytesValueAppender)
	types.RegisterScanner(bv, BytesValueScanner)
}
