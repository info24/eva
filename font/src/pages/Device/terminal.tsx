import React, { useEffect, useState, useRef, useReducer } from 'react';
import { Button, Modal, Form, message, Badge, BadgeProps, Divider } from 'antd';
import { PageContainer, ProCard, ProTable, ProForm, ProFormText, ProFormSelect } from '@ant-design/pro-components';
import {useLocation, useSearchParams, useParams   } from '@umijs/max'
import style from "./terminal.less"
import {getToken, Authorization} from '@/config/token'

import { Terminal } from "xterm"
import { FitAddon } from "xterm-addon-fit"

import 'xterm/css/xterm.css';


const FirewallTerm = (props) => {

    const param = useParams()
    const [searchParams,] = useSearchParams()
    console.log(searchParams.get('id'));

    const key: string = "terminalKey";
    const devRef: any = useRef(null);

    const [status, setStatus] = useState<any>("processing")
    let websocket: WebSocket;
    console.log(Math.floor(window.innerWidth/9), Math.floor(window.innerHeight/20))
    let  terminal = useRef(new Terminal({
        // rendererType: 'canvas',
        disableStdin: false,
        cursorBlink: true, //光标闪烁
        convertEol: true,
        cols: Math.floor(window.innerWidth/9)-1,
        rows: Math.floor(window.innerHeight/20) -1,
        windowsMode: true,
        theme: {
            // foreground: '#000000', //字体
            // background: '#ffffff', //背景色
            foreground: "#ECECEC", //字体
            background: "#000000", //背景色
            black: "\x1b[30m",
        },
    }));

    const InitTerminal = () => {
        console.log("init termainl", getToken())
    

        websocket = new WebSocket(`ws://${window.location.host}/ws/${searchParams.get("id")}?row=${Math.floor(window.innerHeight/20)-1}&col=${Math.floor(window.innerWidth/9-1)}&token=${getToken()}`);
        websocket.addEventListener('beforeSend', function(event: any) {
            console.log("beforeSend")
            event.target?.setRequestHeader(Authorization, getToken())
        })
        websocket.onopen = function(){
            console.log("connection ws done.")
            setStatus("success")
            terminal.current.reset()
            terminal.current.open(devRef.current)
            const fitAddon = new FitAddon()

            terminal.current.loadAddon(fitAddon);
            fitAddon.fit();
            terminal.current.writeln("welcome web terminal \r\ninit done  " + new Date())
        }
        websocket.onmessage = function(e){
            terminal.current.write(e.data)
        }

        function sendPing(){
            websocket.send(new Uint8Array([0x9]));
        }
        setInterval(sendPing, 10000);

        var col = terminal.current.cols
        var row = terminal.current.rows
        console.log(col, row)

        websocket.onclose = function(e) {
            console.log("ws closed, ", e)
            // Modal.warning({
            //   title: 'connecton is closed.',
            //   okText: 'ok'
            // })
            setStatus("error")
            message.error("please check token or contact the administrator. ", 10)
        }
        terminal.current.onData(e => websocket.send(e))


    }

    useEffect(() => {
        document.title = `${searchParams.get("ip")} - ${searchParams.get("name")}`
        InitTerminal()
        return function cleanUp(){
            console.log("clear....")
            websocket.close(1000, "client close request.");
            terminal.current.dispose()
        }
    }, [])

    return (<>

        <div className={style.termContain}  key={key} ref={devRef} />
    </>)
}

export default FirewallTerm;