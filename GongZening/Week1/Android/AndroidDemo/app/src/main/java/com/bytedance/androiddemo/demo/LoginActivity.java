package com.bytedance.androiddemo.demo;

import androidx.appcompat.app.AppCompatActivity;

import android.annotation.SuppressLint;
import android.content.Context;
import android.content.Intent;
import android.content.SharedPreferences;
import android.os.Bundle;
import android.text.TextUtils;
import android.view.View;
import android.widget.Button;
import android.widget.CheckBox;
import android.widget.EditText;
import android.widget.Toast;

import com.bytedance.androiddemo.R;

public class LoginActivity extends AppCompatActivity {

    EditText mEtRoomID;
    EditText mEtUserName;
    Button mBtnJoin;
    CheckBox checkBox;
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        // 1. 关联布局文件
        setContentView(R.layout.activity_login);
        // 2. 初始化控件
        mEtRoomID = findViewById(R.id.et_room_id);
        mEtUserName = findViewById(R.id.et_user_name);
        mBtnJoin = findViewById(R.id.btn_join);
        checkBox = findViewById(R.id.checkBox);
        // 2.1 初始化信息数据
        SharedPreferences data = getSharedPreferences("data", Context.MODE_PRIVATE);
        String roomdata = data.getString("roomdata","");
        String userdata = data.getString("userdata","");
        boolean checkboxdata = data.getBoolean("checkboxdata",false);
        checkBox.setChecked(checkboxdata);
        mEtRoomID.setText(roomdata);
        mEtUserName.setText(userdata);

        // 3. 监听按钮点击
        mBtnJoin.setOnClickListener(new View.OnClickListener() {
            @SuppressLint("ShowToast")
            @Override
            public void onClick(View view) {
                // 4. 获取用户的输入
                String roomId = mEtRoomID.getText().toString();
                String userName = mEtUserName.getText().toString();
                // 4.1 判断输入情况
                if (TextUtils.isEmpty(roomId)) {
                    Toast.makeText(LoginActivity.this, "请输入房间名", Toast.LENGTH_SHORT);
                } else if (TextUtils.isEmpty(userName)){
                    Toast.makeText(LoginActivity.this, "请输入用户名", Toast.LENGTH_SHORT);
                } else {
                    // 4.2 是否保存输入信息
                    if (checkBox.isChecked()) {
                        // 选中状态，保存数据
                        SharedPreferences data = getSharedPreferences("data", Context.MODE_PRIVATE);
                        @SuppressLint("CommitPrefEdits")
                        SharedPreferences.Editor editor = data.edit();

                        editor.putString("roomdata",roomId);
                        editor.putString("userdata",userName);
                        editor.putBoolean("checkboxdata", true);
                        editor.apply();
                    } else {
                        // 未选中状态，清除数据
                        SharedPreferences data = getSharedPreferences("data", Context.MODE_PRIVATE);
                        @SuppressLint("CommitPrefEdits")
                        SharedPreferences.Editor editor = data.edit();
                        editor.remove("roomdata");
                        editor.remove("userdata");
                        editor.putBoolean("checkboxdata", false);
                        editor.apply();
                    }
                    // 5. 创建Intent对象
                    Intent intent = new Intent(LoginActivity.this, ChatActivity.class);
                    intent.putExtra("rid", roomId);
                    intent.putExtra("uid", userName);
                    // 6. 通过Intent对象启动 ChatActivity
                    startActivity(intent);
                }
            }
        });
    }
}
