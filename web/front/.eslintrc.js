module.exports = {
    root: true,
    env: {
        node: true
    },
    'extends': [
        'plugin:vue/essential',
        'eslint:recommended'
    ],
    parserOptions: {
        parser: '@babel/eslint-parser'
    },
    rules: {
        'no-console': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
        'no-debugger': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
        "vue/multi-word-component-names": 0,
        'no-multiple-empty-lines': 1,
        'no-unused-vars': 0,
        'space-before-function-paren': 0,
        'space-before-blocks': 1,
        'no-trailing-spaces': 0,
        'no-multi-spaces': 1,
        'comma-dangle': 0,

        'vue/comment-directive': 0,
        'eol-last': 0,
    }
}
