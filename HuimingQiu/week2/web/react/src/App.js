import React, { useRef, useState } from 'react'

function App() {
  const [roomId, setRoomId] = useState('')
  const [userId, setUserId] = useState('')
  const [buttonDisable, setButtonDisable] = useState(true)
  const enterButton = useRef(null)

  const changeButton = () => {
    if (roomId !== '' && userId !== '') {
      enterButton.current.style.backgroundColor = '#5286ed'
      setButtonDisable(false)
    } else {
      enterButton.current.style.backgroundColor = '#c8c8c8'
      setButtonDisable(true)
    }
  }

  const isValidString = (s) => {
    return /[_a-zA-Z0-9]+/.test(s)
  }

  const onRoomIdInput = (e) => {
    setRoomId(e.target.value)
    changeButton()
  }

  const onUserIdInput = (e) => {
    setUserId(e.target.value)
    changeButton()
  }

  const onButtonClick = () => {
    if (isValidString(roomId) && isValidString(userId)) {
      alert('登录成功')
    } else {
      alert('登录失败')
    }
  }

  return (
    <div className='app'>
      <div className='login-container'>
        <h3 className='login-title'>登录</h3>
        <input
          type='text'
          className='room-id'
          placeholder='房间ID'
          onInput={onRoomIdInput}
        />
        <input
          type='text'
          className='user-id'
          placeholder='用户ID'
          onInput={onUserIdInput}
        />
        <button
          className='enter-button'
          ref={enterButton}
          onClick={onButtonClick}
          disabled={buttonDisable}
        >
          进入房间
        </button>
      </div>
    </div>
  )
}

export default App
