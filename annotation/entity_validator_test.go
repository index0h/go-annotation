package annotation

import (
	"fmt"
	"go/scanner"
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestNewEntityValidator(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := &EntityValidator{}

	actual := NewEntityValidator()

	ctrl.AssertEqual(expected, actual)
}

func TestEntityValidator_Validate_WithSimpleSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
		IsPointer:   true,
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithSimpleSpecAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &SimpleSpec{
		TypeName: "typeName",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithSimpleSpecAndEmptyTypeName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &SimpleSpec{
		TypeName: "",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'TypeName' must be not empty", actual.Error())
}

func TestEntityValidator_Validate_WithSimpleSpecAndInvalidTypeName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &SimpleSpec{
		TypeName: "+invalid",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'TypeName' must be valid identifier, actual value: '+invalid'", actual.Error())
}

func TestEntityValidator_Validate_WithSimpleSpecAndInvalidPackageName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &SimpleSpec{
		TypeName:    "typeName",
		PackageName: "+invalid",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'PackageName' must be valid identifier, actual value: '+invalid'", actual.Error())
}

func TestEntityValidator_Validate_WithArraySpecAndSimpleSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
		Length: "10",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithArraySpecAndArraySpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &ArraySpec{
		Value: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithArraySpecAndMapSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &ArraySpec{
		Value: &MapSpec{
			Key: &SimpleSpec{
				TypeName: "typeName",
			},
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithArraySpecAndStructSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &ArraySpec{
		Value: &StructSpec{},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithArraySpecAndInterfaceSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &ArraySpec{
		Value: &InterfaceSpec{},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithArraySpecAndFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &ArraySpec{
		Value: &FuncSpec{},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithArraySpecAndEllipsisLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
		Length: "...",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithArraySpecAndNilValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &ArraySpec{}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Value' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithArraySpecAndInvalidLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
		Length: "[+invalid",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertType(actual, scanner.ErrorList{})
}

func TestEntityValidator_Validate_WithArraySpecAndInvalidSimpleSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "+invalid",
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'TypeName' must be valid identifier, actual value: '+invalid'", actual.Error())
}

func TestEntityValidator_Validate_WithArraySpecAndInvalidArraySpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &ArraySpec{
		Value: &ArraySpec{
			Value: nil,
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Value' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithArraySpecAndInvalidMapSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &ArraySpec{
		Value: &MapSpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Key' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithArraySpecAndInvalidStructSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &ArraySpec{
		Value: &StructSpec{
			Fields: []*Field{nil},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Fields[0]' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithArraySpecAndInvalidInterfaceSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &ArraySpec{
		Value: &InterfaceSpec{
			Fields: []*Field{nil},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Fields[0]' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithArraySpecAndInvalidFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &ArraySpec{
		Value: &FuncSpec{
			Params: []*Field{nil},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Params[0]' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithArraySpecAndInvalidTypeValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &ArraySpec{
		Value: "invalid",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(fmt.Sprintf("Variable 'Value' has invalid type: '%T'", entity.Value), actual.Error())
}

func TestEntityValidator_Validate_WithMapSpecAndSimpleSpecKey(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithMapSpecAndArraySpecKey(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &MapSpec{
		Key: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithMapSpecAndMapSpecKey(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &MapSpec{
		Key: &MapSpec{
			Key: &SimpleSpec{
				TypeName: "typeName",
			},
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithMapSpecAndStructSpecKey(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &MapSpec{
		Key: &StructSpec{},
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithMapSpecAndInterfaceSpecKey(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &MapSpec{
		Key: &InterfaceSpec{},
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithMapSpecAndFuncSpecKey(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &MapSpec{
		Key: &FuncSpec{},
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithMapSpecAndSimpleSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithMapSpecAndArraySpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithMapSpecAndMapSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: &MapSpec{
			Key: &SimpleSpec{
				TypeName: "typeName",
			},
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithMapSpecAndStructSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: &StructSpec{},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithMapSpecAndInterfaceSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: &InterfaceSpec{},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithMapSpecAndFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: &FuncSpec{},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithMapSpecAndNilKey(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &MapSpec{
		Key: nil,
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Key' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithMapSpecAndNilValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: nil,
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Value' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithMapSpecAndInvalidKey(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "+invalid",
		},
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'TypeName' must be valid identifier, actual value: '+invalid'", actual.Error())
}

func TestEntityValidator_Validate_WithMapSpecAndInvalidValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: &SimpleSpec{
			TypeName: "+invalid",
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'TypeName' must be valid identifier, actual value: '+invalid'", actual.Error())
}

func TestEntityValidator_Validate_WithMapSpecAndInvalidKeyType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &MapSpec{
		Key: "invalid",
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Key' has invalid type: 'string'", actual.Error())
}

func TestEntityValidator_Validate_WithMapSpecAndInvalidValueType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: "invalid",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Value' has invalid type: 'string'", actual.Error())
}

func TestEntityValidator_Validate_WithField(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Field{
		Name: "fieldName",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithFieldAndSimpleSpecSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Field{
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithFieldAndArraySpecSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Field{
		Spec: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithFieldAndMapSpecSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Field{
		Spec: &MapSpec{
			Key: &SimpleSpec{
				TypeName: "typeName",
			},
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithFieldAndStructSpecSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Field{
		Spec: &StructSpec{},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithFieldAndInterfaceSpecSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Field{
		Spec: &InterfaceSpec{},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithFieldAndFuncSpecSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Field{
		Spec: &FuncSpec{},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithFieldAndInvalidName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Field{
		Name: "+invalid",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Name' must be valid identifier, actual value: '+invalid'", actual.Error())
}

func TestEntityValidator_Validate_WithFieldAndNilSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Field{}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Spec' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithFieldAndInvalidSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Field{
		Spec: &SimpleSpec{
			TypeName: "+invalid",
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'TypeName' must be valid identifier, actual value: '+invalid'", actual.Error())
}

func TestEntityValidator_Validate_WithFieldAndInvalidSpecType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Field{
		Spec: "invalid",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Spec' has invalid type: 'string'", actual.Error())
}

func TestEntityValidator_Validate_WithFuncSpecAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &FuncSpec{}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithFuncSpecAndSimpleSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &FuncSpec{
		Params: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
			},
		},
		Results: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithFuncSpecAndArraySpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &FuncSpec{
		Params: []*Field{
			{
				Spec: &ArraySpec{
					Value: &SimpleSpec{
						TypeName: "typeName",
					},
				},
			},
		},
		Results: []*Field{
			{
				Spec: &ArraySpec{
					Value: &SimpleSpec{
						TypeName: "typeName",
					},
				},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithFuncSpecAndArraySpecAndVariadic(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &FuncSpec{
		IsVariadic: true,
		Params: []*Field{
			{
				Spec: &ArraySpec{
					Value: &SimpleSpec{
						TypeName: "typeName",
					},
				},
			},
		},
		Results: []*Field{
			{
				Spec: &ArraySpec{
					Value: &SimpleSpec{
						TypeName: "typeName",
					},
				},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithFuncSpecAndMapSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &FuncSpec{
		Params: []*Field{
			{
				Spec: &MapSpec{
					Key: &SimpleSpec{
						TypeName: "typeName",
					},
					Value: &SimpleSpec{
						TypeName: "typeName",
					},
				},
			},
		},
		Results: []*Field{
			{
				Spec: &MapSpec{
					Key: &SimpleSpec{
						TypeName: "typeName",
					},
					Value: &SimpleSpec{
						TypeName: "typeName",
					},
				},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithFuncSpecAndStructSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &FuncSpec{
		Params: []*Field{
			{
				Spec: &StructSpec{},
			},
		},
		Results: []*Field{
			{
				Spec: &StructSpec{},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithFuncSpecAndInterfaceSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &FuncSpec{
		Params: []*Field{
			{
				Spec: &InterfaceSpec{},
			},
		},
		Results: []*Field{
			{
				Spec: &InterfaceSpec{},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithFuncSpecAndFuncSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &FuncSpec{
		Params: []*Field{
			{
				Spec: &FuncSpec{},
			},
		},
		Results: []*Field{
			{
				Spec: &FuncSpec{},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithFuncSpecAndNilParam(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &FuncSpec{
		Params: []*Field{
			nil,
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Params[0]' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithFuncSpecAndNilResult(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &FuncSpec{
		Results: []*Field{
			nil,
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Results[0]' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithFuncSpecAndInvalidParam(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &FuncSpec{
		Params: []*Field{
			{},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Spec' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithFuncSpecAndInvalidResult(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &FuncSpec{
		Results: []*Field{
			{},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Spec' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithFuncSpecAndEmptyParamsAndIsVariadic(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &FuncSpec{
		IsVariadic: true,
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(fmt.Sprintf("Variable 'Params' must be not empty for variadic %T", entity), actual.Error())
}

func TestEntityValidator_Validate_WithFuncSpecAndInvalidTypeParamSpecAndIsVariadic(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &FuncSpec{
		IsVariadic: true,
		Params: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(fmt.Sprintf("Variable 'Params[0].Spec' has invalid type for variadic '%T'", entity), actual.Error())
}

func TestEntityValidator_Validate_WithFuncSpecAndOneResultWithNameAndSecondResultWithoutName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &FuncSpec{
		Results: []*Field{
			{
				Name: "name",
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Results' must have all fields with names or all without names", actual.Error())
}

func TestEntityValidator_Validate_WithInterfaceSpecAndSimpleSpecFieldSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithInterfaceSpecAndFuncSpecFieldSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &InterfaceSpec{
		Fields: []*Field{
			{
				Name: "name",
				Spec: &FuncSpec{},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithInterfaceSpecAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &InterfaceSpec{}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithInterfaceSpecAndNilField(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &InterfaceSpec{
		Fields: []*Field{
			nil,
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Fields[0]' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithInterfaceSpecAndInvalidField(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &InterfaceSpec{
		Fields: []*Field{
			{},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Spec' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithInterfaceSpecAndSimpleSpecFieldSpecAndFieldName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &InterfaceSpec{
		Fields: []*Field{
			{
				Name: "name",
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Fields[0].Name' must be empty for 'Fields[0].Spec' type *SimpleSpec", actual.Error())
}

func TestEntityValidator_Validate_WithInterfaceSpecAndSimpleSpecFieldSpecAndIsPointer(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName:  "typeName",
					IsPointer: true,
				},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(
		fmt.Sprintf("Variable 'Fields[0].Spec.(%T).IsPointer' must be 'false'", entity.Fields[0].Spec),
		actual.Error(),
	)
}

func TestEntityValidator_Validate_WithInterfaceSpecAndFuncSpecFieldSpecAndEmptyFieldName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: &FuncSpec{},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Fields[0].Name' must be not empty for 'Fields[0].Spec' type *FuncSpec", actual.Error())
}

func TestEntityValidator_Validate_WithInterfaceSpecAndArraySpecFieldSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: &ArraySpec{
					Value: &SimpleSpec{
						TypeName: "typeName",
					},
				},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(fmt.Sprintf("Variable 'Fields[0]' has invalid type '%T'", entity.Fields[0].Spec), actual.Error())
}

func TestEntityValidator_Validate_WithInterfaceSpecAndMapSpecFieldSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: &MapSpec{
					Key: &SimpleSpec{
						TypeName: "typeName",
					},
					Value: &SimpleSpec{
						TypeName: "typeName",
					},
				},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(fmt.Sprintf("Variable 'Fields[0]' has invalid type '%T'", entity.Fields[0].Spec), actual.Error())
}

func TestEntityValidator_Validate_WithInterfaceSpecAndInterfaceSpecFieldSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: &InterfaceSpec{},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(fmt.Sprintf("Variable 'Fields[0]' has invalid type '%T'", entity.Fields[0].Spec), actual.Error())
}

func TestEntityValidator_Validate_WithInterfaceSpecAndStructSpecFieldSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: &StructSpec{},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(fmt.Sprintf("Variable 'Fields[0]' has invalid type '%T'", entity.Fields[0].Spec), actual.Error())
}

func TestEntityValidator_Validate_WithStructSpecAndSimpleSpecFieldSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &StructSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithStructSpecAndArraySpecFieldSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &StructSpec{
		Fields: []*Field{
			{
				Name: "name",
				Spec: &ArraySpec{
					Value: &SimpleSpec{
						TypeName: "typeName",
					},
				},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithStructSpecAndMapSpecFieldSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &StructSpec{
		Fields: []*Field{
			{
				Name: "name",
				Spec: &MapSpec{
					Key: &SimpleSpec{
						TypeName: "typeName",
					},
					Value: &SimpleSpec{
						TypeName: "typeName",
					},
				},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithStructSpecAndStructSpecFieldSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &StructSpec{
		Fields: []*Field{
			{
				Name: "name",
				Spec: &StructSpec{},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithStructSpecAndInterfaceSpecFieldSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &StructSpec{
		Fields: []*Field{
			{
				Name: "name",
				Spec: &InterfaceSpec{},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithStructSpecAndFuncSpecFieldSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &StructSpec{
		Fields: []*Field{
			{
				Name: "name",
				Spec: &FuncSpec{},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithStructSpecAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &StructSpec{}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithStructSpecAndNilField(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &StructSpec{
		Fields: []*Field{
			nil,
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Fields[0]' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithStructSpecAndInvalidFieldSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &StructSpec{
		Fields: []*Field{
			{},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Spec' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithStructSpecAndArraySpecFieldSpecAndEmptyFieldName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &StructSpec{
		Fields: []*Field{
			{
				Spec: &ArraySpec{
					Value: &SimpleSpec{
						TypeName: "typeName",
					},
				},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(
		fmt.Sprintf("Variable 'Fields[0]' with empty 'Name' has invalid type: '%T'", entity.Fields[0].Spec),
		actual.Error(),
	)
}

func TestEntityValidator_Validate_WithStructSpecAndMapSpecFieldSpecAndEmptyFieldName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &StructSpec{
		Fields: []*Field{
			{
				Spec: &MapSpec{
					Key: &SimpleSpec{
						TypeName: "typeName",
					},
					Value: &SimpleSpec{
						TypeName: "typeName",
					},
				},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(
		fmt.Sprintf("Variable 'Fields[0]' with empty 'Name' has invalid type: '%T'", entity.Fields[0].Spec),
		actual.Error(),
	)
}

func TestEntityValidator_Validate_WithStructSpecAndStructSpecFieldSpecAndEmptyFieldName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &StructSpec{
		Fields: []*Field{
			{
				Spec: &StructSpec{},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(
		fmt.Sprintf("Variable 'Fields[0]' with empty 'Name' has invalid type: '%T'", entity.Fields[0].Spec),
		actual.Error(),
	)
}

func TestEntityValidator_Validate_WithStructSpecAndInterfaceSpecFieldSpecAndEmptyFieldName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &StructSpec{
		Fields: []*Field{
			{
				Spec: &InterfaceSpec{},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(
		fmt.Sprintf("Variable 'Fields[0]' with empty 'Name' has invalid type: '%T'", entity.Fields[0].Spec),
		actual.Error(),
	)
}

func TestEntityValidator_Validate_WithStructSpecAndFuncSpecFieldSpecAndEmptyFieldName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &StructSpec{
		Fields: []*Field{
			{
				Spec: &FuncSpec{},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(
		fmt.Sprintf("Variable 'Fields[0]' with empty 'Name' has invalid type: '%T'", entity.Fields[0].Spec),
		actual.Error(),
	)
}

func TestEntityValidator_Validate_WithImport(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Import{
		Alias:     "alias",
		Namespace: "namespace",
		Comment:   "importComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "importAnnotation",
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithImportAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Import{
		Namespace: "namespace",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithImportAndInvalidAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Import{
		Alias:     "+invalid",
		Namespace: "namespace",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Alias' must be valid identifier, actual value: '+invalid'", actual.Error())
}

func TestEntityValidator_Validate_WithImportAndEmptyNamespace(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Import{}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Namespace' must be not empty", actual.Error())
}

func TestEntityValidator_Validate_WithImportGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &ImportGroup{
		Comment: "importGroupComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "importGroupAnnotation",
			},
		},
		Imports: []*Import{
			{
				Namespace: "namespace",
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithImportGroupAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &ImportGroup{}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithImportGroupAndNilImport(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &ImportGroup{
		Imports: []*Import{
			nil,
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Imports[0]' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithImportGroupAndInvalidImport(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &ImportGroup{
		Imports: []*Import{
			{},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Namespace' must be not empty", actual.Error())
}

func TestEntityValidator_Validate_WithConst(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Const{
		Name:    "constName",
		Value:   "constValue",
		Comment: "const\ncomment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "constAnnotation",
			},
		},
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithConstAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Const{
		Name: "constName",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithConstAndEmptyName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Const{
		Name: "",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Name' must be not empty", actual.Error())
}

func TestEntityValidator_Validate_WithConstAndInvalidName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Const{
		Name: "+invalid",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Name' must be valid identifier, actual value: '+invalid'", actual.Error())
}

func TestEntityValidator_Validate_WithConstAndInvalidSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Const{
		Name: "constName",
		Spec: &SimpleSpec{
			TypeName: "+invalid",
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'TypeName' must be valid identifier, actual value: '+invalid'", actual.Error())
}

func TestEntityValidator_Validate_WithConstAndIsPointerSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Const{
		Name: "constName",
		Spec: &SimpleSpec{
			IsPointer: true,
			TypeName:  "typeName",
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(
		fmt.Sprintf("Variable 'Spec.(%T).IsPointer' must be 'false' for %T", &SimpleSpec{}, entity),
		actual.Error(),
	)
}

func TestEntityValidator_Validate_WithConstAndInvalidValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Const{
		Name:  "constName",
		Value: "[invalid",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertType(actual, scanner.ErrorList{})
}

func TestEntityValidator_Validate_WithConstGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &ConstGroup{
		Comment: "comment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "constGroupAnnotation",
			},
		},
		Consts: []*Const{
			{
				Name: "constName",
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithConstGroupAndNilConst(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &ConstGroup{
		Comment: "comment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "constGroupAnnotation",
			},
		},
		Consts: []*Const{
			nil,
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Consts[0]' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithConstGroupAndInvalidConst(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &ConstGroup{
		Comment: "comment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "constGroupAnnotation",
			},
		},
		Consts: []*Const{
			{},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Name' must be not empty", actual.Error())
}

func TestEntityValidator_Validate_WithVarAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Var{
		Name:  "name",
		Value: "value",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithVarAndSimpleSpecSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Var{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithVarAndArraySpecSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Var{
		Name: "name",
		Spec: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithVarAndMapSpecSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Var{
		Name: "name",
		Spec: &MapSpec{
			Key: &SimpleSpec{
				TypeName: "typeName",
			},
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithVarAndStructSpecSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Var{
		Name: "name",
		Spec: &StructSpec{},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithVarAndInterfaceSpecSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Var{
		Name: "name",
		Spec: &InterfaceSpec{},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithVarAndFuncSpecSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Var{
		Name: "name",
		Spec: &FuncSpec{},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithVarAndEmptyName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Var{
		Spec: &FuncSpec{},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Name' must be not empty", actual.Error())
}

func TestEntityValidator_Validate_WithVarAndInvalidName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Var{
		Name: "+invalid",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Name' must be valid identifier, actual value: '+invalid'", actual.Error())
}

func TestEntityValidator_Validate_WithVarAndNilSpecAndEmptyValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Var{
		Name: "name",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(fmt.Sprintf("%T must have not nil 'Spec' or not empty 'Value'", entity), actual.Error())
}

func TestEntityValidator_Validate_WithVarAndInvalidValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Var{
		Name:  "name",
		Value: "[invalid",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertType(actual, scanner.ErrorList{})
}

func TestEntityValidator_Validate_WithVarAndInvalidSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Var{
		Name: "name",
		Spec: &SimpleSpec{},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'TypeName' must be not empty", actual.Error())
}

func TestEntityValidator_Validate_WithVarAndInvalidTypeValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Var{
		Name: "name",
		Spec: "invalid",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(fmt.Sprintf("Variable 'Spec' has invalid type: '%T'", entity.Spec), actual.Error())
}

func TestEntityValidator_Validate_WithVarGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &VarGroup{
		Comment: "varGroupComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "varGroupAnnotation",
			},
		},
		Vars: []*Var{
			{
				Name:  "varName",
				Value: "varValue",
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithVarGroupAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &VarGroup{}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithVarGroupAndNilVar(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &VarGroup{
		Comment: "varGroupComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "varGroupAnnotation",
			},
		},
		Vars: []*Var{
			nil,
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Vars[0]' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithVarGroupAndInvalidVar(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &VarGroup{
		Comment: "varGroupComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "varGroupAnnotation",
			},
		},
		Vars: []*Var{
			{},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Name' must be not empty", actual.Error())
}

func TestEntityValidator_Validate_WithTypeAndSimpleSpecSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Type{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithTypeAndArraySpecSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Type{
		Name: "name",
		Spec: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithTypeAndMapSpecSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Type{
		Name: "name",
		Spec: &MapSpec{
			Key: &SimpleSpec{
				TypeName: "typeName",
			},
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithTypeAndStructSpecSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Type{
		Name: "name",
		Spec: &StructSpec{},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithTypeAndInterfaceSpecSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Type{
		Name: "name",
		Spec: &InterfaceSpec{},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithTypeAndFuncSpecSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Type{
		Name: "name",
		Spec: &FuncSpec{},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithTypeAndEmptyName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Type{
		Spec: &FuncSpec{},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Name' must be not empty", actual.Error())
}

func TestEntityValidator_Validate_WithTypeAndInvalidName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Type{
		Name: "+invalid",
		Spec: &FuncSpec{},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Name' must be valid identifier, actual value: '+invalid'", actual.Error())
}

func TestEntityValidator_Validate_WithTypeAndNilSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Type{
		Name: "name",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Spec' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithTypeAndInvalidSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Type{
		Name: "name",
		Spec: &SimpleSpec{},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'TypeName' must be not empty", actual.Error())
}

func TestEntityValidator_Validate_WithTypeAndInvalidTypeSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Type{
		Name: "name",
		Spec: "invalid",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(fmt.Sprintf("Variable 'Spec' has invalid type: '%T'", entity.Spec), actual.Error())
}

func TestEntityValidator_Validate_WithTypeGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &TypeGroup{
		Comment: "typeGroupComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "typeGroupAnnotation",
			},
		},
		Types: []*Type{
			{
				Name: "name",
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithTypeGroupAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &TypeGroup{}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithTypeGroupAndNilType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &TypeGroup{
		Types: []*Type{
			nil,
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Types[0]' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithTypeGroupAndInvalidType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &TypeGroup{
		Types: []*Type{
			{},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Name' must be not empty", actual.Error())
}

func TestEntityValidator_Validate_WithFunc(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Func{
		Name:    "funcName",
		Content: "funcContent",
		Comment: "funcComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "funcAnnotation",
			},
		},
		Spec: &FuncSpec{},
		Related: &Field{
			Name: "relatedName",
			Spec: &SimpleSpec{
				TypeName: "relatedTypeName",
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithFuncAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Func{
		Name: "funcName",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithFuncAndEmptyName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Func{}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Name' must be not empty", actual.Error())
}

func TestEntityValidator_Validate_WithFuncAndInvalidFuncName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Func{
		Name: "+invalid",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Name' must be valid identifier, actual value: '+invalid'", actual.Error())
}

func TestEntityValidator_Validate_WithFuncAndInvalidSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Func{
		Name: "funcName",
		Spec: &FuncSpec{
			IsVariadic: true,
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(fmt.Sprintf("Variable 'Params' must be not empty for variadic %T", entity.Spec), actual.Error())
}

func TestEntityValidator_Validate_WithFuncAndInvalidRelatedField(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Func{
		Name:    "funcName",
		Spec:    &FuncSpec{},
		Related: &Field{},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Spec' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithFuncAndInvalidRelatedFieldSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Func{
		Name: "funcName",
		Spec: &FuncSpec{},
		Related: &Field{
			Spec: &FuncSpec{},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(
		fmt.Sprintf("Variable 'Related.Spec.(%T)' has invalid type for '%T'", entity.Spec, entity),
		actual.Error(),
	)
}

func TestEntityValidator_Validate_WithFuncAndInvalidRelatedFieldSpecPackageName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Func{
		Name: "funcName",
		Spec: &FuncSpec{},
		Related: &Field{
			Spec: &SimpleSpec{
				PackageName: "packageName",
				TypeName:    "typeName",
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(
		fmt.Sprintf("Variable 'Related.Spec.(%T).PackageName' must be empty for '%T'", entity.Related.Spec, entity),
		actual.Error(),
	)
}

func TestEntityValidator_Validate_WithFuncAndInvalidContent(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Func{
		Name:    "funcName",
		Spec:    &FuncSpec{},
		Content: "[[",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertType(actual, scanner.ErrorList{})
}

func TestEntityValidator_Validate_WithFile(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		Content:     "package filePackage",
		Comment:     "fileComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "fileAnnotation",
			},
		},
		ImportGroups: []*ImportGroup{
			{
				Comment: "importGroup1\ncomment",
			},
			{
				Comment: "importGroup2\ncomment",
			},
		},
		ConstGroups: []*ConstGroup{
			{
				Comment: "constGroup1\ncomment",
			},
			{
				Comment: "constGroup2\ncomment",
			},
		},
		VarGroups: []*VarGroup{
			{
				Comment: "varGroup1\ncomment",
			},
			{
				Comment: "varGroup2\ncomment",
			},
		},
		TypeGroups: []*TypeGroup{
			{
				Comment: "typeGroup1\ncomment",
			},
			{
				Comment: "typeGroup2\ncomment",
			},
		},
		Funcs: []*Func{
			{
				Name: "func1Name",
			},
			{
				Name: "func2Name",
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithFileAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithFileAndEmptyName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &File{
		PackageName: "filePackageName",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Name' must be not empty", actual.Error())
}

func TestEntityValidator_Validate_WithFileAndEmptyPackageName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &File{
		Name: "fileName",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'PackageName' must be not empty", actual.Error())
}

func TestEntityValidator_Validate_WithFileAndInvalidPackageName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &File{
		Name:        "fileName",
		PackageName: "+invalid",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'PackageName' must be valid identifier, actual value: '+invalid'", actual.Error())
}

func TestEntityValidator_Validate_WithFileAndNilImportGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		ImportGroups: []*ImportGroup{
			nil,
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'ImportGroups[0]' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithFileAndInvalidImportGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{},
				},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Namespace' must be not empty", actual.Error())
}

func TestEntityValidator_Validate_WithFileAndNilConstGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		ConstGroups: []*ConstGroup{
			nil,
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'ConstGroups[0]' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithFileAndInvalidConstGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		ConstGroups: []*ConstGroup{
			{
				Consts: []*Const{
					{},
				},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Name' must be not empty", actual.Error())
}

func TestEntityValidator_Validate_WithFileAndNilVarGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		VarGroups: []*VarGroup{
			nil,
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'VarGroups[0]' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithFileAndInvalidVarGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		VarGroups: []*VarGroup{
			{
				Vars: []*Var{
					{},
				},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Name' must be not empty", actual.Error())
}

func TestEntityValidator_Validate_WithFileAndNilTypeGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		TypeGroups: []*TypeGroup{
			nil,
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'TypeGroups[0]' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithFileAndInvalidTypeGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		TypeGroups: []*TypeGroup{
			{
				Types: []*Type{
					{},
				},
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Name' must be not empty", actual.Error())
}

func TestEntityValidator_Validate_WithFileAndNilFunc(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		Funcs: []*Func{
			nil,
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Funcs[0]' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithFileAndInvalidFunc(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		Funcs: []*Func{
			{},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Name' must be not empty", actual.Error())
}

func TestEntityValidator_Validate_WithFileAndInvalidContent(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		Content:     "+invalid",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertType(actual, scanner.ErrorList{})
}

func TestEntityValidator_Validate_WithNamespace(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Namespace{
		Name: "namespace/alias",
		Path: "/namespace/path",
		Files: []*File{
			{
				Name:        "file1Name",
				PackageName: "filePackageName",
			},
			{
				Name:        "file2Name",
				PackageName: "filePackageName",
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithNamespaceAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Namespace{
		Name: "namespace/alias",
		Path: "/namespace/path",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithNamespaceAndEmptyName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Namespace{
		Path: "/namespace/path",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Name' must be not empty", actual.Error())
}

func TestEntityValidator_Validate_WithNamespaceAndEmptyPath(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Namespace{
		Name: "/namespace/alias",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Path' must be not empty", actual.Error())
}

func TestEntityValidator_Validate_WithNamespaceAndInvalidPath(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Namespace{
		Name: "namespace/alias",
		Path: "namespace/path",
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Path' must be absolute path, actual value: 'namespace/path'", actual.Error())
}

func TestEntityValidator_Validate_WithNamespaceAndNilFile(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Namespace{
		Name: "namespace/alias",
		Path: "/namespace/path",
		Files: []*File{
			nil,
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Files[0]' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithNamespaceAndInvalidFile(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Namespace{
		Name: "namespace/alias",
		Path: "/namespace/path",
		Files: []*File{
			{},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Name' must be not empty", actual.Error())
}

func TestEntityValidator_Validate_WithNamespaceAndDuplicateFileName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Namespace{
		Name: "namespace/alias",
		Path: "/namespace/path",
		Files: []*File{
			{
				Name:        "fileName",
				PackageName: "filePackageName",
			},
			{
				Name:        "fileName",
				PackageName: "filePackageName",
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Namespace has duplicate file name: fileName", actual.Error())
}

func TestEntityValidator_Validate_WithNamespaceAndDifferentPackageNames(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Namespace{
		Name: "namespace/alias",
		Path: "/namespace/path",
		Files: []*File{
			{
				Name:        "fileName1",
				PackageName: "filePackageName1",
			},
			{
				Name:        "fileName2",
				PackageName: "filePackageName2",
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Namespace has different packages", actual.Error())
}

func TestEntityValidator_Validate_WithStorage(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace1/alias",
				Path: "/namespace1/path",
			},
			{
				Name: "namespace2/alias",
				Path: "/namespace2/path",
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithStorageAndEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Storage{}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNil(actual)
}

func TestEntityValidator_Validate_WithStorageAndNilNamespace(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Storage{
		Namespaces: []*Namespace{
			nil,
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Namespaces[0]' must be not nil", actual.Error())
}

func TestEntityValidator_Validate_WithStorageAndInvalidNamespace(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Storage{
		Namespaces: []*Namespace{
			{},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Variable 'Name' must be not empty", actual.Error())
}

func TestEntityValidator_Validate_WithStorageAndDuplicateNamespaceName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace/alias",
				Path: "/namespace1/path",
			},
			{
				Name: "namespace/alias",
				Path: "/namespace2/path",
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Storage has duplicate namespace 'Name': 'namespace/alias'", actual.Error())
}

func TestEntityValidator_Validate_WithStorageAndDuplicateNamespacePath(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace1/alias",
				Path: "/namespace/path",
			},
			{
				Name: "namespace2/alias",
				Path: "/namespace/path",
			},
		},
	}

	actual := (&EntityValidator{}).Validate(entity)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame("Storage has duplicate namespace 'Path': '/namespace/path'", actual.Error())
}

func TestEntityValidator_Validate_WithUnexpectedEntity(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := "data"

	ctrl.Subtest("").
		Call((&EntityValidator{}).Validate, entity).
		ExpectPanic(NewErrorMessageConstraint("Can't validate entity with type: '%T'", entity))
}
