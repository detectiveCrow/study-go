## Summary

- Variable declaration
  - `var foo int`
  - `var foo int = 42`
  - `foo := 42`
- Can't redeclare variables, but can shadow them
- All variables must be used
- Visibility
  - lower case first letter for package level scope
  - upper case first letter to export
  - no private scope
- Naming conventions
  - Pascal or camelCase
    - Capitalize acronyms(HTTP, URL)
  - As short as reasonable
    - longer names for longer life spans
- Type conversions
  - `destinationType(variable)`
  - use strconv package for strings
