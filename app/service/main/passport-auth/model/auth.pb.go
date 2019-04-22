// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: auth.proto

/*
	Package model is a generated protocol buffer package.

	It is generated from these files:
		auth.proto

	It has these top-level messages:
		AuthReply
		Cookie
		Token
		Refresh
*/
package model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// AuthReply auth reply
type AuthReply struct {
	// if cookie or token in life time, login is true
	// else login is false and mid csrf expires is empty
	Login bool `protobuf:"varint,1,opt,name=Login,proto3" json:"login"`
	// user identify id
	Mid int64 `protobuf:"varint,2,opt,name=Mid,proto3" json:"mid"`
	// use cookie request this field will return
	// use token request ignore this field
	CSRF string `protobuf:"bytes,3,opt,name=CSRF,proto3" json:"csrf_token"`
	// expiration date
	// unix timestamp
	Expires int64 `protobuf:"varint,4,opt,name=Expires,proto3" json:"expires"`
}

func (m *AuthReply) Reset()                    { *m = AuthReply{} }
func (m *AuthReply) String() string            { return proto.CompactTextString(m) }
func (*AuthReply) ProtoMessage()               {}
func (*AuthReply) Descriptor() ([]byte, []int) { return fileDescriptorAuth, []int{0} }

func (m *AuthReply) GetLogin() bool {
	if m != nil {
		return m.Login
	}
	return false
}

func (m *AuthReply) GetMid() int64 {
	if m != nil {
		return m.Mid
	}
	return 0
}

func (m *AuthReply) GetCSRF() string {
	if m != nil {
		return m.CSRF
	}
	return ""
}

func (m *AuthReply) GetExpires() int64 {
	if m != nil {
		return m.Expires
	}
	return 0
}

type Cookie struct {
	Mid     int64  `protobuf:"varint,1,opt,name=Mid,proto3" json:"mid"`
	Session string `protobuf:"bytes,2,opt,name=Session,proto3" json:"session"`
	CSRF    string `protobuf:"bytes,3,opt,name=CSRF,proto3" json:"csrf"`
	Type    int64  `protobuf:"varint,4,opt,name=Type,proto3" json:"type"`
	Expires int64  `protobuf:"varint,5,opt,name=Expires,proto3" json:"expires"`
}

func (m *Cookie) Reset()                    { *m = Cookie{} }
func (m *Cookie) String() string            { return proto.CompactTextString(m) }
func (*Cookie) ProtoMessage()               {}
func (*Cookie) Descriptor() ([]byte, []int) { return fileDescriptorAuth, []int{1} }

func (m *Cookie) GetMid() int64 {
	if m != nil {
		return m.Mid
	}
	return 0
}

func (m *Cookie) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

func (m *Cookie) GetCSRF() string {
	if m != nil {
		return m.CSRF
	}
	return ""
}

func (m *Cookie) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Cookie) GetExpires() int64 {
	if m != nil {
		return m.Expires
	}
	return 0
}

type Token struct {
	Mid     int64  `protobuf:"varint,1,opt,name=Mid,proto3" json:"mid"`
	AppID   int32  `protobuf:"varint,2,opt,name=AppID,proto3" json:"appid"`
	Token   string `protobuf:"bytes,3,opt,name=Token,proto3" json:"token"`
	Type    int64  `protobuf:"varint,4,opt,name=Type,proto3" json:"type"`
	Expires int64  `protobuf:"varint,5,opt,name=Expires,proto3" json:"expires"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptorAuth, []int{2} }

func (m *Token) GetMid() int64 {
	if m != nil {
		return m.Mid
	}
	return 0
}

func (m *Token) GetAppID() int32 {
	if m != nil {
		return m.AppID
	}
	return 0
}

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Token) GetExpires() int64 {
	if m != nil {
		return m.Expires
	}
	return 0
}

type Refresh struct {
	Mid     int64  `protobuf:"varint,1,opt,name=Mid,proto3" json:"mid"`
	AppID   int32  `protobuf:"varint,2,opt,name=AppID,proto3" json:"appid"`
	Refresh string `protobuf:"bytes,3,opt,name=Refresh,proto3" json:"refresh"`
	Token   string `protobuf:"bytes,4,opt,name=Token,proto3" json:"token"`
	Expires int64  `protobuf:"varint,5,opt,name=Expires,proto3" json:"expires"`
}

func (m *Refresh) Reset()                    { *m = Refresh{} }
func (m *Refresh) String() string            { return proto.CompactTextString(m) }
func (*Refresh) ProtoMessage()               {}
func (*Refresh) Descriptor() ([]byte, []int) { return fileDescriptorAuth, []int{3} }

func (m *Refresh) GetMid() int64 {
	if m != nil {
		return m.Mid
	}
	return 0
}

func (m *Refresh) GetAppID() int32 {
	if m != nil {
		return m.AppID
	}
	return 0
}

func (m *Refresh) GetRefresh() string {
	if m != nil {
		return m.Refresh
	}
	return ""
}

func (m *Refresh) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Refresh) GetExpires() int64 {
	if m != nil {
		return m.Expires
	}
	return 0
}

func init() {
	proto.RegisterType((*AuthReply)(nil), "passport.service.auth.AuthReply")
	proto.RegisterType((*Cookie)(nil), "passport.service.auth.Cookie")
	proto.RegisterType((*Token)(nil), "passport.service.auth.Token")
	proto.RegisterType((*Refresh)(nil), "passport.service.auth.Refresh")
}
func (m *AuthReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Login {
		dAtA[i] = 0x8
		i++
		if m.Login {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Mid != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintAuth(dAtA, i, uint64(m.Mid))
	}
	if len(m.CSRF) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintAuth(dAtA, i, uint64(len(m.CSRF)))
		i += copy(dAtA[i:], m.CSRF)
	}
	if m.Expires != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintAuth(dAtA, i, uint64(m.Expires))
	}
	return i, nil
}

func (m *Cookie) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Cookie) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Mid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintAuth(dAtA, i, uint64(m.Mid))
	}
	if len(m.Session) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAuth(dAtA, i, uint64(len(m.Session)))
		i += copy(dAtA[i:], m.Session)
	}
	if len(m.CSRF) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintAuth(dAtA, i, uint64(len(m.CSRF)))
		i += copy(dAtA[i:], m.CSRF)
	}
	if m.Type != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintAuth(dAtA, i, uint64(m.Type))
	}
	if m.Expires != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintAuth(dAtA, i, uint64(m.Expires))
	}
	return i, nil
}

func (m *Token) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Token) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Mid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintAuth(dAtA, i, uint64(m.Mid))
	}
	if m.AppID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintAuth(dAtA, i, uint64(m.AppID))
	}
	if len(m.Token) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintAuth(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	if m.Type != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintAuth(dAtA, i, uint64(m.Type))
	}
	if m.Expires != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintAuth(dAtA, i, uint64(m.Expires))
	}
	return i, nil
}

func (m *Refresh) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Refresh) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Mid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintAuth(dAtA, i, uint64(m.Mid))
	}
	if m.AppID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintAuth(dAtA, i, uint64(m.AppID))
	}
	if len(m.Refresh) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintAuth(dAtA, i, uint64(len(m.Refresh)))
		i += copy(dAtA[i:], m.Refresh)
	}
	if len(m.Token) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintAuth(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	if m.Expires != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintAuth(dAtA, i, uint64(m.Expires))
	}
	return i, nil
}

func encodeVarintAuth(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AuthReply) Size() (n int) {
	var l int
	_ = l
	if m.Login {
		n += 2
	}
	if m.Mid != 0 {
		n += 1 + sovAuth(uint64(m.Mid))
	}
	l = len(m.CSRF)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	if m.Expires != 0 {
		n += 1 + sovAuth(uint64(m.Expires))
	}
	return n
}

func (m *Cookie) Size() (n int) {
	var l int
	_ = l
	if m.Mid != 0 {
		n += 1 + sovAuth(uint64(m.Mid))
	}
	l = len(m.Session)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	l = len(m.CSRF)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovAuth(uint64(m.Type))
	}
	if m.Expires != 0 {
		n += 1 + sovAuth(uint64(m.Expires))
	}
	return n
}

func (m *Token) Size() (n int) {
	var l int
	_ = l
	if m.Mid != 0 {
		n += 1 + sovAuth(uint64(m.Mid))
	}
	if m.AppID != 0 {
		n += 1 + sovAuth(uint64(m.AppID))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovAuth(uint64(m.Type))
	}
	if m.Expires != 0 {
		n += 1 + sovAuth(uint64(m.Expires))
	}
	return n
}

func (m *Refresh) Size() (n int) {
	var l int
	_ = l
	if m.Mid != 0 {
		n += 1 + sovAuth(uint64(m.Mid))
	}
	if m.AppID != 0 {
		n += 1 + sovAuth(uint64(m.AppID))
	}
	l = len(m.Refresh)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	if m.Expires != 0 {
		n += 1 + sovAuth(uint64(m.Expires))
	}
	return n
}

func sovAuth(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAuth(x uint64) (n int) {
	return sovAuth(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AuthReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AuthReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Login", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Login = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mid", wireType)
			}
			m.Mid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CSRF", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CSRF = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expires", wireType)
			}
			m.Expires = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Expires |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuth
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Cookie) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Cookie: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Cookie: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mid", wireType)
			}
			m.Mid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Session", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Session = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CSRF", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CSRF = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expires", wireType)
			}
			m.Expires = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Expires |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuth
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Token) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Token: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Token: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mid", wireType)
			}
			m.Mid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppID", wireType)
			}
			m.AppID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AppID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expires", wireType)
			}
			m.Expires = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Expires |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuth
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Refresh) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Refresh: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Refresh: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mid", wireType)
			}
			m.Mid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppID", wireType)
			}
			m.AppID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AppID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Refresh", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Refresh = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expires", wireType)
			}
			m.Expires = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Expires |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuth
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAuth(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuth
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthAuth
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAuth
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipAuth(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthAuth = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuth   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("auth.proto", fileDescriptorAuth) }

var fileDescriptorAuth = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0xc6, 0xef, 0xdc, 0x24, 0x4d, 0x33, 0x17, 0xee, 0x22, 0x70, 0xb9, 0x51, 0x4a, 0x52, 0x0a,
	0x85, 0x6e, 0x4c, 0x17, 0x3e, 0x41, 0x53, 0x15, 0x04, 0xdd, 0x4c, 0xbb, 0x72, 0x23, 0x6d, 0x33,
	0x4d, 0x86, 0xfe, 0x99, 0x21, 0x93, 0x88, 0x7d, 0x0d, 0x57, 0xbe, 0x82, 0x0b, 0x77, 0x3e, 0x84,
	0x4b, 0x9f, 0x20, 0x48, 0xdd, 0xe5, 0x29, 0x24, 0x67, 0xd2, 0xa2, 0xa0, 0x45, 0xd0, 0x5d, 0xe6,
	0xfb, 0x0e, 0xdf, 0xf7, 0x3b, 0x87, 0x60, 0x3c, 0xca, 0xd2, 0xd8, 0x17, 0x09, 0x4f, 0xb9, 0xfd,
	0x4f, 0x8c, 0xa4, 0x14, 0x3c, 0x49, 0x7d, 0x49, 0x93, 0x2b, 0x36, 0xa1, 0x7e, 0x69, 0xee, 0x1f,
	0x44, 0x2c, 0x8d, 0xb3, 0xb1, 0x3f, 0xe1, 0x8b, 0x6e, 0xc4, 0x23, 0xde, 0x85, 0xe9, 0x71, 0x36,
	0x85, 0x17, 0x3c, 0xe0, 0x4b, 0xa5, 0xb4, 0x6e, 0x10, 0xb6, 0x7a, 0x59, 0x1a, 0x13, 0x2a, 0xe6,
	0x2b, 0xdb, 0xc3, 0xc6, 0x19, 0x8f, 0xd8, 0xd2, 0x41, 0x4d, 0xd4, 0xa9, 0x07, 0x56, 0x91, 0x7b,
	0xc6, 0xbc, 0x14, 0x88, 0xd2, 0xed, 0x3d, 0xac, 0x9d, 0xb3, 0xd0, 0xf9, 0xdd, 0x44, 0x1d, 0x2d,
	0x30, 0x8b, 0xdc, 0xd3, 0x16, 0x2c, 0x24, 0xa5, 0x66, 0xb7, 0xb0, 0xde, 0x1f, 0x90, 0x13, 0x47,
	0x6b, 0xa2, 0x8e, 0x15, 0xfc, 0x2d, 0x72, 0x0f, 0x4f, 0x64, 0x32, 0xbd, 0x4c, 0xf9, 0x8c, 0x2e,
	0x09, 0x78, 0x76, 0x1b, 0x9b, 0xc7, 0xd7, 0x82, 0x25, 0x54, 0x3a, 0x3a, 0x44, 0xfc, 0x29, 0x72,
	0xcf, 0xa4, 0x4a, 0x22, 0x1b, 0xaf, 0x75, 0x8f, 0x70, 0xad, 0xcf, 0xf9, 0x8c, 0xd1, 0x4d, 0x21,
	0xfa, 0xa0, 0xb0, 0x8d, 0xcd, 0x01, 0x95, 0x92, 0xf1, 0x25, 0xf0, 0x58, 0x2a, 0x4c, 0x2a, 0x89,
	0x6c, 0x3c, 0xbb, 0xf1, 0x8e, 0xab, 0x5e, 0xe4, 0x9e, 0x5e, 0x72, 0x55, 0x44, 0x0d, 0xac, 0x0f,
	0x57, 0x82, 0x56, 0x38, 0xe0, 0xa6, 0x2b, 0x41, 0x09, 0xa8, 0x6f, 0x79, 0x8d, 0x1d, 0xbc, 0x77,
	0x08, 0x1b, 0xc3, 0x72, 0xcd, 0x5d, 0xb8, 0x1e, 0x36, 0x7a, 0x42, 0x9c, 0x1e, 0x01, 0xac, 0xa1,
	0x6e, 0x3b, 0x12, 0x82, 0x85, 0x44, 0xe9, 0xe5, 0x00, 0x84, 0x54, 0xa4, 0x30, 0xa0, 0x8e, 0x57,
	0x85, 0xff, 0x08, 0xeb, 0x03, 0xc2, 0x26, 0xa1, 0xd3, 0x84, 0xca, 0xf8, 0x5b, 0xb4, 0xed, 0x6d,
	0x4c, 0xc5, 0x0b, 0x75, 0x89, 0x92, 0xc8, 0xb6, 0x62, 0xbb, 0x94, 0xfe, 0xc9, 0x52, 0x5f, 0xc3,
	0x0e, 0xfe, 0x3f, 0xae, 0x5d, 0xf4, 0xb4, 0x76, 0xd1, 0xf3, 0xda, 0x45, 0xb7, 0x2f, 0xee, 0xaf,
	0x0b, 0x63, 0xc1, 0x43, 0x3a, 0x1f, 0xd7, 0xe0, 0x3f, 0x3e, 0x7c, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0x69, 0x35, 0xab, 0x86, 0x1b, 0x03, 0x00, 0x00,
}
