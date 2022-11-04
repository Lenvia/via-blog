module.exports = {
  root: true,
  env: {
    node: true
  },
  extends: [
    'plugin:vue/essential',
    '@vue/standard'
  ],
  plugins: [
    'vue'
  ],
  parserOptions: {
    parser: '@babel/eslint-parser'
  },
  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'no-multiple-empty-lines': 1,
    'no-unused-vars': 0,
    'vue/multi-word-component-names': 0,

    'space-before-function-paren': 0,
    'space-before-blocks': 1,
    'no-trailing-spaces': 0,
    'no-multi-spaces': 1,
    'comma-dangle': 0,

    'vue/comment-directive': 0,
    'eol-last': 0,
    // 'vue/comment-directive': [2,
    //   {
    //     // eslint-disable HTML注释,没有开启该禁用的规则则报错
    //     reportUnusedDisableDirectives: true
    //   }
    // ],

  }
}
