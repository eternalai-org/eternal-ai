const resolve = require('@rollup/plugin-node-resolve').default;
const commonjs = require('@rollup/plugin-commonjs');
const typescript = require('@rollup/plugin-typescript');
const { terser } = require('rollup-plugin-terser');
const pkg = require('./package.json');

const json = require('@rollup/plugin-json');
const builtins = require('rollup-plugin-node-builtins');
const globals = require('rollup-plugin-node-globals');
const inject = require('@rollup/plugin-inject');

const { uglify } = require('rollup-plugin-uglify');
import gzipPlugin from 'rollup-plugin-gzip';

module.exports = {
  input: ['./src/index.ts'],
  output: [
    {
      // dir: 'dist',
      format: 'cjs',
      file: 'dist/index.cjs.js',
      sourcemap: false,
    },
    {
      // dir: 'dist',
      format: 'umd',
      file: 'dist/index.umd.js',
      name: 'Interact',
      sourcemap: false,
    },
  ],
  plugins: [
    resolve({}),
    inject({}),
    commonjs(),
    globals(),
    builtins(),
    typescript({ tsconfig: './build.tsconfig.json', declaration: false }),
    terser(),
    json(),
    uglify(),
    gzipPlugin(),
  ],
  external: [...Object.keys(pkg.peerDependencies || {}), 'ethers'],
};
