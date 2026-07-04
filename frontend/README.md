# Repo layout and important directories

TODO

# Tech Stack
- [Typescript](https://github.com/microsoft/typescript): Programming Language
- [Vite](https://github.com/vitejs/vite): Frontend build tool
- [React](https://github.com/react/react): UI library
- [Tanstack Router](https://github.com/tanstack/router): Routing library
- [TanStack Query](https://github.com/tanstack/query): Data fetching and caching library
- [Oxlint](https://github.com/oxc-project/oxc): Linter for TypeScript and JavaScript

# How to run the project

This template provides a minimal setup to get React working in Vite with HMR and some Oxlint rules.

Currently, two official plugins are available:

- [@vitejs/plugin-react](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react) uses [Oxc](https://oxc.rs)
- [@vitejs/plugin-react-swc](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react-swc) uses [SWC](https://swc.rs/)

## React Compiler

The React Compiler is enabled on this template. See [this documentation](https://react.dev/learn/react-compiler) for more information.

Note: This will impact Vite dev & build performances.

## Expanding the Oxlint configuration

If you are developing a production application, we recommend enabling type-aware lint rules by installing `oxlint-tsgolint` and editing `.oxlintrc.json`:

```json
{
  "$schema": "./node_modules/oxlint/configuration_schema.json",
  "plugins": ["react", "typescript", "oxc"],
  "options": {
    "typeAware": true
  },
  "rules": {
    "react/rules-of-hooks": "error",
    "react/only-export-components": ["warn", { "allowConstantExport": true }]
  }
}
```

See the [Oxlint rules documentation](https://oxc.rs/docs/guide/usage/linter/rules) for the full list of rules and categories.

# Build, test, and lint commands

TODO

# Engineering conventions and PR expectations

TODO

# Constraints and do-not rules

TODO

# What done means and how to verify work

TODO
