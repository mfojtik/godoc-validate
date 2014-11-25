godoc-validate
=================

A simple godoc validator that verify all functions in your code
has godoc style comments.

Usage:

```
$ godoc-validate github.com/name/project
```

Sample output:

```
Analyzing template:
  [W] makeParameter() is missing godoc
Processing 'ErrorGenerator' in package template:
  [W] GenerateValue() is missing godoc
Processing 'FooGenerator' in package template:
  [W] GenerateValue() is missing godoc
Processing 'TemplateProcessor' in package template:
  [E] 'BlaProces' godoc name does not match func name Process()
```
