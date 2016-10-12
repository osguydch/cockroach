// Code generated by protoc-gen-gogo.
// source: cockroach/pkg/storage/storagebase/state.proto
// DO NOT EDIT!

/*
	Package storagebase is a generated protocol buffer package.

	It is generated from these files:
		cockroach/pkg/storage/storagebase/state.proto

	It has these top-level messages:
		ReplicaState
		RangeInfo
*/
package storagebase

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import cockroach_storage_engine_enginepb "github.com/cockroachdb/cockroach/pkg/storage/engine/enginepb"
import cockroach_roachpb1 "github.com/cockroachdb/cockroach/pkg/roachpb"
import cockroach_roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"
import cockroach_roachpb2 "github.com/cockroachdb/cockroach/pkg/roachpb"
import cockroach_util_hlc "github.com/cockroachdb/cockroach/pkg/util/hlc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// ReplicaState is the part of the Range Raft state machine which is cached
// in memory and which is manipulated exclusively through consensus.
type ReplicaState struct {
	// The highest (and last) index applied to the state machine.
	RaftAppliedIndex uint64 `protobuf:"varint,1,opt,name=raft_applied_index,json=raftAppliedIndex,proto3" json:"raft_applied_index,omitempty"`
	// The highest (and last) lease index applied to the state machine.
	LeaseAppliedIndex uint64 `protobuf:"varint,2,opt,name=lease_applied_index,json=leaseAppliedIndex,proto3" json:"lease_applied_index,omitempty"`
	// The Range descriptor.
	// The pointer may change, but the referenced RangeDescriptor struct itself
	// must be treated as immutable; it is leaked out of the lock.
	//
	// Changes of the descriptor should always go through one of the
	// (*Replica).setDesc* methods.
	Desc *cockroach_roachpb.RangeDescriptor `protobuf:"bytes,3,opt,name=desc" json:"desc,omitempty"`
	// The latest lease, if any.
	Lease *cockroach_roachpb2.Lease `protobuf:"bytes,4,opt,name=lease" json:"lease,omitempty"`
	// The truncation state of the Raft log.
	TruncatedState *cockroach_roachpb1.RaftTruncatedState `protobuf:"bytes,5,opt,name=truncated_state,json=truncatedState" json:"truncated_state,omitempty"`
	// gcThreshold is the GC threshold of the Range, typically updated when keys
	// are garbage collected. Reads and writes at timestamps <= this time will
	// not be served.
	GCThreshold cockroach_util_hlc.Timestamp                `protobuf:"bytes,6,opt,name=gc_threshold,json=gcThreshold" json:"gc_threshold"`
	Stats       cockroach_storage_engine_enginepb.MVCCStats `protobuf:"bytes,7,opt,name=stats" json:"stats"`
	Frozen      bool                                        `protobuf:"varint,8,opt,name=frozen,proto3" json:"frozen,omitempty"`
	// txn_span_gc_threshold is the (maximum) timestamp below which transaction
	// records may have been garbage collected (as measured by txn.LastActive()).
	// Transaction at lower timestamps must not be allowed to write their initial
	// transaction entry.
	TxnSpanGCThreshold cockroach_util_hlc.Timestamp `protobuf:"bytes,9,opt,name=txn_span_gc_threshold,json=txnSpanGcThreshold" json:"txn_span_gc_threshold"`
}

func (m *ReplicaState) Reset()                    { *m = ReplicaState{} }
func (m *ReplicaState) String() string            { return proto.CompactTextString(m) }
func (*ReplicaState) ProtoMessage()               {}
func (*ReplicaState) Descriptor() ([]byte, []int) { return fileDescriptorState, []int{0} }

type RangeInfo struct {
	ReplicaState `protobuf:"bytes,1,opt,name=state,embedded=state" json:"state"`
	// The highest (and last) index in the Raft log.
	LastIndex  uint64 `protobuf:"varint,2,opt,name=lastIndex,proto3" json:"lastIndex,omitempty"`
	NumPending uint64 `protobuf:"varint,3,opt,name=num_pending,json=numPending,proto3" json:"num_pending,omitempty"`
	NumDropped uint64 `protobuf:"varint,5,opt,name=num_dropped,json=numDropped,proto3" json:"num_dropped,omitempty"`
	// raft_log_size may be initially inaccurate after a server restart.
	// See storage.Replica.mu.raftLogSize.
	RaftLogSize int64 `protobuf:"varint,6,opt,name=raft_log_size,json=raftLogSize,proto3" json:"raft_log_size,omitempty"`
}

func (m *RangeInfo) Reset()                    { *m = RangeInfo{} }
func (m *RangeInfo) String() string            { return proto.CompactTextString(m) }
func (*RangeInfo) ProtoMessage()               {}
func (*RangeInfo) Descriptor() ([]byte, []int) { return fileDescriptorState, []int{1} }

func init() {
	proto.RegisterType((*ReplicaState)(nil), "cockroach.storage.storagebase.ReplicaState")
	proto.RegisterType((*RangeInfo)(nil), "cockroach.storage.storagebase.RangeInfo")
}
func (m *ReplicaState) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ReplicaState) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RaftAppliedIndex != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintState(data, i, uint64(m.RaftAppliedIndex))
	}
	if m.LeaseAppliedIndex != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintState(data, i, uint64(m.LeaseAppliedIndex))
	}
	if m.Desc != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintState(data, i, uint64(m.Desc.Size()))
		n1, err := m.Desc.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Lease != nil {
		data[i] = 0x22
		i++
		i = encodeVarintState(data, i, uint64(m.Lease.Size()))
		n2, err := m.Lease.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.TruncatedState != nil {
		data[i] = 0x2a
		i++
		i = encodeVarintState(data, i, uint64(m.TruncatedState.Size()))
		n3, err := m.TruncatedState.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	data[i] = 0x32
	i++
	i = encodeVarintState(data, i, uint64(m.GCThreshold.Size()))
	n4, err := m.GCThreshold.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	data[i] = 0x3a
	i++
	i = encodeVarintState(data, i, uint64(m.Stats.Size()))
	n5, err := m.Stats.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	if m.Frozen {
		data[i] = 0x40
		i++
		if m.Frozen {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	data[i] = 0x4a
	i++
	i = encodeVarintState(data, i, uint64(m.TxnSpanGCThreshold.Size()))
	n6, err := m.TxnSpanGCThreshold.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	return i, nil
}

func (m *RangeInfo) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RangeInfo) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintState(data, i, uint64(m.ReplicaState.Size()))
	n7, err := m.ReplicaState.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n7
	if m.LastIndex != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintState(data, i, uint64(m.LastIndex))
	}
	if m.NumPending != 0 {
		data[i] = 0x18
		i++
		i = encodeVarintState(data, i, uint64(m.NumPending))
	}
	if m.NumDropped != 0 {
		data[i] = 0x28
		i++
		i = encodeVarintState(data, i, uint64(m.NumDropped))
	}
	if m.RaftLogSize != 0 {
		data[i] = 0x30
		i++
		i = encodeVarintState(data, i, uint64(m.RaftLogSize))
	}
	return i, nil
}

func encodeFixed64State(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32State(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintState(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *ReplicaState) Size() (n int) {
	var l int
	_ = l
	if m.RaftAppliedIndex != 0 {
		n += 1 + sovState(uint64(m.RaftAppliedIndex))
	}
	if m.LeaseAppliedIndex != 0 {
		n += 1 + sovState(uint64(m.LeaseAppliedIndex))
	}
	if m.Desc != nil {
		l = m.Desc.Size()
		n += 1 + l + sovState(uint64(l))
	}
	if m.Lease != nil {
		l = m.Lease.Size()
		n += 1 + l + sovState(uint64(l))
	}
	if m.TruncatedState != nil {
		l = m.TruncatedState.Size()
		n += 1 + l + sovState(uint64(l))
	}
	l = m.GCThreshold.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.Stats.Size()
	n += 1 + l + sovState(uint64(l))
	if m.Frozen {
		n += 2
	}
	l = m.TxnSpanGCThreshold.Size()
	n += 1 + l + sovState(uint64(l))
	return n
}

func (m *RangeInfo) Size() (n int) {
	var l int
	_ = l
	l = m.ReplicaState.Size()
	n += 1 + l + sovState(uint64(l))
	if m.LastIndex != 0 {
		n += 1 + sovState(uint64(m.LastIndex))
	}
	if m.NumPending != 0 {
		n += 1 + sovState(uint64(m.NumPending))
	}
	if m.NumDropped != 0 {
		n += 1 + sovState(uint64(m.NumDropped))
	}
	if m.RaftLogSize != 0 {
		n += 1 + sovState(uint64(m.RaftLogSize))
	}
	return n
}

func sovState(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozState(x uint64) (n int) {
	return sovState(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ReplicaState) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReplicaState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReplicaState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RaftAppliedIndex", wireType)
			}
			m.RaftAppliedIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.RaftAppliedIndex |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaseAppliedIndex", wireType)
			}
			m.LeaseAppliedIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.LeaseAppliedIndex |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Desc == nil {
				m.Desc = &cockroach_roachpb.RangeDescriptor{}
			}
			if err := m.Desc.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lease", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Lease == nil {
				m.Lease = &cockroach_roachpb2.Lease{}
			}
			if err := m.Lease.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TruncatedState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TruncatedState == nil {
				m.TruncatedState = &cockroach_roachpb1.RaftTruncatedState{}
			}
			if err := m.TruncatedState.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GCThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GCThreshold.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Stats.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Frozen", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Frozen = bool(v != 0)
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxnSpanGCThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TxnSpanGCThreshold.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthState
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
func (m *RangeInfo) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RangeInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RangeInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplicaState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReplicaState.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastIndex", wireType)
			}
			m.LastIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.LastIndex |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumPending", wireType)
			}
			m.NumPending = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.NumPending |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumDropped", wireType)
			}
			m.NumDropped = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.NumDropped |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RaftLogSize", wireType)
			}
			m.RaftLogSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.RaftLogSize |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipState(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthState
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
func skipState(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowState
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowState
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowState
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthState
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowState
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipState(data[start:])
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
	ErrInvalidLengthState = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowState   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("cockroach/pkg/storage/storagebase/state.proto", fileDescriptorState) }

var fileDescriptorState = []byte{
	// 602 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x93, 0xd1, 0x6e, 0xd3, 0x3c,
	0x14, 0xc7, 0x97, 0x6f, 0xe9, 0xbe, 0xce, 0x19, 0x30, 0x3c, 0x40, 0xd1, 0xc4, 0xd2, 0xaa, 0x62,
	0x68, 0x88, 0xe1, 0xa0, 0x21, 0x71, 0x4f, 0x37, 0x09, 0x06, 0x03, 0xa1, 0xac, 0x70, 0xc1, 0x4d,
	0xe4, 0x3a, 0x6e, 0x1a, 0x2d, 0xb1, 0xad, 0xd8, 0x45, 0x53, 0x9f, 0x82, 0xc7, 0xaa, 0xc4, 0xcd,
	0x2e, 0xb9, 0xaa, 0x20, 0xbc, 0x08, 0xb2, 0x93, 0xac, 0x29, 0x54, 0x88, 0xab, 0xd8, 0xe7, 0xfc,
	0xce, 0x89, 0xff, 0xc7, 0x7f, 0x83, 0x27, 0x84, 0x93, 0x8b, 0x9c, 0x63, 0x32, 0xf6, 0xc5, 0x45,
	0xec, 0x4b, 0xc5, 0x73, 0x1c, 0xd3, 0xfa, 0x3b, 0xc4, 0x52, 0xaf, 0xb1, 0xa2, 0x48, 0xe4, 0x5c,
	0x71, 0xb8, 0x77, 0x8d, 0xa3, 0x0a, 0x41, 0x0d, 0x74, 0xf7, 0xe9, 0xea, 0x6e, 0x94, 0xc5, 0x09,
	0xab, 0x3f, 0x62, 0xe8, 0x67, 0x9f, 0x09, 0x29, 0x1b, 0xee, 0x3e, 0x5a, 0xae, 0x30, 0x2b, 0x31,
	0xf4, 0x13, 0xa6, 0x68, 0xce, 0x70, 0x1a, 0xe6, 0x78, 0xa4, 0x2a, 0xf4, 0xc1, 0x6a, 0x34, 0xa3,
	0x0a, 0x47, 0x58, 0xe1, 0x8a, 0xea, 0xae, 0xa6, 0x1a, 0xc4, 0xc3, 0x65, 0x62, 0xa2, 0x92, 0xd4,
	0x1f, 0xa7, 0xc4, 0x57, 0x49, 0x46, 0xa5, 0xc2, 0x99, 0xa8, 0xb8, 0x3b, 0x31, 0x8f, 0xb9, 0x59,
	0xfa, 0x7a, 0x55, 0x46, 0x7b, 0x5f, 0x6d, 0xb0, 0x15, 0x50, 0x91, 0x26, 0x04, 0x9f, 0xeb, 0xc1,
	0xc0, 0x43, 0x00, 0xf5, 0x21, 0x43, 0x2c, 0x44, 0x9a, 0xd0, 0x28, 0x4c, 0x58, 0x44, 0x2f, 0x5d,
	0xab, 0x6b, 0x1d, 0xd8, 0xc1, 0xb6, 0xce, 0xbc, 0x28, 0x13, 0xa7, 0x3a, 0x0e, 0x11, 0xd8, 0x49,
	0x29, 0x96, 0xf4, 0x37, 0xfc, 0x3f, 0x83, 0xdf, 0x36, 0xa9, 0x25, 0xfe, 0x39, 0xb0, 0x23, 0x2a,
	0x89, 0xbb, 0xde, 0xb5, 0x0e, 0x9c, 0xa3, 0x1e, 0x5a, 0xcc, 0xbf, 0x52, 0x86, 0x02, 0xcc, 0x62,
	0x7a, 0x42, 0x25, 0xc9, 0x13, 0xa1, 0x78, 0x1e, 0x18, 0x1e, 0x22, 0xd0, 0x32, 0xcd, 0x5c, 0xdb,
	0x14, 0xba, 0x2b, 0x0a, 0xcf, 0x74, 0x3e, 0x28, 0x31, 0xf8, 0x0e, 0xdc, 0x52, 0xf9, 0x84, 0x11,
	0xac, 0x68, 0x14, 0x9a, 0x1b, 0x77, 0x5b, 0xa6, 0x72, 0x7f, 0xe5, 0x2f, 0x47, 0x6a, 0x50, 0xd3,
	0x66, 0x0a, 0xc1, 0x4d, 0xb5, 0xb4, 0x87, 0x1f, 0xc0, 0x56, 0x4c, 0x42, 0x35, 0xce, 0xa9, 0x1c,
	0xf3, 0x34, 0x72, 0x37, 0x4c, 0xb3, 0xbd, 0x46, 0x33, 0x3d, 0x77, 0x34, 0x4e, 0x09, 0x1a, 0xd4,
	0x73, 0xef, 0xef, 0xcc, 0xe6, 0x9d, 0xb5, 0x62, 0xde, 0x71, 0x5e, 0x1e, 0x0f, 0xea, 0xca, 0xc0,
	0x89, 0xc9, 0xf5, 0x06, 0xbe, 0x02, 0x2d, 0x7d, 0x38, 0xe9, 0xfe, 0x6f, 0xfa, 0x1d, 0xa2, 0x3f,
	0xfd, 0x58, 0xba, 0x0c, 0xd5, 0x66, 0x43, 0x6f, 0x3f, 0x1e, 0x1f, 0xeb, 0x33, 0xc9, 0xbe, 0xad,
	0xdb, 0x07, 0x65, 0x03, 0x78, 0x0f, 0x6c, 0x8c, 0x72, 0x3e, 0xa5, 0xcc, 0x6d, 0x77, 0xad, 0x83,
	0x76, 0x50, 0xed, 0x60, 0x0a, 0xee, 0xaa, 0x4b, 0x16, 0x4a, 0x81, 0x59, 0xb8, 0xa4, 0x60, 0xf3,
	0x5f, 0x14, 0xec, 0x56, 0x0a, 0xe0, 0xe0, 0x92, 0x9d, 0x0b, 0xcc, 0x9a, 0x42, 0xa0, 0xaa, 0x62,
	0x0b, 0x3d, 0xbd, 0xc2, 0x02, 0x9b, 0xe6, 0x02, 0x4f, 0xd9, 0x88, 0xc3, 0x37, 0xa5, 0x3a, 0x6a,
	0xdc, 0xe3, 0x1c, 0x3d, 0x46, 0x7f, 0x7d, 0x6d, 0xa8, 0x69, 0xc3, 0x7e, 0x5b, 0xff, 0xf9, 0x6a,
	0xde, 0xb1, 0x4a, 0x81, 0x14, 0xde, 0x07, 0x9b, 0x29, 0x96, 0xea, 0xb4, 0xe1, 0xaf, 0x45, 0x00,
	0x76, 0x80, 0xc3, 0x26, 0x59, 0x28, 0x28, 0x8b, 0x12, 0x16, 0x1b, 0x7b, 0xd9, 0x01, 0x60, 0x93,
	0xec, 0x7d, 0x19, 0xa9, 0x81, 0x28, 0xe7, 0x42, 0xd0, 0xc8, 0x98, 0xa1, 0x04, 0x4e, 0xca, 0x08,
	0xec, 0x81, 0x1b, 0xc6, 0xf7, 0x29, 0x8f, 0x43, 0x99, 0x4c, 0xa9, 0xb9, 0xe2, 0xf5, 0xc0, 0xd1,
	0xc1, 0x33, 0x1e, 0x9f, 0x27, 0x53, 0xfa, 0xda, 0x6e, 0xdb, 0xdb, 0xad, 0xfe, 0xfe, 0xec, 0x87,
	0xb7, 0x36, 0x2b, 0x3c, 0xeb, 0xaa, 0xf0, 0xac, 0x6f, 0x85, 0x67, 0x7d, 0x2f, 0x3c, 0xeb, 0xcb,
	0x4f, 0x6f, 0xed, 0x93, 0xd3, 0x90, 0x33, 0xdc, 0x30, 0x0f, 0xec, 0xd9, 0xaf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x95, 0x58, 0xe2, 0xf8, 0x93, 0x04, 0x00, 0x00,
}