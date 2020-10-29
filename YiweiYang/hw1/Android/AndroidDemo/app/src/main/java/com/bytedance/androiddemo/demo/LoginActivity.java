package com.bytedance.androiddemo.demo;

import androidx.appcompat.app.AppCompatActivity;
import android.content.Context;
import android.content.Intent;
import android.content.SharedPreferences;
import android.os.Bundle;
import android.util.Base64;
import android.view.View;
import android.widget.Button;
import android.widget.EditText;

import com.bytedance.androiddemo.R;

public class LoginActivity extends AppCompatActivity {

    EditText mEtRoomID;
    EditText mEtUserName;
    EditText mEtUserPassword;
    Button mBtnJoin;

    //整合用户信息
    public class UserInfo{
        UserInfo(String rid, String uid, String pw){
            this.rid = rid;
            this.uid = uid;
            this.pw = pw;
        }
        String rid;
        String uid;
        String pw;
    }

    //用于本地存储登录信息
    public class LocalUserInfo {
        SharedPreferences share = null;	//本地存储类
        public LocalUserInfo() {
            share = getSharedPreferences("my_local_data", Context.MODE_PRIVATE);
        }
        //读取本地数据
        public UserInfo getUserInfo(){
            UserInfo userInfo = new UserInfo(
                    new String(Base64.decode(share.getString("rid", "").getBytes(), Base64.DEFAULT)),
                    new String(Base64.decode(share.getString("uid", "").getBytes(), Base64.DEFAULT)),
                    new String(Base64.decode(share.getString("pw", "").getBytes(), Base64.DEFAULT))
            );
            return userInfo;
        }
        //存储本地数据
        public void setUserInfo(UserInfo userInfo){
            SharedPreferences.Editor editor = share.edit();
            editor.putString("rid", Base64.encodeToString(userInfo.rid.getBytes(), Base64.DEFAULT));
            editor.putString("uid", Base64.encodeToString(userInfo.uid.getBytes(), Base64.DEFAULT));
            editor.putString("pw", Base64.encodeToString(userInfo.pw.getBytes(), Base64.DEFAULT));
            editor.commit();
        }
    }


    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        // 1. 关联布局文件
        setContentView(R.layout.activity_login);
        // 2. 获取本地存储信息，初始化控件
        mEtRoomID = findViewById(R.id.et_room_id);
        mEtUserName = findViewById(R.id.et_user_name);
        mEtUserPassword = findViewById(R.id.et_user_password);
        mBtnJoin = findViewById(R.id.btn_join);
        final LocalUserInfo localUserInfo = new LocalUserInfo();
        final UserInfo userInfo = localUserInfo.getUserInfo();
        mEtRoomID.setText(userInfo.rid);
        mEtUserName.setText(userInfo.uid);
        mEtUserPassword.setText(userInfo.pw);
        // 3. 监听按钮点击
        mBtnJoin.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                // 4. 获取用户的输入
                userInfo.rid = mEtRoomID.getText().toString();
                userInfo.uid = mEtUserName.getText().toString();
                userInfo.pw = mEtUserPassword.getText().toString();
                // 5. 创建Intent对象
                Intent intent = new Intent(LoginActivity.this, ChatActivity.class);
                intent.putExtra("rid", userInfo.rid);
                intent.putExtra("uid", userInfo.uid);
                // 6. 存储经过加密后的房间号、用户名、密码信息
                localUserInfo.setUserInfo(userInfo);
                // 7. 通过Intent对象启动 ChatActivity
                startActivity(intent);
            }
        });
    }
}