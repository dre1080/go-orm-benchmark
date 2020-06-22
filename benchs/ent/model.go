// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/dre1080/go-orm-benchmark/benchs/ent/model"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Model is the model entity for the Model schema.
type Model struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Fax holds the value of the "fax" field.
	Fax string `json:"fax,omitempty"`
	// Web holds the value of the "web" field.
	Web string `json:"web,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Right holds the value of the "right" field.
	Right bool `json:"right,omitempty"`
	// Counter holds the value of the "counter" field.
	Counter int64 `json:"counter,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Model) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&sql.NullString{}, // title
		&sql.NullString{}, // fax
		&sql.NullString{}, // web
		&sql.NullInt64{},  // age
		&sql.NullBool{},   // right
		&sql.NullInt64{},  // counter
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Model fields.
func (m *Model) assignValues(values ...interface{}) error {
	if m, n := len(values), len(model.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	m.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		m.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field title", values[1])
	} else if value.Valid {
		m.Title = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field fax", values[2])
	} else if value.Valid {
		m.Fax = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field web", values[3])
	} else if value.Valid {
		m.Web = value.String
	}
	if value, ok := values[4].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field age", values[4])
	} else if value.Valid {
		m.Age = int(value.Int64)
	}
	if value, ok := values[5].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field right", values[5])
	} else if value.Valid {
		m.Right = value.Bool
	}
	if value, ok := values[6].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field counter", values[6])
	} else if value.Valid {
		m.Counter = value.Int64
	}
	return nil
}

// Update returns a builder for updating this Model.
// Note that, you need to call Model.Unwrap() before calling this method, if this Model
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Model) Update() *ModelUpdateOne {
	return (&ModelClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (m *Model) Unwrap() *Model {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Model is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Model) String() string {
	var builder strings.Builder
	builder.WriteString("Model(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", name=")
	builder.WriteString(m.Name)
	builder.WriteString(", title=")
	builder.WriteString(m.Title)
	builder.WriteString(", fax=")
	builder.WriteString(m.Fax)
	builder.WriteString(", web=")
	builder.WriteString(m.Web)
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", m.Age))
	builder.WriteString(", right=")
	builder.WriteString(fmt.Sprintf("%v", m.Right))
	builder.WriteString(", counter=")
	builder.WriteString(fmt.Sprintf("%v", m.Counter))
	builder.WriteByte(')')
	return builder.String()
}

// Models is a parsable slice of Model.
type Models []*Model

func (m Models) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}