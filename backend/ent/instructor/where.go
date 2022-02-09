// Code generated by entc, DO NOT EDIT.

package instructor

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/matsuokashuhei/landin/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// SyllabicCharacters applies equality check predicate on the "syllabic_characters" field. It's identical to SyllabicCharactersEQ.
func SyllabicCharacters(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSyllabicCharacters), v))
	})
}

// Biography applies equality check predicate on the "biography" field. It's identical to BiographyEQ.
func Biography(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBiography), v))
	})
}

// PhoneNumber applies equality check predicate on the "phone_number" field. It's identical to PhoneNumberEQ.
func PhoneNumber(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhoneNumber), v))
	})
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Instructor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instructor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Instructor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instructor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Instructor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instructor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Instructor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instructor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Instructor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instructor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Instructor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instructor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// SyllabicCharactersEQ applies the EQ predicate on the "syllabic_characters" field.
func SyllabicCharactersEQ(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSyllabicCharacters), v))
	})
}

// SyllabicCharactersNEQ applies the NEQ predicate on the "syllabic_characters" field.
func SyllabicCharactersNEQ(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSyllabicCharacters), v))
	})
}

// SyllabicCharactersIn applies the In predicate on the "syllabic_characters" field.
func SyllabicCharactersIn(vs ...string) predicate.Instructor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instructor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSyllabicCharacters), v...))
	})
}

// SyllabicCharactersNotIn applies the NotIn predicate on the "syllabic_characters" field.
func SyllabicCharactersNotIn(vs ...string) predicate.Instructor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instructor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSyllabicCharacters), v...))
	})
}

// SyllabicCharactersGT applies the GT predicate on the "syllabic_characters" field.
func SyllabicCharactersGT(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSyllabicCharacters), v))
	})
}

// SyllabicCharactersGTE applies the GTE predicate on the "syllabic_characters" field.
func SyllabicCharactersGTE(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSyllabicCharacters), v))
	})
}

// SyllabicCharactersLT applies the LT predicate on the "syllabic_characters" field.
func SyllabicCharactersLT(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSyllabicCharacters), v))
	})
}

// SyllabicCharactersLTE applies the LTE predicate on the "syllabic_characters" field.
func SyllabicCharactersLTE(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSyllabicCharacters), v))
	})
}

// SyllabicCharactersContains applies the Contains predicate on the "syllabic_characters" field.
func SyllabicCharactersContains(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSyllabicCharacters), v))
	})
}

// SyllabicCharactersHasPrefix applies the HasPrefix predicate on the "syllabic_characters" field.
func SyllabicCharactersHasPrefix(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSyllabicCharacters), v))
	})
}

// SyllabicCharactersHasSuffix applies the HasSuffix predicate on the "syllabic_characters" field.
func SyllabicCharactersHasSuffix(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSyllabicCharacters), v))
	})
}

// SyllabicCharactersEqualFold applies the EqualFold predicate on the "syllabic_characters" field.
func SyllabicCharactersEqualFold(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSyllabicCharacters), v))
	})
}

// SyllabicCharactersContainsFold applies the ContainsFold predicate on the "syllabic_characters" field.
func SyllabicCharactersContainsFold(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSyllabicCharacters), v))
	})
}

// BiographyEQ applies the EQ predicate on the "biography" field.
func BiographyEQ(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBiography), v))
	})
}

// BiographyNEQ applies the NEQ predicate on the "biography" field.
func BiographyNEQ(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBiography), v))
	})
}

// BiographyIn applies the In predicate on the "biography" field.
func BiographyIn(vs ...string) predicate.Instructor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instructor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBiography), v...))
	})
}

// BiographyNotIn applies the NotIn predicate on the "biography" field.
func BiographyNotIn(vs ...string) predicate.Instructor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instructor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBiography), v...))
	})
}

// BiographyGT applies the GT predicate on the "biography" field.
func BiographyGT(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBiography), v))
	})
}

// BiographyGTE applies the GTE predicate on the "biography" field.
func BiographyGTE(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBiography), v))
	})
}

// BiographyLT applies the LT predicate on the "biography" field.
func BiographyLT(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBiography), v))
	})
}

// BiographyLTE applies the LTE predicate on the "biography" field.
func BiographyLTE(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBiography), v))
	})
}

// BiographyContains applies the Contains predicate on the "biography" field.
func BiographyContains(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBiography), v))
	})
}

// BiographyHasPrefix applies the HasPrefix predicate on the "biography" field.
func BiographyHasPrefix(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBiography), v))
	})
}

// BiographyHasSuffix applies the HasSuffix predicate on the "biography" field.
func BiographyHasSuffix(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBiography), v))
	})
}

// BiographyIsNil applies the IsNil predicate on the "biography" field.
func BiographyIsNil() predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldBiography)))
	})
}

// BiographyNotNil applies the NotNil predicate on the "biography" field.
func BiographyNotNil() predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldBiography)))
	})
}

// BiographyEqualFold applies the EqualFold predicate on the "biography" field.
func BiographyEqualFold(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBiography), v))
	})
}

// BiographyContainsFold applies the ContainsFold predicate on the "biography" field.
func BiographyContainsFold(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBiography), v))
	})
}

// PhoneNumberEQ applies the EQ predicate on the "phone_number" field.
func PhoneNumberEQ(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhoneNumber), v))
	})
}

// PhoneNumberNEQ applies the NEQ predicate on the "phone_number" field.
func PhoneNumberNEQ(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPhoneNumber), v))
	})
}

// PhoneNumberIn applies the In predicate on the "phone_number" field.
func PhoneNumberIn(vs ...string) predicate.Instructor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instructor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPhoneNumber), v...))
	})
}

// PhoneNumberNotIn applies the NotIn predicate on the "phone_number" field.
func PhoneNumberNotIn(vs ...string) predicate.Instructor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instructor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPhoneNumber), v...))
	})
}

// PhoneNumberGT applies the GT predicate on the "phone_number" field.
func PhoneNumberGT(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPhoneNumber), v))
	})
}

// PhoneNumberGTE applies the GTE predicate on the "phone_number" field.
func PhoneNumberGTE(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPhoneNumber), v))
	})
}

// PhoneNumberLT applies the LT predicate on the "phone_number" field.
func PhoneNumberLT(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPhoneNumber), v))
	})
}

// PhoneNumberLTE applies the LTE predicate on the "phone_number" field.
func PhoneNumberLTE(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPhoneNumber), v))
	})
}

// PhoneNumberContains applies the Contains predicate on the "phone_number" field.
func PhoneNumberContains(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPhoneNumber), v))
	})
}

// PhoneNumberHasPrefix applies the HasPrefix predicate on the "phone_number" field.
func PhoneNumberHasPrefix(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPhoneNumber), v))
	})
}

// PhoneNumberHasSuffix applies the HasSuffix predicate on the "phone_number" field.
func PhoneNumberHasSuffix(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPhoneNumber), v))
	})
}

// PhoneNumberIsNil applies the IsNil predicate on the "phone_number" field.
func PhoneNumberIsNil() predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPhoneNumber)))
	})
}

// PhoneNumberNotNil applies the NotNil predicate on the "phone_number" field.
func PhoneNumberNotNil() predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPhoneNumber)))
	})
}

// PhoneNumberEqualFold applies the EqualFold predicate on the "phone_number" field.
func PhoneNumberEqualFold(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPhoneNumber), v))
	})
}

// PhoneNumberContainsFold applies the ContainsFold predicate on the "phone_number" field.
func PhoneNumberContainsFold(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPhoneNumber), v))
	})
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEmail), v))
	})
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Instructor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instructor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEmail), v...))
	})
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Instructor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Instructor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEmail), v...))
	})
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEmail), v))
	})
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEmail), v))
	})
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEmail), v))
	})
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEmail), v))
	})
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEmail), v))
	})
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEmail), v))
	})
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEmail), v))
	})
}

// EmailIsNil applies the IsNil predicate on the "email" field.
func EmailIsNil() predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEmail)))
	})
}

// EmailNotNil applies the NotNil predicate on the "email" field.
func EmailNotNil() predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEmail)))
	})
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEmail), v))
	})
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEmail), v))
	})
}

// HasClasses applies the HasEdge predicate on the "classes" edge.
func HasClasses() predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClassesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ClassesTable, ClassesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClassesWith applies the HasEdge predicate on the "classes" edge with a given conditions (other predicates).
func HasClassesWith(preds ...predicate.Class) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClassesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ClassesTable, ClassesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Instructor) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Instructor) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Instructor) predicate.Instructor {
	return predicate.Instructor(func(s *sql.Selector) {
		p(s.Not())
	})
}