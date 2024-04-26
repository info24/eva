import { setToken } from '@/config/token';
import { ProCard, ProForm, ProFormText } from '@ant-design/pro-components';
import { request, history } from '@umijs/max';
import {Button, Form, Space, Spin, message} from "antd";
import React from 'react';


const LoginPage = () => {
    const [form] = Form.useForm<{username: string, password: string}>();
    return <div style={{display: 'flex', justifyContent: 'center', height: "-webkit-fill-available", background: "gainsboro", padding: 20}}>
        <ProCard style={{width: '50%', height: 300}} boxShadow title="eva ssh">
            <ProForm form={form} onFinish={async (values) => {
                console.log("llll")
                request(`/user/login`, {method: "POST", data: values}).then(res => {
                    if (res.code == 200) {
                        history.push({
                            pathname: "/"
                        })
                        setToken(res.token)
                    }
                }).catch(err => {
                    message.error("username or password error")
                    console.error(err)
                })
            }}>
                <ProFormText name="username" label="username" />
                <ProFormText.Password name="password"  label="password" />
            </ProForm>
        </ProCard>
    </div>
};

export default LoginPage;