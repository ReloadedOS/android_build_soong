package cc

import (
	"github.com/google/blueprint"

	"android/soong/android"
)

type LinkableInterface interface {
	Module() android.Module
	CcLibrary() bool
	CcLibraryInterface() bool

	OutputFile() android.OptionalPath

	IncludeDirs() android.Paths
	SetDepsInLinkOrder([]android.Path)
	GetDepsInLinkOrder() []android.Path

	HasStaticVariant() bool
	GetStaticVariant() LinkableInterface

	NonCcVariants() bool

	StubsVersions() []string
	BuildStubs() bool
	SetBuildStubs()
	SetStubsVersions(string)
	HasStubsVariants() bool
	SelectedStl() string
	ApiLevel() string

	BuildStaticVariant() bool
	BuildSharedVariant() bool
	SetStatic()
	SetShared()
	Static() bool
	Shared() bool
	Toc() android.OptionalPath

	InRecovery() bool
	OnlyInRecovery() bool

	UseVndk() bool
	MustUseVendorVariant() bool
	IsVndk() bool
	HasVendorVariant() bool

	SdkVersion() string

	ToolchainLibrary() bool
	NdkPrebuiltStl() bool
	StubDecorator() bool

	AllStaticDeps() []string
}

type DependencyTag struct {
	blueprint.BaseDependencyTag
	Name    string
	Library bool
	Shared  bool

	ReexportFlags bool

	ExplicitlyVersioned bool
}

var (
	SharedDepTag = DependencyTag{Name: "shared", Library: true, Shared: true}
	StaticDepTag = DependencyTag{Name: "static", Library: true}

	CrtBeginDepTag = DependencyTag{Name: "crtbegin"}
	CrtEndDepTag   = DependencyTag{Name: "crtend"}
)
