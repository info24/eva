export default {
  '/ssh': {
    target: 'http://localhost:9999',
    changeOrigin: true,
    // pathRewrite: { '^': '' },
  },
  '/user': {
    target: 'http://localhost:9999',
    changeOrigin: true,
    // pathRewrite: { '^': '' },
  },
};