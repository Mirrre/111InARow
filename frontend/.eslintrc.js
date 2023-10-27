module.exports = {
    // 表示当前目录即为根目录，ESLint 规则将被限制到该目录下
    root: true,
    // env 表示启用 ESLint 检测的环境
    env: {
      // 在 node 环境下启动 ESLint 检测
      node: true
    },
    // ESLint 中基础配置需要继承的配置
    extends: [
      'plugin:vue/vue3-essential',
      '@vue/standard'
    ],
    parserOptions: {
      parser: '@babel/eslint-parser'
    },
    /**
   * 错误级别分为三种：
   * "off" 或 0 - 关闭规则
   * "warn" 或 1 - 开启规则，使用警告级别的错误：warn (不会导致程序退出)
   * "error" 或 2 - 开启规则，使用错误级别的错误：error (当被触发的时候，程序会退出)
   */
    rules: {
      'no-console': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
      'no-debugger': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
      'space-before-function-paren': 'off'
      //  eslint-plugin-vue  @8 版本中新增了要求组件名称以驼峰格式命名
      // 'vue/multi-word-component-names': 'off',
      // 'quotes': 'off'
    }
  }