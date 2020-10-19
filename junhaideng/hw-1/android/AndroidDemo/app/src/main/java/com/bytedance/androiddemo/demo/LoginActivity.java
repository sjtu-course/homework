package com.bytedance.androiddemo.demo;

import android.content.Intent;
import android.content.SharedPreferences;
import android.os.Bundle;
import android.view.View;
import android.widget.Button;
import android.widget.EditText;

import androidx.appcompat.app.AppCompatActivity;

import com.bytedance.androiddemo.R;
import com.bytedance.androiddemo.utils.InitReaderWriter;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.io.InputStreamReader;

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

        // 从自定义的配置文件中读取数据，然后进行判断，如果有值则添加到控件里
        InitReaderWriter.Info info = InitReaderWriter.initRead(getApplicationContext().getFilesDir().getPath()+File.separator+"webrtc.ini");
            if(!info.uid.isEmpty()){
                mEtUserName.setText(info.uid);
            }
            if(!info.rid.isEmpty()){
                mEtRoomID.setText(info.rid);
            }


        // 看到另外一位同学提到的 SharedPreferences
        // 上网查了一些资料，比较类似于Object-c NSUserDefault
        // 适合少量数据的持久化
        // 下面也贴出相对应的代码
//        SharedPreferences sp = getSharedPreferences("webrtc", MODE_PRIVATE);
//        String rid = sp.getString("rid", "");
//        String uid = sp.getString("uid", "");
//
//        if(!rid.isEmpty()){
//            mEtRoomID.setText(rid);
//        }
//        if(!uid.isEmpty()){
//            mEtUserName.setText(uid);
//        }

        // 3. 监听按钮点击
        mBtnJoin.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                // 4. 获取用户的输入
                String roomId = mEtRoomID.getText().toString();

                // 注意将这里进行修改，之前获取到的内容都是mEtRoomID，bug hhh~~
                String userName = mEtUserName.getText().toString();
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