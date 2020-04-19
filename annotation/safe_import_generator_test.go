package annotation

import (
	"github.com/pkg/errors"
	"strconv"
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestNewSafeImportGenerator(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	validator := NewValidatorMock(ctrl)

	actual := NewSafeImportGenerator(validator)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(validator, actual.validator)
}

func TestNewSafeImportGenerator_WithNilValidator(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	ctrl.Subtest("").
		Call(NewSafeImportGenerator, nil).
		ExpectPanic(NewErrorMessageConstraint("Variable 'validator' must be not nil"))
}

func TestSafeImportGenerator_Generate(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	namespace := "namespace"

	dstFile := &File{}

	expected := &Import{
		Namespace: namespace,
	}

	validator := NewValidatorMock(ctrl)

	safeImportGenerator := &SafeImportGenerator{validator: validator}

	validator.
		EXPECT().
		Validate(ctrl.Equal(&Import{Namespace: namespace})).
		Return(nil)

	actual := safeImportGenerator.Generate(dstFile, namespace)

	ctrl.AssertEqual(expected, actual)
}

func TestSafeImportGenerator_Generate_WithExistsImport(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	alias := "alias"
	namespace := "namespace"

	dstFile := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Alias:     alias,
						Namespace: namespace,
					},
				},
			},
		},
	}

	expected := &Import{
		Alias:     alias,
		Namespace: namespace,
	}

	validator := NewValidatorMock(ctrl)

	safeImportGenerator := &SafeImportGenerator{validator: validator}

	validator.
		EXPECT().
		Validate(ctrl.Equal(&Import{Namespace: namespace})).
		Return(nil)

	actual := safeImportGenerator.Generate(dstFile, namespace)

	ctrl.AssertEqual(expected, actual)
}

func TestSafeImportGenerator_Generate_WithAliasByPath(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	namespace := "namespace/path1/path2"

	dstFile := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Alias:     "path2",
						Namespace: "path2",
					},
				},
			},
			{
				Imports: []*Import{
					{
						Alias:     "path1_path2",
						Namespace: "path1_path2",
					},
				},
			},
		},
	}

	expected := &Import{
		Alias:     "namespace_path1_path2",
		Namespace: namespace,
	}

	validator := NewValidatorMock(ctrl)

	safeImportGenerator := &SafeImportGenerator{validator: validator}

	validator.
		EXPECT().
		Validate(ctrl.Equal(&Import{Namespace: namespace})).
		Return(nil)

	actual := safeImportGenerator.Generate(dstFile, namespace)

	ctrl.AssertEqual(expected, actual)
}

func TestSafeImportGenerator_Generate_WithAliasByIncrement(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	namespace := "namespace"

	dstFile := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Alias:     "namespace",
						Namespace: "another_namespace",
					},
				},
			},
			{
				Imports: []*Import{
					{
						Alias:     "namespace_0",
						Namespace: "another_namespace_0",
					},
					{
						Alias:     "namespace_1",
						Namespace: "another_namespace_1",
					},
					{
						Alias:     "namespace_2",
						Namespace: "another_namespace_2",
					},
				},
			},
		},
	}

	expected := &Import{
		Alias:     "namespace_3",
		Namespace: namespace,
	}

	validator := NewValidatorMock(ctrl)

	safeImportGenerator := &SafeImportGenerator{validator: validator}

	validator.
		EXPECT().
		Validate(ctrl.Equal(&Import{Namespace: namespace})).
		Return(nil)

	actual := safeImportGenerator.Generate(dstFile, namespace)

	ctrl.AssertEqual(expected, actual)
}

func TestSafeImportGenerator_Generate_WithNilDstFile(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	namespace := "namespace"

	validator := NewValidatorMock(ctrl)

	safeImportGenerator := &SafeImportGenerator{validator: validator}

	ctrl.Subtest("").
		Call(safeImportGenerator.Generate, nil, namespace).
		ExpectPanic(NewErrorMessageConstraint("Variable 'dstFile' must be not nil"))
}

func TestSafeImportGenerator_Generate_WithInvalidNamespace(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	namespace := "namespace"
	err := errors.New("message")

	dstFile := &File{}

	validator := NewValidatorMock(ctrl)

	safeImportGenerator := &SafeImportGenerator{validator: validator}

	validator.
		EXPECT().
		Validate(ctrl.Equal(&Import{Namespace: namespace})).
		Return(err)

	ctrl.Subtest("").
		Call(safeImportGenerator.Generate, dstFile, namespace).
		ExpectPanic(ctrl.Same(err))
}

func TestSafeImportGenerator_Generate_WithIncrementLimitOverflow(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	namespace := "namespace"

	dstFile := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Alias:     "namespace",
						Namespace: "another_namespace",
					},
				},
			},
		},
	}

	for i := 0; i <= maxGenerateImportAliasLevel; i++ {
		addImport := &Import{
			Alias:     "namespace_" + strconv.Itoa(i),
			Namespace: "another_namespace_" + strconv.Itoa(i),
		}

		dstFile.ImportGroups[0].Imports = append(dstFile.ImportGroups[0].Imports, addImport)
	}

	validator := NewValidatorMock(ctrl)

	safeImportGenerator := &SafeImportGenerator{validator: validator}

	validator.
		EXPECT().
		Validate(ctrl.Equal(&Import{Namespace: namespace})).
		Return(nil)

	ctrl.Subtest("").
		Call(safeImportGenerator.Generate, dstFile, namespace).
		ExpectPanic(NewErrorMessageConstraint("Can't generate alias for import %s", namespace))
}

func TestSafeImportGenerator_cleanALias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	alias := "~`!@#$%^&*()_+{}[]\\|?<>,.SomeDataHereWith~`!@#$%^&*()_+{}[]\\|?<>,.Symbol~`!@#$%^&*()_+{}[]\\|?<>,."
	expected := "some_data_here_with_symbol"

	validator := NewValidatorMock(ctrl)

	safeImportGenerator := &SafeImportGenerator{validator: validator}

	actual := safeImportGenerator.cleanAlias(alias)

	ctrl.AssertSame(expected, actual)
}

func TestSafeImportGenerator_generateAlias_WithUpperFirstSymbol(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	alias := "SomeDataHereWith~`!@#$%^&*()_+{}[]\\|?<>,.Symbol~`!@#$%^&*()_+{}[]\\|?<>,."
	expected := "some_data_here_with_symbol"

	validator := NewValidatorMock(ctrl)

	safeImportGenerator := &SafeImportGenerator{validator: validator}

	actual := safeImportGenerator.cleanAlias(alias)

	ctrl.AssertSame(expected, actual)
}

func TestSafeImportGenerator_generateAlias_WithLowerFirstSymbol(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	alias := "someDataHereWith~`!@#$%^&*()_+{}[]\\|?<>,.Symbol~`!@#$%^&*()_+{}[]\\|?<>,."
	expected := "some_data_here_with_symbol"

	validator := NewValidatorMock(ctrl)

	safeImportGenerator := &SafeImportGenerator{validator: validator}

	actual := safeImportGenerator.cleanAlias(alias)

	ctrl.AssertSame(expected, actual)
}
