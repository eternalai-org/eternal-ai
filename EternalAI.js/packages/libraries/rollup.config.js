const resolve = require('@rollup/plugin-node-resolve').default;
const commonjs = require('@rollup/plugin-commonjs');
const typescript = require('@rollup/plugin-typescript');
const { terser } = require('rollup-plugin-terser');
const pkg = require('./package.json');

const json = require('@rollup/plugin-json');
const builtins = require('rollup-plugin-node-builtins');
const globals = require('rollup-plugin-node-globals');
const inject = require('@rollup/plugin-inject');

module.exports = {
  input: ['./src/index.ts'],
  output: [
    {
      dir: 'dist',
      format: 'cjs',
      sourcemap: false,
    },
  ],
  plugins: [
    resolve({}),
    inject({}),
    commonjs(),
    globals(),
    builtins(),
    typescript({ tsconfig: './tsconfig-build.json' }),
    terser(),
    json(),
  ],
  external: [...Object.keys(pkg.peerDependencies || {}), 'ethers'],
};
