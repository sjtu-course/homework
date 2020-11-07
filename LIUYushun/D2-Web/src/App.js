import React from 'react';
import styled from "styled-components";

const LoginFrame = styled.div`
    width: 425px;
    height: 380px; 
    margin: 0px auto;
    border-radius: 15px;
    background: #FFFFFF;
    align-self:center;
`

const LoginTitle = styled.div`
    font-size: 32px;
    color: black;
`

const LoginInput = styled.input`
    display: block;
    width: 310px;
    margin: 0px 0px 25px 0px;
    padding: 1px 4px 1px 4px;
    color: #c1c1c1;
    font-size: 16px;
    border-color: #c1c1c1;
    border-style: solid;
    border-width: 1px;
    border-radius: 2px;
`

const LoginButton = styled.button`
    display: block;
    width: 100%;
    height: 35px;
    margin-top: 40px;
    color: white;
    font-size: 18px;
    border-style: none;
    border-width: 0px;
    border-radius: 2px;
    
    background-color: #5268ed;
    :disabled, [disabled]{
        background-color: #c8c8c8;
    }
`

class App extends React.Component {
    constructor(props) {
        super(props)
        this.state = {room: "", user: ""}
        this.loginRef = React.createRef();
    }

    render() {
        return (
            <div style={{"display": "flex", "width": "100vw", "height": "100vh"}}>
                <LoginFrame>
                    <div style={{"padding": "45px 55px 0px 50px"}}>
                        <div style={{"height": "75px"}}><LoginTitle>登录</LoginTitle></div>
                        <LoginInput placeholder={"房间ID"} style={{"height": "31px"}}
                                    onChange={e => this.setState({room: e.target.value})}
                        />
                        <LoginInput placeholder={"用户ID"} style={{"height": "36px"}}
                                    onChange={e => this.setState({user: e.target.value})}
                        />
                        <LoginButton
                            ref={this.loginRef}
                            disabled={((this.state.room === "") || (this.state.user === ""))}
                            onClick={() => this.login()}
                        >
                            进入房间
                        </LoginButton>
                    </div>
                </LoginFrame>
            </div>
        );
    }

    login() {
        let pattern = /[^\w]+/g;
        let fail = (pattern.test(this.state.room) || pattern.test(this.state.user));
        if (!fail) alert("成功"); else alert("失败");
    }
}

export default App;
