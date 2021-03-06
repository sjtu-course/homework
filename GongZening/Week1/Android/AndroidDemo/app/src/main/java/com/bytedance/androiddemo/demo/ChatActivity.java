package com.bytedance.androiddemo.demo;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.appcompat.app.AppCompatActivity;

import android.annotation.SuppressLint;
import android.content.Context;
import android.content.Intent;
import android.os.Bundle;
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
import java.util.List;

public class ChatActivity extends AppCompatActivity implements View.OnClickListener {
    String uid = null;
    String rid = null;

    ListView mLvMsgList;
    EditText mEtMsg;
    Button mBtnSend;
    TextView tv_user_id;

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
        getSupportActionBar().setTitle(rid);

    }

    void initView(){
        mEtMsg = findViewById(R.id.et_msg);
        mBtnSend = findViewById(R.id.btn_send);
        mLvMsgList = findViewById(R.id.lv_msg_list);
        tv_user_id = findViewById(R.id.tv_user_id);
        // ListView 适配器
        adapter = new MsgAdapter(this, R.layout.layout_msg_item, mMsgs);
        mLvMsgList.setAdapter(adapter);
        // 按钮监听器
        mBtnSend.setOnClickListener(this);
    }

    /**
     * 监听点击事件
     * @param view 控件
     */
    @Override
    public void onClick(View view) {
        if (view.getId() == R.id.btn_send) {// 添加消息到消息列表
            mMsgs.add(new Msg(mEtMsg.getText().toString(), uid));
            // 更新消息列表UI
            adapter.notifyDataSetChanged();
            // 清空输入信息
            mEtMsg.setText("");
        }
    }

    public static class Msg{
        Msg(String msg, String uid){
            this.msg = msg;
            this.uid = uid;
        }
        String msg;
        String uid;
    }

    public class MsgAdapter extends ArrayAdapter<Msg>{
        List<Msg> msgs;
        int resource;

        public MsgAdapter(@NonNull Context context, int resource, @NonNull List<Msg> objects) {
            super(context, resource, objects);
            msgs = objects;
            this.resource = resource;
        }


        @SuppressLint("SetTextI18n")
        @NonNull
        @Override
        public View getView(int position, @Nullable View convertView, @NonNull ViewGroup parent) {
            Msg msg = mMsgs.get(position);
            if(convertView == null) {
                convertView = LayoutInflater.from(getContext()).inflate(resource, parent, false);
            }
            TextView tv_user_id = convertView.findViewById(R.id.tv_user_id);
            tv_user_id.setText("From " + uid);
            TextView tv = convertView.findViewById(R.id.tv_msg);
            tv.setText(msg.msg);
            return convertView;
        }
    }
}