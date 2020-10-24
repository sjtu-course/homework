package com.bytedance.androiddemo.utils;

import android.util.Log;

import java.io.*;

public class InitReaderWriter {
    public static void initWrite(String path, String rid, String uid){
        // 读取到文件之后我们就可以进行保存了
        File file = new File(path);
        // 无论我们读取到的内容是否与保存的内容有关，我们都进行修改
        try {
            file.createNewFile();
        } catch (Exception e) {
            e.printStackTrace();
            Log.println(Log.ERROR, "ChatActivity", "创建文件失败");
            return;
        }

        if (file.canWrite()){
            try {
                FileOutputStream output = new FileOutputStream(file);
                output.write("rid=".getBytes());
                output.write(rid.getBytes());
                output.write('\n');
                output.write("uid=".getBytes());
                output.write(uid.getBytes());
                output.close();
            } catch (Exception e) {
                Log.println(Log.ERROR, "ChatActivity", "写入文件失败");
            }
        }
    }
    public static class  Info{
        public String uid;
        public String rid ;
        public Info(String rid, String uid){
            this.rid = rid;
            this.uid = uid;
        }
    }

    public static Info initRead(String path){
        Info info = new Info("", "");
        // 初始化房间名和用户名
        File file = new File(path);
        if (file.exists()) {
            // 由于已经检查了文件是否存在，所以下面这句其实不需要进行异常判断
            FileInputStream input = null;
            try {
                input = new FileInputStream(file);
            } catch (FileNotFoundException e) {
                e.printStackTrace();
            }
            BufferedReader reader = new BufferedReader(new InputStreamReader(input));
            try {
                String line;
                while ((line = reader.readLine()) != null) {
                    String[] list = line.split("=");
                    String key = list[0];
                    String value = list[1];
                    if(key.equals("uid")){
                        info.uid = value;
                    }else if(key.equals("rid")){
                        info.rid = value;
                    }
                }
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
        return info;
    }
}
