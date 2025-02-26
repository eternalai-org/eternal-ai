const resolve = require('@rollup/plugin-node-resolve').default;
const commonjs = require('@rollup/plugin-commonjs');
const typescript = require('@rollup/plugin-typescript');
const { terser } = require('rollup-plugin-terser');
const pkg = require('./package.json');

const json = require('@rollup/plugin-json');
const builtins = require('rollup-plugin-node-builtins');
const globals = require('rollup-plugin-node-globals');

const GLOBAL_PACKAGES = {
  ethers: 'ethers',
};

module.exports = {
  input: 'src/index.ts',
  output: [
    {
      file: pkg.main,
      format: 'cjs',
      sourcemap: true,
      globals: GLOBAL_PACKAGES,
    },
    {
      file: pkg.module,
      format: 'es',
      sourcemap: true,
      globals: GLOBAL_PACKAGES,
    },
    {
      file: 'dist/bundle.iife.js',
      format: 'iife',
      name: 'Interact',
      sourcemap: true,
      globals: GLOBAL_PACKAGES,
    },
    {
      file: 'dist/bundle.umd.js',
      format: 'umd',
      name: 'Interact',
      sourcemap: true,
      globals: GLOBAL_PACKAGES,
    },
  ],
  plugins: [
    resolve({
      browser: true,
      preferBuiltins: true,
    }),
    commonjs(),
    globals(),
    builtins(),
    typescript({ tsconfig: './build.tsconfig.json' }),
    terser(),
    json(),
  ],
  external: [...Object.keys(pkg.peerDependencies || {}), 'ethers'],
};
