package com.bytedance.androiddemo.demo;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.appcompat.app.AppCompatActivity;

import android.content.Context;
import android.content.Intent;
import android.content.SharedPreferences;
import android.os.Bundle;
import android.util.Base64;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ArrayAdapter;
import android.widget.Button;
import android.widget.EditText;
import android.widget.ListView;
import android.widget.TextView;

import com.bytedance.androiddemo.R;

import java.util.ArrayList;
import java.util.Date;
import java.util.List;

public class ChatActivity extends AppCompatActivity implements View.OnClickListener {
    String uid = null;
    String rid = null;

    ListView mLvMsgList;
    EditText mEtMsg;
    Button mBtnSend;

    List<Msg> mMsgs = new ArrayList<>();
    MsgAdapter adapter;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        // 1. 关联布局文件
        setContentView(R.layout.activity_chat);
        // 2. 初始化控件
        initView();
        // 3. 读取Intent内容
        Intent intent = getIntent();
        uid = intent.getStringExtra("uid");
        rid = intent.getStringExtra("rid");
        getSupportActionBar().setTitle("当前房间： " + rid);

    }

    void initView(){
        mEtMsg = findViewById(R.id.et_msg);
        mBtnSend = findViewById(R.id.btn_send);
        mLvMsgList = findViewById(R.id.lv_msg_list);
        // ListView 适配器
        adapter = new MsgAdapter(this, R.layout.layout_msg_item, mMsgs);
        mLvMsgList.setAdapter(adapter);
        // 按钮监听器
        mBtnSend.setOnClickListener(this);
    }

//    //用于本地存储聊天数据
//    public class LocalUserInfo {
//        SharedPreferences share = null;	//本地存储类
//        public LocalUserInfo() {
//            share = getSharedPreferences("my_local_data", Context.MODE_PRIVATE);
//        }
//        //读取本地数据
//        public LoginActivity.UserInfo getUserInfo(){
//            LoginActivity.UserInfo userInfo = new LoginActivity.UserInfo(
//                    new String(Base64.decode(share.getString("rid", "").getBytes(), Base64.DEFAULT)),
//                    new String(Base64.decode(share.getString("uid", "").getBytes(), Base64.DEFAULT)),
//                    new String(Base64.decode(share.getString("pw", "").getBytes(), Base64.DEFAULT))
//            );
//            return userInfo;
//        }
//        //存储本地数据
//        public void setUserInfo(LoginActivity.UserInfo userInfo){
//            SharedPreferences.Editor editor = share.edit();
//            editor.putString("rid", Base64.encodeToString(userInfo.rid.getBytes(), Base64.DEFAULT));
//            editor.putString("uid", Base64.encodeToString(userInfo.uid.getBytes(), Base64.DEFAULT));
//            editor.putString("pw", Base64.encodeToString(userInfo.pw.getBytes(), Base64.DEFAULT));
//            editor.commit();
//        }
//    }

    /**
     * 监听点击事件
     * @param view 控件
     */
    @Override
    public void onClick(View view) {
        switch (view.getId()){
            case R.id.btn_send:
                // 添加消息到消息列表
                Date date = new Date();
                String time = date.toString();

                mMsgs.add(new Msg(mEtMsg.getText().toString(), uid, time));
                //本地保存消息

                // 更新消息列表UI
                adapter.notifyDataSetChanged();
        }
    }

    public static class Msg{
        Msg(String msg, String uid, String time){
            this.msg = msg;
            this.uid = uid;
            this.time = time;
        }
        String msg;
        String uid;
        String time;
    }

    public class MsgAdapter extends ArrayAdapter<Msg>{
        List<Msg> msgs;
        int resource;

        public MsgAdapter(@NonNull Context context, int resource, @NonNull List<Msg> objects) {
            super(context, resource, objects);
            msgs = objects;
            this.resource = resource;
        }


        @NonNull
        @Override
        public View getView(int position, @Nullable View convertView, @NonNull ViewGroup parent) {
            Msg msg = mMsgs.get(position);
            if(convertView == null) {
                convertView = LayoutInflater.from(getContext()).inflate(resource, parent, false);
            }
            TextView tvMsg = convertView.findViewById(R.id.tv_msg);
            tvMsg.setText(msg.msg);
            TextView tvUid = convertView.findViewById(R.id.tv_user_id);
            tvUid.setText("From " + msg.uid);
            TextView tvTime = convertView.findViewById(R.id.tv_time);
            tvTime.setText(msg.time);
            return convertView;
        }
    }
}