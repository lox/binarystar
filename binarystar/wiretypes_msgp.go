package binarystar

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"crypto/sha256"

	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *AddChange) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zxvk uint32
	zxvk, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zxvk > 0 {
		zxvk--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "FileInfo":
			err = z.FileInfo.DecodeMsg(dc)
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *AddChange) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "FileInfo"
	err = en.Append(0x81, 0xa8, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f)
	if err != nil {
		return err
	}
	err = z.FileInfo.EncodeMsg(en)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *AddChange) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "FileInfo"
	o = append(o, 0x81, 0xa8, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f)
	o, err = z.FileInfo.MarshalMsg(o)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *AddChange) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zbzg uint32
	zbzg, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zbzg > 0 {
		zbzg--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "FileInfo":
			bts, err = z.FileInfo.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *AddChange) Msgsize() (s int) {
	s = 1 + 9 + z.FileInfo.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Block) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zcmr uint32
	zcmr, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zcmr > 0 {
		zcmr--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "start":
			z.Start, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "end":
			z.End, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "checksum32":
			z.Checksum32, err = dc.ReadUint32()
			if err != nil {
				return
			}
		case "sha256hash":
			err = dc.ReadExactBytes((z.Sha256hash)[:])
			if err != nil {
				return
			}
		case "hasdata":
			z.HasData, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "rawbytes":
			z.RawBytes, err = dc.ReadBytes(z.RawBytes)
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Block) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 6
	// write "start"
	err = en.Append(0x86, 0xa5, 0x73, 0x74, 0x61, 0x72, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Start)
	if err != nil {
		return
	}
	// write "end"
	err = en.Append(0xa3, 0x65, 0x6e, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.End)
	if err != nil {
		return
	}
	// write "checksum32"
	err = en.Append(0xaa, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x33, 0x32)
	if err != nil {
		return err
	}
	err = en.WriteUint32(z.Checksum32)
	if err != nil {
		return
	}
	// write "sha256hash"
	err = en.Append(0xaa, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x68, 0x61, 0x73, 0x68)
	if err != nil {
		return err
	}
	err = en.WriteBytes((z.Sha256hash)[:])
	if err != nil {
		return
	}
	// write "hasdata"
	err = en.Append(0xa7, 0x68, 0x61, 0x73, 0x64, 0x61, 0x74, 0x61)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.HasData)
	if err != nil {
		return
	}
	// write "rawbytes"
	err = en.Append(0xa8, 0x72, 0x61, 0x77, 0x62, 0x79, 0x74, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.RawBytes)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Block) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "start"
	o = append(o, 0x86, 0xa5, 0x73, 0x74, 0x61, 0x72, 0x74)
	o = msgp.AppendInt64(o, z.Start)
	// string "end"
	o = append(o, 0xa3, 0x65, 0x6e, 0x64)
	o = msgp.AppendInt64(o, z.End)
	// string "checksum32"
	o = append(o, 0xaa, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x33, 0x32)
	o = msgp.AppendUint32(o, z.Checksum32)
	// string "sha256hash"
	o = append(o, 0xaa, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x68, 0x61, 0x73, 0x68)
	o = msgp.AppendBytes(o, (z.Sha256hash)[:])
	// string "hasdata"
	o = append(o, 0xa7, 0x68, 0x61, 0x73, 0x64, 0x61, 0x74, 0x61)
	o = msgp.AppendBool(o, z.HasData)
	// string "rawbytes"
	o = append(o, 0xa8, 0x72, 0x61, 0x77, 0x62, 0x79, 0x74, 0x65, 0x73)
	o = msgp.AppendBytes(o, z.RawBytes)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Block) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zajw uint32
	zajw, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zajw > 0 {
		zajw--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "start":
			z.Start, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "end":
			z.End, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "checksum32":
			z.Checksum32, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		case "sha256hash":
			bts, err = msgp.ReadExactBytes(bts, (z.Sha256hash)[:])
			if err != nil {
				return
			}
		case "hasdata":
			z.HasData, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "rawbytes":
			z.RawBytes, bts, err = msgp.ReadBytesBytes(bts, z.RawBytes)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Block) Msgsize() (s int) {
	s = 1 + 6 + msgp.Int64Size + 4 + msgp.Int64Size + 11 + msgp.Uint32Size + 11 + msgp.ArrayHeaderSize + (sha256.Size * (msgp.ByteSize)) + 8 + msgp.BoolSize + 9 + msgp.BytesPrefixSize + len(z.RawBytes)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ChangeSet) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zxhx uint32
	zxhx, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zxhx > 0 {
		zxhx--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "add":
			var zlqf uint32
			zlqf, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Add) >= int(zlqf) {
				z.Add = (z.Add)[:zlqf]
			} else {
				z.Add = make([]AddChange, zlqf)
			}
			for zwht := range z.Add {
				var zdaf uint32
				zdaf, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				for zdaf > 0 {
					zdaf--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "FileInfo":
						err = z.Add[zwht].FileInfo.DecodeMsg(dc)
						if err != nil {
							return
						}
					default:
						err = dc.Skip()
						if err != nil {
							return
						}
					}
				}
			}
		case "delete":
			var zpks uint32
			zpks, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Delete) >= int(zpks) {
				z.Delete = (z.Delete)[:zpks]
			} else {
				z.Delete = make([]DeleteChange, zpks)
			}
			for zhct := range z.Delete {
				var zjfb uint32
				zjfb, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				for zjfb > 0 {
					zjfb--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "FileInfo":
						err = z.Delete[zhct].FileInfo.DecodeMsg(dc)
						if err != nil {
							return
						}
					default:
						err = dc.Skip()
						if err != nil {
							return
						}
					}
				}
			}
		case "modify":
			var zcxo uint32
			zcxo, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Modify) >= int(zcxo) {
				z.Modify = (z.Modify)[:zcxo]
			} else {
				z.Modify = make([]ModifyChange, zcxo)
			}
			for zcua := range z.Modify {
				var zeff uint32
				zeff, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				for zeff > 0 {
					zeff--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "FileInfo":
						err = z.Modify[zcua].FileInfo.DecodeMsg(dc)
						if err != nil {
							return
						}
					default:
						err = dc.Skip()
						if err != nil {
							return
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ChangeSet) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "add"
	err = en.Append(0x83, 0xa3, 0x61, 0x64, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Add)))
	if err != nil {
		return
	}
	for zwht := range z.Add {
		// map header, size 1
		// write "FileInfo"
		err = en.Append(0x81, 0xa8, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f)
		if err != nil {
			return err
		}
		err = z.Add[zwht].FileInfo.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "delete"
	err = en.Append(0xa6, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Delete)))
	if err != nil {
		return
	}
	for zhct := range z.Delete {
		// map header, size 1
		// write "FileInfo"
		err = en.Append(0x81, 0xa8, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f)
		if err != nil {
			return err
		}
		err = z.Delete[zhct].FileInfo.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "modify"
	err = en.Append(0xa6, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Modify)))
	if err != nil {
		return
	}
	for zcua := range z.Modify {
		// map header, size 1
		// write "FileInfo"
		err = en.Append(0x81, 0xa8, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f)
		if err != nil {
			return err
		}
		err = z.Modify[zcua].FileInfo.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ChangeSet) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "add"
	o = append(o, 0x83, 0xa3, 0x61, 0x64, 0x64)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Add)))
	for zwht := range z.Add {
		// map header, size 1
		// string "FileInfo"
		o = append(o, 0x81, 0xa8, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f)
		o, err = z.Add[zwht].FileInfo.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "delete"
	o = append(o, 0xa6, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Delete)))
	for zhct := range z.Delete {
		// map header, size 1
		// string "FileInfo"
		o = append(o, 0x81, 0xa8, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f)
		o, err = z.Delete[zhct].FileInfo.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "modify"
	o = append(o, 0xa6, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Modify)))
	for zcua := range z.Modify {
		// map header, size 1
		// string "FileInfo"
		o = append(o, 0x81, 0xa8, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f)
		o, err = z.Modify[zcua].FileInfo.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ChangeSet) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zrsw uint32
	zrsw, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zrsw > 0 {
		zrsw--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "add":
			var zxpk uint32
			zxpk, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Add) >= int(zxpk) {
				z.Add = (z.Add)[:zxpk]
			} else {
				z.Add = make([]AddChange, zxpk)
			}
			for zwht := range z.Add {
				var zdnj uint32
				zdnj, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zdnj > 0 {
					zdnj--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "FileInfo":
						bts, err = z.Add[zwht].FileInfo.UnmarshalMsg(bts)
						if err != nil {
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		case "delete":
			var zobc uint32
			zobc, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Delete) >= int(zobc) {
				z.Delete = (z.Delete)[:zobc]
			} else {
				z.Delete = make([]DeleteChange, zobc)
			}
			for zhct := range z.Delete {
				var zsnv uint32
				zsnv, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zsnv > 0 {
					zsnv--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "FileInfo":
						bts, err = z.Delete[zhct].FileInfo.UnmarshalMsg(bts)
						if err != nil {
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		case "modify":
			var zkgt uint32
			zkgt, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Modify) >= int(zkgt) {
				z.Modify = (z.Modify)[:zkgt]
			} else {
				z.Modify = make([]ModifyChange, zkgt)
			}
			for zcua := range z.Modify {
				var zema uint32
				zema, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zema > 0 {
					zema--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "FileInfo":
						bts, err = z.Modify[zcua].FileInfo.UnmarshalMsg(bts)
						if err != nil {
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ChangeSet) Msgsize() (s int) {
	s = 1 + 4 + msgp.ArrayHeaderSize
	for zwht := range z.Add {
		s += 1 + 9 + z.Add[zwht].FileInfo.Msgsize()
	}
	s += 7 + msgp.ArrayHeaderSize
	for zhct := range z.Delete {
		s += 1 + 9 + z.Delete[zhct].FileInfo.Msgsize()
	}
	s += 7 + msgp.ArrayHeaderSize
	for zcua := range z.Modify {
		s += 1 + 9 + z.Modify[zcua].FileInfo.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *DeleteChange) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zpez uint32
	zpez, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zpez > 0 {
		zpez--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "FileInfo":
			err = z.FileInfo.DecodeMsg(dc)
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *DeleteChange) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "FileInfo"
	err = en.Append(0x81, 0xa8, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f)
	if err != nil {
		return err
	}
	err = z.FileInfo.EncodeMsg(en)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *DeleteChange) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "FileInfo"
	o = append(o, 0x81, 0xa8, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f)
	o, err = z.FileInfo.MarshalMsg(o)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DeleteChange) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zqke uint32
	zqke, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zqke > 0 {
		zqke--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "FileInfo":
			bts, err = z.FileInfo.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *DeleteChange) Msgsize() (s int) {
	s = 1 + 9 + z.FileInfo.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *FileInfo) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zqyh uint32
	zqyh, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zqyh > 0 {
		zqyh--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "path":
			z.Path, err = dc.ReadString()
			if err != nil {
				return
			}
		case "size":
			z.Size, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "mode":
			z.Mode, err = dc.ReadUint32()
			if err != nil {
				return
			}
		case "isdeleted":
			z.IsDeleted, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "modtime":
			z.ModTime, err = dc.ReadTime()
			if err != nil {
				return
			}
		case "fingerprint":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Fingerprint = nil
			} else {
				if z.Fingerprint == nil {
					z.Fingerprint = new(Fingerprint)
				}
				err = z.Fingerprint.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *FileInfo) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 6
	// write "path"
	err = en.Append(0x86, 0xa4, 0x70, 0x61, 0x74, 0x68)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Path)
	if err != nil {
		return
	}
	// write "size"
	err = en.Append(0xa4, 0x73, 0x69, 0x7a, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Size)
	if err != nil {
		return
	}
	// write "mode"
	err = en.Append(0xa4, 0x6d, 0x6f, 0x64, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteUint32(z.Mode)
	if err != nil {
		return
	}
	// write "isdeleted"
	err = en.Append(0xa9, 0x69, 0x73, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.IsDeleted)
	if err != nil {
		return
	}
	// write "modtime"
	err = en.Append(0xa7, 0x6d, 0x6f, 0x64, 0x74, 0x69, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteTime(z.ModTime)
	if err != nil {
		return
	}
	// write "fingerprint"
	err = en.Append(0xab, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74)
	if err != nil {
		return err
	}
	if z.Fingerprint == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Fingerprint.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *FileInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "path"
	o = append(o, 0x86, 0xa4, 0x70, 0x61, 0x74, 0x68)
	o = msgp.AppendString(o, z.Path)
	// string "size"
	o = append(o, 0xa4, 0x73, 0x69, 0x7a, 0x65)
	o = msgp.AppendInt64(o, z.Size)
	// string "mode"
	o = append(o, 0xa4, 0x6d, 0x6f, 0x64, 0x65)
	o = msgp.AppendUint32(o, z.Mode)
	// string "isdeleted"
	o = append(o, 0xa9, 0x69, 0x73, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64)
	o = msgp.AppendBool(o, z.IsDeleted)
	// string "modtime"
	o = append(o, 0xa7, 0x6d, 0x6f, 0x64, 0x74, 0x69, 0x6d, 0x65)
	o = msgp.AppendTime(o, z.ModTime)
	// string "fingerprint"
	o = append(o, 0xab, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74)
	if z.Fingerprint == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Fingerprint.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *FileInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zyzr uint32
	zyzr, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zyzr > 0 {
		zyzr--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "path":
			z.Path, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "size":
			z.Size, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "mode":
			z.Mode, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		case "isdeleted":
			z.IsDeleted, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "modtime":
			z.ModTime, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				return
			}
		case "fingerprint":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Fingerprint = nil
			} else {
				if z.Fingerprint == nil {
					z.Fingerprint = new(Fingerprint)
				}
				bts, err = z.Fingerprint.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FileInfo) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Path) + 5 + msgp.Int64Size + 5 + msgp.Uint32Size + 10 + msgp.BoolSize + 8 + msgp.TimeSize + 12
	if z.Fingerprint == nil {
		s += msgp.NilSize
	} else {
		s += z.Fingerprint.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *FileStreamHeaderMessage) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zywj uint32
	zywj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zywj > 0 {
		zywj--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "FileInfo":
			err = z.FileInfo.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "error":
			z.Error, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *FileStreamHeaderMessage) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "FileInfo"
	err = en.Append(0x82, 0xa8, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f)
	if err != nil {
		return err
	}
	err = z.FileInfo.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "error"
	err = en.Append(0xa5, 0x65, 0x72, 0x72, 0x6f, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Error)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *FileStreamHeaderMessage) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "FileInfo"
	o = append(o, 0x82, 0xa8, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f)
	o, err = z.FileInfo.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "error"
	o = append(o, 0xa5, 0x65, 0x72, 0x72, 0x6f, 0x72)
	o = msgp.AppendString(o, z.Error)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *FileStreamHeaderMessage) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zjpj uint32
	zjpj, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zjpj > 0 {
		zjpj--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "FileInfo":
			bts, err = z.FileInfo.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "error":
			z.Error, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FileStreamHeaderMessage) Msgsize() (s int) {
	s = 1 + 9 + z.FileInfo.Msgsize() + 6 + msgp.StringPrefixSize + len(z.Error)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *FileStreamRequestsMessage) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zrfe uint32
	zrfe, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zrfe > 0 {
		zrfe--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "path":
			var zgmo uint32
			zgmo, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Paths) >= int(zgmo) {
				z.Paths = (z.Paths)[:zgmo]
			} else {
				z.Paths = make([]string, zgmo)
			}
			for zzpf := range z.Paths {
				z.Paths[zzpf], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *FileStreamRequestsMessage) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "path"
	err = en.Append(0x81, 0xa4, 0x70, 0x61, 0x74, 0x68)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Paths)))
	if err != nil {
		return
	}
	for zzpf := range z.Paths {
		err = en.WriteString(z.Paths[zzpf])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *FileStreamRequestsMessage) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "path"
	o = append(o, 0x81, 0xa4, 0x70, 0x61, 0x74, 0x68)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Paths)))
	for zzpf := range z.Paths {
		o = msgp.AppendString(o, z.Paths[zzpf])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *FileStreamRequestsMessage) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var ztaf uint32
	ztaf, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for ztaf > 0 {
		ztaf--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "path":
			var zeth uint32
			zeth, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Paths) >= int(zeth) {
				z.Paths = (z.Paths)[:zeth]
			} else {
				z.Paths = make([]string, zeth)
			}
			for zzpf := range z.Paths {
				z.Paths[zzpf], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FileStreamRequestsMessage) Msgsize() (s int) {
	s = 1 + 5 + msgp.ArrayHeaderSize
	for zzpf := range z.Paths {
		s += msgp.StringPrefixSize + len(z.Paths[zzpf])
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Fingerprint) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zrbe uint32
	zrbe, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zrbe > 0 {
		zrbe--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "blocksz":
			z.Blocksz, err = dc.ReadUint32()
			if err != nil {
				return
			}
		case "blockmap":
			var zmfd uint32
			zmfd, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.BlockMap == nil && zmfd > 0 {
				z.BlockMap = make(map[string]map[string]Block, zmfd)
			} else if len(z.BlockMap) > 0 {
				for key, _ := range z.BlockMap {
					delete(z.BlockMap, key)
				}
			}
			for zmfd > 0 {
				zmfd--
				var zsbz string
				var zrjx map[string]Block
				zsbz, err = dc.ReadString()
				if err != nil {
					return
				}
				var zzdc uint32
				zzdc, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				if zrjx == nil && zzdc > 0 {
					zrjx = make(map[string]Block, zzdc)
				} else if len(zrjx) > 0 {
					for key, _ := range zrjx {
						delete(zrjx, key)
					}
				}
				for zzdc > 0 {
					zzdc--
					var zawn string
					var zwel Block
					zawn, err = dc.ReadString()
					if err != nil {
						return
					}
					err = zwel.DecodeMsg(dc)
					if err != nil {
						return
					}
					zrjx[zawn] = zwel
				}
				z.BlockMap[zsbz] = zrjx
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Fingerprint) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "blocksz"
	err = en.Append(0x82, 0xa7, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x7a)
	if err != nil {
		return err
	}
	err = en.WriteUint32(z.Blocksz)
	if err != nil {
		return
	}
	// write "blockmap"
	err = en.Append(0xa8, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6d, 0x61, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.BlockMap)))
	if err != nil {
		return
	}
	for zsbz, zrjx := range z.BlockMap {
		err = en.WriteString(zsbz)
		if err != nil {
			return
		}
		err = en.WriteMapHeader(uint32(len(zrjx)))
		if err != nil {
			return
		}
		for zawn, zwel := range zrjx {
			err = en.WriteString(zawn)
			if err != nil {
				return
			}
			err = zwel.EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Fingerprint) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "blocksz"
	o = append(o, 0x82, 0xa7, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x7a)
	o = msgp.AppendUint32(o, z.Blocksz)
	// string "blockmap"
	o = append(o, 0xa8, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6d, 0x61, 0x70)
	o = msgp.AppendMapHeader(o, uint32(len(z.BlockMap)))
	for zsbz, zrjx := range z.BlockMap {
		o = msgp.AppendString(o, zsbz)
		o = msgp.AppendMapHeader(o, uint32(len(zrjx)))
		for zawn, zwel := range zrjx {
			o = msgp.AppendString(o, zawn)
			o, err = zwel.MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Fingerprint) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zelx uint32
	zelx, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zelx > 0 {
		zelx--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "blocksz":
			z.Blocksz, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		case "blockmap":
			var zbal uint32
			zbal, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.BlockMap == nil && zbal > 0 {
				z.BlockMap = make(map[string]map[string]Block, zbal)
			} else if len(z.BlockMap) > 0 {
				for key, _ := range z.BlockMap {
					delete(z.BlockMap, key)
				}
			}
			for zbal > 0 {
				var zsbz string
				var zrjx map[string]Block
				zbal--
				zsbz, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				var zjqz uint32
				zjqz, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if zrjx == nil && zjqz > 0 {
					zrjx = make(map[string]Block, zjqz)
				} else if len(zrjx) > 0 {
					for key, _ := range zrjx {
						delete(zrjx, key)
					}
				}
				for zjqz > 0 {
					var zawn string
					var zwel Block
					zjqz--
					zawn, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
					bts, err = zwel.UnmarshalMsg(bts)
					if err != nil {
						return
					}
					zrjx[zawn] = zwel
				}
				z.BlockMap[zsbz] = zrjx
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Fingerprint) Msgsize() (s int) {
	s = 1 + 8 + msgp.Uint32Size + 9 + msgp.MapHeaderSize
	if z.BlockMap != nil {
		for zsbz, zrjx := range z.BlockMap {
			_ = zrjx
			s += msgp.StringPrefixSize + len(zsbz) + msgp.MapHeaderSize
			if zrjx != nil {
				for zawn, zwel := range zrjx {
					_ = zwel
					s += msgp.StringPrefixSize + len(zawn) + zwel.Msgsize()
				}
			}
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ModifyChange) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zkct uint32
	zkct, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zkct > 0 {
		zkct--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "FileInfo":
			err = z.FileInfo.DecodeMsg(dc)
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ModifyChange) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "FileInfo"
	err = en.Append(0x81, 0xa8, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f)
	if err != nil {
		return err
	}
	err = z.FileInfo.EncodeMsg(en)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ModifyChange) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "FileInfo"
	o = append(o, 0x81, 0xa8, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f)
	o, err = z.FileInfo.MarshalMsg(o)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ModifyChange) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var ztmt uint32
	ztmt, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for ztmt > 0 {
		ztmt--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "FileInfo":
			bts, err = z.FileInfo.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ModifyChange) Msgsize() (s int) {
	s = 1 + 9 + z.FileInfo.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *SyncMessage) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zana uint32
	zana, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zana > 0 {
		zana--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "files":
			var ztyy uint32
			ztyy, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Files) >= int(ztyy) {
				z.Files = (z.Files)[:ztyy]
			} else {
				z.Files = make([]FileInfo, ztyy)
			}
			for ztco := range z.Files {
				err = z.Files[ztco].DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *SyncMessage) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "files"
	err = en.Append(0x81, 0xa5, 0x66, 0x69, 0x6c, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Files)))
	if err != nil {
		return
	}
	for ztco := range z.Files {
		err = z.Files[ztco].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SyncMessage) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "files"
	o = append(o, 0x81, 0xa5, 0x66, 0x69, 0x6c, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Files)))
	for ztco := range z.Files {
		o, err = z.Files[ztco].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SyncMessage) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zinl uint32
	zinl, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zinl > 0 {
		zinl--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "files":
			var zare uint32
			zare, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Files) >= int(zare) {
				z.Files = (z.Files)[:zare]
			} else {
				z.Files = make([]FileInfo, zare)
			}
			for ztco := range z.Files {
				bts, err = z.Files[ztco].UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SyncMessage) Msgsize() (s int) {
	s = 1 + 6 + msgp.ArrayHeaderSize
	for ztco := range z.Files {
		s += z.Files[ztco].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *SyncResponseMessage) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zljy uint32
	zljy, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zljy > 0 {
		zljy--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "changes":
			err = z.Changes.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "error":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Error = nil
			} else {
				if z.Error == nil {
					z.Error = new(string)
				}
				*z.Error, err = dc.ReadString()
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *SyncResponseMessage) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "changes"
	err = en.Append(0x82, 0xa7, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = z.Changes.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "error"
	err = en.Append(0xa5, 0x65, 0x72, 0x72, 0x6f, 0x72)
	if err != nil {
		return err
	}
	if z.Error == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteString(*z.Error)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SyncResponseMessage) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "changes"
	o = append(o, 0x82, 0xa7, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73)
	o, err = z.Changes.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "error"
	o = append(o, 0xa5, 0x65, 0x72, 0x72, 0x6f, 0x72)
	if z.Error == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendString(o, *z.Error)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SyncResponseMessage) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zixj uint32
	zixj, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zixj > 0 {
		zixj--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "changes":
			bts, err = z.Changes.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "error":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Error = nil
			} else {
				if z.Error == nil {
					z.Error = new(string)
				}
				*z.Error, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SyncResponseMessage) Msgsize() (s int) {
	s = 1 + 8 + z.Changes.Msgsize() + 6
	if z.Error == nil {
		s += msgp.NilSize
	} else {
		s += msgp.StringPrefixSize + len(*z.Error)
	}
	return
}
