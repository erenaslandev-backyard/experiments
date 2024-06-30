// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"demo/ent/instructor"
	"demo/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// InstructorDelete is the builder for deleting a Instructor entity.
type InstructorDelete struct {
	config
	hooks    []Hook
	mutation *InstructorMutation
}

// Where appends a list predicates to the InstructorDelete builder.
func (id *InstructorDelete) Where(ps ...predicate.Instructor) *InstructorDelete {
	id.mutation.Where(ps...)
	return id
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (id *InstructorDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, id.sqlExec, id.mutation, id.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (id *InstructorDelete) ExecX(ctx context.Context) int {
	n, err := id.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (id *InstructorDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(instructor.Table, sqlgraph.NewFieldSpec(instructor.FieldID, field.TypeInt))
	if ps := id.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, id.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	id.mutation.done = true
	return affected, err
}

// InstructorDeleteOne is the builder for deleting a single Instructor entity.
type InstructorDeleteOne struct {
	id *InstructorDelete
}

// Where appends a list predicates to the InstructorDelete builder.
func (ido *InstructorDeleteOne) Where(ps ...predicate.Instructor) *InstructorDeleteOne {
	ido.id.mutation.Where(ps...)
	return ido
}

// Exec executes the deletion query.
func (ido *InstructorDeleteOne) Exec(ctx context.Context) error {
	n, err := ido.id.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{instructor.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ido *InstructorDeleteOne) ExecX(ctx context.Context) {
	if err := ido.Exec(ctx); err != nil {
		panic(err)
	}
}
