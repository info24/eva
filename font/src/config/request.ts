import React from 'react';
import { notification, Modal, message as msg } from 'antd';
import { ConsoleSqlOutlined } from '@ant-design/icons';
import {Authorization} from "./token";
import { history } from '@umijs/max';

const codeMessage = {
  200: '服务器成功返回请求的数据。',
  201: '新建或修改数据成功。',
  202: '一个请求已经进入后台排队（异步任务）。',
  204: '删除数据成功。',
  400: '发出的请求有错误，服务器没有进行新建或修改数据的操作。',
  401: '用户没有权限（令牌、用户名、密码错误）。',
  403: '用户得到授权，但是访问是被禁止的。',
  404: '发出的请求针对的是不存在的记录，服务器没有进行操作。',
  406: '请求的格式不可得。',
  410: '请求的资源被永久删除，且不会再得到的。',
  422: '当创建一个对象时，发生一个验证错误。',
  500: '服务器发生错误，请检查服务器。',
  502: '网关错误。',
  503: '服务不可用，服务器暂时过载或维护。',
  504: '网关超时。',
};

const reg = new RegExp(
  /^(http(s)?:\/\/)([a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+|\w+)(:[0-9]{1,5})?\/api/,
);
const description = (
  code: React.ReactText,
  url: string,
  log: { name: string; detailCode: string; detail: string },
) =>
  React.createElement('span', null, [
    codeMessage[code],
    React.createElement(
      'a',
      {
        key: 'error',
        onClick: () => {
          Modal.error({
            title: `${log?.name}: ${log?.detailCode}`,
            content: [
              React.createElement(
                'span',
                {
                  key: 'url',
                  style: {
                    'word-wrap': 'break-word',
                    'word-break': 'break-all',
                  },
                },
                ['URL: ', url],
              ),
              React.createElement(
                'p',
                {
                  key: 'detail',
                },
                log.detail,
              ),
            ],
            centered: true,
          });
        },
      },
      '详情',
    ),
  ]);

// 请求中间件 就是发起请求和响应之后需要统一操作数据就写这
// https://github.com/umijs/umi-request#example-1
const middleware = async (ctx, next: any) => {
  const { params } = ctx.req.options;
  
  // 分页处理
  const { current, pageSize, showSizeChanger, ...rest } = params as {
    current?: number;
    pageSize?: number;
    showSizeChanger?: boolean;
  };

  if (current || pageSize) {
    ctx.req.options.params = {
      ...rest,
      page: current ? current - 1 : 0,
      size: pageSize || 10,
    };
  }

  await next();
  
  const res = await ctx.res;
  if (res.data && res.data.list) {
    // eslint-disable-next-line no-shadow
    const { list, pageSize, count, ...data } = res.data as {
      list: any[];
      pageSize: number;
      count: number;
    };
    const { page } = ctx.req.options?.params as { page?: number };
    res.data = {
      list,
      current: page ? page + 1 : 1,
      pageSize,
      total: count,
      ...data,
    };
  }
  ctx.res = res;
};


// request拦截器, 改变 url 或 options.
const requestInterceptors = (url, options) => {
  console.log("test...", url)
  options.headers["aaa"] = "bbb"
  var auth = sessionStorage.getItem(Authorization);
  if (!auth && url != "/user/login") {
    history.push({
      pathname: "/login"
    })
    return
  } else {
    options.headers[Authorization] = `Bearer ${auth}`
  }
  return {
    url,
    options: { ...options, interceptors: true },
  };
};



const responseInterceptors = async (response) => {
  console.log("response ", response)
  if (response && response.status) {
    const errorText = codeMessage[response.status] || response.statusText;
    const { status, url = '' } = response;
    const ApiURL = "url?.replace(reg, '')";
    const data = response;
    // const data = await response?.clone()?.json() || {};
    if (response.status === 401) {
      history.push({
        pathname: "/login"
      })
      console.log("401 login.")
    } else if (data.code && data.code !== 200) {
      const { code, message, log } = data;
      if ([409].includes(code)) {
        msg.error(message);
      } else {
        notification.error({
          message: `请求错误 ${code}: ${message}`,
          description: description(code, ApiURL, log),
        });
      }
      throw response;
    } else if (![200, 201, 202, 204].includes(response.status)) {
      notification.error({
        message: `请求错误 ${status}: ${ApiURL}`,
        description: errorText,
      });
      throw response;
    }
    console.log("url: ", ApiURL)
  } else if (!response) {
    notification.error({
      description: '您的网络发生异常，无法连接服务器',
      message: '网络异常',
    });
    throw response;
  }

  return response;
  // console.log(result)
  // var {code="-1", msg="rest error"} = result
  // if (code == -1) {
  //   notification.error({
  //     message: `request error ${code}: ${msg}`,
  //     description: `server error: ${msg}`
  //   })
  //   throw response
  // }
  // return result;
};

const request = {
  credentials: 'include', // 默认请求是否带上cookie
  prefix: '', // 统一的请求头
  errorHandler: (error) => {
    console.log(error)
    // 集中处理错误
    throw error;
  },
  middlewares: [middleware],
  requestInterceptors: [requestInterceptors],
  responseInterceptors: [responseInterceptors],
};

export default request;

