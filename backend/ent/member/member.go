// Code generated by entc, DO NOT EDIT.

package member

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

const (
	// Label holds the string label denoting the member type in the database.
	Label = "member"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldNumber holds the string denoting the number field in the database.
	FieldNumber = "number"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldKana holds the string denoting the kana field in the database.
	FieldKana = "kana"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldDateOfBirth holds the string denoting the date_of_birth field in the database.
	FieldDateOfBirth = "date_of_birth"
	// FieldPhoneNumber holds the string denoting the phone_number field in the database.
	FieldPhoneNumber = "phone_number"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldDateOfAdmission holds the string denoting the date_of_admission field in the database.
	FieldDateOfAdmission = "date_of_admission"
	// FieldDateOfWithdrawal holds the string denoting the date_of_withdrawal field in the database.
	FieldDateOfWithdrawal = "date_of_withdrawal"
	// FieldMemo holds the string denoting the memo field in the database.
	FieldMemo = "memo"
	// EdgeMembersClasses holds the string denoting the members_classes edge name in mutations.
	EdgeMembersClasses = "members_classes"
	// Table holds the table name of the member in the database.
	Table = "members"
	// MembersClassesTable is the table that holds the members_classes relation/edge.
	MembersClassesTable = "members_classes"
	// MembersClassesInverseTable is the table name for the MembersClass entity.
	// It exists in this package in order to avoid circular dependency with the "membersclass" package.
	MembersClassesInverseTable = "members_classes"
	// MembersClassesColumn is the table column denoting the members_classes relation/edge.
	MembersClassesColumn = "member_members_classes"
)

// Columns holds all SQL columns for member fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldNumber,
	FieldName,
	FieldKana,
	FieldGender,
	FieldDateOfBirth,
	FieldPhoneNumber,
	FieldEmail,
	FieldDateOfAdmission,
	FieldDateOfWithdrawal,
	FieldMemo,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)

// Gender defines the type for the "gender" enum field.
type Gender string

// Gender values.
const (
	GenderMALE   Gender = "MALE"
	GenderFEMALE Gender = "FEMALE"
	GenderOTHER  Gender = "OTHER"
)

func (ge Gender) String() string {
	return string(ge)
}

// GenderValidator is a validator for the "gender" field enum values. It is called by the builders before save.
func GenderValidator(ge Gender) error {
	switch ge {
	case GenderMALE, GenderFEMALE, GenderOTHER:
		return nil
	default:
		return fmt.Errorf("member: invalid enum value for gender field: %q", ge)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (ge Gender) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(ge.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (ge *Gender) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*ge = Gender(str)
	if err := GenderValidator(*ge); err != nil {
		return fmt.Errorf("%s is not a valid Gender", str)
	}
	return nil
}
