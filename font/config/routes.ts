import { layout } from '../src/app';
export default [
  {
    path: '/',
    redirect: '/device',
  },
  {
    path: '/device',
    name: 'device',
    component: './Device/device'
  },
  {
    path: '/login',
    name: 'login',
    component: './Login',
    layout: false
  },

  {
    path: '/device/terminal',
    name: 'deviceTerminal',
    component: "./Device/terminal",
    layout: false
  }

];