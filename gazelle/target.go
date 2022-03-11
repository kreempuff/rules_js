package gazelle

import (
	"github.com/bazelbuild/bazel-gazelle/config"
	"github.com/bazelbuild/bazel-gazelle/rule"
	"github.com/emirpasic/gods/sets/treeset"
	godsutils "github.com/emirpasic/gods/utils"
)

// targetBuilder builds targets to be generated by Gazelle.
type targetBuilder struct {
	kind         string
	name         string
	projectRoot  string
	bzlPackage   string
	srcs         *treeset.Set
	deps         *treeset.Set
	resolvedDeps *treeset.Set
	visibility   *treeset.Set
}

// newTargetBuilder constructs a new targetBuilder.
func newTargetBuilder(kind, name, projectRoot, bzlPackage string) *targetBuilder {
	return &targetBuilder{
		kind:         kind,
		name:         name,
		projectRoot:  projectRoot,
		bzlPackage:   bzlPackage,
		srcs:         treeset.NewWith(godsutils.StringComparator),
		deps:         treeset.NewWith(moduleComparator),
		resolvedDeps: treeset.NewWith(godsutils.StringComparator),
		visibility:   treeset.NewWith(godsutils.StringComparator),
	}
}

// addSrcs copies all values from the provided srcs to the target.
func (t *targetBuilder) addSrcs(srcs *treeset.Set) *targetBuilder {
	it := srcs.Iterator()
	for it.Next() {
		t.srcs.Add(it.Value().(string))
	}
	return t
}

// addModuleDependencies copies all values from the provided deps to the target.
func (t *targetBuilder) addModuleDependencies(deps *treeset.Set) *targetBuilder {
	it := deps.Iterator()
	for it.Next() {
		//TODO t.deps.Add(it.Value().(module))
	}
	return t
}

// addVisibility adds a visibility to the target.
func (t *targetBuilder) addVisibility(visibility string) *targetBuilder {
	t.visibility.Add(visibility)
	return t
}

// build returns the assembled *rule.Rule for the target.
func (t *targetBuilder) build() *rule.Rule {
	r := rule.NewRule(t.kind, t.name)
	if !t.srcs.Empty() {
		r.SetAttr("srcs", t.srcs.Values())
	}
	if !t.visibility.Empty() {
		r.SetAttr("visibility", t.visibility.Values())
	}
	if !t.deps.Empty() {
		r.SetPrivateAttr(config.GazelleImportsKey, t.deps)
	}
	r.SetPrivateAttr(resolvedDepsKey, t.resolvedDeps)
	return r
}

// module represents a fully-qualified, dot-separated, Python module as seen on
// the import statement, alongside the line number where it happened.
type module struct {
	// The fully-qualified, dot-separated, Python module name as seen on import
	// statements.
	Name string `json:"name"`
	// The line number where the import happened.
	LineNumber uint32 `json:"lineno"`
	// The path to the module file relative to the Bazel workspace root.
	Filepath string `json:"filepath"`
}

// moduleComparator compares modules by name.
func moduleComparator(a, b interface{}) int {
	return godsutils.StringComparator(a.(module).Name, b.(module).Name)
}
