module.exports = {
  root: true,
  env: {
    node: true
  },
  parser: "vue-eslint-parser",
  // extends: ["plugin:vue/essential", "@vue/airbnb"],
  parserOptions: {
    parser: "babel-eslint",
    sourceType: 'module'
  },
  rules: {
    "no-console": process.env.NODE_ENV === "production" ? "warn" : "off",
    "no-debugger": process.env.NODE_ENV === "production" ? "warn" : "off",
  }
};
