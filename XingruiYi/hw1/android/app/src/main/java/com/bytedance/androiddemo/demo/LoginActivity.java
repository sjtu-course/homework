package com.bytedance.androiddemo.demo;

import androidx.appcompat.app.AppCompatActivity;

import android.content.Context;
import android.content.Intent;
import android.content.SharedPreferences;
import android.os.Bundle;
import android.view.View;
import android.widget.Button;
import android.widget.EditText;

import com.bytedance.androiddemo.R;

public class LoginActivity extends AppCompatActivity {

    EditText mEtRoomID;
    EditText mEtUserName;
    Button mBtnJoin;
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        // 1. 关联布局文件
        setContentView(R.layout.activity_login);
        // 2. 初始化控件
        mEtRoomID = findViewById(R.id.et_room_id);
        mEtUserName = findViewById(R.id.et_user_name);
        mBtnJoin = findViewById(R.id.btn_join);

        // add. 读取preference
        SharedPreferences sharedPref = getPreferences(Context.MODE_PRIVATE);
        mEtRoomID.setText(sharedPref.getString("mEtRoomID", null));
        mEtUserName.setText(sharedPref.getString("mEtUserName", null));

        // 3. 监听按钮点击
        mBtnJoin.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                // 4. 获取用户的输入
                String roomId = mEtRoomID.getText().toString();
                String userName = mEtUserName.getText().toString();

                // add. 写入Preference
                SharedPreferences sharedPref = getPreferences(Context.MODE_PRIVATE);
                SharedPreferences.Editor editor = sharedPref.edit();
                editor.putString("mEtRoomID", roomId);
                editor.putString("mEtUserName", userName);
                editor.apply();

                // 5. 创建Intent对象
                Intent intent = new Intent(LoginActivity.this, ChatActivity.class);
                intent.putExtra("rid", roomId);
                intent.putExtra("uid", userName);
                // 6. 通过Intent对象启动 ChatActivity
                startActivity(intent);
            }
        });
    }
}