// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/succinctlabs/op-succinct-go/proposer/db/ent/proofrequest"
)

// ProofRequestCreate is the builder for creating a ProofRequest entity.
type ProofRequestCreate struct {
	config
	mutation *ProofRequestMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (prc *ProofRequestCreate) SetType(pr proofrequest.Type) *ProofRequestCreate {
	prc.mutation.SetType(pr)
	return prc
}

// SetStartBlock sets the "start_block" field.
func (prc *ProofRequestCreate) SetStartBlock(u uint64) *ProofRequestCreate {
	prc.mutation.SetStartBlock(u)
	return prc
}

// SetEndBlock sets the "end_block" field.
func (prc *ProofRequestCreate) SetEndBlock(u uint64) *ProofRequestCreate {
	prc.mutation.SetEndBlock(u)
	return prc
}

// SetStatus sets the "status" field.
func (prc *ProofRequestCreate) SetStatus(pr proofrequest.Status) *ProofRequestCreate {
	prc.mutation.SetStatus(pr)
	return prc
}

// SetRequestAddedTime sets the "request_added_time" field.
func (prc *ProofRequestCreate) SetRequestAddedTime(u uint64) *ProofRequestCreate {
	prc.mutation.SetRequestAddedTime(u)
	return prc
}

// SetProverRequestID sets the "prover_request_id" field.
func (prc *ProofRequestCreate) SetProverRequestID(s string) *ProofRequestCreate {
	prc.mutation.SetProverRequestID(s)
	return prc
}

// SetNillableProverRequestID sets the "prover_request_id" field if the given value is not nil.
func (prc *ProofRequestCreate) SetNillableProverRequestID(s *string) *ProofRequestCreate {
	if s != nil {
		prc.SetProverRequestID(*s)
	}
	return prc
}

// SetProofRequestTime sets the "proof_request_time" field.
func (prc *ProofRequestCreate) SetProofRequestTime(u uint64) *ProofRequestCreate {
	prc.mutation.SetProofRequestTime(u)
	return prc
}

// SetNillableProofRequestTime sets the "proof_request_time" field if the given value is not nil.
func (prc *ProofRequestCreate) SetNillableProofRequestTime(u *uint64) *ProofRequestCreate {
	if u != nil {
		prc.SetProofRequestTime(*u)
	}
	return prc
}

// SetL1BlockNumber sets the "l1_block_number" field.
func (prc *ProofRequestCreate) SetL1BlockNumber(u uint64) *ProofRequestCreate {
	prc.mutation.SetL1BlockNumber(u)
	return prc
}

// SetNillableL1BlockNumber sets the "l1_block_number" field if the given value is not nil.
func (prc *ProofRequestCreate) SetNillableL1BlockNumber(u *uint64) *ProofRequestCreate {
	if u != nil {
		prc.SetL1BlockNumber(*u)
	}
	return prc
}

// SetL1BlockHash sets the "l1_block_hash" field.
func (prc *ProofRequestCreate) SetL1BlockHash(s string) *ProofRequestCreate {
	prc.mutation.SetL1BlockHash(s)
	return prc
}

// SetNillableL1BlockHash sets the "l1_block_hash" field if the given value is not nil.
func (prc *ProofRequestCreate) SetNillableL1BlockHash(s *string) *ProofRequestCreate {
	if s != nil {
		prc.SetL1BlockHash(*s)
	}
	return prc
}

// SetProof sets the "proof" field.
func (prc *ProofRequestCreate) SetProof(b []byte) *ProofRequestCreate {
	prc.mutation.SetProof(b)
	return prc
}

// Mutation returns the ProofRequestMutation object of the builder.
func (prc *ProofRequestCreate) Mutation() *ProofRequestMutation {
	return prc.mutation
}

// Save creates the ProofRequest in the database.
func (prc *ProofRequestCreate) Save(ctx context.Context) (*ProofRequest, error) {
	return withHooks(ctx, prc.sqlSave, prc.mutation, prc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (prc *ProofRequestCreate) SaveX(ctx context.Context) *ProofRequest {
	v, err := prc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (prc *ProofRequestCreate) Exec(ctx context.Context) error {
	_, err := prc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (prc *ProofRequestCreate) ExecX(ctx context.Context) {
	if err := prc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (prc *ProofRequestCreate) check() error {
	if _, ok := prc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "ProofRequest.type"`)}
	}
	if v, ok := prc.mutation.GetType(); ok {
		if err := proofrequest.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "ProofRequest.type": %w`, err)}
		}
	}
	if _, ok := prc.mutation.StartBlock(); !ok {
		return &ValidationError{Name: "start_block", err: errors.New(`ent: missing required field "ProofRequest.start_block"`)}
	}
	if _, ok := prc.mutation.EndBlock(); !ok {
		return &ValidationError{Name: "end_block", err: errors.New(`ent: missing required field "ProofRequest.end_block"`)}
	}
	if _, ok := prc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "ProofRequest.status"`)}
	}
	if v, ok := prc.mutation.Status(); ok {
		if err := proofrequest.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "ProofRequest.status": %w`, err)}
		}
	}
	if _, ok := prc.mutation.RequestAddedTime(); !ok {
		return &ValidationError{Name: "request_added_time", err: errors.New(`ent: missing required field "ProofRequest.request_added_time"`)}
	}
	return nil
}

func (prc *ProofRequestCreate) sqlSave(ctx context.Context) (*ProofRequest, error) {
	if err := prc.check(); err != nil {
		return nil, err
	}
	_node, _spec := prc.createSpec()
	if err := sqlgraph.CreateNode(ctx, prc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	prc.mutation.id = &_node.ID
	prc.mutation.done = true
	return _node, nil
}

func (prc *ProofRequestCreate) createSpec() (*ProofRequest, *sqlgraph.CreateSpec) {
	var (
		_node = &ProofRequest{config: prc.config}
		_spec = sqlgraph.NewCreateSpec(proofrequest.Table, sqlgraph.NewFieldSpec(proofrequest.FieldID, field.TypeInt))
	)
	if value, ok := prc.mutation.GetType(); ok {
		_spec.SetField(proofrequest.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := prc.mutation.StartBlock(); ok {
		_spec.SetField(proofrequest.FieldStartBlock, field.TypeUint64, value)
		_node.StartBlock = value
	}
	if value, ok := prc.mutation.EndBlock(); ok {
		_spec.SetField(proofrequest.FieldEndBlock, field.TypeUint64, value)
		_node.EndBlock = value
	}
	if value, ok := prc.mutation.Status(); ok {
		_spec.SetField(proofrequest.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := prc.mutation.RequestAddedTime(); ok {
		_spec.SetField(proofrequest.FieldRequestAddedTime, field.TypeUint64, value)
		_node.RequestAddedTime = value
	}
	if value, ok := prc.mutation.ProverRequestID(); ok {
		_spec.SetField(proofrequest.FieldProverRequestID, field.TypeString, value)
		_node.ProverRequestID = value
	}
	if value, ok := prc.mutation.ProofRequestTime(); ok {
		_spec.SetField(proofrequest.FieldProofRequestTime, field.TypeUint64, value)
		_node.ProofRequestTime = value
	}
	if value, ok := prc.mutation.L1BlockNumber(); ok {
		_spec.SetField(proofrequest.FieldL1BlockNumber, field.TypeUint64, value)
		_node.L1BlockNumber = value
	}
	if value, ok := prc.mutation.L1BlockHash(); ok {
		_spec.SetField(proofrequest.FieldL1BlockHash, field.TypeString, value)
		_node.L1BlockHash = value
	}
	if value, ok := prc.mutation.Proof(); ok {
		_spec.SetField(proofrequest.FieldProof, field.TypeBytes, value)
		_node.Proof = value
	}
	return _node, _spec
}

// ProofRequestCreateBulk is the builder for creating many ProofRequest entities in bulk.
type ProofRequestCreateBulk struct {
	config
	err      error
	builders []*ProofRequestCreate
}

// Save creates the ProofRequest entities in the database.
func (prcb *ProofRequestCreateBulk) Save(ctx context.Context) ([]*ProofRequest, error) {
	if prcb.err != nil {
		return nil, prcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(prcb.builders))
	nodes := make([]*ProofRequest, len(prcb.builders))
	mutators := make([]Mutator, len(prcb.builders))
	for i := range prcb.builders {
		func(i int, root context.Context) {
			builder := prcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProofRequestMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, prcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, prcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, prcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (prcb *ProofRequestCreateBulk) SaveX(ctx context.Context) []*ProofRequest {
	v, err := prcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (prcb *ProofRequestCreateBulk) Exec(ctx context.Context) error {
	_, err := prcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (prcb *ProofRequestCreateBulk) ExecX(ctx context.Context) {
	if err := prcb.Exec(ctx); err != nil {
		panic(err)
	}
}
