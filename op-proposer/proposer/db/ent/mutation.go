// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/moongate-forks/kona-sp1/op-proposer/proposer/db/ent/predicate"
	"github.com/moongate-forks/kona-sp1/op-proposer/proposer/db/ent/proofrequest"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeProofRequest = "ProofRequest"
)

// ProofRequestMutation represents an operation that mutates the ProofRequest nodes in the graph.
type ProofRequestMutation struct {
	config
	op                    Op
	typ                   string
	id                    *int
	_type                 *proofrequest.Type
	start_block           *uint64
	addstart_block        *int64
	end_block             *uint64
	addend_block          *int64
	status                *proofrequest.Status
	request_added_time    *uint64
	addrequest_added_time *int64
	prover_request_id     *string
	proof_request_time    *uint64
	addproof_request_time *int64
	l1_block_number       *uint64
	addl1_block_number    *int64
	l1_block_hash         *string
	proof                 *[]byte
	clearedFields         map[string]struct{}
	done                  bool
	oldValue              func(context.Context) (*ProofRequest, error)
	predicates            []predicate.ProofRequest
}

var _ ent.Mutation = (*ProofRequestMutation)(nil)

// proofrequestOption allows management of the mutation configuration using functional options.
type proofrequestOption func(*ProofRequestMutation)

// newProofRequestMutation creates new mutation for the ProofRequest entity.
func newProofRequestMutation(c config, op Op, opts ...proofrequestOption) *ProofRequestMutation {
	m := &ProofRequestMutation{
		config:        c,
		op:            op,
		typ:           TypeProofRequest,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withProofRequestID sets the ID field of the mutation.
func withProofRequestID(id int) proofrequestOption {
	return func(m *ProofRequestMutation) {
		var (
			err   error
			once  sync.Once
			value *ProofRequest
		)
		m.oldValue = func(ctx context.Context) (*ProofRequest, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().ProofRequest.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withProofRequest sets the old ProofRequest of the mutation.
func withProofRequest(node *ProofRequest) proofrequestOption {
	return func(m *ProofRequestMutation) {
		m.oldValue = func(context.Context) (*ProofRequest, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ProofRequestMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ProofRequestMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ProofRequestMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ProofRequestMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().ProofRequest.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetType sets the "type" field.
func (m *ProofRequestMutation) SetType(pr proofrequest.Type) {
	m._type = &pr
}

// GetType returns the value of the "type" field in the mutation.
func (m *ProofRequestMutation) GetType() (r proofrequest.Type, exists bool) {
	v := m._type
	if v == nil {
		return
	}
	return *v, true
}

// OldType returns the old "type" field's value of the ProofRequest entity.
// If the ProofRequest object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProofRequestMutation) OldType(ctx context.Context) (v proofrequest.Type, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldType is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldType requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldType: %w", err)
	}
	return oldValue.Type, nil
}

// ResetType resets all changes to the "type" field.
func (m *ProofRequestMutation) ResetType() {
	m._type = nil
}

// SetStartBlock sets the "start_block" field.
func (m *ProofRequestMutation) SetStartBlock(u uint64) {
	m.start_block = &u
	m.addstart_block = nil
}

// StartBlock returns the value of the "start_block" field in the mutation.
func (m *ProofRequestMutation) StartBlock() (r uint64, exists bool) {
	v := m.start_block
	if v == nil {
		return
	}
	return *v, true
}

// OldStartBlock returns the old "start_block" field's value of the ProofRequest entity.
// If the ProofRequest object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProofRequestMutation) OldStartBlock(ctx context.Context) (v uint64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStartBlock is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStartBlock requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStartBlock: %w", err)
	}
	return oldValue.StartBlock, nil
}

// AddStartBlock adds u to the "start_block" field.
func (m *ProofRequestMutation) AddStartBlock(u int64) {
	if m.addstart_block != nil {
		*m.addstart_block += u
	} else {
		m.addstart_block = &u
	}
}

// AddedStartBlock returns the value that was added to the "start_block" field in this mutation.
func (m *ProofRequestMutation) AddedStartBlock() (r int64, exists bool) {
	v := m.addstart_block
	if v == nil {
		return
	}
	return *v, true
}

// ResetStartBlock resets all changes to the "start_block" field.
func (m *ProofRequestMutation) ResetStartBlock() {
	m.start_block = nil
	m.addstart_block = nil
}

// SetEndBlock sets the "end_block" field.
func (m *ProofRequestMutation) SetEndBlock(u uint64) {
	m.end_block = &u
	m.addend_block = nil
}

// EndBlock returns the value of the "end_block" field in the mutation.
func (m *ProofRequestMutation) EndBlock() (r uint64, exists bool) {
	v := m.end_block
	if v == nil {
		return
	}
	return *v, true
}

// OldEndBlock returns the old "end_block" field's value of the ProofRequest entity.
// If the ProofRequest object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProofRequestMutation) OldEndBlock(ctx context.Context) (v uint64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEndBlock is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEndBlock requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEndBlock: %w", err)
	}
	return oldValue.EndBlock, nil
}

// AddEndBlock adds u to the "end_block" field.
func (m *ProofRequestMutation) AddEndBlock(u int64) {
	if m.addend_block != nil {
		*m.addend_block += u
	} else {
		m.addend_block = &u
	}
}

// AddedEndBlock returns the value that was added to the "end_block" field in this mutation.
func (m *ProofRequestMutation) AddedEndBlock() (r int64, exists bool) {
	v := m.addend_block
	if v == nil {
		return
	}
	return *v, true
}

// ResetEndBlock resets all changes to the "end_block" field.
func (m *ProofRequestMutation) ResetEndBlock() {
	m.end_block = nil
	m.addend_block = nil
}

// SetStatus sets the "status" field.
func (m *ProofRequestMutation) SetStatus(pr proofrequest.Status) {
	m.status = &pr
}

// Status returns the value of the "status" field in the mutation.
func (m *ProofRequestMutation) Status() (r proofrequest.Status, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the ProofRequest entity.
// If the ProofRequest object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProofRequestMutation) OldStatus(ctx context.Context) (v proofrequest.Status, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// ResetStatus resets all changes to the "status" field.
func (m *ProofRequestMutation) ResetStatus() {
	m.status = nil
}

// SetRequestAddedTime sets the "request_added_time" field.
func (m *ProofRequestMutation) SetRequestAddedTime(u uint64) {
	m.request_added_time = &u
	m.addrequest_added_time = nil
}

// RequestAddedTime returns the value of the "request_added_time" field in the mutation.
func (m *ProofRequestMutation) RequestAddedTime() (r uint64, exists bool) {
	v := m.request_added_time
	if v == nil {
		return
	}
	return *v, true
}

// OldRequestAddedTime returns the old "request_added_time" field's value of the ProofRequest entity.
// If the ProofRequest object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProofRequestMutation) OldRequestAddedTime(ctx context.Context) (v uint64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRequestAddedTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRequestAddedTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRequestAddedTime: %w", err)
	}
	return oldValue.RequestAddedTime, nil
}

// AddRequestAddedTime adds u to the "request_added_time" field.
func (m *ProofRequestMutation) AddRequestAddedTime(u int64) {
	if m.addrequest_added_time != nil {
		*m.addrequest_added_time += u
	} else {
		m.addrequest_added_time = &u
	}
}

// AddedRequestAddedTime returns the value that was added to the "request_added_time" field in this mutation.
func (m *ProofRequestMutation) AddedRequestAddedTime() (r int64, exists bool) {
	v := m.addrequest_added_time
	if v == nil {
		return
	}
	return *v, true
}

// ResetRequestAddedTime resets all changes to the "request_added_time" field.
func (m *ProofRequestMutation) ResetRequestAddedTime() {
	m.request_added_time = nil
	m.addrequest_added_time = nil
}

// SetProverRequestID sets the "prover_request_id" field.
func (m *ProofRequestMutation) SetProverRequestID(s string) {
	m.prover_request_id = &s
}

// ProverRequestID returns the value of the "prover_request_id" field in the mutation.
func (m *ProofRequestMutation) ProverRequestID() (r string, exists bool) {
	v := m.prover_request_id
	if v == nil {
		return
	}
	return *v, true
}

// OldProverRequestID returns the old "prover_request_id" field's value of the ProofRequest entity.
// If the ProofRequest object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProofRequestMutation) OldProverRequestID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldProverRequestID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldProverRequestID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldProverRequestID: %w", err)
	}
	return oldValue.ProverRequestID, nil
}

// ClearProverRequestID clears the value of the "prover_request_id" field.
func (m *ProofRequestMutation) ClearProverRequestID() {
	m.prover_request_id = nil
	m.clearedFields[proofrequest.FieldProverRequestID] = struct{}{}
}

// ProverRequestIDCleared returns if the "prover_request_id" field was cleared in this mutation.
func (m *ProofRequestMutation) ProverRequestIDCleared() bool {
	_, ok := m.clearedFields[proofrequest.FieldProverRequestID]
	return ok
}

// ResetProverRequestID resets all changes to the "prover_request_id" field.
func (m *ProofRequestMutation) ResetProverRequestID() {
	m.prover_request_id = nil
	delete(m.clearedFields, proofrequest.FieldProverRequestID)
}

// SetProofRequestTime sets the "proof_request_time" field.
func (m *ProofRequestMutation) SetProofRequestTime(u uint64) {
	m.proof_request_time = &u
	m.addproof_request_time = nil
}

// ProofRequestTime returns the value of the "proof_request_time" field in the mutation.
func (m *ProofRequestMutation) ProofRequestTime() (r uint64, exists bool) {
	v := m.proof_request_time
	if v == nil {
		return
	}
	return *v, true
}

// OldProofRequestTime returns the old "proof_request_time" field's value of the ProofRequest entity.
// If the ProofRequest object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProofRequestMutation) OldProofRequestTime(ctx context.Context) (v uint64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldProofRequestTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldProofRequestTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldProofRequestTime: %w", err)
	}
	return oldValue.ProofRequestTime, nil
}

// AddProofRequestTime adds u to the "proof_request_time" field.
func (m *ProofRequestMutation) AddProofRequestTime(u int64) {
	if m.addproof_request_time != nil {
		*m.addproof_request_time += u
	} else {
		m.addproof_request_time = &u
	}
}

// AddedProofRequestTime returns the value that was added to the "proof_request_time" field in this mutation.
func (m *ProofRequestMutation) AddedProofRequestTime() (r int64, exists bool) {
	v := m.addproof_request_time
	if v == nil {
		return
	}
	return *v, true
}

// ClearProofRequestTime clears the value of the "proof_request_time" field.
func (m *ProofRequestMutation) ClearProofRequestTime() {
	m.proof_request_time = nil
	m.addproof_request_time = nil
	m.clearedFields[proofrequest.FieldProofRequestTime] = struct{}{}
}

// ProofRequestTimeCleared returns if the "proof_request_time" field was cleared in this mutation.
func (m *ProofRequestMutation) ProofRequestTimeCleared() bool {
	_, ok := m.clearedFields[proofrequest.FieldProofRequestTime]
	return ok
}

// ResetProofRequestTime resets all changes to the "proof_request_time" field.
func (m *ProofRequestMutation) ResetProofRequestTime() {
	m.proof_request_time = nil
	m.addproof_request_time = nil
	delete(m.clearedFields, proofrequest.FieldProofRequestTime)
}

// SetL1BlockNumber sets the "l1_block_number" field.
func (m *ProofRequestMutation) SetL1BlockNumber(u uint64) {
	m.l1_block_number = &u
	m.addl1_block_number = nil
}

// L1BlockNumber returns the value of the "l1_block_number" field in the mutation.
func (m *ProofRequestMutation) L1BlockNumber() (r uint64, exists bool) {
	v := m.l1_block_number
	if v == nil {
		return
	}
	return *v, true
}

// OldL1BlockNumber returns the old "l1_block_number" field's value of the ProofRequest entity.
// If the ProofRequest object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProofRequestMutation) OldL1BlockNumber(ctx context.Context) (v uint64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldL1BlockNumber is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldL1BlockNumber requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldL1BlockNumber: %w", err)
	}
	return oldValue.L1BlockNumber, nil
}

// AddL1BlockNumber adds u to the "l1_block_number" field.
func (m *ProofRequestMutation) AddL1BlockNumber(u int64) {
	if m.addl1_block_number != nil {
		*m.addl1_block_number += u
	} else {
		m.addl1_block_number = &u
	}
}

// AddedL1BlockNumber returns the value that was added to the "l1_block_number" field in this mutation.
func (m *ProofRequestMutation) AddedL1BlockNumber() (r int64, exists bool) {
	v := m.addl1_block_number
	if v == nil {
		return
	}
	return *v, true
}

// ClearL1BlockNumber clears the value of the "l1_block_number" field.
func (m *ProofRequestMutation) ClearL1BlockNumber() {
	m.l1_block_number = nil
	m.addl1_block_number = nil
	m.clearedFields[proofrequest.FieldL1BlockNumber] = struct{}{}
}

// L1BlockNumberCleared returns if the "l1_block_number" field was cleared in this mutation.
func (m *ProofRequestMutation) L1BlockNumberCleared() bool {
	_, ok := m.clearedFields[proofrequest.FieldL1BlockNumber]
	return ok
}

// ResetL1BlockNumber resets all changes to the "l1_block_number" field.
func (m *ProofRequestMutation) ResetL1BlockNumber() {
	m.l1_block_number = nil
	m.addl1_block_number = nil
	delete(m.clearedFields, proofrequest.FieldL1BlockNumber)
}

// SetL1BlockHash sets the "l1_block_hash" field.
func (m *ProofRequestMutation) SetL1BlockHash(s string) {
	m.l1_block_hash = &s
}

// L1BlockHash returns the value of the "l1_block_hash" field in the mutation.
func (m *ProofRequestMutation) L1BlockHash() (r string, exists bool) {
	v := m.l1_block_hash
	if v == nil {
		return
	}
	return *v, true
}

// OldL1BlockHash returns the old "l1_block_hash" field's value of the ProofRequest entity.
// If the ProofRequest object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProofRequestMutation) OldL1BlockHash(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldL1BlockHash is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldL1BlockHash requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldL1BlockHash: %w", err)
	}
	return oldValue.L1BlockHash, nil
}

// ClearL1BlockHash clears the value of the "l1_block_hash" field.
func (m *ProofRequestMutation) ClearL1BlockHash() {
	m.l1_block_hash = nil
	m.clearedFields[proofrequest.FieldL1BlockHash] = struct{}{}
}

// L1BlockHashCleared returns if the "l1_block_hash" field was cleared in this mutation.
func (m *ProofRequestMutation) L1BlockHashCleared() bool {
	_, ok := m.clearedFields[proofrequest.FieldL1BlockHash]
	return ok
}

// ResetL1BlockHash resets all changes to the "l1_block_hash" field.
func (m *ProofRequestMutation) ResetL1BlockHash() {
	m.l1_block_hash = nil
	delete(m.clearedFields, proofrequest.FieldL1BlockHash)
}

// SetProof sets the "proof" field.
func (m *ProofRequestMutation) SetProof(b []byte) {
	m.proof = &b
}

// Proof returns the value of the "proof" field in the mutation.
func (m *ProofRequestMutation) Proof() (r []byte, exists bool) {
	v := m.proof
	if v == nil {
		return
	}
	return *v, true
}

// OldProof returns the old "proof" field's value of the ProofRequest entity.
// If the ProofRequest object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProofRequestMutation) OldProof(ctx context.Context) (v []byte, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldProof is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldProof requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldProof: %w", err)
	}
	return oldValue.Proof, nil
}

// ClearProof clears the value of the "proof" field.
func (m *ProofRequestMutation) ClearProof() {
	m.proof = nil
	m.clearedFields[proofrequest.FieldProof] = struct{}{}
}

// ProofCleared returns if the "proof" field was cleared in this mutation.
func (m *ProofRequestMutation) ProofCleared() bool {
	_, ok := m.clearedFields[proofrequest.FieldProof]
	return ok
}

// ResetProof resets all changes to the "proof" field.
func (m *ProofRequestMutation) ResetProof() {
	m.proof = nil
	delete(m.clearedFields, proofrequest.FieldProof)
}

// Where appends a list predicates to the ProofRequestMutation builder.
func (m *ProofRequestMutation) Where(ps ...predicate.ProofRequest) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the ProofRequestMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *ProofRequestMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.ProofRequest, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *ProofRequestMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *ProofRequestMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (ProofRequest).
func (m *ProofRequestMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ProofRequestMutation) Fields() []string {
	fields := make([]string, 0, 10)
	if m._type != nil {
		fields = append(fields, proofrequest.FieldType)
	}
	if m.start_block != nil {
		fields = append(fields, proofrequest.FieldStartBlock)
	}
	if m.end_block != nil {
		fields = append(fields, proofrequest.FieldEndBlock)
	}
	if m.status != nil {
		fields = append(fields, proofrequest.FieldStatus)
	}
	if m.request_added_time != nil {
		fields = append(fields, proofrequest.FieldRequestAddedTime)
	}
	if m.prover_request_id != nil {
		fields = append(fields, proofrequest.FieldProverRequestID)
	}
	if m.proof_request_time != nil {
		fields = append(fields, proofrequest.FieldProofRequestTime)
	}
	if m.l1_block_number != nil {
		fields = append(fields, proofrequest.FieldL1BlockNumber)
	}
	if m.l1_block_hash != nil {
		fields = append(fields, proofrequest.FieldL1BlockHash)
	}
	if m.proof != nil {
		fields = append(fields, proofrequest.FieldProof)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ProofRequestMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case proofrequest.FieldType:
		return m.GetType()
	case proofrequest.FieldStartBlock:
		return m.StartBlock()
	case proofrequest.FieldEndBlock:
		return m.EndBlock()
	case proofrequest.FieldStatus:
		return m.Status()
	case proofrequest.FieldRequestAddedTime:
		return m.RequestAddedTime()
	case proofrequest.FieldProverRequestID:
		return m.ProverRequestID()
	case proofrequest.FieldProofRequestTime:
		return m.ProofRequestTime()
	case proofrequest.FieldL1BlockNumber:
		return m.L1BlockNumber()
	case proofrequest.FieldL1BlockHash:
		return m.L1BlockHash()
	case proofrequest.FieldProof:
		return m.Proof()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ProofRequestMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case proofrequest.FieldType:
		return m.OldType(ctx)
	case proofrequest.FieldStartBlock:
		return m.OldStartBlock(ctx)
	case proofrequest.FieldEndBlock:
		return m.OldEndBlock(ctx)
	case proofrequest.FieldStatus:
		return m.OldStatus(ctx)
	case proofrequest.FieldRequestAddedTime:
		return m.OldRequestAddedTime(ctx)
	case proofrequest.FieldProverRequestID:
		return m.OldProverRequestID(ctx)
	case proofrequest.FieldProofRequestTime:
		return m.OldProofRequestTime(ctx)
	case proofrequest.FieldL1BlockNumber:
		return m.OldL1BlockNumber(ctx)
	case proofrequest.FieldL1BlockHash:
		return m.OldL1BlockHash(ctx)
	case proofrequest.FieldProof:
		return m.OldProof(ctx)
	}
	return nil, fmt.Errorf("unknown ProofRequest field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ProofRequestMutation) SetField(name string, value ent.Value) error {
	switch name {
	case proofrequest.FieldType:
		v, ok := value.(proofrequest.Type)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetType(v)
		return nil
	case proofrequest.FieldStartBlock:
		v, ok := value.(uint64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStartBlock(v)
		return nil
	case proofrequest.FieldEndBlock:
		v, ok := value.(uint64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEndBlock(v)
		return nil
	case proofrequest.FieldStatus:
		v, ok := value.(proofrequest.Status)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	case proofrequest.FieldRequestAddedTime:
		v, ok := value.(uint64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRequestAddedTime(v)
		return nil
	case proofrequest.FieldProverRequestID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetProverRequestID(v)
		return nil
	case proofrequest.FieldProofRequestTime:
		v, ok := value.(uint64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetProofRequestTime(v)
		return nil
	case proofrequest.FieldL1BlockNumber:
		v, ok := value.(uint64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetL1BlockNumber(v)
		return nil
	case proofrequest.FieldL1BlockHash:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetL1BlockHash(v)
		return nil
	case proofrequest.FieldProof:
		v, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetProof(v)
		return nil
	}
	return fmt.Errorf("unknown ProofRequest field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ProofRequestMutation) AddedFields() []string {
	var fields []string
	if m.addstart_block != nil {
		fields = append(fields, proofrequest.FieldStartBlock)
	}
	if m.addend_block != nil {
		fields = append(fields, proofrequest.FieldEndBlock)
	}
	if m.addrequest_added_time != nil {
		fields = append(fields, proofrequest.FieldRequestAddedTime)
	}
	if m.addproof_request_time != nil {
		fields = append(fields, proofrequest.FieldProofRequestTime)
	}
	if m.addl1_block_number != nil {
		fields = append(fields, proofrequest.FieldL1BlockNumber)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ProofRequestMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case proofrequest.FieldStartBlock:
		return m.AddedStartBlock()
	case proofrequest.FieldEndBlock:
		return m.AddedEndBlock()
	case proofrequest.FieldRequestAddedTime:
		return m.AddedRequestAddedTime()
	case proofrequest.FieldProofRequestTime:
		return m.AddedProofRequestTime()
	case proofrequest.FieldL1BlockNumber:
		return m.AddedL1BlockNumber()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ProofRequestMutation) AddField(name string, value ent.Value) error {
	switch name {
	case proofrequest.FieldStartBlock:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddStartBlock(v)
		return nil
	case proofrequest.FieldEndBlock:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddEndBlock(v)
		return nil
	case proofrequest.FieldRequestAddedTime:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddRequestAddedTime(v)
		return nil
	case proofrequest.FieldProofRequestTime:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddProofRequestTime(v)
		return nil
	case proofrequest.FieldL1BlockNumber:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddL1BlockNumber(v)
		return nil
	}
	return fmt.Errorf("unknown ProofRequest numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ProofRequestMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(proofrequest.FieldProverRequestID) {
		fields = append(fields, proofrequest.FieldProverRequestID)
	}
	if m.FieldCleared(proofrequest.FieldProofRequestTime) {
		fields = append(fields, proofrequest.FieldProofRequestTime)
	}
	if m.FieldCleared(proofrequest.FieldL1BlockNumber) {
		fields = append(fields, proofrequest.FieldL1BlockNumber)
	}
	if m.FieldCleared(proofrequest.FieldL1BlockHash) {
		fields = append(fields, proofrequest.FieldL1BlockHash)
	}
	if m.FieldCleared(proofrequest.FieldProof) {
		fields = append(fields, proofrequest.FieldProof)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ProofRequestMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ProofRequestMutation) ClearField(name string) error {
	switch name {
	case proofrequest.FieldProverRequestID:
		m.ClearProverRequestID()
		return nil
	case proofrequest.FieldProofRequestTime:
		m.ClearProofRequestTime()
		return nil
	case proofrequest.FieldL1BlockNumber:
		m.ClearL1BlockNumber()
		return nil
	case proofrequest.FieldL1BlockHash:
		m.ClearL1BlockHash()
		return nil
	case proofrequest.FieldProof:
		m.ClearProof()
		return nil
	}
	return fmt.Errorf("unknown ProofRequest nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ProofRequestMutation) ResetField(name string) error {
	switch name {
	case proofrequest.FieldType:
		m.ResetType()
		return nil
	case proofrequest.FieldStartBlock:
		m.ResetStartBlock()
		return nil
	case proofrequest.FieldEndBlock:
		m.ResetEndBlock()
		return nil
	case proofrequest.FieldStatus:
		m.ResetStatus()
		return nil
	case proofrequest.FieldRequestAddedTime:
		m.ResetRequestAddedTime()
		return nil
	case proofrequest.FieldProverRequestID:
		m.ResetProverRequestID()
		return nil
	case proofrequest.FieldProofRequestTime:
		m.ResetProofRequestTime()
		return nil
	case proofrequest.FieldL1BlockNumber:
		m.ResetL1BlockNumber()
		return nil
	case proofrequest.FieldL1BlockHash:
		m.ResetL1BlockHash()
		return nil
	case proofrequest.FieldProof:
		m.ResetProof()
		return nil
	}
	return fmt.Errorf("unknown ProofRequest field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ProofRequestMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ProofRequestMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ProofRequestMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ProofRequestMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ProofRequestMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ProofRequestMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ProofRequestMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown ProofRequest unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ProofRequestMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown ProofRequest edge %s", name)
}
