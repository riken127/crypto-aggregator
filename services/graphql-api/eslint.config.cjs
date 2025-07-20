const js = require("@eslint/js");
const globals = require("globals");

module.exports = [
  {
    files: ["**/*.{js,cjs}"],
    plugins: { js },
    languageOptions: {
      sourceType: "commonjs",
      globals: globals.node,
      parserOptions: js.configs.recommended.parserOptions,
    },
    rules: {
      ...js.configs.recommended.rules,
    },
  },
];
