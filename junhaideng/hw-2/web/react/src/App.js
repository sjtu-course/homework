import React, {createRef} from 'react';
import './App.css';

const ROOM = "room"
const USERNAME = "username"

class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            room: "",
            username: "",
            disable: true,
        }

        this.btn = createRef()
        // 绑定方法
        this.handleClick = this.handleClick.bind(this)
        this.onInput = this.onInput.bind(this)
        this.validate = this.validate.bind(this)
    }

    handleClick(){
        let regexp = /^([a-zA-Z]|[0-9]|_)+$/
        if(regexp.test(this.state.room) && regexp.test(this.state.username)){
            alert("成功")
        }else{
            alert("失败")
        }
    }

    onInput(typ, e){
        if(typ === USERNAME){
            this.setState({
                username: e.target.value,
            },  this.validate)
        }else if (typ === ROOM){
            this.setState({
                room: e.target.value
            },this.validate )
        }
    }

    validate(){
        // 判断是否均有输入
        if(this.state.room && this.state.username){
            this.btn.current.className = "btn success"
            this.setState({
                disable: false
            })
        }else{
            this.btn.current.className = "btn default"
            this.setState({
                disable: true
            })
        }
    }


    render() {
        return <>
            <div className={"container"}>
                <div className={"content"}>
                    <div className="title">
                        登录
                    </div>
                    <div className="form">
                        <div>
                            <input onInput={this.onInput.bind(this, ROOM)} className="input room" type="text" placeholder="房间ID"/>
                        </div>

                        <div>
                            <input onInput={this.onInput.bind(this, USERNAME)} className="input username" type="text" placeholder="用户ID"/>
                        </div>
                        <div>
                            <button onClick={this.handleClick.bind(this)} ref={this.btn} className="btn default" disabled={this.state.disable}>进入房间</button>
                        </div>
                    </div>
                </div>
            </div>
        </>
    }
}

export default App;
