<!-- Generated with Stardoc: http://skydoc.bazel.build -->


Rules for linking npm dependencies and packaging and linking first-party deps.

Load these with,

```starlark
load("@aspect_rules_js//npm:defs.bzl", "npm_package")
```


<a id="npm_package"></a>

## npm_package

<pre>
npm_package(<a href="#npm_package-name">name</a>, <a href="#npm_package-include_sources">include_sources</a>, <a href="#npm_package-include_transitive_sources">include_transitive_sources</a>, <a href="#npm_package-include_declarations">include_declarations</a>,
            <a href="#npm_package-include_transitive_declarations">include_transitive_declarations</a>, <a href="#npm_package-include_runfiles">include_runfiles</a>, <a href="#npm_package-kwargs">kwargs</a>)
</pre>

TODO

**PARAMETERS**


| Name  | Description | Default Value |
| :------------- | :------------- | :------------- |
| <a id="npm_package-name"></a>name |  Unique name for this target.   |  none |
| <a id="npm_package-include_sources"></a>include_sources |  When True, <code>sources</code> from <code>JsInfo</code> providers in data targets are included in the list of available files to copy.   |  <code>True</code> |
| <a id="npm_package-include_transitive_sources"></a>include_transitive_sources |  When True, <code>transitive_sources</code> from <code>JsInfo</code> providers in data targets are included in the list of available files to copy.   |  <code>True</code> |
| <a id="npm_package-include_declarations"></a>include_declarations |  When True, <code>declarations</code> from <code>JsInfo</code> providers in data targets are included in the list of available files to copy.   |  <code>True</code> |
| <a id="npm_package-include_transitive_declarations"></a>include_transitive_declarations |  When True, <code>transitive_declarations</code> from <code>JsInfo</code> providers in data targets are included in the list of available files to copy.   |  <code>True</code> |
| <a id="npm_package-include_runfiles"></a>include_runfiles |  When True, default runfiles from <code>srcs</code> targets are included in the list of available files to copy.<br><br>This may be needed in a few cases:<br><br>- to work-around issues with rules that don't provide everything needed in sources, transitive_sources, declarations & transitive_declarations - to depend on the runfiles targets that don't use JsInfo<br><br>NB: The default value will be flipped to False in the next major release as runfiles are not needed in the general case and adding them to the list of files available to copy can add noticeable overhead to the analysis phase in a large repository with many npm_package targets.   |  <code>True</code> |
| <a id="npm_package-kwargs"></a>kwargs |  Other attributes   |  none |


<a id="npm_package_lib.implementation"></a>

## npm_package_lib.implementation

<pre>
npm_package_lib.implementation(<a href="#npm_package_lib.implementation-ctx">ctx</a>)
</pre>



**PARAMETERS**


| Name  | Description | Default Value |
| :------------- | :------------- | :------------- |
| <a id="npm_package_lib.implementation-ctx"></a>ctx |  <p align="center"> - </p>   |  none |


<a id="stamped_package_json"></a>

## stamped_package_json

<pre>
stamped_package_json(<a href="#stamped_package_json-name">name</a>, <a href="#stamped_package_json-stamp_var">stamp_var</a>, <a href="#stamped_package_json-kwargs">kwargs</a>)
</pre>

Convenience wrapper to set the "version" property in package.json with the git tag.

In unstamped builds (typically those without `--stamp`) the version will be set to `0.0.0`.
This ensures that actions which use the package.json file can get cache hits.

For more information on stamping, read https://docs.aspect.build/rules/aspect_bazel_lib/docs/stamping.

Using this rule requires that you register the jq toolchain in your WORKSPACE:

```starlark
load("@aspect_bazel_lib//lib:repositories.bzl", "register_jq_toolchains")

register_jq_toolchains()
```


**PARAMETERS**


| Name  | Description | Default Value |
| :------------- | :------------- | :------------- |
| <a id="stamped_package_json-name"></a>name |  name of the resulting <code>jq</code> target, must be "package"   |  none |
| <a id="stamped_package_json-stamp_var"></a>stamp_var |  a key from the bazel-out/stable-status.txt or bazel-out/volatile-status.txt files   |  none |
| <a id="stamped_package_json-kwargs"></a>kwargs |  additional attributes passed to the jq rule, see https://docs.aspect.build/rules/aspect_bazel_lib/docs/jq   |  none |


