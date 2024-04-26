import {ModalForm, PageContainer, ProColumnType, ProFormSelect, ProFormText, ProTable, TableDropdown} from "@ant-design/pro-components";
import React, { useRef, useState } from "react";
import { request, history } from '@umijs/max';
import {useSearchParams} from "@@/exports";
import {Button, Form, Space, Spin, message} from "antd";
import { CodeOutlined } from "@ant-design/icons";
import Icon from "@ant-design/icons/lib/components/Icon";
import IconFont from "@ant-design/icons/lib/components/IconFont";



export default () => {
    const [searchParams,] = useSearchParams()
    const [form] = Form.useForm<{ID: number; name: string; username: string; password: string; description: string; pty: string}>();
    const ref = useRef<any>();
    const [modal, setModal] = useState<boolean>(false);
    const [loading, setLoading] = useState<boolean>(false);

    const columns = [{
        title: "name",
        dataIndex: "name",
        key: "name"
    }, {
        title: "ip",
        dataIndex: "ip",
        key: "ip"
    }, {
        title: "Username",
        dataIndex: "username"
    }, {
        title: "pty",
        dataIndex: "pty"
    }, {
        title: "Description",
        dataIndex: "description"
    }, {
        title: "action",
        valueType: "action",
        render: (_, record) => <Space size="middle">
                <a onClick={() => {
                    console.log(record.ID)
                    history.push({
                        pathname: `/device/terminal?id=${record.ID}&name=${record.name}&ip=${record.ip}`
                    })
                }}>shell</a>
                <a onClick={() => {
                    form.setFieldsValue(record)
                    setModal(true)
                }}>
                    modify
                </a>
                <TableDropdown key="actionGroup" onSelect={(key) => {
                    if (key == 'delete') {
                        setLoading(true)
                        request(`ssh/device/${record.ID}`,{method: "DELETE"}).then(res => {
                            message.info(record.name + " delete done")
                            ref.current.reload()
                        }).finally(() => {
                            setLoading(false)
                        })
                    }
                }} menus={[{key: 'delete', name: 'delete'}]} />
            </Space>
    }]
    console.log(ref)
    console.log(searchParams.get('id'));
    return <PageContainer title="Device Manager">
        <Spin spinning={loading}>
        <ModalForm open={modal} title="Add Device" form={form} 
            trigger={
                <Button type="primary" onClick={() => {
                    form.resetFields()
                    setModal(true)}
                }> Add Device</Button>
            } 
            onOpenChange={setModal}
            onFinish={async (values) => {
            console.log(values)
            var data;
            try {
                setLoading(true)
                if (values.ID) {
                    data = await request(`/ssh/device/${values.ID}`, {method: "POST", data: values})
                } else {
                    data =await request("/ssh/device", {method: "POST", data: values})
                }
            } catch (err) {
                message.error(" error : " + err)
                console.error("err", err)
                return false
            } finally {
                setLoading(false)
            }
            
            
            if (data.code != 200) {
                message.warning("error: " + data.msg)
                return false;
            }
            
            console.log(data)
            message.success("add success")
            if (ref.current) {
                ref.current.reload()
            }
            return true
        }}>
            <ProFormText name="ID" label="ID" hidden />
            <ProFormText name="name" label="Name" />
            <ProFormText name="ip" label="Host" />
            <ProFormText name="username" label="UserName" />
            <ProFormText name="password" label="PassWord" />
            <ProFormSelect name="pty" label="Pty" options={[{
                value: "xterm",
                label: "xterm"
            }, {
                value: "xterm-256color",
                label: "xterm-256color"
            }, {
                value: "pt100",
                label: "pt100"
            }]} />
            <ProFormText name="description" label="Description" />

        </ModalForm>
        
        <ProTable actionRef={ref} key="ID" columns={columns} request={async (params, sort, filter) => {
            
            var data = await request(`/ssh/device`, {method: "GET",  params: {
                ...params,
              }})
              console.log(data)
            return {
                data: data.data,
                success: true,
                total: data.total
            }
        }} />
        </Spin>
    </PageContainer>
}