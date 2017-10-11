package binarystar

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
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
		case "Start":
			z.Start, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "End":
			z.End, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Checksum32":
			z.Checksum32, err = dc.ReadUint32()
			if err != nil {
				return
			}
		case "Sha256hash":
			err = dc.ReadExactBytes((z.Sha256hash)[:])
			if err != nil {
				return
			}
		case "HasData":
			z.HasData, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "RawBytes":
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
	// write "Start"
	err = en.Append(0x86, 0xa5, 0x53, 0x74, 0x61, 0x72, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Start)
	if err != nil {
		return
	}
	// write "End"
	err = en.Append(0xa3, 0x45, 0x6e, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.End)
	if err != nil {
		return
	}
	// write "Checksum32"
	err = en.Append(0xaa, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x33, 0x32)
	if err != nil {
		return err
	}
	err = en.WriteUint32(z.Checksum32)
	if err != nil {
		return
	}
	// write "Sha256hash"
	err = en.Append(0xaa, 0x53, 0x68, 0x61, 0x32, 0x35, 0x36, 0x68, 0x61, 0x73, 0x68)
	if err != nil {
		return err
	}
	err = en.WriteBytes((z.Sha256hash)[:])
	if err != nil {
		return
	}
	// write "HasData"
	err = en.Append(0xa7, 0x48, 0x61, 0x73, 0x44, 0x61, 0x74, 0x61)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.HasData)
	if err != nil {
		return
	}
	// write "RawBytes"
	err = en.Append(0xa8, 0x52, 0x61, 0x77, 0x42, 0x79, 0x74, 0x65, 0x73)
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
	// string "Start"
	o = append(o, 0x86, 0xa5, 0x53, 0x74, 0x61, 0x72, 0x74)
	o = msgp.AppendInt64(o, z.Start)
	// string "End"
	o = append(o, 0xa3, 0x45, 0x6e, 0x64)
	o = msgp.AppendInt64(o, z.End)
	// string "Checksum32"
	o = append(o, 0xaa, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x33, 0x32)
	o = msgp.AppendUint32(o, z.Checksum32)
	// string "Sha256hash"
	o = append(o, 0xaa, 0x53, 0x68, 0x61, 0x32, 0x35, 0x36, 0x68, 0x61, 0x73, 0x68)
	o = msgp.AppendBytes(o, (z.Sha256hash)[:])
	// string "HasData"
	o = append(o, 0xa7, 0x48, 0x61, 0x73, 0x44, 0x61, 0x74, 0x61)
	o = msgp.AppendBool(o, z.HasData)
	// string "RawBytes"
	o = append(o, 0xa8, 0x52, 0x61, 0x77, 0x42, 0x79, 0x74, 0x65, 0x73)
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
		case "Start":
			z.Start, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "End":
			z.End, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "Checksum32":
			z.Checksum32, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		case "Sha256hash":
			bts, err = msgp.ReadExactBytes(bts, (z.Sha256hash)[:])
			if err != nil {
				return
			}
		case "HasData":
			z.HasData, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "RawBytes":
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
	s = 1 + 6 + msgp.Int64Size + 4 + msgp.Int64Size + 11 + msgp.Uint32Size + 11 + msgp.ArrayHeaderSize + (32 * (msgp.ByteSize)) + 8 + msgp.BoolSize + 9 + msgp.BytesPrefixSize + len(z.RawBytes)
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
		case "modify":
			var zpks uint32
			zpks, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Modify) >= int(zpks) {
				z.Modify = (z.Modify)[:zpks]
			} else {
				z.Modify = make([]ModifyChange, zpks)
			}
			for zhct := range z.Modify {
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
					case "from":
						err = z.Modify[zhct].From.DecodeMsg(dc)
						if err != nil {
							return
						}
					case "to":
						err = z.Modify[zhct].To.DecodeMsg(dc)
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
			var zcxo uint32
			zcxo, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Delete) >= int(zcxo) {
				z.Delete = (z.Delete)[:zcxo]
			} else {
				z.Delete = make([]DeleteChange, zcxo)
			}
			for zcua := range z.Delete {
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
					case "from":
						err = z.Delete[zcua].From.DecodeMsg(dc)
						if err != nil {
							return
						}
					case "to":
						err = z.Delete[zcua].To.DecodeMsg(dc)
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
	// write "modify"
	err = en.Append(0xa6, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Modify)))
	if err != nil {
		return
	}
	for zhct := range z.Modify {
		// map header, size 2
		// write "from"
		err = en.Append(0x82, 0xa4, 0x66, 0x72, 0x6f, 0x6d)
		if err != nil {
			return err
		}
		err = z.Modify[zhct].From.EncodeMsg(en)
		if err != nil {
			return
		}
		// write "to"
		err = en.Append(0xa2, 0x74, 0x6f)
		if err != nil {
			return err
		}
		err = z.Modify[zhct].To.EncodeMsg(en)
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
	for zcua := range z.Delete {
		// map header, size 2
		// write "from"
		err = en.Append(0x82, 0xa4, 0x66, 0x72, 0x6f, 0x6d)
		if err != nil {
			return err
		}
		err = z.Delete[zcua].From.EncodeMsg(en)
		if err != nil {
			return
		}
		// write "to"
		err = en.Append(0xa2, 0x74, 0x6f)
		if err != nil {
			return err
		}
		err = z.Delete[zcua].To.EncodeMsg(en)
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
	// string "modify"
	o = append(o, 0xa6, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Modify)))
	for zhct := range z.Modify {
		// map header, size 2
		// string "from"
		o = append(o, 0x82, 0xa4, 0x66, 0x72, 0x6f, 0x6d)
		o, err = z.Modify[zhct].From.MarshalMsg(o)
		if err != nil {
			return
		}
		// string "to"
		o = append(o, 0xa2, 0x74, 0x6f)
		o, err = z.Modify[zhct].To.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "delete"
	o = append(o, 0xa6, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Delete)))
	for zcua := range z.Delete {
		// map header, size 2
		// string "from"
		o = append(o, 0x82, 0xa4, 0x66, 0x72, 0x6f, 0x6d)
		o, err = z.Delete[zcua].From.MarshalMsg(o)
		if err != nil {
			return
		}
		// string "to"
		o = append(o, 0xa2, 0x74, 0x6f)
		o, err = z.Delete[zcua].To.MarshalMsg(o)
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
		case "modify":
			var zobc uint32
			zobc, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Modify) >= int(zobc) {
				z.Modify = (z.Modify)[:zobc]
			} else {
				z.Modify = make([]ModifyChange, zobc)
			}
			for zhct := range z.Modify {
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
					case "from":
						bts, err = z.Modify[zhct].From.UnmarshalMsg(bts)
						if err != nil {
							return
						}
					case "to":
						bts, err = z.Modify[zhct].To.UnmarshalMsg(bts)
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
			var zkgt uint32
			zkgt, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Delete) >= int(zkgt) {
				z.Delete = (z.Delete)[:zkgt]
			} else {
				z.Delete = make([]DeleteChange, zkgt)
			}
			for zcua := range z.Delete {
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
					case "from":
						bts, err = z.Delete[zcua].From.UnmarshalMsg(bts)
						if err != nil {
							return
						}
					case "to":
						bts, err = z.Delete[zcua].To.UnmarshalMsg(bts)
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
	for zhct := range z.Modify {
		s += 1 + 5 + z.Modify[zhct].From.Msgsize() + 3 + z.Modify[zhct].To.Msgsize()
	}
	s += 7 + msgp.ArrayHeaderSize
	for zcua := range z.Delete {
		s += 1 + 5 + z.Delete[zcua].From.Msgsize() + 3 + z.Delete[zcua].To.Msgsize()
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
		case "from":
			err = z.From.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "to":
			err = z.To.DecodeMsg(dc)
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
	// map header, size 2
	// write "from"
	err = en.Append(0x82, 0xa4, 0x66, 0x72, 0x6f, 0x6d)
	if err != nil {
		return err
	}
	err = z.From.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "to"
	err = en.Append(0xa2, 0x74, 0x6f)
	if err != nil {
		return err
	}
	err = z.To.EncodeMsg(en)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *DeleteChange) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "from"
	o = append(o, 0x82, 0xa4, 0x66, 0x72, 0x6f, 0x6d)
	o, err = z.From.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "to"
	o = append(o, 0xa2, 0x74, 0x6f)
	o, err = z.To.MarshalMsg(o)
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
		case "from":
			bts, err = z.From.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "to":
			bts, err = z.To.UnmarshalMsg(bts)
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
	s = 1 + 5 + z.From.Msgsize() + 3 + z.To.Msgsize()
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
func (z *Fingerprint) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zgmo uint32
	zgmo, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zgmo > 0 {
		zgmo--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Blocksz":
			z.Blocksz, err = dc.ReadUint32()
			if err != nil {
				return
			}
		case "BlockMap":
			var ztaf uint32
			ztaf, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.BlockMap == nil && ztaf > 0 {
				z.BlockMap = make(map[string]map[string]Block, ztaf)
			} else if len(z.BlockMap) > 0 {
				for key, _ := range z.BlockMap {
					delete(z.BlockMap, key)
				}
			}
			for ztaf > 0 {
				ztaf--
				var zywj string
				var zjpj map[string]Block
				zywj, err = dc.ReadString()
				if err != nil {
					return
				}
				var zeth uint32
				zeth, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				if zjpj == nil && zeth > 0 {
					zjpj = make(map[string]Block, zeth)
				} else if len(zjpj) > 0 {
					for key, _ := range zjpj {
						delete(zjpj, key)
					}
				}
				for zeth > 0 {
					zeth--
					var zzpf string
					var zrfe Block
					zzpf, err = dc.ReadString()
					if err != nil {
						return
					}
					err = zrfe.DecodeMsg(dc)
					if err != nil {
						return
					}
					zjpj[zzpf] = zrfe
				}
				z.BlockMap[zywj] = zjpj
			}
		case "Source":
			z.Source, err = dc.ReadString()
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
func (z *Fingerprint) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Blocksz"
	err = en.Append(0x83, 0xa7, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x7a)
	if err != nil {
		return err
	}
	err = en.WriteUint32(z.Blocksz)
	if err != nil {
		return
	}
	// write "BlockMap"
	err = en.Append(0xa8, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x61, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.BlockMap)))
	if err != nil {
		return
	}
	for zywj, zjpj := range z.BlockMap {
		err = en.WriteString(zywj)
		if err != nil {
			return
		}
		err = en.WriteMapHeader(uint32(len(zjpj)))
		if err != nil {
			return
		}
		for zzpf, zrfe := range zjpj {
			err = en.WriteString(zzpf)
			if err != nil {
				return
			}
			err = zrfe.EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	// write "Source"
	err = en.Append(0xa6, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Source)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Fingerprint) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Blocksz"
	o = append(o, 0x83, 0xa7, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x7a)
	o = msgp.AppendUint32(o, z.Blocksz)
	// string "BlockMap"
	o = append(o, 0xa8, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x61, 0x70)
	o = msgp.AppendMapHeader(o, uint32(len(z.BlockMap)))
	for zywj, zjpj := range z.BlockMap {
		o = msgp.AppendString(o, zywj)
		o = msgp.AppendMapHeader(o, uint32(len(zjpj)))
		for zzpf, zrfe := range zjpj {
			o = msgp.AppendString(o, zzpf)
			o, err = zrfe.MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "Source"
	o = append(o, 0xa6, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65)
	o = msgp.AppendString(o, z.Source)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Fingerprint) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zsbz uint32
	zsbz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zsbz > 0 {
		zsbz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Blocksz":
			z.Blocksz, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		case "BlockMap":
			var zrjx uint32
			zrjx, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.BlockMap == nil && zrjx > 0 {
				z.BlockMap = make(map[string]map[string]Block, zrjx)
			} else if len(z.BlockMap) > 0 {
				for key, _ := range z.BlockMap {
					delete(z.BlockMap, key)
				}
			}
			for zrjx > 0 {
				var zywj string
				var zjpj map[string]Block
				zrjx--
				zywj, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				var zawn uint32
				zawn, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if zjpj == nil && zawn > 0 {
					zjpj = make(map[string]Block, zawn)
				} else if len(zjpj) > 0 {
					for key, _ := range zjpj {
						delete(zjpj, key)
					}
				}
				for zawn > 0 {
					var zzpf string
					var zrfe Block
					zawn--
					zzpf, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
					bts, err = zrfe.UnmarshalMsg(bts)
					if err != nil {
						return
					}
					zjpj[zzpf] = zrfe
				}
				z.BlockMap[zywj] = zjpj
			}
		case "Source":
			z.Source, bts, err = msgp.ReadStringBytes(bts)
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
func (z *Fingerprint) Msgsize() (s int) {
	s = 1 + 8 + msgp.Uint32Size + 9 + msgp.MapHeaderSize
	if z.BlockMap != nil {
		for zywj, zjpj := range z.BlockMap {
			_ = zjpj
			s += msgp.StringPrefixSize + len(zywj) + msgp.MapHeaderSize
			if zjpj != nil {
				for zzpf, zrfe := range zjpj {
					_ = zrfe
					s += msgp.StringPrefixSize + len(zzpf) + zrfe.Msgsize()
				}
			}
		}
	}
	s += 7 + msgp.StringPrefixSize + len(z.Source)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ModifyChange) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zwel uint32
	zwel, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zwel > 0 {
		zwel--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "from":
			err = z.From.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "to":
			err = z.To.DecodeMsg(dc)
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
	// map header, size 2
	// write "from"
	err = en.Append(0x82, 0xa4, 0x66, 0x72, 0x6f, 0x6d)
	if err != nil {
		return err
	}
	err = z.From.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "to"
	err = en.Append(0xa2, 0x74, 0x6f)
	if err != nil {
		return err
	}
	err = z.To.EncodeMsg(en)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ModifyChange) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "from"
	o = append(o, 0x82, 0xa4, 0x66, 0x72, 0x6f, 0x6d)
	o, err = z.From.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "to"
	o = append(o, 0xa2, 0x74, 0x6f)
	o, err = z.To.MarshalMsg(o)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ModifyChange) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zrbe uint32
	zrbe, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zrbe > 0 {
		zrbe--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "from":
			bts, err = z.From.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "to":
			bts, err = z.To.UnmarshalMsg(bts)
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
	s = 1 + 5 + z.From.Msgsize() + 3 + z.To.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Patch) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zzdc uint32
	zzdc, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zzdc > 0 {
		zzdc--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Blocks":
			var zelx uint32
			zelx, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Blocks) >= int(zelx) {
				z.Blocks = (z.Blocks)[:zelx]
			} else {
				z.Blocks = make([]Block, zelx)
			}
			for zmfd := range z.Blocks {
				err = z.Blocks[zmfd].DecodeMsg(dc)
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
func (z *Patch) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "Blocks"
	err = en.Append(0x81, 0xa6, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Blocks)))
	if err != nil {
		return
	}
	for zmfd := range z.Blocks {
		err = z.Blocks[zmfd].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Patch) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Blocks"
	o = append(o, 0x81, 0xa6, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Blocks)))
	for zmfd := range z.Blocks {
		o, err = z.Blocks[zmfd].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Patch) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zbal uint32
	zbal, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zbal > 0 {
		zbal--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Blocks":
			var zjqz uint32
			zjqz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Blocks) >= int(zjqz) {
				z.Blocks = (z.Blocks)[:zjqz]
			} else {
				z.Blocks = make([]Block, zjqz)
			}
			for zmfd := range z.Blocks {
				bts, err = z.Blocks[zmfd].UnmarshalMsg(bts)
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
func (z *Patch) Msgsize() (s int) {
	s = 1 + 7 + msgp.ArrayHeaderSize
	for zmfd := range z.Blocks {
		s += z.Blocks[zmfd].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *SyncMessage) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var ztmt uint32
	ztmt, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for ztmt > 0 {
		ztmt--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "files":
			var ztco uint32
			ztco, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Files) >= int(ztco) {
				z.Files = (z.Files)[:ztco]
			} else {
				z.Files = make([]FileInfo, ztco)
			}
			for zkct := range z.Files {
				err = z.Files[zkct].DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "changes":
			err = z.Changes.DecodeMsg(dc)
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
func (z *SyncMessage) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "files"
	err = en.Append(0x82, 0xa5, 0x66, 0x69, 0x6c, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Files)))
	if err != nil {
		return
	}
	for zkct := range z.Files {
		err = z.Files[zkct].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "changes"
	err = en.Append(0xa7, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = z.Changes.EncodeMsg(en)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SyncMessage) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "files"
	o = append(o, 0x82, 0xa5, 0x66, 0x69, 0x6c, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Files)))
	for zkct := range z.Files {
		o, err = z.Files[zkct].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "changes"
	o = append(o, 0xa7, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73)
	o, err = z.Changes.MarshalMsg(o)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SyncMessage) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zana uint32
	zana, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zana > 0 {
		zana--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "files":
			var ztyy uint32
			ztyy, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Files) >= int(ztyy) {
				z.Files = (z.Files)[:ztyy]
			} else {
				z.Files = make([]FileInfo, ztyy)
			}
			for zkct := range z.Files {
				bts, err = z.Files[zkct].UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "changes":
			bts, err = z.Changes.UnmarshalMsg(bts)
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
func (z *SyncMessage) Msgsize() (s int) {
	s = 1 + 6 + msgp.ArrayHeaderSize
	for zkct := range z.Files {
		s += z.Files[zkct].Msgsize()
	}
	s += 8 + z.Changes.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *SyncResponseMessage) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zinl uint32
	zinl, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zinl > 0 {
		zinl--
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
	var zare uint32
	zare, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zare > 0 {
		zare--
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
