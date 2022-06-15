// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/rs/xid"
	"github.com/softilium/mb4/ent/investaccount"
	"github.com/softilium/mb4/ent/investaccountcashflow"
	"github.com/softilium/mb4/ent/investaccountvaluation"
	"github.com/softilium/mb4/ent/schema"
	"github.com/softilium/mb4/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	investaccountFields := schema.InvestAccount{}.Fields()
	_ = investaccountFields
	// investaccountDescDescr is the schema descriptor for Descr field.
	investaccountDescDescr := investaccountFields[1].Descriptor()
	// investaccount.DescrValidator is a validator for the "Descr" field. It is called by the builders before save.
	investaccount.DescrValidator = func() func(string) error {
		validators := investaccountDescDescr.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(_Descr string) error {
			for _, fn := range fns {
				if err := fn(_Descr); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// investaccountDescID is the schema descriptor for id field.
	investaccountDescID := investaccountFields[0].Descriptor()
	// investaccount.DefaultID holds the default value on creation for the id field.
	investaccount.DefaultID = investaccountDescID.Default.(func() xid.ID)
	investaccountcashflowFields := schema.InvestAccountCashflow{}.Fields()
	_ = investaccountcashflowFields
	// investaccountcashflowDescID is the schema descriptor for id field.
	investaccountcashflowDescID := investaccountcashflowFields[0].Descriptor()
	// investaccountcashflow.DefaultID holds the default value on creation for the id field.
	investaccountcashflow.DefaultID = investaccountcashflowDescID.Default.(func() xid.ID)
	investaccountvaluationFields := schema.InvestAccountValuation{}.Fields()
	_ = investaccountvaluationFields
	// investaccountvaluationDescID is the schema descriptor for id field.
	investaccountvaluationDescID := investaccountvaluationFields[0].Descriptor()
	// investaccountvaluation.DefaultID holds the default value on creation for the id field.
	investaccountvaluation.DefaultID = investaccountvaluationDescID.Default.(func() xid.ID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUserName is the schema descriptor for UserName field.
	userDescUserName := userFields[1].Descriptor()
	// user.UserNameValidator is a validator for the "UserName" field. It is called by the builders before save.
	user.UserNameValidator = func() func(string) error {
		validators := userDescUserName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(_UserName string) error {
			for _, fn := range fns {
				if err := fn(_UserName); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescAuthType is the schema descriptor for AuthType field.
	userDescAuthType := userFields[2].Descriptor()
	// user.DefaultAuthType holds the default value on creation for the AuthType field.
	user.DefaultAuthType = userDescAuthType.Default.(int32)
	// user.AuthTypeValidator is a validator for the "AuthType" field. It is called by the builders before save.
	user.AuthTypeValidator = userDescAuthType.Validators[0].(func(int32) error)
	// userDescAdmin is the schema descriptor for Admin field.
	userDescAdmin := userFields[4].Descriptor()
	// user.DefaultAdmin holds the default value on creation for the Admin field.
	user.DefaultAdmin = userDescAdmin.Default.(bool)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() xid.ID)
}
