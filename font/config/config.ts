import { defineConfig } from '@umijs/max';
import routes from './routes';
import proxy from './proxy';

const publicPath = '/web';

export default defineConfig({
  layout: {
    title: 'EVA',
    locale: false,
    logo: null,
    
  },
  // alias: {
  //   "@": "src"
  // },
  model: {},
  initialState: {},
  request: {},
  // 路由配置
  routes,
  // 本地代理：只在dev环境有效
  proxy,
  // 构建配置
  codeSplitting: {
    jsStrategy: 'granularChunks'
  },
  // compilerOptions: {
  //   "baseUrl": "."
  // },
  hash: true,
  npmClient: 'pnpm',
  antd: {
    theme: { 
      token:{
        "colorPrimary": "#eb2f96",
        "colorSuccess": "#eb2f96",
        "wireframe": false,
        // "colorBgBase": "#fff1f9",
        "colorTextBase": "#000000"
      }
     },
  }

});